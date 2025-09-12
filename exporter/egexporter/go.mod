module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/egexporter

go 1.24

require (
	github.com/ClickHouse/clickhouse-go/v2 v2.40.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.133.0
	github.com/testcontainers/testcontainers-go v0.38.0
    go.opentelemetry.io/collector/component v1.39.0
    go.opentelemetry.io/collector/component/componenttest v0.133.0
    go.opentelemetry.io/collector/config/configopaque v1.39.0
    go.opentelemetry.io/collector/config/configretry v1.39.0
    go.opentelemetry.io/collector/confmap v1.39.0
    go.opentelemetry.io/collector/confmap/xconfmap v0.133.0
    go.opentelemetry.io/collector/exporter v0.133.0
    go.opentelemetry.io/collector/exporter/exportertest v0.133.0
    go.opentelemetry.io/collector/featuregate v1.39.0
    go.opentelemetry.io/collector/pdata v1.39.0
    go.opentelemetry.io/otel v1.37.0
    go.uber.org/goleak v1.3.0
    go.uber.org/zap v1.27.0
)