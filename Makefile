build: 
	go build -o bin/hcloud-cost-service

run: 
	go run .

test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
