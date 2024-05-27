package eginnovations

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type egExporter struct {
	config        *Config
	traceExporter ptraceotlp.GRPCClient
	clientConn    *grpc.ClientConn
	callOptions   []grpc.CallOption
	url           string
	settings      component.TelemetrySettings
	userAgent     string
}

func (e *egExporter) Start(ctx context.Context, host component.Host) (err error) {

	opts := e.configureDialOpts()
	tls := e.config.TLSSetting.Insecure
	if tls {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if e.clientConn, err = e.config.ToClientConn(ctx, host, e.settings, opts...); err != nil {
			return err
		}
	} else {
		if e.clientConn, err = e.config.ToClientConn(ctx, host, e.settings, opts...); err != nil {
			return err
		}
	}

	e.traceExporter = ptraceotlp.NewGRPCClient(e.clientConn)
	if e.config.Headers == nil {
		e.config.Headers = make(map[string]configopaque.String)
	}
	e.callOptions = []grpc.CallOption{
		grpc.WaitForReady(e.config.WaitForReady),
	}

	return nil
}

func NewEgExporter(cfg component.Config, set exporter.CreateSettings) *egExporter {
	iCfg := cfg.(*Config)
	userAgent := fmt.Sprintf("%s/%s (%s/%s)", set.BuildInfo.Description, set.BuildInfo.Version, runtime.GOOS, runtime.GOARCH)
	return &egExporter{
		config:    iCfg,
		url:       iCfg.Endpoint,
		settings:  set.TelemetrySettings,
		userAgent: userAgent,
	}
}

var _ consumer.ConsumeTracesFunc = (*egExporter)(nil).ConsumeTraces

func (e *egExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	rs := td.ResourceSpans()
	for i := 0; i < rs.Len(); i++ {
		rs := rs.At(i)
		ils := rs.ScopeSpans()
		for j := 0; j < ils.Len(); j++ {
			ils := ils.At(j)
			spans := ils.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				fmt.Printf("Span %s:\n", span.Name())
				fmt.Printf("  TraceID: %s\n", span.TraceID())
				fmt.Printf("  SpanID: %s\n", span.SpanID())
				fmt.Printf("  StartTime: %s\n", time.Unix(0, int64(span.StartTimestamp())).UTC().Format(time.RFC3339Nano))
				fmt.Printf("  EndTime: %s\n", time.Unix(0, int64(span.EndTimestamp())).UTC().Format(time.RFC3339Nano))
			}
		}
	}

	_, err := e.traceExporter.Export(e.outgoingContext(ctx), ptraceotlp.NewExportRequestFromTraces(td), e.callOptions...)
	if err != nil {
		return err
	}

	return nil
}

func (e *egExporter) outgoingContext(ctx context.Context) context.Context {
	md := metadata.New(nil)
	for k, v := range e.config.Headers {
		md.Set(k, string(v))
	}
	return metadata.NewOutgoingContext(ctx, md)
}

func (e *egExporter) Shutdown(context.Context) error {
	if e.clientConn == nil {
		return nil
	}
	return e.clientConn.Close()
}

type loginCreds struct {
	Userid string
	token  configopaque.String
}

func (c *loginCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"userid": c.Userid,
		"token":  string(c.token),
	}, nil
}

func (c *loginCreds) RequireTransportSecurity() bool {
	return true
}

func (e *egExporter) configureDialOpts() []grpc.DialOption {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithPerRPCCredentials(&loginCreds{
		Userid: e.config.UserID,
		token:  e.config.Token,
	}))
	opts = append(opts, grpc.WithUserAgent(e.userAgent))
	opts = append(opts, grpc.WithTimeout(15*time.Second))
	opts = append(opts, grpc.WithUserAgent(e.userAgent))
	return opts
}
