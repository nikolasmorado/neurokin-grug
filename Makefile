build:
	@go build -o bin/app .

run: build
	@./bin/app

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch

clean:
	rm -rf bin/*

test:
	go test -v ./...
