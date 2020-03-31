#include "emscripten.h"

EMSCRIPTEN_KEEPALIVE

int int_sqrt(int x) {
  return x * x;
}
// to create wasm, run: emcc sqrt.c -O3 -s WASM=1 -s EXTRA_EXPORTED_RUNTIME_METHODS='["cwrap"]' -o sqrt.js