package main

import (
	"fmt"

	"github.com/IsabelF98/CMSC701_PA2/pkgbitvectors"
	"github.com/dropbox/godropbox/container/bitvector"
)

func main() {
	data := "bananasareyummy"
	n := 100
	bv := bitvector.NewBitVector([]byte(data), n)

	// for i := 0; i < n; i++ {
	// 	fmt.Println(bv.Element(i))
	// }
	// fmt.Println(" ")

	r := pkgbitvectors.RankSupport(bv)
	fmt.Println(r.Rank(43))

}
