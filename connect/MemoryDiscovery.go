package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Discovery service that keeps connections in memory.

Configuration parameters
  [connection key 1]:
  ... connection parameters for key 1
  [connection key 2]:
  ... connection parameters for key N
see
IDiscovery

see
ConnectionParams

Example
  config := NewConfigParamsFromTuples(
      "key1.host", "10.1.1.100",
      "key1.port", "8080",
      "key2.host", "10.1.1.100",
      "key2.port", "8082"
  );

  discovery := NewMemoryDiscovery();
  discovery.ReadConnections(config);

  discovery.Resolve("123", "key1", (err, connection) => {
      // Result: host=10.1.1.100;port=8080
  });
*/
type MemoryDiscovery struct {
	items map[string][]*ConnectionParams
}

// Creates a new instance of discovery service.
// Returns *MemoryDiscovery
func NewEmptyMemoryDiscovery() *MemoryDiscovery {
	return &MemoryDiscovery{
		items: map[string][]*ConnectionParams{},
	}
}

// Creates a new instance of discovery service.
// Parameters:
// 			- config *config.ConfigParams
// 			configuration with connection parameters.
// Returns *MemoryDiscovery
func NewMemoryDiscovery(config *config.ConfigParams) *MemoryDiscovery {
	c := &MemoryDiscovery{
		items: map[string][]*ConnectionParams{},
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

// Configures component by passing configuration parameters.
// Parameters:
// 			- config *config.ConfigParams
// configuration parameters to be set.
func (c *MemoryDiscovery) Configure(config *config.ConfigParams) {
	c.ReadConnections(config)
}

// Reads connections from configuration parameters. Each section represents an individual Connectionparams
// Parameters:
// 			- config *configure.ConfigParams
// configuration parameters to be read
func (c *MemoryDiscovery) ReadConnections(config *config.ConfigParams) {
	c.items = map[string][]*ConnectionParams{}

	keys := config.Keys()
	for _, key := range keys {
		value := config.GetAsString(key)
		connection := NewConnectionParamsFromString(value)
		c.items[key] = []*ConnectionParams{connection}
	}
}

// Registers connection parameters into the discovery service.
// Parameters:
// 			- correlationId string
// 			transaction id to trace execution through call chain.
// 			- key string
// 			a key to uniquely identify the connection parameters.
// 			- connection *ConnectionParams
// Returns  *ConnectionParams, error
// registered connection or error.
func (c *MemoryDiscovery) Register(correlationId string, key string,
	connection *ConnectionParams) (result *ConnectionParams, err error) {

	if connection != nil {
		connections, _ := c.items[key]
		if connections == nil {
			connections = []*ConnectionParams{connection}
			c.items[key] = connections
		} else {
			connections = append(connections, connection)
		}
	}

	return connection, nil
}

// Resolves a single connection parameters by its key.
// Parameters:
// 			- correlationId: string
// 			 transaction id to trace execution through call chain.
// 			- key: string
// 			a key to uniquely identify the connection.
// Returns  *ConnectionParams, error
// receives found connection or error.
func (c *MemoryDiscovery) ResolveOne(correlationId string,
	key string) (result *ConnectionParams, err error) {

	connections, _ := c.ResolveAll(correlationId, key)
	if len(connections) > 0 {
		return connections[0], nil
	}

	return nil, nil
}

func (c *MemoryDiscovery) ResolveAll(correlationId string,
	key string) (result []*ConnectionParams, err error) {
	connections, _ := c.items[key]

	if connections == nil {
		connections = []*ConnectionParams{}
	}

	return connections, nil
}
