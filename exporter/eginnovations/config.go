// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package eginnovations

import (
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
)

type Config struct {
	UserID                  string              `mapstructure:"userId"`
	Token                   configopaque.String `mapstructure:"token"`
	Debug                   bool                `mapstructure:"debug"`
	configgrpc.ClientConfig `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

func (c *Config) Validate() error {
	if c.Endpoint == "" {
		return fmt.Errorf("endpoint not specified, please fix the configuration")
	}
	return nil

}
