.DEFAULT_GOAL := up
.PHONY:	up down log build run clean gen tailwind templ doc

# run entire goath-stack application at :8080 (docker)
up:
	docker compose -p goath-stack -f docker/docker-compose.yaml up -d

# stop and clean up entire goath-stack application
down:
	docker compose -p goath-stack down

log:
	docker compose -p goath-stack logs -f

build:
	go build -o goath-stack ./cmd/server/

# running information is defined by .env file
run:
	npm install && go tool air

clean:
	rm -rf ./node_modules ./tmp

gen: templ doc tailwind

# generate tailwind css
tailwind:
	npm run tailwind

# generate templ go
templ:
	go tool templ generate

# build swag docs
doc:
	go tool swag init -g cmd/server/main.go

