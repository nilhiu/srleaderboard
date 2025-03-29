.PHONY: run templ build-css docker-build get-ip build restart

run:
	@echo "Starting all services..."
	@docker compose up --detach --build

templ:
	@echo "Generating Go templates..."
	@templ generate

build-css:
	@echo "Building Tailwind CSS..."
	@npx @tailwindcss/cli -i internal/app/static/css/tailwind.css -o internal/app/static/css/tailwind_gen.css

docker-build:
	@echo "Rebuilding srleaderboard container..."
	@docker compose up --detach --build srleaderboard

restart: templ build-css docker-build
