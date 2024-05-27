package eginnovations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	config := Config{}
	config.Endpoint = ""
	expected := fmt.Errorf("endpoint not specified, please fix the configuration")
	err := config.Validate()
	assert.Equal(t, expected, err, "missing endpoint")
	config.Endpoint = "localhost:8080"
	noErr := config.Validate()
	assert.Equal(t, nil, noErr, "valid config")
}
