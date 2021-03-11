package test_info

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-gox/pip-services3-components-gox/info"
	"github.com/stretchr/testify/assert"
)

func TestContextInfo(t *testing.T) {
	contextInfo := info.NewContextInfo()

	assert.Equal(t, "unknown", contextInfo.Name)
	assert.Equal(t, "", contextInfo.Description)
	assert.True(t, len(contextInfo.ContextId) > 0)

	contextInfo.Name = "new name"
	contextInfo.Description = "new description"
	contextInfo.ContextId = "new context id"

	assert.Equal(t, "new name", contextInfo.Name)
	assert.Equal(t, "new description", contextInfo.Description)
	assert.Equal(t, "new context id", contextInfo.ContextId)
}

func TestContextInfoFromConfig(t *testing.T) {
	cfg := config.NewConfigParamsFromTuples(
		"info.name", "new name",
		"info.description", "new description",
		"properties.access_key", "key",
		"properties.store_key", "store key",
	)

	contextInfo := info.NewContextInfoFromConfig(cfg)
	assert.Equal(t, "new name", contextInfo.Name)
	assert.Equal(t, "new description", contextInfo.Description)
}
