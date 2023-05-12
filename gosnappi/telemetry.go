package gosnappi

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type telemetry struct {
	transport     string
	endpoint      string
	serviceName   string
	rootCtx       context.Context
	spanProcessor sdktrace.SpanProcessor
	traceProvider *sdktrace.TracerProvider
}

type Telemetry interface {
	isOTLPEnabled() bool
	getRootContext() context.Context
	WithHTTP() Telemetry
	WithGRPC() Telemetry
	WithExporterEndPoint(endpoint string) Telemetry
	WithRootContext(ctx context.Context) Telemetry
	WithCustomSpanProcess(spanProcessor sdktrace.SpanProcessor) Telemetry
	WithServiceName(serviceName string) Telemetry
	Start() (Telemetry, error)
	Stop()
	NewSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span)
	SetSpanStatus(span trace.Span, code codes.Code, description string)
	SetSpanAttributes(span trace.Span, attrs []attribute.KeyValue)
	SetSpanEvent(span trace.Span, eventStr string)
	CloseSpan(span trace.Span)
	View()
}

var tracer = otel.Tracer("gosnappi-tracer")

// just a debug function to view the telemetry struct
// TODO: Remove this.
func (t *telemetry) View() {
	fmt.Println("fetching the tracer")
	fmt.Printf("tel is %v\n", t)
	fmt.Println(t.transport)
}

// Internal fucntion to check wheather telemetry is enabled or not.
// Used by rest of the functions to become no-ops.
func (t *telemetry) isOTLPEnabled() bool {
	return t.endpoint != ""
}

// Internal function to fetch the root context if provided by user.
// Default context is background
func (t *telemetry) getRootContext() context.Context {
	if t.rootCtx != nil {
		return t.rootCtx
	}

	return context.Background()
}

// Sets the transport to HTTP which uses REST communication
func (t *telemetry) WithHTTP() Telemetry {
	t.transport = "HTTP"
	return t
}

// Sets the transport to GRPC
func (t *telemetry) WithGRPC() Telemetry {
	t.transport = "GRPC"
	return t
}

// Sets The endpoint of the OTLP collector which will collect the data and visualize it.
// For HTTP the endpoint shoule be like "IP:PORT" e.g. "127.0.0.1:4138"
func (t *telemetry) WithExporterEndPoint(endpoint string) Telemetry {
	t.endpoint = endpoint
	return t
}

// Sets The Root Context.
// If the user wants all the spans to be the child of a single span, this method should be used.
func (t *telemetry) WithRootContext(ctx context.Context) Telemetry {
	t.rootCtx = ctx
	return t
}

// Gives ability to the user to set a custom span processor for processing the spans
func (t *telemetry) WithCustomSpanProcess(spanProcessor sdktrace.SpanProcessor) Telemetry {
	t.spanProcessor = spanProcessor
	return t
}

// Gives ability to the user to set name for the OTLP service
func (t *telemetry) WithServiceName(serviceName string) Telemetry {
	t.serviceName = serviceName
	return t
}

// Initiates the trace provider with proper resources, exporter information
// and span processors
func (t *telemetry) Start() (Telemetry, error) {

	if t.isOTLPEnabled() {

		var exporter *otlptrace.Exporter
		var err error

		if t.transport == "HTTP" {

			// creating exporter which communicates with the OTLP collector
			exporter, err = otlptrace.New(
				context.Background(),
				otlptracehttp.NewClient(
					otlptracehttp.WithInsecure(),
					otlptracehttp.WithEndpoint(t.endpoint),
				),
			)

			// raising error if exporter creation had some issues
			if err != nil {
				return nil, fmt.Errorf("Error creating OTLP trace exporter: %v\n", err)
			}

		} else if t.transport == "GRPC" {

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			conn, err := grpc.DialContext(ctx, t.endpoint,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithBlock(),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
			}

			exporter, err = otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))

			// raising error if exporter creation had some issues
			if err != nil {
				return nil, fmt.Errorf("Error creating OTLP trace exporter: %v\n", err)
			}

		} else {
			return nil, fmt.Errorf("transport %s is not supported", t.transport)
		}

		// defining the service name
		resources, err := resource.New(
			context.Background(),
			resource.WithAttributes(
				attribute.String("service.name", t.serviceName),
				attribute.String("application", t.serviceName),
			),
		)

		if err != nil {
			return nil, fmt.Errorf("Error setting resources for OTLP trace: %v", err)
		}

		var spanProcessor sdktrace.SpanProcessor

		// by default we use BatchSpanProcessor
		if t.spanProcessor != nil {
			spanProcessor = t.spanProcessor
		} else {
			spanProcessor = sdktrace.NewBatchSpanProcessor(exporter)
		}

		// Creating the traceProvider
		traceProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSpanProcessor(spanProcessor),
			sdktrace.WithResource(resources),
		)

		// a TextMapPropagator is a component that injects values ​​into and extracts values ​​
		// from transport as string key/value pairs
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{}),
		)

		otel.SetTracerProvider(traceProvider)
		t.traceProvider = traceProvider

		return t, nil
	}

	return nil, nil
}

// Gracefully shuts down the trace provider and flushes the collector streams.
func (t *telemetry) Stop() {
	if t.isOTLPEnabled() {
		if err := t.traceProvider.Shutdown(context.Background()); err != nil {
			GlobalLogger.Error().Msg("Failed shutting down trace provider")
		}
		GlobalLogger.Info().Msg("shut down successful !!!!")
	}
}

// Creates a new span , if a parent context is passed then it will be child span.
// By default the span will be root span.
func (t *telemetry) NewSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if t.isOTLPEnabled() {
		return tracer.Start(ctx, name, opts...)
	}

	return nil, nil
}

// Closes a Span.
func (t *telemetry) CloseSpan(span trace.Span) {
	if t.isOTLPEnabled() {
		span.End()
	}
}

// This part is generally used to set errors and stuff with proper code and desc.
func (t *telemetry) SetSpanStatus(span trace.Span, code codes.Code, description string) {
	if t.isOTLPEnabled() {
		span.SetStatus(code, description)
	}
}

// This method is used to set attributes to a particular span.
func (t *telemetry) SetSpanAttributes(span trace.Span, attrs []attribute.KeyValue) {
	if t.isOTLPEnabled() {
		span.SetAttributes(attrs...)
	}
}

// This method is used to set event to a particular span.
func (t *telemetry) SetSpanEvent(span trace.Span, eventStr string) {
	if t.isOTLPEnabled() {
		span.AddEvent(eventStr)
	}
}
