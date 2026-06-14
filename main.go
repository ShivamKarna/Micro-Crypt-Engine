package main

import (
	"encoding/hex"
	"syscall/js"
)

// generateInitialSeed turns a string key into a 32-bit starting number (seed) using FNV-1a
func generateInitialSeed(key string) uint32 {
	var hashValue uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		hashValue = hashValue ^ uint32(key[i])
		hashValue = hashValue * 16777619
	}
	return hashValue
}

// nextByte mutates the state and returns one pseudo-random byte via LCG
func generatePseudoRandomByte(initialSeed *uint32)byte{
	const linearMultiplier uint32 = 1103515245
	const linearIncrement uint32  = 12345
	const bitmaskThirtyOneBits uint32  = 0x7FFFFFFF

	*initialSeed = (linearMultiplier * *initialSeed + linearIncrement) &  bitmaskThirtyOneBits

	middleNoisyByte := byte(*initialSeed >> 16)
	return middleNoisyByte
}

func performCrypto(input, key, mode string)string{
	var buffer []byte
	var err error

	if mode == "decrypt"{
		buffer,err = hex.DecodeString(input)
		if err !=nil{
			return "ERROR : Invalid HexaDecimal Ciphertext"
		}
	}else{
		buffer = []byte(input)
	}

	structuralSeed := generateInitialSeed(key)

	for i := 0; i<len(buffer); i++{
		buffer[i] ^= generatePseudoRandomByte(&structuralSeed)
	}

	if mode == "decrypt"{
		/* if mode is decrypt then just return the decrypted message from
		bytes back to normal human readable string
		*/
		return string(buffer)
	}
	/* if mode is encrypt then just return the hex form
	of the bytes
	*/
	return hex.EncodeToString(buffer)
}

func executeCryptographicEngine(this js.Value, javascriptArguments []js.Value) any{
	// first take the inputs from the Frontend with emptyness checks
	const minimumRequiredArguments = 3

	if len(javascriptArguments)< 3{
		return "Error : Missing operational arguments from Frontend ! "
	}

	return performCrypto(javascriptArguments[0].String(),javascriptArguments[1].String(),javascriptArguments[2].String())
}

func main(){
	// Bind the function
	encFunc := js.FuncOf(executeCryptographicEngine)
	defer encFunc.Release() 

	js.Global().Set("encryptTextNative", encFunc)

	// Block forever 
	select {}
}