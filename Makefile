CGO_ENABLED = 1

-include .makefiles/Makefile
-include .makefiles/pkg/go/v1/Makefile

run: artifacts/build/debug/$(GOHOSTOS)/$(GOHOSTARCH)/dogmash
	$< --load-plugin "../../dogmatiq/example/artifacts/build/debug/darwin/amd64/bank.so" $(RUN_ARGS)

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"
