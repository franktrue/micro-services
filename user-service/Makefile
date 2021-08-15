build:
	protoc -I. --micro_out=. --go_out=. ./proto/user/user.proto
	docker build -t user-service .

run:
	docker run -d -p 9999:9999 \ 
		-e MICRO_SERVER_ADDRESS=:9999 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=user-db:3306 \
      	-e DB_NAME=micro_shop_user \
      	-e DB_USER=test \
      	-e DB_PASSWORD=test \
		user-service