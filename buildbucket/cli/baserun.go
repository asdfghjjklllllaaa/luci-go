// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/maruel/subcommands"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/buildbucket/protoutil"
	"go.chromium.org/luci/cipd/version"
	"go.chromium.org/luci/common/data/text"
	"go.chromium.org/luci/common/lhttp"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/grpc/prpc"

	pb "go.chromium.org/luci/buildbucket/proto"
)

var expectedCodeRPCOption = prpc.ExpectedCode(
	codes.InvalidArgument, codes.NotFound, codes.PermissionDenied)

func doc(doc string) string {
	return text.Doc(doc)
}

// baseCommandRun provides common command run functionality.
// All bb subcommands must embed it directly or indirectly.
type baseCommandRun struct {
	subcommands.CommandRunBase
	authFlags authcli.Flags
	host      string
	json      bool
	noColor   bool

	httpClient *http.Client
	client     pb.BuildsClient
}

func (r *baseCommandRun) RegisterDefaultFlags(p Params) {
	r.Flags.StringVar(&r.host, "host", p.DefaultBuildbucketHost, doc(`
		Host for the buildbucket service instance.
	`))
	r.Flags.BoolVar(&r.noColor, "nocolor", false, doc(`
		Strip ANSI color codes in the output.
	`))
	r.authFlags.Register(&r.Flags, p.Auth)
}

func (r *baseCommandRun) RegisterJSONFlag() {
	r.Flags.BoolVar(&r.json, "json", false, doc(`
		Print objects in JSON format, one after another (not an array).

		Intended for "jq" tool. If using bb from scripts, consider "batch" subcommand.
	`))
}

// initClients validates -host flag and initializes r.httpClient and r.client.
func (r *baseCommandRun) initClients(ctx context.Context) error {
	// Create HTTP Client.
	authOpts, err := r.authFlags.Options()
	if err != nil {
		return err
	}
	r.httpClient, err = auth.NewAuthenticator(ctx, auth.SilentLogin, authOpts).Client()
	if err != nil {
		return err
	}

	// Validate -host
	if r.host == "" {
		return fmt.Errorf("a host for the buildbucket service must be provided")
	}
	if strings.ContainsRune(r.host, '/') {
		return fmt.Errorf("invalid host %q", r.host)
	}

	// Create Buildbucket client.
	rpcOpts := prpc.DefaultOptions()
	rpcOpts.Insecure = lhttp.IsLocalHost(r.host)
	info, err := version.GetCurrentVersion()
	if err != nil {
		return err
	}
	rpcOpts.UserAgent = fmt.Sprintf("buildbucket CLI, instanceID=%q", info.InstanceID)
	r.client = pb.NewBuildsPRPCClient(&prpc.Client{
		C:       r.httpClient,
		Host:    r.host,
		Options: rpcOpts,
	})
	return nil
}

func (r *baseCommandRun) done(ctx context.Context, err error) int {
	if err != nil {
		logging.Errorf(ctx, "%s", err)
		return 1
	}
	return 0
}

// retrieveBuildIDs converts build arguments to int64 build ids,
// where a build argument can be an int64 build or a
// "<project>/<bucket>/<builder>/<build_number>" string.
func (r *baseCommandRun) retrieveBuildIDs(ctx context.Context, builds []string) (buildIDs []int64, err error) {
	return retrieveBuildIDs(builds, func(req *pb.BatchRequest) (*pb.BatchResponse, error) {
		return r.client.Batch(ctx, req)
	})
}

func retrieveBuildIDs(builds []string, callBatch func(*pb.BatchRequest) (*pb.BatchResponse, error)) (buildIDs []int64, err error) {
	buildIDs = make([]int64, len(builds))
	batchReq := &pb.BatchRequest{
		Requests: make([]*pb.BatchRequest_Request, 0, len(builds)),
	}
	indexes := make([]int, 0, len(builds))
	idFieldMask := &field_mask.FieldMask{Paths: []string{"id"}}
	for i, b := range builds {
		getBuild, err := protoutil.ParseGetBuildRequest(b)
		if err != nil {
			return nil, fmt.Errorf("invalid build %q: %s", b, err)
		}
		if getBuild.Builder == nil {
			buildIDs[i] = getBuild.Id
		} else {
			getBuild.Fields = idFieldMask
			batchReq.Requests = append(batchReq.Requests, &pb.BatchRequest_Request{
				Request: &pb.BatchRequest_Request_GetBuild{GetBuild: getBuild},
			})
			indexes = append(indexes, i)
		}
	}

	if len(batchReq.Requests) == 0 {
		return buildIDs, nil
	}

	res, err := callBatch(batchReq)
	for i, res := range res.Responses {
		j := indexes[i]
		switch codes.Code(res.GetError().GetCode()) {
		case codes.OK:
			buildIDs[j] = res.GetGetBuild().Id
		case codes.NotFound:
			return nil, fmt.Errorf("build %q not found", builds[j])
		default:
			return nil, err
		}
	}
	return buildIDs, nil
}

// retrieveBuildID converts a build string into a build id.
// May make a GetBuild RPC.
func (r *baseCommandRun) retrieveBuildID(ctx context.Context, build string) (int64, error) {
	getBuild, err := protoutil.ParseGetBuildRequest(build)
	if err != nil {
		return 0, err
	}

	if getBuild.Id != 0 {
		return getBuild.Id, nil
	}

	res, err := r.client.GetBuild(ctx, getBuild)
	if err != nil {
		return 0, err
	}
	return res.Id, nil
}