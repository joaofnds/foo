package config

func JaegerCollectorEndpoint() string {
	return GetString("jaeger.reporter.collectorEndpoint")
}
