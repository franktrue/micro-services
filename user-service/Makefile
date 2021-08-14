build:
	protoc -I. --micro_out=. --go_out=. ./proto/user/user.proto
	docker build -t user-service .

run:
	docker run -d -p 9999:9999 -e MICRO_SERVER_ADDRESS=:9999 -e MICRO_REGISTRY=mdns user-service