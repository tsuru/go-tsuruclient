.PHONY: generate-cmd process-files generate docker-generate-cmd docker-generate pre-generate post-generate

GENERATE_ARGS = generate -t ./templates -i ./api.yaml -g go -c gen-config.json
GENERATOR_VERSION = v3.3.4

generate-dev: pre-generate-dev docker-generate-cmd-dev process-files post-generate
generate: pre-generate-prod docker-generate-cmd process-files post-generate

process-files:
	rm -rf pkg/tsuru && mkdir -p pkg/tsuru
	mv *.go pkg/tsuru
	gofmt -s -w pkg/tsuru/
	goimports -w pkg/tsuru/

docker-generate-cmd:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/app -w /app openapitools/openapi-generator-cli:$(GENERATOR_VERSION) $(GENERATE_ARGS)

docker-generate: pre-generate docker-generate-cmd process-files post-generate

pre-generate-dev:
	cp ../tsuru/docs/reference/api.yaml .

pre-generate-prod:
	curl -O https://raw.githubusercontent.com/tsuru/tsuru/master/docs/reference/api.yaml

post-generate:
	rm api.yaml
