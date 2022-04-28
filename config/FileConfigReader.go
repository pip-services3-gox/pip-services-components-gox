package config

import cconfig "github.com/pip-services3-gox/pip-services3-commons-gox/config"

// FileConfigReader is an abstract config reader that reads configuration from a file.
// Child classes add support for config files in their specific format like JSON, YAML or property files.
//	Configuration parameters:
//		path: path to configuration file
//		parameters: this entire section is used as template parameters
type FileConfigReader struct {
	ConfigReader
	path string
}

// FileConfigReaderPathKey is a constant for path key
const FileConfigReaderPathKey = "path"

// NewEmptyFileConfigReader creates a new instance of the config reader.
//	Returns: *FileConfigReader
func NewEmptyFileConfigReader() *FileConfigReader {
	return &FileConfigReader{
		ConfigReader: *NewConfigReader(),
	}
}

// NewFileConfigReader creates a new instance of the config reader.
//	Parameters: path string a path to configuration file.
//	Returns: *FileConfigReader
func NewFileConfigReader(path string) *FileConfigReader {
	return &FileConfigReader{
		ConfigReader: *NewConfigReader(),
		path:         path,
	}
}

// Configure component by passing configuration parameters.
//	Parameters: config *cconfig.ConfigParams configuration parameters to be set.
func (c *FileConfigReader) Configure(config *cconfig.ConfigParams) {
	c.ConfigReader.Configure(config)
	c.path = config.GetAsStringWithDefault(FileConfigReaderPathKey, c.path)
}

// Path get the path to configuration file..
//	Returns: string the path to configuration file.
func (c *FileConfigReader) Path() string {
	return c.path
}

// SetPath set the path to configuration file.
//	Parameters: path string a new path to configuration file.
func (c *FileConfigReader) SetPath(path string) {
	c.path = path
}
