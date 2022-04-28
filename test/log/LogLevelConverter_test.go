package test_log

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/log"
	"github.com/stretchr/testify/assert"
)

func TestLogLevelConverter(t *testing.T) {
	level := log.LevelConverter.ToLogLevel("info")
	assert.Equal(t, log.LevelInfo, level)

	level = log.LevelConverter.ToLogLevel("4")
	assert.Equal(t, log.LevelInfo, level)

	str := log.LevelConverter.ToString(level)
	assert.Equal(t, "INFO", str)
}
