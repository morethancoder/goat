# Original commands
run:
	go run main.go

compile: tailwindcss templ run

vite:
	bun run vite

watch:
	ls ./components/*.templ | entr -r make compile

test:
	go test -v ./... -count=1 

tailwindcss:
	bun run tailwindcss --config tailwind.config.js -i static/css/input.css -o styles.css

templ:
	~/go/bin/templ generate ./components

.PHONY: run compile vite watch test tailwindcss templ
