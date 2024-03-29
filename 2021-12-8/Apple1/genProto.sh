PROTO_PATH=./auth/api
GO_OUT_PATH=./auth/api/gen/v1
mkdir -p $GO_OUT_PATH

protoc -I=. --go-grpc_out=paths=source_relative:gen/go trip.proto --go_out=paths=source_relative:gen/go trip.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go trip.proto