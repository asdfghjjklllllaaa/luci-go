// Copyright 2018 The LUCI Authors.
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

package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
)

// SwarmingClient is a Swarming server client.
type SwarmingClient struct {
	// Client is the *http.Client to use to communicate with the Swarming server.
	*http.Client
	// PlatformStrategy is the platform-specific strategy to use.
	PlatformStrategy
	// Server is the Swarming server URL.
	server string
}

// cliKey is the key to a *SwarmingClient in the context.
var cliKey = "cli"

// withClient returns a new context with the given *SwarmingClient installed.
func withClient(c context.Context, cli *SwarmingClient) context.Context {
	return context.WithValue(c, &cliKey, cli)
}

// getClient returns the *SwarmingClient installed in the current context.
func getClient(c context.Context) *SwarmingClient {
	return c.Value(&cliKey).(*SwarmingClient)
}

// Fetch fetches the Swarming bot code.
func (s *SwarmingClient) Fetch(c context.Context, dir, user string) error {
	botCode := s.server + "/bot_code"
	logging.Infof(c, "downloading: %s", botCode)
	rsp, err := s.Get(botCode)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return errors.Reason("server returned %q", rsp.Status).Err()
	}

	// 0755 allows the directory structure to be read and listed by all users.
	// Useful when SSHing fo the instance.
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	zip := filepath.Join(dir, "swarming_bot.zip")
	logging.Infof(c, "installing: %s", zip)
	// 0644 allows the zip to be read by all users.
	// Useful when SSHing to the instance.
	out, err := os.OpenFile(zip, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, rsp.Body)
	if err != nil {
		return err
	}
	if err := s.chown(c, zip, user); err != nil {
		return err
	}
	return nil
}
