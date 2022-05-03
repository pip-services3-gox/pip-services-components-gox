package test_test

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/test"
	"github.com/stretchr/testify/assert"
)

func TestShutdown(t *testing.T) {
	sd := test.NewShutdown()

	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	sd.Shutdown(context.Background())
}
