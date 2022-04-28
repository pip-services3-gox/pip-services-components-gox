package log

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
)

// CompositeLogger aggregates all loggers from component references under a single component.
// It allows logging messages and conveniently send them to multiple destinations.
//	References:
//		*:logger:*:*:1.0 (optional) ILogger components to pass log messages
//	see ILogger
//	Example:
//		type MyComponent {
//			_logger CompositeLogger
//		}
//		func (mc* MyComponent) Configure(config: ConfigParams): void {
//			mc._logger.Configure(config);
//			...
//		}
//
//		func (mc* MyComponent) SetReferences(references: IReferences): void {
//			mc._logger.SetReferences(references);
//			...
//		}
//
//		func (mc* MyComponent) myMethod(ctx context.Context, string correlationId): void {
//			mc._logger.Debug(ctx context.Context, correlationId, "Called method mycomponent.mymethod");
//			...
//		}
//		var mc MyComponent = MyComponent{}
//		mc._logger = NewCompositeLogger();
type CompositeLogger struct {
	Logger
	loggers []ILogger
}

// NewCompositeLogger creates a new instance of the logger.
//	Returns: *CompositeLogger
func NewCompositeLogger() *CompositeLogger {
	c := &CompositeLogger{
		loggers: []ILogger{},
	}
	c.Logger = *InheritLogger(c)
	c.SetLevel(LevelTrace)
	return c
}

// NewCompositeLoggerFromReferences creates a new instance of the logger.
//	Parameters: refer.IReferences references to locate the component dependencies.
//	Returns: CompositeLogger
func NewCompositeLoggerFromReferences(references refer.IReferences) *CompositeLogger {
	c := NewCompositeLogger()
	c.SetReferences(references)
	return c
}

// SetReferences sets references to dependent components.
//	Parameters: refer.IReferences references to locate the component dependencies.
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

		if logger, ok := l.(ILogger); ok {
			c.loggers = append(c.loggers, logger)
		}
	}
}

// Writes a log message to the logger destination(s).
// Parameters:
//		- ctx context.Context
//		- level LogLevel a log level.
//		- correlationId string transaction id to trace execution through call chain.
//		- err error an error object associated with this message.
//		- message string a human-readable message to log.
func (c *CompositeLogger) Write(ctx context.Context, level LevelType, correlationId string, err error, message string) {
	if c.loggers == nil && len(c.loggers) == 0 {
		return
	}

	for _, logger := range c.loggers {
		logger.Log(ctx, level, correlationId, err, message)
	}
}
