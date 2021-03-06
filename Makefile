HOST_GOOS=$(shell go env GOOS)
HOST_GOARCH=$(shell go env GOARCH)

build: vendor
	go build -o ./bin/app ./app/cmd/app
	go build -o ./bin/auth ./auth/cmd/auth
	go build -o ./bin/dbaccessor ./dbaccessor/cmd/dbaccessor

test: vendor
	go test -cover -race $(shell ./tools/glide novendor)

vet: tools/glide
	go vet $(shell ./tools/glide novendor)

fmt: tools/glide
	go fmt $(shell ./tools/glide novendor)

vendor: tools/glide
	./tools/glide install

protoc:
	protoc -I./app/apppb --go_out=plugins=micro:app/apppb app/apppb/app.proto
	protoc -I./auth/authpb --go_out=plugins=micro:auth/authpb auth/authpb/auth.proto
	protoc -I./dbaccessor/dbaccessorpb --go_out=plugins=micro:dbaccessor/dbaccessorpb dbaccessor/dbaccessorpb/dbaccessor.proto

tools/glide:
	@echo "Downloading glide"
	mkdir -p tools
	curl -L https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-$(HOST_GOOS)-$(HOST_GOARCH).tar.gz | tar -xz -C tools
	mv tools/$(HOST_GOOS)-$(HOST_GOARCH)/glide tools/glide
	rm -r tools/$(HOST_GOOS)-$(HOST_GOARCH)
