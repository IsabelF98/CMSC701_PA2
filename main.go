package main

import (
	"github.com/IsabelF98/CMSC701_PA2/pkgbitvectors"
	"github.com/dropbox/godropbox/container/bitvector"
)

func main() {
	//data := "bananasareyummy"
	data := "BitchBetterHaveMyMoneyPayMeWhatYouWantRihannaIsAGod"
	n := 300
	bv := bitvector.NewBitVector([]byte(data), n)

	// for i := 0; i < n; i++ {
	// 	fmt.Println(bv.Element(i))
	// }

	r := pkgbitvectors.RankSupport(bv)
	//fmt.Println(r.Rank(23))
	pkgbitvectors.SelectSuport(r)
}
