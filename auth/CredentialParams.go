package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Contains credentials to authenticate against external services. They are used together with connection parameters, but usually stored in a separate store, protected from unauthorized access.

Configuration parameters
  store_key: key to retrieve parameters from credential store
  username: user name
  user: alternative to username
  password: user password
  pass: alternative to password
  access_id: application access id
  client_id: alternative to access_id
  access_key: application secret key
  client_key: alternative to access_key
  secret_key: alternative to access_key
In addition to standard parameters CredentialParams may contain any number of custom parameters

see
ConfigParams

see
ConnectionParams

see
CredentialResolver

see
ICredentialStore

Example
  credential := NewCredentialParamsFromTuples(
      "user", "jdoe",
      "pass", "pass123",
      "pin", "321"
  );

  username := credential.Username();             // Result: "jdoe"
  password := credential.Password();             // Result: "pass123"
*/
type CredentialParams struct {
	config.ConfigParams
}

// Creates a new credential parameters and fills it with values.
// Returns *CredentialParams
func NewEmptyCredentialParams() *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewEmptyConfigParams(),
	}
}

// Creates a new credential parameters and fills it with values.
// Parameters:
//  - values map[string]string
//  an object to be converted into key-value pairs to initialize these credentials.
// Returns *CredentialParams
func NewCredentialParams(values map[string]string) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParams(values),
	}
}

// Method that creates a ConfigParams object based on the values that are stored in the 'value' object's properties.
// Parameters:
//   - value interface{}
//   configuration parameters in the form of an object with properties.
// Returns *ConfigParams
// generated ConfigParams.
func NewCredentialParamsFromValue(value interface{}) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromValue(value),
	}
}

// Creates a new CredentialParams object filled with provided key-value pairs called tuples. Tuples parameters contain a sequence of key1, value1, key2, value2, ... pairs.
// Parameters:
//   - tuples ...interface{}
//   the tuples to fill a new CredentialParams object.
// Returns *CredentialParams
// a new CredentialParams object.
func NewCredentialParamsFromTuples(tuples ...interface{}) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

// Static method for creating a CredentialParams from an array of tuples.
// Parameters:
//   - tuples []interface{}
//   the key-value tuples array to initialize the new StringValueMap with.
// Returns CredentialParams
// the CredentialParams created and filled by the 'tuples' array provided.
func NewCredentialParamsFromTuplesArray(tuples []interface{}) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

// Creates a new CredentialParams object filled with key-value pairs serialized as a string.
// Parameters:
//   - line string
//   a string with serialized key-value pairs as "key1=value1;key2=value2;..." Example: "Key1=123;Key2=ABC;Key3=2016-09-16T00:00:00.00Z"
// Returns *CredentialParams
// a new CredentialParams object.
func NewCredentialParamsFromString(line string) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromString(line),
	}
}

// Static method for creating a CredentialParams using the maps passed as parameters.
// Parameters:
//   - maps ...map[string]string
//   the maps passed to this method to create a StringValueMap with.
// Returns *CredentialParams
// the CredentialParams created.
func NewCredentialParamsFromMaps(maps ...map[string]string) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromMaps(maps...),
	}
}

// Retrieves all CredentialParams from configuration parameters from "credentials" section. If "credential" section is present instead, than it returns a list with only one CredentialParams.
// Parameters:
//   - config *config.ConfigParams
//   a configuration parameters to retrieve credentials
// Returns []*CredentialParams
// a list of retrieved CredentialParams
func NewManyCredentialParamsFromConfig(config *config.ConfigParams) []*CredentialParams {
	result := []*CredentialParams{}

	credentials := config.GetSection("credentials")

	if credentials.Len() > 0 {
		for _, section := range credentials.GetSectionNames() {
			credential := credentials.GetSection(section)
			result = append(result, NewCredentialParams(credential.Value()))
		}
	} else {
		credential := config.GetSection("credential")
		if credential.Len() > 0 {
			result = append(result, NewCredentialParams(credential.Value()))
		}
	}

	return result
}

// Retrieves a single CredentialParams from configuration parameters from "credential" section. If "credentials" section is present instead, then is returns only the first credential element.
// Parameters:
//   - config *config.ConfigParams
//   ConfigParams, containing a section named "credential(s)".
// Returns []*CredentialParams
// the generated CredentialParams object.
func NewCredentialParamsFromConfig(config *config.ConfigParams) *CredentialParams {
	credentials := NewManyCredentialParamsFromConfig(config)
	if len(credentials) > 0 {
		return credentials[0]
	}
	return nil
}

// Checks if these credential parameters shall be retrieved from CredentialStore. The credential parameters are redirected to CredentialStore when store_key parameter is set.
// Returns bool
// true if credentials shall be retrieved from CredentialStore
func (c *CredentialParams) UseCredentialStore() bool {
	return c.GetAsString("store_key") != ""
}

// Gets the key to retrieve these credentials from CredentialStore. If this key is null, than all parameters are already present.
// Returns string
// the store key to retrieve credentials.
func (c *CredentialParams) StoreKey() string {
	return c.GetAsString("store_key")
}

// Sets the key to retrieve these parameters from CredentialStore.
// Parameters:
//   - value string
//   a new key to retrieve credentials.
func (c *CredentialParams) SetStoreKey(value string) {
	c.Put("store_key", value)
}

// Gets the user name. The value can be stored in parameters "username" or "user".
// Returns string
// the user name.
func (c *CredentialParams) Username() string {
	username := c.GetAsString("username")
	if username == "" {
		username = c.GetAsString("user")
	}
	return username
}

// Sets the user name.
// Parameters:
//   - value string
//   a new user name.
func (c *CredentialParams) SetUsername(value string) {
	c.Put("username", value)
}

// Get the user password. The value can be stored in parameters "password" or "pass".
// Returns string
// the user password.
func (c *CredentialParams) Password() string {
	password := c.GetAsString("password")
	if password == "" {
		password = c.GetAsString("pass")
	}
	return password
}

// Sets the user password.
// Parameters:
//   - value string
//   a new user password.
func (c *CredentialParams) SetPassword(value string) {
	c.Put("password", value)
}

// Gets the application access id. The value can be stored in parameters "access_id" pr "client_id"
// Returns string
// the application access id.
func (c *CredentialParams) AccessId() string {
	accessId := c.GetAsString("access_id")
	if accessId == "" {
		accessId = c.GetAsString("client_id")
	}
	return accessId
}

// Sets the application access id.
// Parameters:
//   - value: string
//   a new application access id.
func (c *CredentialParams) SetAccessId(value string) {
	c.Put("access_id", value)
}

// Gets the application secret key. The value can be stored in parameters "access_key", "client_key" or "secret_key".
// Returns string
// the application secret key.
func (c *CredentialParams) AccessKey() string {
	accessKey := c.GetAsString("access_key")
	if accessKey == "" {
		accessKey = c.GetAsString("client_key")
	}
	return accessKey
}

// Sets the application secret key.
// Parameters
//   - value string
//   a new application secret key.
func (c *CredentialParams) SetAccessKey(value string) {
	c.Put("access_key", value)
}
