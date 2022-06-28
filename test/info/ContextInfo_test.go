package test_info

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-components-gox/info"
	"github.com/stretchr/testify/assert"
)

func TestContextInfo(t *testing.T) {
	contextInfo := info.NewContextInfo()

	assert.Equal(t, info.ContextInfoNameUnknown, contextInfo.Name)
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
		info.ContextInfoParameterInfoName, "new name",
		info.ContextInfoParameterInfoDescription, "new description",
		info.ContextInfoSectionNameProperties+".access_key", "key",
		info.ContextInfoSectionNameProperties+".store_key", "store key",
	)

	contextInfo := info.NewContextInfoFromConfig(context.Background(), cfg)
	assert.Equal(t, "new name", contextInfo.Name)
	assert.Equal(t, "new description", contextInfo.Description)
}
