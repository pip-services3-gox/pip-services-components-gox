package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Contains connection parameters to connect to external services. They are used together with credential parameters, but usually stored separately from more protected sensitive values.

Configuration parameters
  discovery_key: key to retrieve parameters from discovery service
  protocol: connection protocol like http, https, tcp, udp
  host: host name or IP address
  port: port number
  uri: resource URI or connection string with all parameters in it
In addition to standard parameters ConnectionParams may contain any number of custom parameters

see
ConfigParams

see
CredentialParams

see
ConnectionResolver

see
IDiscovery

Example
Example ConnectionParams object usage:

  connection := NewConnectionParamsFromTuples(
      "protocol", "http",
      "host", "10.1.1.100",
      "port", "8080",
      "cluster", "mycluster"
  );

  host := connection.Host();                             // Result: "10.1.1.100"
  port := connection.Port();                             // Result: 8080
  cluster := connection.GetAsNullableString("cluster");     // Result: "mycluster"
*/
type ConnectionParams struct {
	config.ConfigParams
}

// Creates a new connection parameters and fills it with values.
// Returns *ConnectionParams
func NewEmptyConnectionParams() *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewEmptyConfigParams(),
	}
}

// Creates a new connection parameters and fills it with values.
// Parameters:
//   - values map[string]string
//   an object to be converted into key-value pairs to initialize this connection.
// Returns *ConnectionParams
func NewConnectionParams(values map[string]string) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParams(values),
	}
}

// Method that creates a ConfigParams object based on the values that are stored in the 'value' object's properties.
// see
// RecursiveObjectReader.getProperties
// Parameters:
//   - value interface{}
//   configuration parameters in the form of an object with properties.
// Returns ConnectionParams
// generated ConnectionParams.
func NewConnectionParamsFromValue(value interface{}) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromValue(value),
	}
}

// Creates a new ConnectionParams object filled with provided key-value pairs called tuples. Tuples parameters contain a sequence of key1, value1, key2, value2, ... pairs.
// Parameters:
//   - tuples ...interface{}
//   the tuples to fill a new ConnectionParams object.
// Returns *ConnectionParams
// a new ConnectionParams object.
func NewConnectionParamsFromTuples(tuples ...interface{}) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

// Method for creating a StringValueMap from an array of tuples.
// Parameters:
//   - tuples []interface{}
//   the key-value tuples array to initialize the new StringValueMap with.
// Returns *ConnectionParams
// the ConnectionParams created and filled by the 'tuples' array provided.
func NewConnectionParamsFromTuplesArray(tuples []interface{}) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

// Creates a new ConnectionParams object filled with key-value pairs serialized as a string.
// Parameters:
//   - line string
//   a string with serialized key-value pairs as "key1=value1;key2=value2;..." Example: "Key1=123;Key2=ABC;Key3=2016-09-16T00:00:00.00Z"
// Returns *ConnectionParams
// a new ConnectionParams object.
func NewConnectionParamsFromString(line string) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromString(line),
	}
}

// Static method for creating a StringValueMap using the maps passed as parameters.
// Parameters:
//   - maps ...map[string]string
//   the maps passed to this method to create a StringValueMap with.
// Returns ConnectionParams
// the ConnectionParams created.
func NewConnectionParamsFromMaps(maps ...map[string]string) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromMaps(maps...),
	}
}

// Retrieves all ConnectionParams from configuration parameters from "connections" section. If "connection" section is present instead, than it returns a list with only one ConnectionParams.
// Parameters:
//   - config *config.ConfigParams
//   a configuration parameters to retrieve connections
// Returns []*ConnectionParams
// a list of retrieved ConnectionParams
func NewManyConnectionParamsFromConfig(config *config.ConfigParams) []*ConnectionParams {
	result := []*ConnectionParams{}

	connections := config.GetSection("connections")

	if connections.Len() > 0 {
		for _, section := range connections.GetSectionNames() {
			connection := connections.GetSection(section)
			result = append(result, NewConnectionParams(connection.Value()))
		}
	} else {
		connection := config.GetSection("connection")
		if connection.Len() > 0 {
			result = append(result, NewConnectionParams(connection.Value()))
		}
	}

	return result
}

// Retrieves a single ConnectionParams from configuration parameters from "connection" section. If "connections" section is present instead, then is returns only the first connection element.
// Parameters:
//   - config *config.ConfigParams
//   ConnectionParams, containing a section named "connection(s)".
// Returns *ConnectionParams
// the generated ConnectionParams object.
func NewConnectionParamsFromConfig(config *config.ConfigParams) *ConnectionParams {
	connections := NewManyConnectionParamsFromConfig(config)
	if len(connections) > 0 {
		return connections[0]
	}
	return nil
}

// Checks if these connection parameters shall be retrieved from DiscoveryService.
// The connection parameters are redirected to DiscoveryService when discovery_key parameter is set.
// Returns bool
// true if connection shall be retrieved from DiscoveryService
func (c *ConnectionParams) UseDiscovery() bool {
	return c.GetAsString("discovery_key") != ""
}

// Gets the key to retrieve this connection from DiscoveryService. If this key is null, than all parameters are already present.
// see
// UseDiscovery
// Returns string
// the discovery key to retrieve connection.
func (c *ConnectionParams) DiscoveryKey() string {
	return c.GetAsString("discovery_key")
}

// Sets the key to retrieve these parameters from DiscoveryService.
// Parameters:
//   - value string
//   a new key to retrieve connection.
func (c *ConnectionParams) SetDiscoveryKey(value string) {
	c.Put("discovery_key", value)
}

// Gets the connection protocol.
// Returns string
// the connection protocol or the default value if it's not set.
func (c *ConnectionParams) Protocol() string {
	return c.GetAsString("protocol")
}

// Gets the connection protocol.
// Parameters:
//   - defaultValue string
//   the default protocol
// Returns string
// the connection protocol or the default value if it's not set.
func (c *ConnectionParams) ProtocolWithDefault(defaultValue string) string {
	return c.GetAsStringWithDefault("protocol", defaultValue)
}

// Sets the connection protocol.
// Parameters:
//   - value string
//   a new connection protocol.
func (c *ConnectionParams) SetProtocol(value string) {
	c.Put("protocol", value)
}

// Gets the host name or IP address.
// Returns string
// the host name or IP address.
func (c *ConnectionParams) Host() string {
	host := c.GetAsString("host")
	if host != "" {
		return host
	}
	return c.GetAsString("ip")
}

// Sets the host name or IP address.
// Parameters:
//   - value string
//   a new host name or IP address.
func (c *ConnectionParams) SetHost(value string) {
	c.Put("host", value)
}

// Gets the port number.
// Returns int
// the port number.
func (c *ConnectionParams) Port() int {
	return c.GetAsInteger("port")
}

// Gets the port number.
// Parameters:
//  - defaultValue int
//  default port number
// Returns int
// the port number.
func (c *ConnectionParams) PortWithDefault(defaultValue int) int {
	return c.GetAsIntegerWithDefault("port", defaultValue)
}

// Sets the port number.
// see
// Host
// Parameters:
//   - value int
//   a new port number.
func (c *ConnectionParams) SetPort(value int) {
	c.Put("port", value)
}

// Gets the resource URI or connection string. Usually it includes all connection parameters in it.
// Returns string
// the resource URI or connection string.
func (c *ConnectionParams) Uri() string {
	return c.GetAsString("uri")
}

// Sets the resource URI or connection string.
// Parameters:
//   - value string
//   a new resource URI or connection string.
func (c *ConnectionParams) SetUri(value string) {
	c.Put("uri", value)
}
