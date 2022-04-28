package connect

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

// MemoryDiscovery discovery service that keeps connections in memory.
//	Configuration parameters
//		[connection key 1]:
//		... connection parameters for key 1
//		[connection key 2]:
//		... connection parameters for key N
//	see IDiscovery
//	see ConnectionParams
//	Example
//		config := NewConfigParamsFromTuples(
//			"key1.host", "10.1.1.100",
//			"key1.port", "8080",
//			"key2.host", "10.1.1.100",
//			"key2.port", "8082"
//		);
//		discovery := NewMemoryDiscovery();
//		discovery.ReadConnections(config);
//		discovery.Resolve("123", "key1", (err, connection) => {
//			// Result: host=10.1.1.100;port=8080
//		});
type MemoryDiscovery struct {
	items map[string][]*ConnectionParams
}

// NewEmptyMemoryDiscovery creates a new instance of discovery service.
//	Returns: *MemoryDiscovery
func NewEmptyMemoryDiscovery() *MemoryDiscovery {
	return &MemoryDiscovery{
		items: map[string][]*ConnectionParams{},
	}
}

// NewMemoryDiscovery creates a new instance of discovery service.
//	Parameters: config *config.ConfigParams configuration with connection parameters.
//	Returns: *MemoryDiscovery
func NewMemoryDiscovery(config *config.ConfigParams) *MemoryDiscovery {
	c := &MemoryDiscovery{
		items: map[string][]*ConnectionParams{},
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

// Configure component by passing configuration parameters.
//	Parameters: config *config.ConfigParams configuration parameters to be set.
func (c *MemoryDiscovery) Configure(config *config.ConfigParams) {
	c.ReadConnections(config)
}

// ReadConnections from configuration parameters. Each section represents an individual Connectionparams
//	Parameters: config *configure.ConfigParams configuration parameters to be read
func (c *MemoryDiscovery) ReadConnections(config *config.ConfigParams) {
	c.items = make(map[string][]*ConnectionParams)

	keys := config.Keys()
	for _, key := range keys {
		value := config.GetAsString(key)
		connection := NewConnectionParamsFromString(value)
		c.items[key] = []*ConnectionParams{connection}
	}
}

// Register connection parameters into the discovery service.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- key string a key to uniquely identify the connection parameters.
//		- connection *ConnectionParams
//	Returns: *ConnectionParams, error registered connection or error.
func (c *MemoryDiscovery) Register(ctx context.Context, correlationId string, key string,
	connection *ConnectionParams) (result *ConnectionParams, err error) {

	if connection != nil {
		if connections, ok := c.items[key]; ok && connections == nil {
			connections = []*ConnectionParams{connection}
			c.items[key] = connections
		} else {
			connections = append(connections, connection)
		}
	}

	return connection, nil
}

// ResolveOne a single connection parameters by its key.
//	Parameters:
//		- ctx context.Context
//		- correlationId: string transaction id to trace execution through call chain.
//		- key: string a key to uniquely identify the connection.
//	Returns: *ConnectionParams, error receives found connection or error.
func (c *MemoryDiscovery) ResolveOne(ctx context.Context, correlationId string,
	key string) (result *ConnectionParams, err error) {

	connections, _ := c.ResolveAll(ctx, correlationId, key)
	if len(connections) > 0 {
		return connections[0], nil
	}

	return nil, nil
}

// ResolveAll connection parameters by its key.
//	Parameters:
//		- ctx context.Context
//		- correlationId: string transaction id to trace execution through call chain.
//		- key: string a key to uniquely identify the connection.
//	Returns: *ConnectionParams, error receives found connection or error.
func (c *MemoryDiscovery) ResolveAll(ctx context.Context, correlationId string,
	key string) (result []*ConnectionParams, err error) {
	connections, _ := c.items[key]

	if connections == nil {
		connections = []*ConnectionParams{}
	}

	return connections, nil
}
