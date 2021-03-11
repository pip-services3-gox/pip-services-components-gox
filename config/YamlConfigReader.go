package config

import (
	"fmt"
	"io/ioutil"

	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"gopkg.in/yaml.v2"
)

/*
Config reader that reads configuration from YAML file.

The reader supports parameterization using Handlebars template engine.

Configuration parameters
  path: path to configuration file
  parameters: this entire section is used as template parameters
  ...
see
IConfigReader

see
FileConfigReader

Example
  ======== config.yml ======
  key1: "{{KEY1_VALUE}}"
  key2: "{{KEY2_VALUE}}"
  ===========================

  configReader := NewYamlConfigReader("config.yml")

  parameters := NewConfigParamsFromTuples("KEY1_VALUE", 123, "KEY2_VALUE", "ABC");
  res, err := configReader.ReadConfig("123", parameters);
  // Result: key1=123;key2=ABC

*/
type YamlConfigReader struct {
	FileConfigReader
}

// Creates a new instance of the config reader.
// Returns *YamlConfigReader
func NewEmptyYamlConfigReader() *YamlConfigReader {
	return &YamlConfigReader{
		FileConfigReader: *NewEmptyFileConfigReader(),
	}
}

// Creates a new instance of the config reader.
// Parameters:
//   - path string
//   a path to configuration file.
// Returns *YamlConfigReader
func NewYamlConfigReader(path string) *YamlConfigReader {
	return &YamlConfigReader{
		FileConfigReader: *NewFileConfigReader(path),
	}
}

// Reads configuration file, parameterizes its content and converts it into JSON object.
// Parameters:
//   - correlationId string
//    transaction id to trace execution through call chain.
//   - parameters *cconfig.ConfigParams
//   values to parameters the configuration.
// Returns interface{}, error
// a JSON object with configuration and error.
func (c *YamlConfigReader) ReadObject(correlationId string,
	parameters *cconfig.ConfigParams) (interface{}, error) {

	if c.Path() == "" {
		return nil, errors.NewConfigError(correlationId, "NO_PATH", "Missing config file path")
	}

	b, err := ioutil.ReadFile(c.Path())
	if err != nil {
		err = errors.NewFileError(
			correlationId,
			"READ_FAILED",
			"Failed reading configuration "+c.Path()+": "+err.Error(),
		).WithDetails("path", c.Path()).WithCause(err)
		return nil, err
	}

	data := string(b)
	data, err = c.Parameterize(data, parameters)
	if err != nil {
		return nil, err
	}

	var m interface{}
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, err
	}

	//return convert.MapConverter.ToMap(m), nil
	return m, err
}

// Reads configuration from a file, parameterize it with given values and returns a new ConfigParams object.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - path string
//   - parameters *cconfig.ConfigParams
//   values to parameters the configuration or null to skip parameterization.
// Returns *cconfgi.ConfigParams, error
func (c *YamlConfigReader) ReadConfig(correlationId string,
	parameters *cconfig.ConfigParams) (result *cconfig.ConfigParams, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	value, err := c.ReadObject(correlationId, parameters)
	if err != nil {
		return nil, err
	}

	config := cconfig.NewConfigParamsFromValue(value)
	return config, err
}

// Reads configuration file, parameterizes its content and converts it into JSON object.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - path string
//   - parameters *cconfig.ConfigParams
//   values to parameters the configuration.
// Returns interface{}, error
// a JSON object with configuration.
func ReadYamlObject(correlationId string, path string,
	parameters *cconfig.ConfigParams) (interface{}, error) {

	reader := NewYamlConfigReader(path)
	return reader.ReadObject(correlationId, parameters)
}

// Reads configuration from a file, parameterize it with given values and returns a new ConfigParams object.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - path string
//   - parameters *cconfig.ConfigParams
//   values to parameters the configuration or null to skip parameterization.
// Returns *cconfig.ConfigParams, error
func ReadYamlConfig(correlationId string, path string,
	parameters *cconfig.ConfigParams) (*cconfig.ConfigParams, error) {

	reader := NewYamlConfigReader(path)
	return reader.ReadConfig(correlationId, parameters)
}
