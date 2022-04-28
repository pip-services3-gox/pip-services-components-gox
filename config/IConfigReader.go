package config

import (
	"context"
	cconfig "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

// IConfigReader Interface for configuration readers that retrieve configuration from
// various sources and make it available for other components.
// Some IConfigReader implementations may support configuration parameterization.
// The parameterization allows using configuration as a template and inject there dynamic values.
// The values may come from application command like arguments or environment variables.
type IConfigReader interface {

	// ReadConfig reads configuration and parameterize it with given values.
	ReadConfig(ctx context.Context, correlationId string, parameters *cconfig.ConfigParams) (*cconfig.ConfigParams, error)
}
