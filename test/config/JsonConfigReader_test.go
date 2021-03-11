package test_config

import (
	"testing"

	pconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-gox/pip-services3-components-gox/config"
	"github.com/stretchr/testify/assert"
)

func TestJsonConfigReader(t *testing.T) {
	parameters := pconfig.NewConfigParamsFromTuples(
		"param1", "Test Param 1",
		"param2", "Test Param 2",
	)
	config, err := config.ReadJsonConfig("", "./config.json", parameters)

	assert.Nil(t, err)
	assert.Equal(t, 9, config.Len())
	assert.Equal(t, 123, config.GetAsInteger("field1.field11"))
	assert.Equal(t, "ABC", config.GetAsString("field1.field12"))
	assert.Equal(t, 123, config.GetAsInteger("field2.0"))
	assert.Equal(t, "ABC", config.GetAsString("field2.1"))
	assert.Equal(t, 543, config.GetAsInteger("field2.2.field21"))
	assert.Equal(t, "XYZ", config.GetAsString("field2.2.field22"))
	assert.Equal(t, true, config.GetAsBoolean("field3"))
	assert.Equal(t, "Test Param 1", config.GetAsString("field4"))
	assert.Equal(t, "Test Param 2", config.GetAsString("field5"))
}
