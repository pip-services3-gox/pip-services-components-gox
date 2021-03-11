package connect

/*
Interface for discovery services which are used to store and resolve connection parameters to connect to external services.
*/
type IDiscovery interface {
	// Registers connection parameters into the discovery service.
	Register(correlationId string, key string,
		connection *ConnectionParams) (result *ConnectionParams, err error)
	// Resolves a single connection parameters by its key.
	ResolveOne(correlationId string, key string) (result *ConnectionParams, err error)
	// Resolves all connection parameters by their key.
	ResolveAll(correlationId string, key string) (result []*ConnectionParams, err error)
}
