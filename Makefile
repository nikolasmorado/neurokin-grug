# Paths
OUTPUT_DIR = bin
CSS_INPUT = views/css/app.css
CSS_OUTPUT = public/styles.css

# Commands
build: templ css
	@go build -o $(OUTPUT_DIR)/app .

run: build
	@./$(OUTPUT_DIR)/app

templ:
	~/go/bin/templ generate

css:
	@tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT) --minify

clean:
	@rm -rf $(OUTPUT_DIR)/* $(CSS_OUTPUT)

test:
	@go test -v ./...

# New target for compiling all at once
all: clean templ css build

# Tailwind watch mode without rebuild
watch-css:
	@tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT) --minify --watch

