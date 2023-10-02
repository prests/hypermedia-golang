run:
	make css
	air

build:
	make css
	go build -o ./build/app ./cmd/main.go

css:
	./tailwindcss -i internal/input.css -o web/assets/output.css --minify

css-watch:
	./tailwindcss -i internal/input.css -o web/assets/output.css --watch
