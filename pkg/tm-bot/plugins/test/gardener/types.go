// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
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

package gardener

import (
	"github.com/gardener/test-infra/pkg/common"
	"github.com/gardener/test-infra/pkg/hostscheduler"
	"github.com/gardener/test-infra/pkg/tm-bot/github/ghval"
)

// DefaultsConfig is the defaults configuration that can be configured using the repository configuration for the tests command
type DefaultsConfig struct {
	HostProvider             *hostscheduler.Provider `json:"hostprovider"`
	BaseClusterCloudProvider *common.CloudProvider   `json:"baseClusterCloudprovider"`

	GardenSetup *struct {
		Revision *ghval.StringOrGitHubValue `json:"revision"`
	} `json:"gardensetup"`
	Gardener *struct {
		Version *ghval.StringOrGitHubValue `json:"version"`
		Commit  *ghval.StringOrGitHubValue `json:"commit"`
	} `json:"gardener"`

	// github.com/gardener/gardener-extension version
	// has to be a yaml of a map[extension-name]{version: "0.0.0",repo:"github.com/x/x"}
	GardenerExtensions *ghval.GitHubValue `json:"gardener-extensions"`

	ShootFlavors *[]*common.ShootFlavor `json:"shootFlavors"`
}
