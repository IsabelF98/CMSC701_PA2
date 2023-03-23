package pkgbitvectors

// import (
// 	"github.com/dropbox/godropbox/container/bitvector"
// )

// type SelectStruct struct {
// 	BV                  *bitvector.BitVector
// 	ChunkOffset         []int
// 	SparChunk1Offset    []int
// 	SubChunkOffset      []int
// 	SparSubChunk1Offset []int
// }

// func SelectSuport(r *RankStruct) *SelectStruct {
// 	bv := r.BV       // bit vector
// 	n := bv.Length() // bitvector length

// 	chunk_weight := 10   // int(math.Pow(math.Log2(float64(n)), 2))   // chunk weight
// 	chunk_spar_len := 23 // int(math.Pow(math.Log2(float64(n)), 4)) // min length of sparse chunk

// 	max_weight := r.Rank(n - 1)                         // max weight of bitvector
// 	n_chunks := int(float64(max_weight / chunk_weight)) // number of chunks

// 	sub_chunk_weight := 4                                         // sub-chunk weight
// 	sub_chunk_spar_len := 10                                      // min length of sparse sub-chunks
// 	n_sub_chunks := int(float64(chunk_weight / sub_chunk_weight)) // number of sub-chunks

// 	chunk_offset := []int{}        // array of chunk offsets
// 	spar_chunk_1offsets := []int{} // array of sparse chunk 1 offsets

// 	sub_chunk_offset := []int{}        // array of sub-chunk offsets
// 	spar_sub_chunk_1offsets := []int{} // array of sparse chunk 1 offsets

// 	chunk_weight_idx := 0   // index for chunk weight
// 	chunk_position_idx := 0 // index of chunk position

// 	for i := 0; i < n; i++ {
// 		chunk_rank := r.Rank(i) // find rank at position i

// 		sub_chunk_weight_idx := 0                    // index for sub-chunk weight
// 		sub_chunk_position_idx := chunk_position_idx // index of sub-chunk position

// 		if chunk_rank == (chunk_weight_idx+1)*chunk_weight { // if the rank coresponds to chunk weight
// 			chunk_offset = append(chunk_offset, chunk_position_idx) // add index to chunk offset array

// 			if i-chunk_position_idx-1 >= chunk_spar_len { // if chunk is sparse
// 				// find 1 offseet in sparse chunk
// 				for j := chunk_position_idx; j < i; j++ {
// 					if bv.Element(j) == 1 {
// 						spar_chunk_1offsets = append(spar_chunk_1offsets, j+1) // add index to sparse chunk 1 offset array
// 					}
// 				}

// 			} else { // if the chunk is dense
// 				for j := chunk_position_idx; j < i; j++ {
// 					sub_chunk_rank := r.Rank(j) // find rank at position j
// 					aux_weight := (chunk_weight_idx)*chunk_weight + (sub_chunk_weight_idx+1)*sub_chunk_weight

// 					if sub_chunk_rank == aux_weight { // if the rank coresponds to sub-chunk weight
// 						sub_chunk_offset = append(sub_chunk_offset, sub_chunk_position_idx) // add index to sub-chunk offset array

// 						if j-sub_chunk_position_idx-1 >= sub_chunk_spar_len { // if sub-chunk is sparse
// 							// find 1 offseet in sparse sub-chunk
// 							for k := sub_chunk_position_idx; k < j; k++ {
// 								if bv.Element(k) == 1 {
// 									spar_sub_chunk_1offsets = append(spar_sub_chunk_1offsets, k+1) // add index to sparse sub-chunk 1 offset array
// 								}
// 							}
// 						}
// 						sub_chunk_weight_idx += 1  // increase sub-chunk weigh index by 1
// 						sub_chunk_position_idx = j // new sub-chunk start index

// 					} else if sub_chunk_weight_idx == n_sub_chunks { // last sub-chunk might not have total sub-chunk weight
// 						sub_chunk_offset = append(sub_chunk_offset, sub_chunk_position_idx) // add index to sub-chunk offset array

// 						if j-sub_chunk_position_idx-1 >= sub_chunk_spar_len { // if sub-chunk is sparse
// 							// find 1 offseet in sparse sub-chunk
// 							for k := sub_chunk_position_idx; k < j; k++ {
// 								if bv.Element(k) == 1 {
// 									spar_sub_chunk_1offsets = append(spar_sub_chunk_1offsets, k+1) // add index to sparse sub-chunk 1 offset array
// 								}
// 							}
// 						}
// 						sub_chunk_weight_idx += 1  // increase sub-chunk weigh index by 1
// 						sub_chunk_position_idx = j // new sub-chunk start index
// 					}
// 				}
// 			}
// 			chunk_weight_idx += 1  // increase chunk weigh index by 1
// 			chunk_position_idx = i // new chunk start index

