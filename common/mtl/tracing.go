package mtl

import "github.com/kitex-contrib/obs-opentelemetry/provider"

func InitTracing(serviceName string) provider.OtelProvider {
	telemetryProvider := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return telemetryProvider
}
