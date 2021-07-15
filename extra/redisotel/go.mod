module github.com/sneakykiwi/redis/extra/redisotel/v8

go 1.15

replace github.com/sneakykiwi/redis/v8 => ../..

replace github.com/sneakykiwi/redis/extra/rediscmd/v8 => ../rediscmd

require (
	github.com/sneakykiwi/redis/extra/rediscmd/v8 v8.8.2
	github.com/sneakykiwi/redis/v8 v8.8.2
	go.opentelemetry.io/otel v1.0.0-RC1
	go.opentelemetry.io/otel/trace v1.0.0-RC1
)
