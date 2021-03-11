package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Credential store that keeps credentials in memory.

Configuration parameters
  [credential key 1]:
  ... credential parameters for key 1
  [credential key 2]:
  ... credential parameters for key N
  ...
see
ICredentialStore

see
CredentialParams

Example
  config := NewConfigParamsFromTuples(
      "key1.user", "jdoe",
      "key1.pass", "pass123",
      "key2.user", "bsmith",
      "key2.pass", "mypass"
  );

  credentialStore := NewEmptyMemoryCredentialStore();
  credentialStore.ReadCredentials(config);
  res, err := credentialStore.Lookup("123", "key1");
*/
type MemoryCredentialStore struct {
	items map[string]*CredentialParams
}

// Creates a new instance of the credential store.
// Returns *MemoryCredentialStore
func NewEmptyMemoryCredentialStore() *MemoryCredentialStore {
	return &MemoryCredentialStore{
		items: map[string]*CredentialParams{},
	}
}

// Creates a new instance of the credential store.
// Parameters:
//   - config *config.ConfigParams
//   configuration with credential parameters.
// Returns *MemoryCredentialStore
func NewMemoryCredentialStore(config *config.ConfigParams) *MemoryCredentialStore {
	c := &MemoryCredentialStore{
		items: map[string]*CredentialParams{},
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
// configuration parameters to be set.
func (c *MemoryCredentialStore) Configure(config *config.ConfigParams) {
	c.ReadCredentials(config)
}

// Reads credentials from configuration parameters. Each section represents an individual CredentialParams
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be read
func (c *MemoryCredentialStore) ReadCredentials(config *config.ConfigParams) {
	c.items = map[string]*CredentialParams{}

	keys := config.Keys()
	for _, key := range keys {
		value := config.GetAsString(key)
		credential := NewCredentialParamsFromString(value)
		c.items[key] = credential
	}
}

// Stores credential parameters into the store.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - key string
//   a key to uniquely identify the credential parameters.
//   - credential *CredentialParams
//   a credential parameters to be stored.
// Return error
func (c *MemoryCredentialStore) Store(correlationId string, key string,
	credential *CredentialParams) error {

	if credential != nil {
		c.items[key] = credential
	} else {
		delete(c.items, key)
	}

	return nil
}

// Lookups credential parameters by its key.
// Parameters:
//   - correlationId string
//    transaction id to trace execution through call chain.
//   - key string
//   a key to uniquely identify the credential parameters.
// Return result *CredentialParams, err error
// result of lookup and error message
func (c *MemoryCredentialStore) Lookup(correlationId string,
	key string) (result *CredentialParams, err error) {
	credential, _ := c.items[key]
	return credential, nil
}
