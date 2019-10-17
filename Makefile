.PHONY: generate-cmd process-files generate docker-generate-cmd docker-generate pre-generate post-generate

GENERATE_ARGS = generate -i ./api.yaml -g go -c gen-config.json
GENERATOR_VERSION = v4.1.3

generate: pre-generate docker-generate-cmd process-files post-generate

process-files:
	rm -rf pkg/tsuru && mkdir -p pkg/tsuru
	mv *.go pkg/tsuru
	gofmt -s -w pkg/tsuru/
	goimports -w pkg/tsuru/

docker-generate-cmd:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/app -w /app openapitools/openapi-generator-cli:$(GENERATOR_VERSION) $(GENERATE_ARGS)

docker-generate: pre-generate docker-generate-cmd process-files post-generate

pre-generate:
	cp ../tsuru/docs/reference/api.yaml .

post-generate:
	rm api.yaml
