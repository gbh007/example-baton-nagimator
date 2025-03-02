JSONNET = github.com/google/go-jsonnet/cmd/jsonnet@latest
JSONNET_JB = github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb@latest

JSONNET_EXT_VARS = 

.PHONY: install
install:
	go run $(JSONNET_JB) install github.com/grafana/grafonnet/gen/grafonnet-latest@main

.PHONY: build
build: install
	go run $(JSONNET) $(JSONNET_EXT_VARS) -J vendor dashboard.jsonnet > ../deployment/dashboards/example-baton-nagimator.json