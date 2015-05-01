// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package archiver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/luci/luci-go/client/isolatedclient"
	"github.com/luci/luci-go/client/isolatedclient/isolatedfake"
	"github.com/luci/luci-go/common/isolated"
	"github.com/maruel/ut"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

type int64Slice []int64

func (a int64Slice) Len() int           { return len(a) }
func (a int64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64Slice) Less(i, j int) bool { return a[i] < a[j] }

func TestArchiverEmpty(t *testing.T) {
	t.Parallel()
	a := New(isolatedclient.New("https://localhost:1", "default-gzip"))
	ut.AssertEqual(t, &Stats{[]int64{}, []int64{}, []*UploadStat{}}, a.Stats())
}

func TestArchiverFile(t *testing.T) {
	t.Parallel()
	server := isolatedfake.New()
	ts := httptest.NewServer(server)
	defer ts.Close()
	a := New(isolatedclient.New(ts.URL, "default-gzip"))

	fEmpty, err := ioutil.TempFile("", "archiver")
	ut.AssertEqual(t, nil, err)
	future1 := a.PushFile(fEmpty.Name(), fEmpty.Name())
	ut.AssertEqual(t, fEmpty.Name(), future1.DisplayName())
	fFoo, err := ioutil.TempFile("", "archiver")
	ut.AssertEqual(t, nil, err)
	ut.AssertEqual(t, nil, ioutil.WriteFile(fFoo.Name(), []byte("foo"), 0600))
	future2 := a.PushFile(fFoo.Name(), fFoo.Name())
	// Push the same file another time. It'll get linked to the first.
	future3 := a.PushFile(fFoo.Name(), fFoo.Name())
	future1.WaitForHashed()
	future2.WaitForHashed()
	future3.WaitForHashed()
	ut.AssertEqual(t, nil, a.Close())

	stats := a.Stats()
	ut.AssertEqual(t, []int64{}, stats.Hits)
	misses := int64Slice(stats.Misses)
	sort.Sort(misses)
	// Only 2 lookups, not 3.
	ut.AssertEqual(t, int64Slice{0, 3}, misses)
	ut.AssertEqual(t, int64(3), stats.TotalPushed())
	ut.AssertEqual(t, int64(0), stats.TotalHits())
	ut.AssertEqual(t, int64(3), stats.TotalMisses())
	expected := map[isolated.HexDigest][]byte{
		"0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33": []byte("foo"),
		"da39a3ee5e6b4b0d3255bfef95601890afd80709": {},
	}
	ut.AssertEqual(t, expected, server.Contents())
	ut.AssertEqual(t, isolated.HexDigest("da39a3ee5e6b4b0d3255bfef95601890afd80709"), future1.Digest())
	ut.AssertEqual(t, nil, future1.Error())
	ut.AssertEqual(t, isolated.HexDigest("0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"), future2.Digest())
	ut.AssertEqual(t, nil, future2.Error())
	ut.AssertEqual(t, isolated.HexDigest("0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"), future3.Digest())
	ut.AssertEqual(t, nil, future3.Error())
	ut.AssertEqual(t, nil, server.Error())
}

func TestArchiverDirectory(t *testing.T) {
	// Uploads a real directory. 2 times the same file.
	t.Parallel()
	server := isolatedfake.New()
	ts := httptest.NewServer(server)
	defer ts.Close()
	a := New(isolatedclient.New(ts.URL, "default-gzip"))

	tmpDir, err := ioutil.TempDir("", "archiver")
	ut.AssertEqual(t, nil, err)
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fail()
		}
	}()

	ut.AssertEqual(t, nil, ioutil.WriteFile(filepath.Join(tmpDir, "bar"), []byte("foo"), 0600))
	ut.AssertEqual(t, nil, ioutil.WriteFile(filepath.Join(tmpDir, "bar_dupe"), []byte("foo"), 0600))
	future := PushDirectory(a, tmpDir, nil)
	ut.AssertEqual(t, filepath.Base(tmpDir)+".isolated", future.DisplayName())
	future.WaitForHashed()
	ut.AssertEqual(t, nil, a.Close())

	isolatedData := isolated.Isolated{
		Algo: "sha-1",
		Files: map[string]isolated.File{
			"bar":      {Digest: "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", Mode: newInt(384), Size: newInt64(3)},
			"bar_dupe": {Digest: "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", Mode: newInt(384), Size: newInt64(3)},
		},
		Version: isolated.IsolatedFormatVersion,
	}
	encoded, err := json.Marshal(isolatedData)
	ut.AssertEqual(t, nil, err)
	isolatedEncoded := string(encoded) + "\n"
	stats := a.Stats()
	ut.AssertEqual(t, []int64{}, stats.Hits)
	misses := int64Slice(stats.Misses)
	sort.Sort(misses)
	// There's 3 cache misses even if the same content is looked up twice.
	ut.AssertEqual(t, int64Slice{3, 3, int64(len(isolatedEncoded))}, misses)
	expected := map[string]string{
		"0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33": "foo",
		"be8b1d344ab7129b7a80ab1977f5c32bd0c093c5": isolatedEncoded,
	}
	actual := map[string]string{}
	for k, v := range server.Contents() {
		actual[string(k)] = string(v)
	}
	ut.AssertEqual(t, expected, actual)
	ut.AssertEqual(t, nil, server.Error())
}
