package test_auth

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/auth"
	"github.com/stretchr/testify/assert"
)

func TestGetAndSetStoreKey(t *testing.T) {
	сredential := auth.NewEmptyCredentialParams()
	assert.Equal(t, "", сredential.StoreKey())

	сredential.SetStoreKey("Store key")
	assert.Equal(t, "Store key", сredential.StoreKey())
	assert.True(t, сredential.UseCredentialStore())
}

func TestGetAndSetUsername(t *testing.T) {
	сredential := auth.NewEmptyCredentialParams()
	assert.Equal(t, "", сredential.Username())

	сredential.SetUsername("Kate Negrienko")
	assert.Equal(t, "Kate Negrienko", сredential.Username())
}

func TestGetAndSetPassword(t *testing.T) {
	сredential := auth.NewEmptyCredentialParams()
	assert.Equal(t, "", сredential.Password())

	сredential.SetPassword("qwerty")
	assert.Equal(t, "qwerty", сredential.Password())
}

func TestGetAndSetAccessKey(t *testing.T) {
	сredential := auth.NewEmptyCredentialParams()
	assert.Equal(t, "", сredential.AccessKey())

	сredential.SetAccessKey("key")
	assert.Equal(t, "key", сredential.AccessKey())
}
