package pkgbitvectors

import (
	"math"
)

func (r *RankStruct) Select(rank int) int {
	/*
		This function finds the index of a rank for a given bitvector using binary search.
		The function uses the rank data created using the Jacobson rank.
	*/

	// if rank is 0 the first instancee is the 0th index
	if rank == 0 {
		return 0
	}

	bv := r.BV       // bit vector
	n := bv.Length() // bitvector length

	left := 0      // initialize left position
	right := n - 1 // initialize right position

	for left <= right {
		aux := float64(left+right) / 2.0
		center := int(math.Floor(aux)) // cneter position
		c_rank := r.Rank(center)

		if rank < c_rank { // look at top half
			right = int(center) - 1

		} else if rank > c_rank { // look at bottom half
			left = int(center) + 1
			if center == right { // rank happens at last index
				return right + 1 // last index + 1
			}

		} else { // found rank
			for center > 0 && rank == c_rank { // find first instance of rank
				center -= 1
				c_rank = r.Rank(center)
			}
			return center + 1 // index (not inclusive)
		}

	}
	return -1 // return -1 if error occured (not a possible index)
}
