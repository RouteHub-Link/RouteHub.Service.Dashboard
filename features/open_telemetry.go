package features

import (
	"context"
	"log"
	"log/slog"
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

// Tracer
func (ots OpenTelemetryService) InitTracer() func(context.Context) error {

	var secureOption otlptracegrpc.Option
	otc := ots.config

	if strings.ToLower(otc.Insecure) == "false" || otc.Insecure == "0" || strings.ToLower(otc.Insecure) == "f" {
		secureOption = otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	} else {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(otc.CollectorURL),
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
		),
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
