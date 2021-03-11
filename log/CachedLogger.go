package log

import (
	"sync"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

/*
Abstract logger that caches captured log messages in memory and periodically dumps them. Child classes implement saving cached messages to their specified destinations.

Configuration parameters
  level: maximum log level to capture
  source: source (context) name
  options:
    interval: interval in milliseconds to save log messages (default: 10 seconds)
    max_cache_size: maximum number of messages stored in this cache (default: 100)
References
*:context-info:*:*:1.0 (optional) ContextInfo to detect the context id and specify counters source
*/
type ICachedLogSaver interface {
	Save(messages []*LogMessage) error
}

type CachedLogger struct {
	Logger
	Cache        []*LogMessage
	Updated      bool
	LastDumpTime time.Time
	MaxCacheSize int
	Interval     int
	Lock         *sync.Mutex
	saver        ICachedLogSaver
}

// Creates a new instance of the logger from ICachedLogSaver
// Parameters:
//  - saver ICachedLogSaver
// Returns CachedLogger
func InheritCachedLogger(saver ICachedLogSaver) *CachedLogger {
	c := &CachedLogger{
		Cache:        []*LogMessage{},
		Updated:      false,
		LastDumpTime: time.Now(),
		MaxCacheSize: 100,
		Interval:     10000,
		Lock:         &sync.Mutex{},
		saver:        saver,
	}
	c.Logger = *InheritLogger(c)
	return c
}

// Writes a log message to the logger destination.
// Parameters:
//   - level LogLevel
//   a log level.
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
func (c *CachedLogger) Write(level int, correlationId string, err error, message string) {
	logMessage := &LogMessage{
		Time:          time.Now().UTC(),
		Level:         level,
		Source:        c.source,
		Message:       message,
		CorrelationId: correlationId,
	}

	if err != nil {
		errorDescription := errors.NewErrorDescription(err)
		logMessage.Error = *errorDescription
	}

	c.Lock.Lock()
	c.Cache = append(c.Cache, logMessage)
	c.Lock.Unlock()

	c.Update()
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be set.
func (c *CachedLogger) Configure(cfg *config.ConfigParams) {
	c.Logger.Configure(cfg)

	c.Interval = cfg.GetAsIntegerWithDefault("options.interval", c.Interval)
	c.MaxCacheSize = cfg.GetAsIntegerWithDefault("options.max_cache_size", c.MaxCacheSize)
}

// Clears (removes) all cached log messages.
func (c *CachedLogger) Clear() {
	c.Lock.Lock()
	c.Cache = []*LogMessage{}
	c.Updated = false
	c.Lock.Unlock()
}

// Dumps (writes) the currently cached log messages.
func (c *CachedLogger) Dump() error {
	if c.Updated {
		if !c.Updated {
			return nil
		}

		var messages []*LogMessage
		c.Lock.Lock()

		messages = c.Cache
		c.Cache = []*LogMessage{}

		c.Lock.Unlock()

		err := c.saver.Save(messages)
		if err != nil {
			c.Lock.Lock()

			// Put failed messages back to cache
			c.Cache = append(messages, c.Cache...)

			// Truncate cache to max size
			if len(c.Cache) > c.MaxCacheSize {
				c.Cache = c.Cache[len(c.Cache)-c.MaxCacheSize:]
			}

			c.Lock.Unlock()
		}

		c.Updated = false
		c.LastDumpTime = time.Now()
		return err
	}
	return nil
}

// Makes message cache as updated and dumps it when timeout expires.
func (c *CachedLogger) Update() {
	c.Updated = true

	elapsed := int(time.Since(c.LastDumpTime).Seconds() * 1000)

	if elapsed > c.Interval {
		// Todo: Decide what to do with the error
		c.Dump()
	}
}
