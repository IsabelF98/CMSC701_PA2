# CMSC 701 Programing Assignment 2: Bitvectors
Isabel Fernandez
March 28, 2022

## Overall Infromation
* Language: Go 
* External Packages Used: github.com/dropbox/godropbox/container/bitvector

## Rank
* RankSupport(bv): Jacobson rank structure, input is a bit-vecotor (bv) created using the package linked above and returns a rank support structure (rank_struct). *rank_struct=RankSupport(bv)*

* Rank(idx): Returns the rank as an index (rank) at a given index (not inclusive). Operated on a rank support structure (rank_struct). *rank=rank_struct.Rank(idx)*

* Overhead: Returns the amount of overhead space (overhead) as an integer stored in rank support structure (rank_struct). *overhead=rank_struct.Overhead()*

## Select
* Select(rank): Returns the index (idx) as an integer of the first instance of the given rank value (rank). Operated on a rank support structure (rank_struct). Call: *idx=rank_struct.Select(rank)*

## Sparse Array:
* Create(size): Creates an empty sparse array (spar_array) structure of a given size (size) as an integer. The sparse array is stored as a appendable sparse array structure with a bit vector and string array. Call: *spar_arr=Create(size)*

* Append(elem, pos): Appends a string (elem) to the saprse array structure (spar_array) at a position (pos) as an integer. Returns the updated spar_array structure. Call: *spar_array=spar_array.Append(elem, pos)*

* Finalize(): Finalizes a sparse array structure (spar_array). Cannot append elements after finalized. Returns a new type of structure that contains rank supposrt structure (fin_spar_array). Call: *fin_spar_array=spar_array.Finalize()*

* GetElement(idx): Gets the element (elem) as a string at an index (idx) in the finalized sparse array structure (fin_spar_array). Call: *elem=fin_spar_array.GetElement(idx)*

* GetAtIdx(idx,elem): Checks if an element (elem) exists at a given index (idx). If element does exist returns true, otherwise false (boolian value). Call: *bool_val=fin_spar_array.GetAtIdx(idx,elem)*

* GetAtRank(rank,elem): Checks if an element (elem) exists at a given rank (rank). If element does exist returns true, otherwise false (boolian value). Call: *bool_val=fin_spar_array.GetAtRank(rank,elem)*

* GetIndexOf(rank): Retruns the index as an integer at the instanse of a given rank (inclusive). Call: *idx=fin_spar_array.GetIndexOf(rank)*

* NumElementAt(idx): Returns the number of elements as an integer at a giveen index (inclusive). Call: *N_elements=fin_spar_array.NumElementAt(idx)*

* Size(): Returns size of sparse array as an integer. Will be equal to the size given when sparse array was created. Call *size=fin_spar_array.Size()*

* NumElem(): Returns the number of elements in the sparse array as an integer. Will be equal to the number of elements appended into the sparse array.. Call *N_elements=fin_spar_array.NumElem()*
