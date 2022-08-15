ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

main.wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

wasm_exec.js:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" .

go_js_wasm_exec:
	cp "$(GOROOT)/misc/wasm/go_js_wasm_exec" .

wasm_exec_node.js:
	cp "$(GOROOT)/misc/wasm/wasm_exec_node.js" .

.PHONY site: main.wasm wasm_exec.js

.PHONY test: go_js_wasm_exec wasm_exec_node.js wasm_exec.js
	PATH="$$PATH:$(shell pwd)" GOOS=js GOARCH=wasm go test -v ./...
