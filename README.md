# mikro

Example microservice application build on top of [go-micro][1].

## Requirements

* Go 1.7
* make

## Running Examples

Examples are running with usage of mDNS registry to avoid setting up another discovery service.

1. Build the project with `make build`.
2. Run services with [goreman][2] (or another tool supporting [Procfile][3]'s), e.g. `goreman start`.
3. From another terminal run the example command `go run examples/example.go --registry mdns`.

[1]: https://github.com/micro/go-micro
[2]: https://github.com/mattn/goreman
[3]: https://devcenter.heroku.com/articles/procfile
