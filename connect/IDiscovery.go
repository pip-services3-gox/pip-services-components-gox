package connect

// IDiscovery interface for discovery services which are used to store and resolve
// connection parameters to connect to external services.
type IDiscovery interface {

	// Register connection parameters into the discovery service.
	Register(correlationId string, key string,
		connection *ConnectionParams) (result *ConnectionParams, err error)

	// ResolveOne a single connection parameters by its key.
	ResolveOne(correlationId string, key string) (result *ConnectionParams, err error)

	// ResolveAll all connection parameters by their key.
	ResolveAll(correlationId string, key string) (result []*ConnectionParams, err error)
}
