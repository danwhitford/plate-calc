ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

main.wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

wasm_exec.js:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" .

.PHONY site: main.wasm wasm_exec.js
