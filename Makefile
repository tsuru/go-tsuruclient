.PHONY: generate-cmd process-files generate docker-generate-cmd docker-generate pre-generate post-generate

DOCKER ?= docker

GENERATE_ARGS = generate -t ./templates -i ./api.yaml -g go -c gen-config.json
GENERATOR_VERSION = v3.3.4

uname_m := $(shell uname -m)
ifeq ($(uname_m),arm64)
  IMAGE ?= tsuru/openapi-generator-cli
else
  IMAGE ?= openapitools/openapi-generator-cli
endif

generate-dev: pre-generate-dev docker-generate-cmd process-files post-generate
generate: pre-generate-prod docker-generate-cmd process-files post-generate

process-files:
	find pkg/tsuru/ -name "*.go" ! -name 'custom_*' | xargs rm
	mv *.go pkg/tsuru
	gofmt -s -w pkg/tsuru/
	git apply patchs/01.diff
	git apply patchs/02.diff
	goimports -w pkg/tsuru/
	go mod tidy

docker-generate-cmd:
	$(DOCKER) run -it --rm -u `id -u`:`id -g` -v `pwd`:/app -w /app $(IMAGE):$(GENERATOR_VERSION) $(GENERATE_ARGS)

docker-generate: pre-generate docker-generate-cmd process-files post-generate

pre-generate-dev: clean-docs
	cp ../tsuru/docs/reference/api.yaml .

pre-generate-prod: clean-docs
	curl -O https://raw.githubusercontent.com/tsuru/tsuru/main/docs/reference/api.yaml

post-generate:
	rm api.yaml

clean-docs:
	rm -f docs/*.md
