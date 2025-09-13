generate-swagger:
	swag init -g cmd/app/main.go -o docs/api/
	
run:
	docker build -t bookingtogo-app . && docker image prune -f && docker run -it --rm \
  	--network app-network \
  	-p 8334:8334 \
  	-v $(pwd)/config.json:/config.json \
  	bookingtogo-app


build:
	go build -o app cmd/app/main.go