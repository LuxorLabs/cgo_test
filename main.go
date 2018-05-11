package main

import (
	"fmt"

	"github.com/LuxorLabs/cgo_test/cgo/cnh_hashing"
)

func main() {
	var bytes []byte
	cnh_hashing.Hash(bytes)
	fmt.Printf("Hello, world.\n")
}
