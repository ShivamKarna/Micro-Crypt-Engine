# Micro-Crypt Engine

A fun, experimental Go-to-Wasm stream cipher utilizing custom FNV-1a and LCG algorithms.
<br>

## Go try it out at: [**https://microcrypt.pages.dev**](https://microcrypt.pages.dev)

<br>

> ⚠️ **Disclaimer:** This project was built for educational purposes. I wanted to explore encryption and decryption and code a mini version from scratch to better understand stream ciphers and WebAssembly integration.

---

## Core Mechanics

The engine implements foundational computer science concepts to achieve its streaming properties without external dependencies:

- **FNV-1a Hashing Engine:** Processes the input passphrase string byte-by-byte into a stable, pseudo-random 32-bit integer seed via bitwise XOR operations and prime integer multiplications.
- **Linear Congruential Generator (LCG):** Uses standard glibc tracking constants to mutate the initial seed state and shift bits to isolate high-entropy middle bytes.
- **Symmetric XOR Slicing:** Runs a continuous bitwise XOR stream cipher loop where identical operations handle both encryption and decryption execution blocks.

## Local Compilation and Running

To compile this engine down to WebAssembly and run it locally, execute the following commands in your terminal:

### 1. Compile the Go Engine to WebAssembly

Target the browser architecture explicitly by setting the environmental build variables before running the compiler:

```bash
GOOS=js GOARCH=wasm go build -o web/crypto.wasm main.go
```

## 2. Run a Local Web Server

Because WebAssembly requires strict CORS headers to load files securely via JavaScript, you must serve the directory through a web server rather than opening the HTML file directly:

```bash
# Using Bun
bunx serve web

# Using Python
python3 -m http.server --directory web
```
