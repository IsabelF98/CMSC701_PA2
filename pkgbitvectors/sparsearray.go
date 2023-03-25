package pkgbitvectors

import (
	"github.com/dropbox/godropbox/container/bitvector"
)

// Structure of sparse array to be appended to
type ApendSparArr struct {
	BV          *bitvector.BitVector
	StringArray []string
}

// Structure of finalized sparse array
type FinSparArr struct {
	StringArray []string
	R           *RankStruct
}

func Create(size int) *ApendSparArr {
	/*
		This fucntion creates a empty sparse array structure of a given size
	*/
	data := []byte{} // empty byte array
	for i := 0; i < size; i++ {
		data = append(data, 0) // append 0 to byte array until size is reached
	}
	var spar_arr ApendSparArr                        // create sparse array structure
	spar_arr.BV = bitvector.NewBitVector(data, size) // add bit vecotr
	spar_arr.StringArray = []string{}                // add empty string array
	return &spar_arr
}

func (spar_arr *ApendSparArr) Append(elem string, pos int) *ApendSparArr {
	/*
		This function adds an emement to the sparse array at a spacific index.
		If an element already exists at the position an error will occur.
	*/
	bv := spar_arr.BV                   // current bitvector
	aux_str_arr := spar_arr.StringArray // current string array
	size := bv.Length()                 // size of bit vector

	// update bitvector
	if bv.Element(pos) == 1 { // an element already exists at this position
		panic("There is already an element at this position")
	} else { // add a 1 to the bit vecotr at this position
		if pos == size-1 {
			bv.Delete(pos)
			bv.Append(1)
		} else {
			bv.Delete(pos)
			bv.Insert(1, pos)
		}
	}

	// update string array
	str_arr := []string{}
	if len(aux_str_arr) == 0 { // array is empty add element
		str_arr = append(aux_str_arr, elem)
	} else { // array already has elements
		r := RankSupport(bv)
		rank := r.Rank(pos) // rank of bit vector at that positon

		if rank == len(aux_str_arr) { // add element to the end of array
			str_arr = append(aux_str_arr, elem)
		} else if rank == 0 { // add element to the beging of array
			str_arr = append(str_arr, elem)
			str_arr = append(str_arr, aux_str_arr...)
		} else { // add element at position within the array
			str_arr = append(str_arr, aux_str_arr[:rank]...)
			str_arr = append(str_arr, elem)
			str_arr = append(str_arr, aux_str_arr[rank:]...)
		}
	}
	// update sparse array structure
	var spar_arr2 ApendSparArr
	spar_arr2.BV = bv
	spar_arr2.StringArray = str_arr

	return &spar_arr2
}

func (spar_arr *ApendSparArr) Finalize() *FinSparArr {
	/*
		This function finalizes the sparse array structure. After finalized no other elements
		can be appended. A rank structure for the bit vecotr is creeeated. If the array is
		empty an error occurs.
	*/
	bv := spar_arr.BV               // current bit vecotr
	str_arr := spar_arr.StringArray // current sparse array

	if len(str_arr) == 0 { // array is empty
		panic("This is an empty sparse array!")
	}
	// create finalized sparse array structure
	var fin_spar_arr FinSparArr
	fin_spar_arr.StringArray = str_arr
	fin_spar_arr.R = RankSupport(bv)
	return &fin_spar_arr
}

func (fin_spar_arr *FinSparArr) GetElement(idx int) string {
	/*
		This function gets the element in the array at a given index.
		If an element does not exist at the given index and empty string is returned.
	*/
	str_arr := fin_spar_arr.StringArray // string array
	r := fin_spar_arr.R                 // rank structuree
	bit := r.BV.Element(idx)            // bit in bitvector at the index
	if bit == 0 {                       // element does not exist at that position
		return ""
	} else {
		rank := r.Rank(idx)  // rank at position
		return str_arr[rank] // retrun element at position
	}
}

func (fin_spar_arr *FinSparArr) GetAtIdx(idx int, elem string) bool {
	/*
		This function checks if an element is present at a given INDEX.
		If the element does exist at the index it returns true, false otherwise.
	*/
	check_elem := fin_spar_arr.GetElement(idx) // element at the index given

	if check_elem == elem { // element is a match
		return true
	} else { // emlement is not a mathc
		return false
	}
}

func (fin_spar_arr *FinSparArr) GetAtRank(rank int, elem string) bool {
	/*
		This function checks if an element is present at a given RANK.
		If the element does exist at the index it returns true, false otherwise.
	*/
	check_elem := fin_spar_arr.StringArray[rank] // element at the given rank

	if check_elem == elem { // element is a match
		return true
	} else { // element is not a match
		return false
	}
}

func (fin_spar_arr *FinSparArr) GetIndexOf(rank int) int {
	/*
		This function retruns the index at the instanse of a given rank (inclusive).
		If the rank is bigger than the number of elements -1 is returned.
	*/
	if rank > fin_spar_arr.NumElem() { // rank is bigger than number of elements (not possible)
		return -1
	} else {
		r := fin_spar_arr.R
		idx := r.Select(rank) // index of rank (not inclusive)
		return idx - 1
	}
}

func (fin_spar_arr *FinSparArr) NumElementAt(idx int) int {
	/*
		This function returns the number of elements at a giveen index (inclusive).
		If index is out or range of bitvector error occurs.
	*/
	r := fin_spar_arr.R     // rank support of bitvector
	bv := fin_spar_arr.R.BV // bitvector

	if idx+1 == bv.Length() { // if at last index
		rank := r.Rank(idx)          // rank at index (not inclusive)
		rank += int(bv.Element(idx)) // plus rank at index (inclusive)
		return rank
	} else if idx+1 > bv.Length() { // index is out of range
		panic("Index is out of range.")
	} else {
		rank := r.Rank(idx + 1) // rank at index (inclusive)
		return rank
	}
}

func (fin_spar_arr *FinSparArr) Size() int {
	/*
		This function returns the size of the sparse array (length of bit vector).
	*/
	return int(fin_spar_arr.R.BV.Length())
}

func (fin_spar_arr *FinSparArr) NumElem() int {
	/*
		This function returns thee number of non empty elements in the sparse array (length of string array).
	*/
	return len(fin_spar_arr.StringArray)
}
