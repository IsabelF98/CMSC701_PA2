package pkgbitvectors

import (
	"math"

	"github.com/dropbox/godropbox/container/bitvector"
)

type RankStruct struct {
	BV             *bitvector.BitVector
	CumulativeRank *bitvector.BitVector
	RelCumulRank   *bitvector.BitVector
}

func RankSupport(bv *bitvector.BitVector) *RankStruct {
	/* This function creates two data sets for finding the 1 rank of a bit vector.
	The rank of the chunks and the rank of the sub-chunks based on Jacobson's rank.
	*/

	n := bv.Length()                                     // bitvector length
	chunk_len := int(math.Pow(math.Log2(float64(n)), 2)) // chunk length
	sub_chunk_len := int(0.5 * math.Log2(float64(n)))    // sub-chunk length

	cumulative_rank := []byte{0}   // array of chunk ranks
	relative_cumu_rank := []byte{} // array of sub-chunk rank arrays

	chunk_start := 0 // start of chunk index

	for i := 0; i < int(n/chunk_len)+1; i++ { // iterate through cunks
		sub_chunk_start := chunk_start // start of sub-chunk index

		chunk_rank := byte(0)                                    // initiate chunk rank as 0
		relative_cumu_rank = append(relative_cumu_rank, byte(0)) // array sub-chunk ranks
		idx := (i * (int(chunk_len/sub_chunk_len) + 1))

		if i < int(n/chunk_len) { // check that index is not out of range for chunk
			for j := 0; j < chunk_len; j++ {
				chunk_rank += byte(bv.Element(chunk_start + j)) // sum up chunk rank
			}
			cumulative_rank = append(cumulative_rank, cumulative_rank[i]+chunk_rank) // append chunk rank to cumulative_rank
			chunk_start += chunk_len                                                 // update chunk start index
		}

		for k := 0; k < int(chunk_len/sub_chunk_len); k++ { // iterate through sub chunks
			sub_chunk_rank := byte(0) // initiate sub-chunk rank as 0

			if sub_chunk_start+sub_chunk_len < n { // check that that index is not out of range for sub-chunk
				for l := 0; l < sub_chunk_len; l++ {
					sub_chunk_rank += byte(bv.Element(sub_chunk_start + l)) // sum up sub-chunk rank
				}
			} else {
				break
			}
			relative_cumu_rank = append(relative_cumu_rank, relative_cumu_rank[idx+k]+sub_chunk_rank) // append sub-chunk rank to sub_chunk_cumu_rank
			sub_chunk_start += sub_chunk_len                                                          // update sub-chunk start index
		}

	}

	var r RankStruct
	r.BV = bv
	r.CumulativeRank = bitvector.NewBitVector(cumulative_rank, len(cumulative_rank))
	r.RelCumulRank = bitvector.NewBitVector(relative_cumu_rank, len(relative_cumu_rank))

	return &r
}

func (r *RankStruct) Rank(idx int) int {
	/*
		This function finds the 1 rank of of a bitvector at index idx.
		The function uses the rank data created using the previous function.
	*/

	bv := r.BV                                           // bitvector
	n := bv.Length()                                     // bitvector length
	chunk_len := int(math.Pow(math.Log2(float64(n)), 2)) // chunk length
	sub_chunk_len := int(0.5 * math.Log2(float64(n)))    // sub-chunk length

	idx1 := idx / chunk_len // index of chunk
	idx_aux := (idx1 * (int(chunk_len/sub_chunk_len) + 1))
	idx2 := (idx - idx1*chunk_len) / sub_chunk_len // index of sub-chunk
	rank1 := r.CumulativeRank.Bytes()[idx1]        // rank of chunk
	rank2 := r.RelCumulRank.Bytes()[idx_aux+idx2]  // rank of sub-chunk

	idx3 := chunk_len*idx1 + sub_chunk_len*idx2 // start index for sub-chunk
	rank3 := byte(0)

	for k := idx3; k < idx; k++ {
		rank3 += byte(bv.Element(k)) // sum up remaining rank
	}
	return int(rank1 + rank2 + rank3)
}
