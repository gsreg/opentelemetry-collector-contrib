package eginnovations

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var Type = component.MustNewType("eginnovations")

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelBeta),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		UserID: "",
		Token:  "",
		ClientConfig: configgrpc.ClientConfig{
			Compression: configcompression.TypeGzip,
		},
	}
}

func createTracesExporter(

	ctx context.Context,
	set exporter.CreateSettings,
	config component.Config,

) (exporter.Traces, error) {
	cfg := config.(*Config)
	egExporter := NewEgExporter(config, set)

	return exporterhelper.NewTracesExporter(
		ctx,
		set,
		cfg,
		egExporter.ConsumeTraces,
		exporterhelper.WithStart(egExporter.Start),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(configretry.BackOffConfig{Enabled: false}),
		exporterhelper.WithQueue(exporterhelper.QueueSettings{Enabled: false}),
	)
}
