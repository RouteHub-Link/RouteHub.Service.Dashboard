lint:
	revive -config ./linter/revive.toml -formatter friendly ./...

watch:
	templ generate --watch --proxy="http://localhost:8080" --cmd="make run"

tailwind-init:
	npx tailwindcss init

tailwind:
	tailwindcss -i web/public/css/input.css -o web/public/css/output.css

build: 
	@echo "Building..."
	@templ generate
	@make tailwind
	@go build -o main .

run:
	@echo "Running..."
	@make tailwind
	@go run .