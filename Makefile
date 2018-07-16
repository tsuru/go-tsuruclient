SWAGGER_CODEGEN=swagger-codegen
generate:
	$(SWAGGER_CODEGEN) generate -t ./templates -i ../tsuru/docs/reference/api.yaml -l go -c gen-config.json
	rm -rf pkg/tsuru && mkdir -p pkg/tsuru
	mv *.go pkg/tsuru
	gofmt -s -w pkg/tsuru/
	goimports -w pkg/tsuru/
