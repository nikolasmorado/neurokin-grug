build:
	go build -o bin/neurokin-grug ./main.go

run: build
	./bin/neurokin-grug

test:
	go test -v ./...
