package test

import (
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type Shutdown struct {
	started    bool
	runCode    int
	mode       string
	minTimeout int
	maxTimeout int
}

func NewShutdown() *Shutdown {
	return &Shutdown{
		started:    false,
		runCode:    0,
		mode:       "exception",
		minTimeout: 300000,
		maxTimeout: 900000,
	}
}

func (c *Shutdown) Configure(config config.ConfigParams) {
	c.mode = config.GetAsStringWithDefault("mode", c.mode)
	c.minTimeout = config.GetAsIntegerWithDefault("min_timeout", c.minTimeout)
	c.maxTimeout = config.GetAsIntegerWithDefault("max_timeout", c.maxTimeout)
}

func (c *Shutdown) IsOpen() bool {
	return c.started
}

func (c *Shutdown) Open(correlationId string) error {
	if c.started {
		return nil
	}

	delay := int(float32(c.maxTimeout-c.minTimeout)*rand.Float32() + float32(c.minTimeout))
	c.runCode++
	go c.doShutdown(delay, c.runCode)
	c.started = true

	return nil
}

func (c *Shutdown) Close(correlationId string) error {
	// Todo: Properly interrupt the go proc
	c.started = false
	return nil
}

func (c *Shutdown) Shutdown() {
	if c.mode == "null" || c.mode == "nullpointer" {
		var obj io.Writer
		obj.Write([]byte{})
	} else if c.mode == "zero" || c.mode == "dividebyzero" {
		_ = 0 / 100
	} else if c.mode == "exit" || c.mode == "processexit" {
		os.Exit(1)
	} else {
		panic("Crash test exception")
	}
}

func (c *Shutdown) doShutdown(delay int, runCode int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)

	if c.started && c.runCode == runCode {
		c.Shutdown()
	}
}
