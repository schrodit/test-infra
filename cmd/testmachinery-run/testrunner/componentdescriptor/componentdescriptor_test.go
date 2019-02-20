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
package componentdescriptor

import (
	"io/ioutil"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ComponentDescriptor Suite")
}

var _ = Describe("componentdescriptor test", func() {
	It("Should parse a component descriptor and return 2 dependencies", func() {
		input, err := ioutil.ReadFile("./testdata/component_descriptor_1")
		Expect(err).ToNot(HaveOccurred(), "Cannot read json file from ./testdata/component_descriptor_1")

		dependencies, err := GetComponents(input)
		Expect(err).ToNot(HaveOccurred())

		Expect(len(dependencies)).To(Equal(2))
	})

	It("Should parse a component descriptor and ignore duplicates", func() {
		input, err := ioutil.ReadFile("./testdata/component_descriptor_2")
		Expect(err).ToNot(HaveOccurred(), "Cannot read json file from ./testdata/component_descriptor_2")

		result := []*Component{
			&Component{
				Name:    "repo1",
				Version: "0.17.0",
			},
			&Component{
				Name:    "repo2",
				Version: "1.27.0",
			},
		}

		dependencies, err := GetComponents(input)
		Expect(err).ToNot(HaveOccurred())

		Expect(len(dependencies)).To(Equal(2), "There should be 2 dependencies")
		Expect(reflect.DeepEqual(dependencies, result)).To(BeTrue())
	})
})