This sample repository contains two sample wasm projects, one using Emscripten and the other using Go.

## Emscripten
To compile `square.c` execute `emcc square.c -O3 -s WASM=1 -s EXTRA_EXPORTED_RUNTIME_METHODS='["cwrap"]' -o square.js`.

> Note, you need to have [Emscripten](http://emscripten.org) [setup and configured](https://emscripten.org/docs/getting_started/downloads.html).

After the compilation you should have a `square.js` and a `square.wasm` file. Launch `index.html` in the browser to see it in action.

## Go
To compile `cube.go` execute `GOOS=js GOARCH=wasm go build -o cube.wasm`

> Note, you need to have Go [setup and configured](https://golang.org/doc/install)

After compilation you should have a `cube.go` file available. Lanuch `index-go.html` in your browser to see it in action.