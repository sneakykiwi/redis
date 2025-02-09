package main

import (
	"context"

	"github.com/sneakykiwi/redis/extra/redisotel/v8"
	"github.com/sneakykiwi/redis/v8"
	"go.opentelemetry.io/otel"
	stdoutexporter "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("app_or_package_name")

func main() {
	exporter, err := stdoutexporter.New(stdoutexporter.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	provider := sdktrace.NewTracerProvider()
	provider.RegisterSpanProcessor(sdktrace.NewSimpleSpanProcessor(exporter))

	otel.SetTracerProvider(provider)

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	defer rdb.Close()

	rdb.AddHook(redisotel.NewTracingHook())

	ctx, span := tracer.Start(context.TODO(), "main")
	defer span.End()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}
