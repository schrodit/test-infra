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

package secrets

import (
	"github.com/gardener/gardener/pkg/utils"
	"k8s.io/client-go/rest"
)

// GenerateKubeconfigFromRestConfig generates a kubernetes kubeconfig from a rest client
func GenerateKubeconfigFromRestConfig(cfg *rest.Config, name string) ([]byte, error) {
	values := map[string]interface{}{
		"APIServerURL":      cfg.Host,
		"CACertificate":     utils.EncodeBase64(cfg.TLSClientConfig.CAData),
		"ClientCertificate": utils.EncodeBase64(cfg.TLSClientConfig.CertData),
		"ClientKey":         utils.EncodeBase64(cfg.TLSClientConfig.KeyData),
		"ClusterName":       name,
	}

	if cfg.Username != "" && cfg.Password != "" {
		values["BasicAuthUsername"] = cfg.Username
		values["BasicAuthPassword"] = cfg.Password
	}

	return utils.RenderLocalTemplate(kubeconfigTemplate, values)
}

const kubeconfigTemplate = `---
apiVersion: v1
kind: Config
current-context: {{ .ClusterName }}
clusters:
- name: {{ .ClusterName }}
  cluster:
    certificate-authority-data: {{ .CACertificate }}
    server: https://{{ .APIServerURL }}
contexts:
- name: {{ .ClusterName }}
  context:
    cluster: {{ .ClusterName }}
{{- if and .ClientCertificate .ClientKey }}
    user: {{ .ClusterName }}
{{- else }}
    user: {{ .ClusterName }}-basic-auth
{{- end}}
users:
{{- if and .ClientCertificate .ClientKey }}
- name: {{ .ClusterName }}
  user:
    client-certificate-data: {{ .ClientCertificate }}
    client-key-data: {{ .ClientKey }}
{{- end}}
{{- if and .BasicAuthUsername .BasicAuthPassword }}
- name: {{ .ClusterName }}-basic-auth
  user:
    username: {{ .BasicAuthUsername }}
    password: {{ .BasicAuthPassword }}
{{- end}}`
