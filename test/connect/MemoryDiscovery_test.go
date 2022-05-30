package test_connect

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-components-gox/connect"
	"github.com/stretchr/testify/assert"
)

func TestMemoryDiscoveryReadConnections(t *testing.T) {
	config := config.NewConfigParamsFromTuples(
		"connections.key1.host", "10.1.1.100",
		"connections.key1.port", "8080",
		"connections.key2.host", "10.1.1.101",
		"connections.key2.port", "8082",
	)

	discovery := connect.NewEmptyMemoryDiscovery()
	discovery.Configure(context.Background(), config)

	// Resolve one
	connection, err := discovery.ResolveOne("123", "key1")

	assert.Equal(t, err, nil)
	assert.Equal(t, "10.1.1.100", connection.Host())
	assert.Equal(t, 8080, connection.Port())

	connection, err = discovery.ResolveOne("123", "key2")

	assert.Equal(t, err, nil)
	assert.Equal(t, "10.1.1.101", connection.Host())
	assert.Equal(t, 8082, connection.Port())

	// Resolve all
	_, err = discovery.Register("123", "key1", connect.NewConnectionParamsFromTuples(
		"host", "10.3.3.151",
	))
	assert.Equal(t, err, nil)

	connections, err := discovery.ResolveAll("123", "key1")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(connections) > 1, true)
}
