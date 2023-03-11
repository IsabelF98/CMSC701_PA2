package pkgbitvectors

import (
	"math"

	"github.com/dropbox/godropbox/container/bitvector"
)

func rankData(bv *bitvector.BitVector) ([]int, [][]int) {
	/* This function creates two data sets for finding the rank of a bit vector.
	The rank of the chunks and the rank of the sub-chunks based on Jacobson's rank.
	*/

	n := bv.Length()                                     // bitvector length
	chunk_len := int(math.Pow(math.Log2(float64(n)), 2)) // chunk length
	sub_chunk_len := int(0.5 * math.Log2(float64(n)))    // sub-chunk length

	cumulative_rank := []int{0}     // array of chunk ranks
	relative_cumu_rank := [][]int{} // array of sub-chunk rank arrays

	chunk_start := 0 // start of chunk index

	for i := 0; i < int(n/chunk_len)+1; i++ { // iterate through cunks
		sub_chunk_start := chunk_start // start of sub-chunk index

		chunk_rank := 0                 // initiate chunk rank as 0
		sub_chunk_cumu_rank := []int{0} // array sub-chunk ranks

		if i < int(n/chunk_len) { // check that index is not out of range for chunk
			for j := 0; j < chunk_len; j++ {
				chunk_rank += int(bv.Element(chunk_start + j)) // sum up chunk rank
			}
			cumulative_rank = append(cumulative_rank, chunk_rank) // append chunk rank to cumulative_rank
			chunk_start += chunk_len                              // update chunk start index
		}

		for k := 0; k < int(chunk_len/sub_chunk_len)+1; k++ { // iterate through sub chunks
			sub_chunk_rank := 0 // initiate sub-chunk rank as 0

			if sub_chunk_start+sub_chunk_len < n { // check that that index is not out of range for sub-chunk
				for l := 0; l < sub_chunk_len; l++ {
					sub_chunk_rank += int(bv.Element(sub_chunk_start + l)) // sum up sub-chunk rank
				}
			} else {
				break
			}
			sub_chunk_cumu_rank = append(sub_chunk_cumu_rank, sub_chunk_rank) // append sub-chunk rank to sub_chunk_cumu_rank
			sub_chunk_start += sub_chunk_len                                  // update sub-chunk start index
		}
		relative_cumu_rank = append(relative_cumu_rank, sub_chunk_cumu_rank) // append sub-chunk rank array to relative_cumu_rank
	}
	return cumulative_rank, relative_cumu_rank
}

func Rank(bv *bitvector.BitVector, idx int) int {
	n := bv.Length()                                     // bitvector length
	chunk_len := int(math.Pow(math.Log2(float64(n)), 2)) // chunk length
	sub_chunk_len := int(0.5 * math.Log2(float64(n)))    // sub-chunk length

	cumulative_rank, relative_cumu_rank := rankData(bv) // rank data structures
	idx1 := idx / chunk_len                             // index of chunk
	idx2 := (idx - idx1*chunk_len) / sub_chunk_len      // index of sub-chunk

	rank := 0 // initialize rank
	for i := 0; i < idx1+1; i++ {
		rank += cumulative_rank[i] // sum up chunk ranks
	}
	for j := 0; j < idx2+1; j++ {
		rank += relative_cumu_rank[idx1][j] // sum up sub-chunk ranks
	}
	start := chunk_len*idx1 + sub_chunk_len*idx2 // start index for sub-chunk
	for k := start; k < idx; k++ {
		rank += int(bv.Element(k)) // sum up remaining rank
	}
	return rank
}
