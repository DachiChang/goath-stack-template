.DEFAULT_GOAL := up
.PHONY:	init up down log build docker-image docker-run run clean gen tailwind templ doc

TEMPLATE_FROM := github.com/dachichang/goath-stack-template
MODULE_NAME := $(shell go list -m)
UNAME := $(shell uname -s)

ifeq ($(UNAME),Darwin)
	SED_INPLACE := -i ''
else
	SED_INPLACE := -i
endif

# apply all necessary modifications to the project.
init:
	@echo "Replacing $(TEMPLATE_FROM) -> $(MODULE_NAME) in necessary files."
	@grep -rl $(TEMPLATE_FROM) --include="*.templ" | xargs sed $(SED_INPLACE) -e "s|$(TEMPLATE_FROM)|$(MODULE_NAME)|g"
	@echo "Create .env from env.example"
	@cp env.example .env

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

# run in docker
docker-image:
	docker build -t $(MODULE_NAME):latest -f docker/Dockerfile .

docker-run:
	docker run --rm -it --env-file .env -p 3000:3000 $(MODULE_NAME)

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
