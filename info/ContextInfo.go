package info

import (
	"os"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Context information component that provides detail information about execution context: container or/and process.

Most often ContextInfo is used by logging and performance counters to identify source of the collected logs and metrics.

Configuration parameters
  name: the context (container or process) name
  description: human-readable description of the context
  properties: entire section of additional descriptive properties
  ...
Example
  contextInfo := NewContextInfo();
  contextInfo.Configure(NewConfigParamsFromTuples(
      "name", "MyMicroservice",
      "description", "My first microservice"
  ));

  context.Name;            // Result: "MyMicroservice"
  context.ContextId;        // Possible result: "mylaptop"
  context.StartTime;        // Possible result: 2018-01-01:22:12:23.45Z
*/
type ContextInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ContextId   string            `json:"context_id"`
	StartTime   time.Time         `json:"start_time"`
	Properties  map[string]string `json:"properties"`
}

// Creates a new instance of this context info.
// Returns *ContextInfo
func NewContextInfo() *ContextInfo {
	c := &ContextInfo{
		Name:       "unknown",
		StartTime:  time.Now(),
		Properties: map[string]string{},
	}
	c.ContextId, _ = os.Hostname()
	return c
}

// Calculates the context uptime as from the start time.
// Returns int64
// number of milliseconds from the context start time.
func (c *ContextInfo) Uptime() int64 {
	return time.Now().Unix() - c.StartTime.Unix()
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be set.
func (c *ContextInfo) Configure(cfg *config.ConfigParams) {
	c.Name = cfg.GetAsStringWithDefault("name", c.Name)
	c.Name = cfg.GetAsStringWithDefault("info.name", c.Name)

	c.Description = cfg.GetAsStringWithDefault("description", c.Description)
	c.Description = cfg.GetAsStringWithDefault("info.description", c.Description)

	c.Properties = cfg.GetSection("properties").InnerValue().(map[string]string)
}

// Creates a new instance of this context info.
// Parameters:
//   - ncfg *config.ConfigParams
//   a context configuration parameters.
// Returns *ContextInfo
func NewContextInfoFromConfig(cfg *config.ConfigParams) *ContextInfo {
	result := NewContextInfo()
	result.Configure(cfg)
	return result
}
