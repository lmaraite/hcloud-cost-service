build: 
	go build -o bin/hcloud-cost-service

run: 
	go run .

test:
	go test -v -coverpkg=./... ./...

coverage:
	go test -v -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
