build:
	go build -o bin/neurokin-grug cmd/api/main.go

run: build
	./bin/neurokin-grug

test:
	go test -v ./...
