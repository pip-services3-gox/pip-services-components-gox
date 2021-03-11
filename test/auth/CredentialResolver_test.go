package test_auth

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-gox/pip-services3-components-gox/auth"
	"github.com/stretchr/testify/assert"
)

func TestCredentialResolverConfigure(t *testing.T) {
	restConfig := config.NewConfigParamsFromTuples(
		"credential.username", "Negrienko",
		"credential.password", "qwerty",
		"credential.access_key", "key",
		"credential.store_key", "store key",
	)
	credentialResolver := auth.NewCredentialResolver(restConfig, nil)
	credentials := credentialResolver.GetAll()
	assert.Len(t, credentials, 1)

	credential := credentials[0]
	assert.Equal(t, "Negrienko", credential.Username())
	assert.Equal(t, "qwerty", credential.Password())
	assert.Equal(t, "key", credential.AccessKey())
	assert.Equal(t, "store key", credential.StoreKey())
}

func TestCredentialResolverLookup(t *testing.T) {
	credentialResolver := auth.NewEmptyCredentialResolver()

	credential, err := credentialResolver.Lookup("")
	assert.Nil(t, err)
	assert.Nil(t, credential)

	restConfigWithoutStoreKey := config.NewConfigParamsFromTuples(
		"credential.username", "Negrienko",
		"credential.password", "qwerty",
		"credential.access_key", "key",
	)
	credentialResolver = auth.NewCredentialResolver(restConfigWithoutStoreKey, nil)

	credential, err = credentialResolver.Lookup("")
	assert.Nil(t, err)
	assert.NotNil(t, credential)
	assert.Equal(t, "Negrienko", credential.Username())
	assert.Equal(t, "qwerty", credential.Password())
	assert.Equal(t, "key", credential.AccessKey())
	assert.Equal(t, "", credential.StoreKey())

	restConfig := config.NewConfigParamsFromTuples(
		"credential.username", "Negrienko",
		"credential.password", "qwerty",
		"credential.access_key", "key",
		"credential.store_key", "store key",
	)
	credentialResolver = auth.NewCredentialResolver(restConfig, nil)

	credential, err = credentialResolver.Lookup("")
	assert.Nil(t, err)
	assert.Nil(t, credential)
}
