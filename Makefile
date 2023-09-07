GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
INTERNAL_PROTO_FILES=$(shell  cd app/admin && find internal -name *.proto)
API_PROTO_FILES=$(shell cd api/admin/v1/ && find . -name "*.proto")
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "go-kratos/beer-" $$0 ":0.1.0"}')

.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/envoyproxy/protoc-gen-validate@latest


.PHONY: grpc
# generate grpc code
grpc:
	 cd api/admin/v1 && protoc --proto_path=. \
           --proto_path=../../../third_party \
           --proto_path=../../../pkg/proto \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: http
# generate http code
http:
	cd api/admin/v1 && protoc --proto_path=. \
           --proto_path=../../../third_party \
           --proto_path=../../../pkg/proto \
           --go_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
           --validate_out=paths=source_relative,lang=go:. \
           --openapi_out=fq_schema_naming=true,default_response=false:../../../. \
           $(API_PROTO_FILES)

.PHONY: errors
# generate errors code
errors:
	cd api/admin/v1 && protoc --proto_path=. \
           --proto_path=../../../third_party \
           --proto_path=../../../pkg/proto \
           --go_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	cd app/admin && protoc --proto_path=./internal \
		--proto_path=../../third_party \
		--proto_path=../../pkg/proto \
		--go_out=paths=source_relative:./internal \
		$(INTERNAL_PROTO_FILES)

.PHONY: swagger
# generate swagger
swagger:
	cd api/admin/v1 && protoc --proto_path=. \
	       --proto_path=../../../third_party \
	       --proto_path=../../../pkg/proto \
	       --openapiv2_out . \
	       --openapiv2_opt logtostderr=true \
	       --openapiv2_opt allow_merge=true \
	       --openapiv2_opt merge_file_name=api.proto \
	       --openapiv2_opt json_names_for_fields=false \
	       --openapiv2_opt=openapi_configuration=swagger.config.yaml \
           $(API_PROTO_FILES)

.PHONY: proto
# generate internal proto struct
proto:
	cd app/admin && protoc --proto_path=. \
           --proto_path=../../third_party \
           --proto_path=../../pkg/proto \
           --go_out=paths=source_relative:. \
           $(INTERNAL_PROTO_FILES)

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin ./...

build_server:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/im-admin app/admin/cmd/server/main.go app/admin/cmd/server/wire_gen.go

.PHONY: test
# test
test:
	go test -v ./... -cover

.PHONY: run
run:
	cd cmd/server/ && go run .

.PHONY: docker
docker:
	cd ../../.. && docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(DOCKER_IMAGE) .

.PHONY: wire
# generate wire
wire:
	cd app/admin/cmd/server && wire

.PHONY: api
# generate api proto
api: grpc http swagger errors



.PHONY: all
# generate all
all: grpc http config proto wire generate swagger errors build test

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
