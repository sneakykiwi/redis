module github.com/sneakykiwi/redis/extra/rediscensus/v8

go 1.15

replace github.com/sneakykiwi/redis/extra/rediscmd/v8 => ../rediscmd

require (
	github.com/sneakykiwi/redis/extra/rediscmd/v8 v8.8.2
	github.com/sneakykiwi/redis/v8 v8.8.2
	go.opencensus.io v0.22.6
)
