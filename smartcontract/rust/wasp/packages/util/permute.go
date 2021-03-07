package util

import (
	"bytes"
	"github.com/iotaledger/wasp/packages/hashing"
	"sort"
)

// Permutation16 deterministic permutation of integers from 0 to size-1
type Permutation16 struct {
	size        uint16
	permutation []uint16
	curSeqIndex uint16
}

func NewPermutation16(size uint16, seed []byte) *Permutation16 {
	ret := &Permutation16{
		size: size,
	}
	return ret.Shuffle(seed)
}

type idxToPermute struct {
	idx  uint16
	hash hashing.HashValue
}

func (perm *Permutation16) Shuffle(seed []byte) *Permutation16 {
	tosort := make([]*idxToPermute, perm.size)
	data := make([]byte, len(seed)+2)
	copy(data, seed)

	for i := range tosort {
		copy(data[len(data)-2:], Uint16To2Bytes(uint16(i)))
		tosort[i] = &idxToPermute{
			idx:  uint16(i),
			hash: hashing.HashData(data),
		}
	}
	sort.Slice(tosort, func(i, j int) bool {
		return bytes.Compare(tosort[i].hash[:], tosort[j].hash[:]) < 0
	})

	perm.permutation = make([]uint16, perm.size)
	for i := range perm.permutation {
		perm.permutation[i] = tosort[i].idx
	}
	perm.curSeqIndex = 0
	return perm
}

func (perm *Permutation16) Current() uint16 {
	return perm.permutation[perm.curSeqIndex]
}

func (perm *Permutation16) Next() uint16 {
	ret := perm.permutation[perm.curSeqIndex]
	perm.curSeqIndex = (perm.curSeqIndex + 1) % perm.size
	return ret
}

func (perm *Permutation16) GetArray() []uint16 {
	ret := make([]uint16, len(perm.permutation))
	copy(ret, perm.permutation)
	return ret
}

func ValidPermutation(perm []uint16) bool {
	n := uint16(len(perm))

	// check if every value exists
	for i := uint16(0); i < n; i++ {
		if _, found := findIndexOf(i, perm); !found {
			return false
		}
	}

	// check if all value are different
	for i, v1 := range perm {
		for j, v2 := range perm {
			if i != j && v1 == v2 {
				return false
			}
		}
	}
	return true
}

func findIndexOf(val uint16, sequence []uint16) (uint16, bool) {
	for i, s := range sequence {
		if s == val {
			return uint16(i), true
		}
	}
	return 0, false
}
