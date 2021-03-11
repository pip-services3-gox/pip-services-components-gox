package test_log

import (
	"errors"
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/log"
	"github.com/stretchr/testify/assert"
)

type LoggerFixture struct {
	logger log.ILogger
}

func NewLoggerFixture(logger log.ILogger) *LoggerFixture {
	return &LoggerFixture{
		logger: logger,
	}
}

func (c *LoggerFixture) TestLogLevel(t *testing.T) {
	assert.True(t, c.logger.Level() >= log.None)
	assert.True(t, c.logger.Level() <= log.Trace)
}

func (c *LoggerFixture) TestSimpleLogging(t *testing.T) {
	c.logger.SetLevel(log.Trace)

	c.logger.Fatal("", nil, "Fatal error message")
	c.logger.Error("", nil, "Error message")
	c.logger.Warn("", "Warning message")
	c.logger.Info("", "Information message")
	c.logger.Debug("", "Debug message")
	c.logger.Trace("", "Trace message")
}

func (c *LoggerFixture) TestErrorLogging(t *testing.T) {
	err := errors.New("Test error")

	c.logger.Fatal("123", err, "Fatal error")
	c.logger.Error("123", err, "Recoverable error")
}
