build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/johnwoz123/payrock-mock-api-service/user-auth-service \
		proto/user/user.proto
	docker build -t user-auth-service .

run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-auth-service
