// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	cfg "github.com/luci/luci-go/common/config/impl/memory"
)

const project1CronCfg = `
job {
  id: "noop-job-1"
  schedule: "*/10 * * * * * *"
  task: {
    noop: {}
  }
}

job {
  id: "noop-job-2"
  schedule: "*/10 * * * * * *"
  task: {
    noop: {}
  }
}
`

const project2CronCfg = `
job {
  id: "noop-job-1"
  schedule: "*/10 * * * * * *"
  task: {
    noop: {}
  }
}

job {
  id: "noop-job-2"
  schedule: "*/10 * * * * * *"
  task: {
    noop: {}
  }
}
`

// devServerConfig returns mocked luci-config configs to use locally on dev
// server (to avoid talking to real luci-config service).
func devServerConfigs() map[string]cfg.ConfigSet {
	return map[string]cfg.ConfigSet{
		"projects/project1": {
			"cron.cfg": project1CronCfg,
		},
		"projects/project2": {
			"cron.cfg": project2CronCfg,
		},
	}
}
