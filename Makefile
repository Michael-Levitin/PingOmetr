first:
	go mod tidy

server_start:
	cd ./cmd/server/ && go run main.go
# wait for "server is running"

client1_start:
	cd ./cmd/client1/ && go run main.go
client2_start:
	cd ./cmd/client2/ && go run main.go
client3_start:
	cd ./cmd/client3/ && go run main.go
admin_start:
	cd ./cmd/admin/ && go run main.go

tests:
	go test ./...

tests2:
	go test -v ./...