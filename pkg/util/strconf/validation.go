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

package strconf

import "fmt"

// Validate validates a testrun config element.
func Validate(identifier string, source *ConfigSource) error {
	if source.ConfigMapKeyRef == nil && source.SecretKeyRef == nil {
		return fmt.Errorf("%s.(configMapKeyRef or secretMapKeyRef): Required configMapKeyRef or secretMapKeyRef: Either a configmap ref or a secretmap ref have to be defined", identifier)
	}
	if source.ConfigMapKeyRef != nil {
		if source.ConfigMapKeyRef.Key == "" {
			return fmt.Errorf("%s.configMapKeyRef.key: Required value", identifier)
		}
		if source.ConfigMapKeyRef.Name == "" {
			return fmt.Errorf("%s.configMapKeyRef.name: Required value", identifier)
		}
	}
	if source.SecretKeyRef != nil {
		if source.SecretKeyRef.Key == "" {
			return fmt.Errorf("%s.secretKeyRef.key: Required value", identifier)
		}
		if source.SecretKeyRef.Name == "" {
			return fmt.Errorf("%s.secretKeyRef.name: Required value", identifier)
		}
	}
	return nil
}
