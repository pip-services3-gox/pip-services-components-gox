package build

/*
Basic component factory that creates components using registered types and factory functions.

Example
  factory := NewFactory();
  factory.RegisterAsType(
      NewDescriptor("mygroup", "mycomponent1", "default", "*", "1.0"),
      MyComponent1
  );
  factory.Register(
      NewDescriptor("mygroup", "mycomponent2", "default", "*", "1.0"),
      (locator){
          return NewMyComponent2();
      }
  );

  factory.Create(NewDescriptor("mygroup", "mycomponent1", "default", "name1", "1.0"))
  factory.Create(NewDescriptor("mygroup", "mycomponent2", "default", "name2", "1.0"))
*/

import (
	refl "reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type registration struct {
	locator interface{}
	factory func() interface{}
}

type Factory struct {
	registrations []*registration
}

// Create new factory
// Return *Factory
func NewFactory() *Factory {
	return &Factory{
		registrations: []*registration{},
	}
}

// Registers a component using a factory method.
// Parameters:
//   - locator interface{}
//   a locator to identify component to be created.
//   factory func() interface{}
//   a factory function that receives a locator and returns a created component.
func (c *Factory) Register(locator interface{}, factory func() interface{}) {
	if locator == nil {
		panic("Locator cannot be nil")
	}
	if factory == nil {
		panic("Factory cannot be nil")
	}

	c.registrations = append(c.registrations, &registration{
		locator: locator,
		factory: factory,
	})
}

// Registers a component using its type (a constructor function).
// Parameters:
//   - locator interface{}
//   a locator to identify component to be created.
//   - factory interface{}
//   a factory.
func (c *Factory) RegisterType(locator interface{}, factory interface{}) {
	if locator == nil {
		panic("Locator cannot be nil")
	}
	if factory == nil {
		panic("Factory cannot be nil")
	}

	val := refl.ValueOf(factory)
	if val.Kind() != refl.Func {
		panic("Factory must be parameterless function")
	}

	c.Register(locator, func() interface{} {
		return val.Call([]refl.Value{})[0].Interface()
	})
}

// Checks if this factory is able to create component by given locator.
// This method searches for all registered components and returns a locator for component it is able to create that matches the given locator. If the factory is not able to create a requested component is returns null.
// Parameters:
//   - locator interface{}
//   a locator to identify component to be created.
// Returns interface{}
// a locator for a component that the factory is able to create.
func (c *Factory) CanCreate(locator interface{}) interface{} {
	for _, registration := range c.registrations {
		thisLocator := registration.locator

		equatable, ok := thisLocator.(data.IEquatable)
		if ok && equatable.Equals(locator) {
			return thisLocator
		}

		if thisLocator == locator {
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
// the created component and a CreateError if the factory is not able to create the component.
func (c *Factory) Create(locator interface{}) (interface{}, error) {
	var factory func() interface{}

	for _, registration := range c.registrations {
		thisLocator := registration.locator

		equatable, ok := thisLocator.(data.IEquatable)
		if ok && equatable.Equals(locator) {
			factory = registration.factory
			break
		}

		if thisLocator == locator {
			factory = registration.factory
			break
		}
	}

	if factory == nil {
		return nil, NewCreateErrorByLocator("", locator)
	}

	var err error

	obj := func() interface{} {
		defer func() {
			if r := recover(); r != nil {
				tempMessage := convert.StringConverter.ToString(r)
				tempError := NewCreateError("", tempMessage)

				cause, ok := r.(error)
				if ok {
					tempError.WithCause(cause)
				}

				err = tempError
			}
		}()

		return factory()
	}()

	return obj, err
}
