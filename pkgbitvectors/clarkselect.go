package pkgbitvectors

import (
	"fmt"

	"github.com/dropbox/godropbox/container/bitvector"
)

type SelectStruct struct {
	BV *bitvector.BitVector
}

func SelectSuport(r *RankStruct) {
	bv := r.BV       // bit vector
	n := bv.Length() // bitvector length

	chunk_weight := 10   // int(math.Pow(math.Log2(float64(n)), 2))   // chunk weight
	chunk_spar_len := 23 // int(math.Pow(math.Log2(float64(n)), 4)) // min length of sparse chunk

	max_weight := r.Rank(n - 1)
	n_chunks := int(float64(max_weight / chunk_weight))

	spar_chunk_offset := []int{}   // array of sparse chunk offsets
	spar_chunk_1offsets := []int{} // array of sparse chunk 1 offsets
	dense_chunk_offset := []int{}  // array of densee chunk offsets

	chunk_weight_idx := 0   // index for chunk weight
	chunk_position_idx := 0 // index of chunk position

	for i := 0; i < n; i++ {
		rank := r.Rank(i) // find rank at position i

		if rank == (chunk_weight_idx+1)*chunk_weight { // if the rank coresponds to chunk weight
			if i-chunk_position_idx-1 >= chunk_spar_len { // if chunk is sparse
				spar_chunk_offset = append(spar_chunk_offset, chunk_position_idx) // add index to sparse chunk offset array
				// find 1 offseet in sparse chunk
				for j := chunk_position_idx; j < i; j++ {
					if bv.Element(j) == 1 {
						spar_chunk_1offsets = append(spar_chunk_1offsets, j+1) // add index to sparse chunk 1 offset array
					}
				}
				chunk_weight_idx += 1  // increase chunk weigh index by 1
				chunk_position_idx = i // new chunk start index
			} else { // if the chunk is dense
				dense_chunk_offset = append(dense_chunk_offset, chunk_position_idx)
				chunk_weight_idx += 1
				chunk_position_idx = i
			}
		} else if chunk_weight_idx == n_chunks { // last chunk might not have total chunk weight
			if i-chunk_position_idx-1 >= chunk_spar_len {
				spar_chunk_offset = append(spar_chunk_offset, chunk_position_idx)
				chunk_weight_idx += 1
				chunk_position_idx = i
			} else {
				dense_chunk_offset = append(dense_chunk_offset, chunk_position_idx)
				chunk_weight_idx += 1
				chunk_position_idx = i
			}
		}
	}
	fmt.Println(spar_chunk_offset)
	fmt.Println(spar_chunk_1offsets)
	fmt.Println(dense_chunk_offset)

}