// 		} else if chunk_weight_idx == n_chunks { // last chunk might not have total chunk weight
// 			chunk_offset = append(chunk_offset, chunk_position_idx) // add index to chunk offset array

// 			if i-chunk_position_idx-1 >= chunk_spar_len { // if chunk is sparse
// 				// find 1 offseet in sparse chunk
// 				for j := chunk_position_idx; j < n; j++ {
// 					if bv.Element(j) == 1 {
// 						spar_chunk_1offsets = append(spar_chunk_1offsets, j+1) // add index to sparse chunk 1 offset array
// 					}
// 				}

// 			} else { // if chunk is dense
// 				for j := chunk_position_idx; j < n; j++ {
// 					sub_chunk_rank := r.Rank(j) // find rank at position j
// 					aux_weight := (chunk_weight_idx)*chunk_weight + (sub_chunk_weight_idx+1)*sub_chunk_weight

// 					if sub_chunk_rank == aux_weight { // if the rank coresponds to sub-chunk weight
// 						sub_chunk_offset = append(sub_chunk_offset, sub_chunk_position_idx) // add index to sub-chunk offset array

// 						if j-sub_chunk_position_idx-1 >= sub_chunk_spar_len { // if sub-chunk is sparse
// 							// find 1 offseet in sparse sub-chunk
// 							for k := sub_chunk_position_idx; k < j; k++ {
// 								if bv.Element(k) == 1 {
// 									spar_sub_chunk_1offsets = append(spar_sub_chunk_1offsets, k+1) // add index to sparse sub-chunk 1 offset array
// 								}
// 							}
// 						}
// 						sub_chunk_weight_idx += 1  // increase sub-chunk weigh index by 1
// 						sub_chunk_position_idx = j // new sub-chunk start index

// 					} else if sub_chunk_weight_idx == n_sub_chunks { // last sub-chunk might not have total sub-chunk weight
// 						sub_chunk_offset = append(sub_chunk_offset, sub_chunk_position_idx)
// 						if j-sub_chunk_position_idx-1 >= sub_chunk_spar_len { // if sub-chunk is sparse
// 							for k := sub_chunk_position_idx; k < n; k++ {
// 								if bv.Element(k) == 1 {
// 									spar_sub_chunk_1offsets = append(spar_sub_chunk_1offsets, k+1) // add index to sparse sub-chunk 1 offset array
// 								}
// 							}
// 						}
// 						sub_chunk_weight_idx += 1  // increase sub-chunk weigh index by 1
// 						sub_chunk_position_idx = j // new sub-chunk start index
// 					}
// 				}
// 			}
// 			chunk_weight_idx += 1  // increase chunk weigh index by 1
// 			chunk_position_idx = i // new chunk start index
// 		}
// 	}
// 	var s SelectStruct
// 	s.BV = bv
// 	s.ChunkOffset = chunk_offset
// 	s.SparChunk1Offset = spar_chunk_1offsets
// 	s.SubChunkOffset = sub_chunk_offset
// 	s.SparSubChunk1Offset = spar_sub_chunk_1offsets

// 	return &s
// }

// func (s *SelectStruct) Select(rank int) int {
// 	bv := s.BV // bit vector
// 	// n := bv.Length() // bitvector length

// 	chunk_weight := 10   // int(math.Pow(math.Log2(float64(n)), 2))   // chunk weight
// 	chunk_spar_len := 23 // int(math.Pow(math.Log2(float64(n)), 4)) // min length of sparse chunk

// 	sub_chunk_weight := 4    // int(math.Sqrt(math.Log2(float64(n)))) // sub-chunk weight
// 	sub_chunk_spar_len := 10 // int(0.5*math.Log2(float64(n))) // min length of sparse sub-chunks

// 	chunk_offset := s.ChunkOffset
// 	spar_chunk_1offsets := s.SparChunk1Offset
// 	sub_chunk_offset := s.SubChunkOffset
// 	spar_sub_chunk_1offsets := s.SparSubChunk1Offset

// 	idx1 := int(rank / chunk_weight)

// 	if chunk_offset[idx1+1]-chunk_offset[idx1] >= chunk_spar_len {
// 		for i := 0; i < len(spar_chunk_1offsets); i++ {
// 			if spar_chunk_1offsets[i] >= chunk_offset[idx1] {
// 				idx2 := i + (rank - (chunk_weight * idx1)) - 1
// 				return spar_chunk_1offsets[idx2]
// 			}
// 		}
// 	}
// }
