package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

/*
Helper class to retrieve component connections.

If connections are configured to be retrieved from IDiscovery, it automatically locates IDiscovery in component references and retrieve connections from there using discovery_key parameter.

Configuration parameters
  connection:

    discovery_key: (optional) a key to retrieve the connection from IDiscovery
      ... other connection parameters
    connections: alternative to connection

      [connection params 1]: first connection parameters
      ... connection parameters for key 1
      [connection params N]: Nth connection parameters
      ... connection parameters for key N
References
*:discovery:*:*:1.0 (optional) IDiscovery services to resolve connections
see
ConnectionParams

see
IDiscovery

Example
  config = NewConfigParamsFromTuples(
      "connection.host", "10.1.1.100",
      "connection.port", 8080
  );

  connectionResolver := NewConnectionResolver();
  connectionResolver.Configure(config);
  connectionResolver.SetReferences(references);

  res, err := connectionResolver.Resolve("123");
*/
type ConnectionResolver struct {
	connections []*ConnectionParams
	references  refer.IReferences
}

// Creates a new instance of connection resolver.
// Returns *ConnectionResolver
func NewEmptyConnectionResolver() *ConnectionResolver {
	return &ConnectionResolver{
		connections: []*ConnectionParams{},
		references:  nil,
	}
}

// Creates a new instance of connection resolver.
// Parameters:
//   - config *config.ConfigParams
//   component configuration parameters
//   - references refer.IReferences
//   component references
// Returns *ConnectionResolver
func NewConnectionResolver(config *config.ConfigParams,
	references refer.IReferences) *ConnectionResolver {
	c := &ConnectionResolver{
		connections: []*ConnectionParams{},
		references:  references,
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be set.
func (c *ConnectionResolver) Configure(config *config.ConfigParams) {
	connections := NewManyConnectionParamsFromConfig(config)

	for _, connection := range connections {
		c.connections = append(c.connections, connection)
	}
}

// Sets references to dependent components.
// Parameters:
//   - references refer.IReferences
//   references to locate the component dependencies.
func (c *ConnectionResolver) SetReferences(references refer.IReferences) {
	c.references = references
}

// Gets all connections configured in component configuration.
// Redirect to Discovery services is not done at this point. If you need fully fleshed connection use resolve method instead.
// Returns []*ConnectionParams
// a list with connection parameters
func (c *ConnectionResolver) GetAll() []*ConnectionParams {
	return c.connections
}

// Adds a new connection to component connections
// Parameters:
//   - connection *ConnectionParams
//   new connection parameters to be added
func (c *ConnectionResolver) Add(connection *ConnectionParams) {
	c.connections = append(c.connections, connection)
}

func (c *ConnectionResolver) resolveInDiscovery(correlationId string,
	connection *ConnectionParams) (result *ConnectionParams, err error) {

	if !connection.UseDiscovery() {
		return connection, nil
	}

	key := connection.DiscoveryKey()
	if c.references == nil {
		return nil, nil
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	components := c.references.GetOptional(discoveryDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, discoveryDescriptor)
		return nil, err
	}

	for _, component := range components {
		discovery, _ := component.(IDiscovery)
		if discovery != nil {
			connection, err = discovery.ResolveOne(correlationId, key)
			if connection != nil || err != nil {
				return connection, err
			}
		}
	}

	return nil, nil
}

// Resolves a single component connection. If connections are configured to be retrieved from Discovery service it finds a IDiscovery and resolves the connection there.
// see
// IDiscovery
// Parameters:
//   - correlationId: string
//   transaction id to trace execution through call chain.
// Returns *ConnectionParams, error
// resolved connection or error.
func (c *ConnectionResolver) Resolve(correlationId string) (*ConnectionParams, error) {
	if len(c.connections) == 0 {
		return nil, nil
	}

	resolveConnections := []*ConnectionParams{}

	for _, connection := range c.connections {
		if !connection.UseDiscovery() {
			return connection, nil
		}

		resolveConnections = append(resolveConnections, connection)
	}

	for _, connection := range resolveConnections {
		c, err := c.resolveInDiscovery(correlationId, connection)
		if c != nil || err != nil {
			return c, err
		}
	}

	return nil, nil
}

func (c *ConnectionResolver) resolveAllInDiscovery(correlationId string,
	connection *ConnectionParams) (result []*ConnectionParams, err error) {

	if !connection.UseDiscovery() {
		return []*ConnectionParams{connection}, nil
	}

	key := connection.DiscoveryKey()
	if c.references == nil {
		return nil, nil
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	components := c.references.GetOptional(discoveryDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, discoveryDescriptor)
		return nil, err
	}

	resolvedConnections := []*ConnectionParams{}

	for _, component := range components {
		discovery, _ := component.(IDiscovery)
		if discovery != nil {
			connections, err := discovery.ResolveAll(correlationId, key)
			if err != nil {
				return nil, err
			}
			if connections != nil {
				for _, c := range connections {
					resolvedConnections = append(resolvedConnections, c)
				}
			}
		}
	}

	return resolvedConnections, nil
}

// Resolves all component connection. If connections are configured to be retrieved from Discovery service it finds a IDiscovery and resolves the connection there.
// see
// IDiscovery
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns []*ConnectionParams, error
// resolved connections or error.
func (c *ConnectionResolver) ResolveAll(correlationId string) ([]*ConnectionParams, error) {
	resolvedConnections := []*ConnectionParams{}
	resolveConnections := []*ConnectionParams{}

	for _, connection := range c.connections {
		if !connection.UseDiscovery() {
			resolvedConnections = append(resolvedConnections, connection)
		} else {
			resolveConnections = append(resolveConnections, connection)
		}
	}

	for _, connection := range resolveConnections {
		connections, err := c.resolveAllInDiscovery(correlationId, connection)
		if err != nil {
			return nil, err
		}
		for _, c := range connections {
			resolvedConnections = append(resolvedConnections, c)
		}
	}

	return resolvedConnections, nil
}

func (c *ConnectionResolver) registerInDiscovery(correlationId string,
	connection *ConnectionParams) (result bool, err error) {

	if !connection.UseDiscovery() {
		return false, nil
	}

	key := connection.DiscoveryKey()
	if c.references == nil {
		return false, nil
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	components := c.references.GetOptional(discoveryDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, discoveryDescriptor)
		return false, err
	}

	registered := false

	for _, component := range components {
		discovery, _ := component.(IDiscovery)
		if discovery != nil {
			_, err = discovery.Register(correlationId, key, connection)
			if err != nil {
				return false, err
			}
			registered = true
		}
	}

	return registered, nil
}

// Registers the given connection in all referenced discovery services. This method can be used for dynamic service discovery.
// see
// IDiscovery
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - connection *ConnectionParams
//   a connection to register.
// Returns error
func (c *ConnectionResolver) Register(correlationId string, connection *ConnectionParams) error {
	registered, err := c.registerInDiscovery(correlationId, connection)
	if registered {
		c.connections = append(c.connections, connection)
	}
	return err
}
