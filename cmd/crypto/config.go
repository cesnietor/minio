// MinIO Cloud Storage, (C) 2015, 2016, 2017, 2018 MinIO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crypto

const (
	EnvKMSMode         = "MINIO_SSE_VAULT_MODE" // Setting this to "KV" enables the key-value store
	EnvKMSKeyStorePath = "MINIO_SSE_VAULT_KEY_STORE_PATH"
)

var GlobalKeyStore KeyStore

// KMSConfig has the KMS config for hashicorp vault
type KMSConfig struct {
	AutoEncryption bool        `json:"-"`
	Vault          VaultConfig `json:"vault"`
}
