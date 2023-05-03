Se necesita instalar protoc

Se necesita validar paquetes como:

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/glog v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)

$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

Se rquieren unos archivos de google/api/:
	annotations.proto
	http.proto


protoc -I . --go_out ./ --go-grpc_out ./ --grpc-gateway_out ./ --go-grpc_opt paths=source_relative --grpc-gateway_opt paths=source_relative ./proto/pruebas.proto



https://medium.com/@busranurguner/create-simple-grpc-gateway-in-go-d5a49d8466e1