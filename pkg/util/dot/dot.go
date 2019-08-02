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

package dot

import (
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
)

func GenerateDotFileFromSpec(tr *v1beta1.Testrun) (string, error) {
	graphName := "no_name"

	if tr.Name != "" {
		graphName = tr.Name
	} else if tr.GenerateName != "" {
		graphName = tr.GenerateName
	}

	graph := gographviz.NewGraph()
	graph.Directed = true
	if err := graph.SetName(escape(graphName)); err != nil {
		return "", err
	}
	for _, step := range tr.Spec.TestFlow {
		for _, d := range step.DependsOn {
			if !graph.IsNode(d) {
				if err := graph.AddNode(graphName, escape(d), nil); err != nil {
					return "", err
				}
			}
			if err := graph.AddNode(graphName, escape(step.Name), nil); err != nil {
				return "", err
			}
			if err := graph.AddEdge(escape(d), escape(step.Name), true, nil); err != nil {
				return "", err
			}
		}
	}

	return graph.String(), nil
}

func GenerateDotFileFromStatus(tr *v1beta1.Testrun) (string, error) {
	graphName := "no_name"

	if tr.Name != "" {
		graphName = tr.Name
	} else if tr.GenerateName != "" {
		graphName = tr.GenerateName
	}

	graph := gographviz.NewGraph()
	graph.Directed = true
	if err := graph.SetName(escape(graphName)); err != nil {
		return "", err
	}
	for _, step := range tr.Status.Steps {
		for _, d := range step.Position.DependsOn {
			if !graph.IsNode(d) {
				if err := graph.AddNode(graphName, escape(d), nil); err != nil {
					return "", err
				}
			}
			if err := graph.AddNode(graphName, escape(step.Name), nil); err != nil {
				return "", err
			}
			if err := graph.AddEdge(escape(d), escape(step.Name), true, nil); err != nil {
				return "", err
			}
		}
	}

	return graph.String(), nil
}

func escape(s string) string {
	return strings.ReplaceAll(s, "-", "_")
}
