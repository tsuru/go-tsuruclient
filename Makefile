generate:
	swagger-codegen generate -i ../tsuru/docs/reference/api.yaml -l go -c gen-config.json
	rm -rf pkg/tsuru && mkdir -p pkg/tsuru
	mv *.go pkg/tsuru
	cp real.travis.yml .travis.yml
	gofmt -s -w pkg/tsuru/
