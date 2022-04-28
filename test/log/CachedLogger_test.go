package test_log

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/log"
)

type cachedLoggerSaver struct {
	counter uint32
}

func (c *cachedLoggerSaver) Save(ctx context.Context, messages []log.LogMessage) error {
	c.counter += uint32(len(messages))
	return nil
}

var saver = &cachedLoggerSaver{}

func newCustomCachedLogger() *log.CachedLogger {
	logger := log.InheritCachedLogger(saver)
	logger.Configure(config.NewConfigParamsFromTuples(
		log.ConfigParameterOptionsInterval, 100,
		log.ConfigParameterOptionsMaxCacheSize, 1,
	))
	return logger
}

func newCachedLoggerFixture() *LoggerFixture {
	logger := newCustomCachedLogger()
	fixture := NewLoggerFixture(logger)
	return fixture
}

func TestCachedLogLevel(t *testing.T) {
	fixture := newCachedLoggerFixture()
	fixture.TestLogLevel(t)
}

func TestCachedSimpleLogging(t *testing.T) {
	fixture := newCachedLoggerFixture()
	fixture.TestSimpleLogging(t)
}

func TestCachedErrorLogging(t *testing.T) {
	fixture := newCachedLoggerFixture()
	fixture.TestErrorLogging(t)
}
