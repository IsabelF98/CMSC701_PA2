package pkgbitvectors

import (
	"math"
)

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

func (r *RankStruct) Overhead() int {
	bv := r.BV
	cumulative_rank := r.CumulativeRank
	relative_cumu_rank := r.RelCumulRank

	size := bv.Length() + cumulative_rank.Length() + relative_cumu_rank.Length()

	return size
}

// func (r *RankStruct) Save(filename string) {
// 	buf := &bytes.Buffer{}
// 	if err := gob.NewEncoder(buf).Encode(r); err != nil {
// 		panic(err)
// 	}

// 	f, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatal("Couldn't open file")
// 	}
// 	defer f.Close()

// 	w := bufio.NewWriter(f)
// 	_, err = buf.WriteTo(w)
// 	if err != nil {
// 		log.Fatal("Write failed")
// 	}
// }

// func Load(filename string) *RankStruct {
// 	content, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Couldn't read file", err)
// 	}

// 	var r RankStruct
// 	err = json.Unmarshal(content, &r)
// 	if err != nil {
// 		fmt.Println("Couldn't load data", err)
// 	}

// 	return &r
// }
