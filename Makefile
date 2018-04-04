generate:
	swagger-codegen generate -i ../tsuru/docs/reference/api.yaml -l go -c gen-config.json
