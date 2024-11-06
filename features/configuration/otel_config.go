package configuration

import (
	"fmt"
	"strconv"

	"github.com/caarlos0/env"
)

type OTELConfig struct {
	CollectorEndpoint string `env:"OTEL_COLLECTOR_ENDPOINT" envDefault:"localhost:4317"`
	ServiceName       string `env:"OTEL_SERVICE_NAME" envDefault:"RouteHub.Dashboard"`
	Enabled           bool   `env:"OTEL_ENABLED" envDefault:"false"`
	Insecure          bool   `env:"OTEL_INSECURE" envDefault:"true"`
}

func (o *OTELConfig) String() string {
	return fmt.Sprintf("OTELConfig; CollectorEndpoint: %s ServiceName: %s Enabled: %t Insecure: %t ", o.GetCollectorEndpoint(), o.GetServiceName(), o.IsEnabled(), o.IsInsecure())
}
func (o *OTELConfig) IsEmpty() bool {
	if o == nil {
		return true
	}

	return o.CollectorEndpoint == "" && o.ServiceName == "" && !o.Enabled && !o.Insecure
}

func (o *OTELConfig) GetCollectorEndpoint() string {
	return o.CollectorEndpoint
}

func (o *OTELConfig) GetServiceName() string {
	return o.ServiceName
}

func (o *OTELConfig) IsEnabled() bool {
	return o.Enabled
}

func (o *OTELConfig) IsInsecure() bool {
	return o.Insecure
}

func (o *OTELConfig) IsInsecureAsString() string {
	return strconv.FormatBool(o.Insecure)
}

func (o *OTELConfig) Parse(config *Config) {
	if err := env.Parse(o); err != nil {
		logger.Warn("Failed to parse otel configuration from enviroment", "error", err)
	}
	if !o.IsEmpty() {
		config.OTEL = o
	}
}
