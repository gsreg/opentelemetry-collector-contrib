package egexporter

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/egexporter/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	//"go.opentelemetry.io/collector/featuregate"
)

//var featureGateJSON = featuregate.GlobalRegistry().MustRegister("clickhouse.json", featuregate.StageAlpha)

// NewFactory creates a factory for the ClickHouse exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithMetrics(createMetricExporter, metadata.MetricsStability),
	)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	c := cfg.(*Config)
	c.collectorVersion = set.BuildInfo.Version

	//if featureGateJSON.IsEnabled() {
	//	exp := newLogsJSONExporter(set.Logger, c)
	//
	//	return exporterhelper.NewLogs(
	//		ctx,
	//		set,
	//		cfg,
	//		exp.pushLogsData,
	//		exporterhelper.WithStart(exp.start),
	//		exporterhelper.WithShutdown(exp.shutdown),
	//		exporterhelper.WithTimeout(c.TimeoutSettings),
	//		exporterhelper.WithQueue(c.QueueSettings),
	//		exporterhelper.WithRetry(c.BackOffConfig),
	//	)
	//}

	exp := newLogsExporter(set.Logger, c)

	return exporterhelper.NewLogs(
		ctx,
		set,
		cfg,
		exp.pushLogsData,
		exporterhelper.WithStart(exp.start),
		exporterhelper.WithShutdown(exp.shutdown),
		exporterhelper.WithTimeout(c.TimeoutSettings),
		exporterhelper.WithQueue(c.QueueSettings),
		exporterhelper.WithRetry(c.BackOffConfig),
	)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	c := cfg.(*Config)
	c.collectorVersion = set.BuildInfo.Version

	//if featureGateJSON.IsEnabled() {
	//	exp := newTracesJSONExporter(set.Logger, c)
	//
	//	return exporterhelper.NewTraces(
	//		ctx,
	//		set,
	//		cfg,
	//		exp.pushTraceData,
	//		exporterhelper.WithStart(exp.start),
	//		exporterhelper.WithShutdown(exp.shutdown),
	//		exporterhelper.WithTimeout(c.TimeoutSettings),
	//		exporterhelper.WithQueue(c.QueueSettings),
	//		exporterhelper.WithRetry(c.BackOffConfig),
	//	)
	//}

	exp := newTracesExporter(set.Logger, c)

	return exporterhelper.NewTraces(
		ctx,
		set,
		cfg,
		exp.pushTraceData,
		exporterhelper.WithStart(exp.start),
		exporterhelper.WithShutdown(exp.shutdown),
		exporterhelper.WithTimeout(c.TimeoutSettings),
		exporterhelper.WithQueue(c.QueueSettings),
		exporterhelper.WithRetry(c.BackOffConfig),
	)
}

func createMetricExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	c := cfg.(*Config)
	c.collectorVersion = set.BuildInfo.Version
	exp := newMetricsExporter(set.Logger, c)

	return exporterhelper.NewMetrics(
		ctx,
		set,
		cfg,
		exp.pushMetricsData,
		exporterhelper.WithStart(exp.start),
		exporterhelper.WithShutdown(exp.shutdown),
		exporterhelper.WithTimeout(c.TimeoutSettings),
		exporterhelper.WithQueue(c.QueueSettings),
		exporterhelper.WithRetry(c.BackOffConfig),
	)
}
