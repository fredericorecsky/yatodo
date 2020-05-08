
APP = yatodo

BUILD_DIR ?= build

DOCKER_IMAGE ?= "fredericorecsky/$(APP)"

E2E_TEST_DIR ?= test/e2e

VERSION ?= $(shell cat ./version)

.PHONY: clean

default: build

dep:
	go get -d -v ./...

build: clean
	go build -v -o $(BUILD_DIR)/$(APP)

build-docker:
	docker build -t $(DOCKER_IMAGE):$(VERSION) .

clean:
	rm -rf $(BUILD_DIR) > /dev/null

clean-devdb:
	rm test.db db/test.db

run:
	go run -v $(APP).go

run-docker:
	 docker run -e DB_BACKEND=$(DB_BACKEND) -e DB_DSN="$(DB_DSN)" -p 9000:9000 -it --rm --name yadayada $(DOCKER_IMAGE):$(VERSION)

test-curl:
	bash test/curl.sh

test-db:
	go test -v github.com/fredericorecsky/yatodo/db