build:
	@go build -o bin/app .

run: build
	@./bin/app

css:
	tailwindcss -i views/css/app.css -o public/styles.css --minify --watch

clean:
	rm -rf bin/* public/styles.css

test:
	go test -v ./...
