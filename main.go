package main

import (
	"encoding/hex"
	"syscall/js"
)

// hash turns a string key into a 32-bit starting number (seed)
// h *= 16777619
func hashKey(key string) uint32 {
	var hashKey uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		hashKey = hashKey ^ uint32(key[i])
		hashKey = hashKey * 16777619
	}
	return hashKey
}

// nextByte mutates the state and returns one pseudo-random byte
func nextByte(state *uint32) byte {
	*state = (1103515245**state + 12345) & 0x7FFFFFFF
	return byte(*state >> 16)
}

// encrypt is the main function exposed to JavaScript
// encrypt handles both real-time encryption and decryption based on a mode flag
func encrypt(this js.Value, args []js.Value) any {
	if len(args) < 3 {
		return "Error: Missing arguments"
	}

	msg := args[0].String()
	key := args[1].String()
	mode := args[2].String() // "encrypt" or "decrypt"

	if msg == "" {
		return ""
	}
	if key == "" {
		return "ERROR : Enter a Key"
	}

	// 1. Initialize our same mathematical state using the key hash
	state := hash(key)
	var src []byte
	var err error

	// 2. Prepare the input bytes depending on the operation mode
	if mode == "decrypt" {
		// If decrypting, the input is a hex string. We must turn it back to raw bytes.
		src, err = hex.DecodeString(msg)
		if err != nil {
			return "ERROR : Invalid Hexadecimal Ciphertext"
		}
	} else {
		// If encrypting, just convert raw text directly to bytes
		src = []byte(msg)
	}

	out := make([]byte, len(src))

	// 3. The exact same symmetric XOR loop runs for both actions!
	for i := 0; i < len(src); i++ {
		b := nextByte(&state)
		out[i] = src[i] ^ b
	}

	// 4. Return formatted data back to the webpage
	if mode == "decrypt" {
		return string(out) // Return readable human text string
	}
	return hex.EncodeToString(out) // Return unreadable hex payload string
}

func main() {
	keepAlive := make(chan struct{}, 0)
	js.Global().Set("encryptTextNative", js.FuncOf(encrypt))
	<-keepAlive
}
