package log

import "github.com/pip-services3-go/pip-services3-commons-go/refer"

/*
Aggregates all loggers from component references under a single component.

It allows to log messages and conveniently send them to multiple destinations.

References
*:logger:*:*:1.0 (optional) ILogger components to pass log messages
see
ILogger

Example
  type MyComponent {
      _logger CompositeLogger
  }
      func (mc* MyComponent) Configure(config: ConfigParams): void {
          mc._logger.Configure(config);
          ...
      }

      func (mc* MyComponent) SetReferences(references: IReferences): void {
          mc._logger.SetReferences(references);
          ...
      }

      func (mc* MyComponent)myMethod(string correlationId): void {
          mc._logger.Debug(correlationId, "Called method mycomponent.mymethod");
          ...
      }
  var mc MyComponent = MyComponent{}
  mc._logger = NewCompositeLogger();
*/
type CompositeLogger struct {
	Logger
	loggers []ILogger
}

// Creates a new instance of the logger.
// Returns *CompositeLogger
func NewCompositeLogger() *CompositeLogger {
	c := &CompositeLogger{
		loggers: []ILogger{},
	}
	c.Logger = *InheritLogger(c)
	c.SetLevel(Trace)
	return c
}

// Creates a new instance of the logger.
// Parameters:
//   - references refer.IReferences
//   references to locate the component dependencies.
// Returns CompositeLogger
func NewCompositeLoggerFromReferences(references refer.IReferences) *CompositeLogger {
	c := NewCompositeLogger()
	c.SetReferences(references)
	return c
}

// Sets references to dependent components.
// Parameters:
//   - references IReferences
//   references to locate the component dependencies.
func (c *CompositeLogger) SetReferences(references refer.IReferences) {
	c.Logger.SetReferences(references)

	if c.loggers == nil {
		c.loggers = []ILogger{}
	}

	loggers := references.GetOptional(refer.NewDescriptor("*", "logger", "*", "*", "*"))
	for _, l := range loggers {
		if l == c {
			continue
		}

		logger, ok := l.(ILogger)
		if ok {
			c.loggers = append(c.loggers, logger)
		}
	}
}

// Writes a log message to the logger destination(s).
// Parameters:
//   - level int
//   a log level.
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
func (c *CompositeLogger) Write(level int, correlationId string, err error, message string) {
	if c.loggers == nil && len(c.loggers) == 0 {
		return
	}

	for _, logger := range c.loggers {
		logger.Log(level, correlationId, err, message)
	}
}
