package main

import (
	"fmt"

	"github.com/IsabelF98/CMSC701_PA2/pkgbitvectors"
)

func main() {

	// PART 1 AND 2
	// file, err := ioutil.ReadFile("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := string(file)

	// data_size := []int{}
	// for k := 1; k < 101; k++ {
	// 	data_size = append(data_size, 100*k)
	// }

	// for j := 0; j < len(data_size); j++ { // data size
	// 	n := data_size[j]
	// 	bv := bitvector.NewBitVector([]byte(data), n) // bitvector of data

	// 	N := 50              // number of ranks tested
	// 	idx_array := []int{} // array of indexes to be tested
	// 	rank_arry := []int{} // array of ranks to be tested
	// 	for i := 0; i < N; i++ {
	// 		idx := rand.Intn(n)
	// 		idx_array = append(idx_array, idx)
	// 		rank := rand.Intn(int(float32(n) * 0.5))
	// 		rank_arry = append(rank_arry, rank)
	// 	}

	// 	// start := time.Now()
	// 	r := pkgbitvectors.RankSupport(bv) // rank structure
	// 	overhead := r.Overhead()
	// 	fmt.Println(overhead)
	// for i := 0; i < N; i++ {
	// 	r.Rank(idx_array[i]) // compute rank for each index
	// }
	// duration := time.Since(start)
	// fmt.Println(duration.Seconds())

	// start := time.Now()
	// for i := 0; i < N; i++ {
	// 	r.Select(rank_arry[i]) // compute select for each rank
	// }
	// duration := time.Since(start)
	// fmt.Println(duration.Seconds())

	//}

	// PART 3
	spar_arr := pkgbitvectors.Create(10)
	spar_arr = spar_arr.Append("foo", 1)
	spar_arr = spar_arr.Append("baz", 9)
	spar_arr = spar_arr.Append("bar", 5)
	spar_arr = spar_arr.Append("cat", 0)

	fin_spar_arr := spar_arr.Finalize()
	fmt.Println(fin_spar_arr.NumElementAt(9))
	// bv := fin_spar_arr.R.BV
	// for i := 0; i < bv.Length(); i++ {
	// 	fmt.Print(bv.Element(i))
	// }
	// fmt.Println("")
	// fmt.Println(fin_spar_arr.GetIndexOf(2))

}
