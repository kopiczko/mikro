HOST_GOOS=$(shell go env GOOS)
HOST_GOARCH=$(shell go env GOARCH)

test: vendor
	go test -cover -race $(shell ./tools/glide novendor)

vet: tools/glide
	go vet $(shell ./tools/glide novendor)

fmt: tools/glide
	go fmt $(shell ./tools/glide novendor)

vendor: tools/glide
	./tools/glide install

protoc:
	protoc -I./dbaccessor/dbaccessorpb --go_out=plugins=micro:dbaccessor/dbaccessorpb dbaccessor/dbaccessorpb/dbaccessor.proto

tools/glide:
	@echo "Downloading glide"
	mkdir -p tools
	curl -L https://github.com/Masterminds/glide/releases/download/0.10.2/glide-0.10.2-$(HOST_GOOS)-$(HOST_GOARCH).tar.gz | tar -xz -C tools
	mv tools/$(HOST_GOOS)-$(HOST_GOARCH)/glide tools/glide
	rm -r tools/$(HOST_GOOS)-$(HOST_GOARCH)
