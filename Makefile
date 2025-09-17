run:
	docker build -t otp-app . && docker image prune -f && docker run -it --rm \
  	--network app-network \
  	-p 8335:8335 \
  	-v $(pwd)/config.json:/config.json \
  	otp-app


build:
	go build -o app cmd/app/main.go

test: 
	go test -v -cover -coverprofile=coverage.out ./internal/usecases/otp/... && go tool cover -html=coverage.out -o coverage.html