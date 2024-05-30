// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eginnovations // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/eginnovations"

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exportertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/eginnovations/internal/metadata"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateUnmarshalConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	tests := struct {
		id       component.ID
		expected component.Config
	}{
		id: component.NewIDWithName(metadata.Type, "all"),
		expected: &Config{
			UserID: "user1",
			Token:  "xxxxxxxxx",
			ClientConfig: configgrpc.ClientConfig{
				Endpoint:    "eginnovations.com",
				Compression: configcompression.TypeGzip,
				TLSSetting: configtls.ClientConfig{
					Config:             configtls.Config{},
					Insecure:           false,
					InsecureSkipVerify: false,
					ServerName:         "",
				},
				ReadBufferSize:  0,
				WriteBufferSize: 0,
				WaitForReady:    false,
				Headers: map[string]configopaque.String{
					"hdr1": "value1",
				},
				BalancerName: "",
			},
		},
	}

	t.Run(tests.id.String(), func(t *testing.T) {
		factory := NewFactory()
		cfg := factory.CreateDefaultConfig()
		sub, err := cm.Sub(tests.id.String())
		require.NoError(t, err)
		require.NoError(t, component.UnmarshalConfig(sub, cfg))
	})

}

func TestCreateTestTraceExporter(t *testing.T) {
	cfg := createDefaultConfig()
	egCfg := cfg.(*Config)
	egCfg.Endpoint = "http://localhost:9999"
	eg, err := createTracesExporter(context.Background(), exportertest.NewNopCreateSettings(), cfg)
	assert.NoError(t, err)
	assert.NotNil(t, eg)
}
