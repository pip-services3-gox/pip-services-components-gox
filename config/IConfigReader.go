package config

import c "github.com/pip-services3-go/pip-services3-commons-go/config"

/*
Interface for configuration readers that retrieve configuration from various sources and make it available for other components.

Some IConfigReader implementations may support configuration parameterization. The parameterization allows to use configuration as a template and inject there dynamic values. The values may come from application command like arguments or environment variables.
*/
type IConfigReader interface {
	// Reads configuration and parameterize it with given values.
	ReadConfig(correlationId string, parameters *c.ConfigParams) (*c.ConfigParams, error)
}
