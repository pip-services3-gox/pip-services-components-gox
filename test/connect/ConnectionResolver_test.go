package test_connect

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/connect"
	"github.com/stretchr/testify/assert"
)

func TestConnectionResolverConfigure(t *testing.T) {
	restConfig := config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", "localhost",
		"connection.port", 3000,
	)
	connectionResolver := connect.NewConnectionResolver(context.Background(), restConfig, nil)
	connections := connectionResolver.GetAll()
	assert.Len(t, connections, 1)

	connection := connections[0]
	assert.Equal(t, "http", connection.Protocol())
	assert.Equal(t, "localhost", connection.Host())
	assert.Equal(t, 3000, connection.Port())
}

func TestConnectionResolverRegister(t *testing.T) {
	connection := connect.NewEmptyConnectionParams()
	connectionResolver := connect.NewEmptyConnectionResolver()

	err := connectionResolver.Register("", connection)
	assert.Nil(t, err)

	connections := connectionResolver.GetAll()
	assert.Len(t, connections, 0)

	err = connectionResolver.Register("", connection)
	assert.Nil(t, err)

	connections = connectionResolver.GetAll()
	assert.Len(t, connections, 0)

	connection.SetDiscoveryKey("Discovery key")
	err = connectionResolver.Register("", connection)
	assert.Nil(t, err)

	connections = connectionResolver.GetAll()
	assert.Len(t, connections, 0)
}

func TestConnectionResolverResolve(t *testing.T) {
	restConfig := config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", "localhost",
		"connection.port", 3000,
	)
	connectionResolver := connect.NewConnectionResolver(context.Background(), restConfig, nil)

	connection, err := connectionResolver.Resolve("")
	assert.Nil(t, err)
	assert.NotNil(t, connection)

	restConfigDiscovery := config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", "localhost",
		"connection.port", 3000,
		"connection.discovery_key", "Discovery key",
	)
	references := refer.NewEmptyReferences()
	connectionResolver = connect.NewConnectionResolver(context.Background(), restConfigDiscovery, references)

	connection, err = connectionResolver.Resolve("")
	assert.NotNil(t, err)
	assert.Nil(t, connection)
}
