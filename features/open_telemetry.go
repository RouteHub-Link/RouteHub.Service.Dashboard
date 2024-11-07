package features

import (
	"context"
	"log"
	"log/slog"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

type OpenTelemetryService struct {
	config OpenTelemetryConfig
	Tracer *sdktrace.TracerProvider
}

func NewOpenTelemetryService(options ...OpenTelemetryOption) *OpenTelemetryService {
	oct := &OpenTelemetryConfig{}
	for _, option := range options {
		option(oct)
	}

	ots := &OpenTelemetryService{
		config: *oct,
	}

	return ots
}

type OpenTelemetryConfig struct {
	ServiceName  string
	CollectorURL string
	Insecure     string
	Logger       *slog.Logger
	Headers      map[string]string // Add this field to store headers
}

func (c *OpenTelemetryConfig) SetServiceName(serviceName string) {
	c.ServiceName = serviceName
}

func (c *OpenTelemetryConfig) SetCollectorURL(collectorURL string) {
	c.CollectorURL = collectorURL
}

func (c *OpenTelemetryConfig) SetInsecure(insecure string) {
	c.Insecure = insecure
}

func (c *OpenTelemetryConfig) SetHeaders(headers string) {
	//  Headers: uptrace-dsn=http://project_home_secret_token@localhost:14318?grpc=14317
	rawHeaders := strings.Split(headers, ";")
	c.Headers = make(map[string]string)
	for _, rawHeader := range rawHeaders {
		header := strings.Split(rawHeader, "=")
		c.Headers[header[0]] = strings.Join(header[1:], "")
	}

}

// Options

type OpenTelemetryOption func(*OpenTelemetryConfig)

func WithServiceName(serviceName string) OpenTelemetryOption {
	return func(c *OpenTelemetryConfig) {
		c.SetServiceName(serviceName)
	}
}

func WithCollectorURL(collectorURL string) OpenTelemetryOption {
	return func(c *OpenTelemetryConfig) {
		c.SetCollectorURL(collectorURL)
	}
}

func WithInsecure(insecure string) OpenTelemetryOption {
	return func(c *OpenTelemetryConfig) {
		c.SetInsecure(insecure)
	}
}

func WithLogger(logger *slog.Logger) OpenTelemetryOption {
	return func(c *OpenTelemetryConfig) {
		c.Logger = logger
	}
}

func WithHeaders(headers string) OpenTelemetryOption {
	return func(c *OpenTelemetryConfig) {
		c.SetHeaders(headers)
	}
}

// Tracer
func (ots OpenTelemetryService) InitTracer() func(context.Context) error {

	var secureOption otlptracegrpc.Option
	otc := ots.config

	if strings.ToLower(otc.Insecure) == "false" || otc.Insecure == "0" || strings.ToLower(otc.Insecure) == "f" {
		secureOption = otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	} else {
		secureOption = otlptracegrpc.WithInsecure()
	}

	// Convert headers map to otlptracegrpc.Option
	var headersOption otlptracegrpc.Option
	if len(otc.Headers) > 0 {
		headersOption = otlptracegrpc.WithHeaders(ots.config.Headers)
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(otc.CollectorURL),
			headersOption,
		),
	)

	if err != nil {
		log.Fatalf("Failed to create exporter: %v", err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", otc.ServiceName),
			attribute.String("library.language", "go"),
			attribute.String("library.version", "0.1.0"),
			attribute.String("deployment.environment", os.Getenv("ENVIRONMENT")),
			attribute.String("host.name", os.Getenv("HOSTNAME")),
		),
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithTelemetrySDK(),
	)
	if err != nil {
		log.Fatalf("Could not set resources: %v", err)
	}

	ots.Tracer = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)

	otel.SetTracerProvider(
		ots.Tracer,
	)

	log.Println("Tracer initialized")
	return exporter.Shutdown
}
