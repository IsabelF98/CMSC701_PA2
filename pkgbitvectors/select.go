package pkgbitvectors

import (
	"math"
)

func (r *RankStruct) Select(rank int) int {
	/*

	 */

	bv := r.BV       // bit vector
	n := bv.Length() // bitvector length

	left := 0
	right := n - 1
	aux := float64(left+right) / 2.0
	center := int(math.Floor(aux))

	for i := 0; i < n; i++ {
		c_rank := r.Rank(center)

		// stopping criterial
		if rank < c_rank && center == left+1 {
			return center
		}
		if rank > c_rank && center == right-1 {
			return right
		}
		if rank == c_rank { // rank is a direct match
			return center
		}

		if rank < c_rank { // look at top half
			right = int(center)
			aux = float64(left+right) / 2.0
			center = int(math.Floor(aux))

		} else if rank > c_rank { // look at bottom half
			left = int(center)
			aux = float64(left+right) / 2.0
			center = int(math.Floor(aux))
		}
	}
	return -1 // return -1 if error occured (not a possible index)
}
