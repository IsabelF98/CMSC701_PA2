package main

import (
	"fmt"
	"math/rand"
	"time"

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

	// 	r := pkgbitvectors.RankSupport(bv) // rank structure
	// 	overhead := r.Overhead()
	// 	fmt.Println(overhead)
	// 	start := time.Now()
	// 	for i := 0; i < N; i++ {
	// 		r.Rank(idx_array[i]) // compute rank for each index
	// 	}
	// 	duration := time.Since(start)
	// 	fmt.Println(duration.Seconds())

	// 	start := time.Now()
	// 	for i := 0; i < N; i++ {
	// 		r.Select(rank_arry[i]) // compute select for each rank
	// 	}
	// 	duration := time.Since(start)
	// 	fmt.Println(duration.Seconds())

	// }

	// PART 3
	size_array := []int{1000, 10000, 100000, 1000000}
	rand_str := []string{"cat", "foo", "bat", "car", "yes", "bar", "jar", "far", "mmm", "lit", "kit", "mit",
		"top", "mop", "hop", "lot", "cot", "hot", "not", "grr", "hee", "she", "oui", "tea", "the", "wow",
		"now", "moo", "ste", "hip", "hop", "mar", "non", "nan", "fam", "fan", "brr", "sir", "fur", "dir", "lol", "wtf"}
	sparsity_level := []float64{0.01, 0.05, 0.1}
	for i := 0; i < len(size_array); i++ {
		for k := 0; k < len(sparsity_level); k++ {
			spar_arr := pkgbitvectors.Create(size_array[i])
			fmt.Println("size:", size_array[i])
			fmt.Println("sparsity:", sparsity_level[k])
			aux := int(float64(size_array[i]) * sparsity_level[k])
			for j := 0; j < aux; j++ {
				pos := (size_array[i] / aux) * j
				elm := rand.Intn(len(rand_str))
				spar_arr = spar_arr.Append(rand_str[elm], pos)
			}
			fin_spar_arr := spar_arr.Finalize()

			start := time.Now()
			fin_spar_arr.NumElementAt((size_array[i] / aux) * 5)
			duration := time.Since(start)
			fmt.Println(duration.Seconds())
			start = time.Now()
			fin_spar_arr.GetIndexOf(5)
			duration = time.Since(start)
			fmt.Println(duration.Seconds())
			fmt.Println(" ")
		}
	}

	// spar_arr := pkgbitvectors.Create(10)
	// spar_arr = spar_arr.Append("foo", 1)
	// spar_arr = spar_arr.Append("baz", 9)
	// spar_arr = spar_arr.Append("bar", 5)
	// spar_arr = spar_arr.Append("cat", 0)

	// fin_spar_arr := spar_arr.Finalize()
	// fmt.Println(fin_spar_arr.NumElementAt(9))
	// bv := fin_spar_arr.R.BV
	// for i := 0; i < bv.Length(); i++ {
	// 	fmt.Print(bv.Element(i))
	// }
	// fmt.Println("")
	// fmt.Println(fin_spar_arr.GetIndexOf(2))

}
