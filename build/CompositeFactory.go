package build

/*

Aggregates multiple factories into a single factory component. When a new component is requested, it iterates through factories to locate the one able to create the requested component.

This component is used to conveniently keep all supported factories in a single place.

Example
  factory := NewCompositeFactory();
  factory.Add(NewDefaultLoggerFactory());
  factory.Add(NewDefaultCountersFactory());

  loggerLocator := NewDescriptor("*", "logger", "*", "*", "1.0");
  factory.CanCreate(loggerLocator);         // Result: Descriptor("pip-service", "logger", "null", "default", "1.0")
  factory.Create(loggerLocator);             // Result: created NullLogger
*/
type CompositeFactory struct {
	factories []IFactory
}

// Creates a new instance of the factory.
// Returns *CompositeFactory
func NewCompositeFactory() *CompositeFactory {
	return &CompositeFactory{
		factories: []IFactory{},
	}
}

// Creates a new instance of the factory.
// Parameters:
//   - factories ...IFactory
//   a list of factories to embed into this factory.
// Returns *CompositeFactory
func NewCompositeFactoryFromFactories(factories ...IFactory) *CompositeFactory {
	return &CompositeFactory{
		factories: factories,
	}
}

// Adds a factory into the list of embedded factories.
// Parameters:
//   - factory IFactory
//   a factory to be added.
func (c *CompositeFactory) Add(factory IFactory) {
	if factory == nil {
		panic("Factory cannot be nil")
	}

	c.factories = append(c.factories, factory)
}

// Removes a factory from the list of embedded factories.
// Parameters:
//   - factory IFactory
//   the factory to remove.
func (c *CompositeFactory) Remove(factory IFactory) {
	for i, thisFactory := range c.factories {
		if thisFactory == factory {
			c.factories = append(c.factories[:i], c.factories[i+1:]...)
			break
		}
	}
}

// Checks if this factory is able to create component by given locator.
// This method searches for all registered components and returns a locator for component it is able to create that matches the given locator. If the factory is not able to create a requested component is returns null.
// Parameters:
//   - locator interface{}
//   a locator to identify component to be created.
// Returns interface{}
// a locator for a component that the factory is able to create.
func (c *CompositeFactory) CanCreate(locator interface{}) interface{} {
	if locator == nil {
		panic("Locator cannot be null")
	}

	// Iterate from the latest factories
	for i := len(c.factories) - 1; i >= 0; i-- {
		thisLocator := c.factories[i].CanCreate(locator)
		if thisLocator != nil {
			return thisLocator
		}
	}

	return nil
}

// Creates a component identified by given locator.
// Parameters:
//   - locator interface{}
//   a locator to identify component to be created.
// Returns interface{}, error
// the created component and a CreateError if the factory is not able to create the component..
func (c *CompositeFactory) Create(locator interface{}) (interface{}, error) {
	if locator == nil {
		panic("Locator cannot be null")
	}

	// Iterate from the latest factories
	for i := len(c.factories) - 1; i >= 0; i-- {
		factory := c.factories[i]
		if factory.CanCreate(locator) != nil {
			return factory.Create(locator)
		}
	}

	return nil, NewCreateErrorByLocator("", locator)
}
