### gererate proto

protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		gatewaypb/protofiles/gw.proto

protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		gatewaypb/protofiles/gw.proto
		
protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		gatewaypb/protofiles/gw.proto
	
### APIs	
 go run server/main.go & # Run gRPC server
 go run main.go & # Run Gateway

##### 1. curl localhost:8080/get?name=live
{"message":"Received GET method live"}

##### 2. curl -X POST localhost:8080/post -d '{"name": "live"}'
{"message":"Received POST method live"}