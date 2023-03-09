package main

import (
	"fmt"

	"github.com/dropbox/godropbox/container/bitvector"
)

func main() {
	data := "banana"
	n := 15
	bv := bitvector.NewBitVector([]byte(data), n)
	for i := 0; i < n; i++ {
		fmt.Println(bv.Element(i))
	}
}
