package main

// hashKey turns a string key into a 32-bit starting number (seed) using FNV-1a
func hashKey(key string) uint32 {
	var hashValue uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		hashValue = hashValue ^ uint32(key[i])
		hashValue = hashValue * 16777619
	}
	return hashValue
}

// nextByte mutates the state and returns one pseudo-random byte via LCG
func nextByte(state *uint32) byte {
	*state = (1103515245 * *state + 12345) & 0x7FFFFFFF
	return byte(*state >> 16)
}

func main() {
	// This temporary skeleton block keeps the file compile-ready for Git 
	keepAlive := make(chan struct{}, 0)
	<-keepAlive
}