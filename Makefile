.PHONY: generate-cmd process-files generate docker-generate-cmd docker-generate

OPENAPI_CODEGEN = openapi-generator

generate: generate-cmd process-files

generate-cmd:
	$(OPENAPI_CODEGEN) generate -t ./templates -i ../tsuru/docs/reference/api.yaml -g go -c gen-config.json

process-files:
	rm -rf pkg/tsuru && mkdir -p pkg/tsuru
	mv *.go pkg/tsuru
	gofmt -s -w pkg/tsuru/
	goimports -w pkg/tsuru/

docker-generate-cmd:
	cp ../tsuru/docs/reference/api.yaml .
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/app -w /app --entrypoint java swaggerapi/swagger-codegen-cli:2.4.0 -jar /opt/swagger-codegen-cli/swagger-codegen-cli.jar generate -t ./templates -i api.yaml -l go -c gen-config.json
	rm api.yaml

docker-generate: docker-generate-cmd process-files
