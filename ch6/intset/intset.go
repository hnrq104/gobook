package main

import (
	"bytes"
	"fmt"
)

var uintSize = 32 << (^uint(0) >> 63)

// An intset is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the small non negative element x
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, x%uintSize
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x ...int) {
	for _, num := range x {
		word, bit := num/uintSize, num%uintSize
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}

		s.words[word] |= (1 << bit)
	}
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

/*Exercise 6.1*/
func (s *IntSet) Len() int {
	size := 0
	for _, word := range s.words {
		for j := 0; j < uintSize; j++ {
			if word>>j == 0 {
				break
			}
			if word&(1<<uint(j)) != 0 {
				size++
			}
		}
	}
	return size
}

func (s *IntSet) Remove(x int) {
	word, bit := x/uintSize, x%uintSize

	if word >= len(s.words) {
		return
	}

	s.words[word] &^= (1 << uint(bit))

}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	var cp *IntSet = new(IntSet)
	cp.words = make([]uint, len(s.words))
	copy(cp.words, s.words)
	return cp
}

//Exercise 6.2

// Sets to s intersection between s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	maxlenght := min(len(s.words), len(t.words))
	biggestWord := 0
	for i := 0; i < maxlenght; i++ {
		s.words[i] &= t.words[i]
		if s.words[i] != 0 {
			biggestWord++
		}
	}

	s.words = s.words[:biggestWord+1]
}

// Sets to s difference between s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	maxlenght := min(len(s.words), len(t.words))
	for i := 0; i < maxlenght; i++ {
		s.words[i] &^= t.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elem() []uint {
	elems := make([]uint, 0)
	for i, word := range s.words {
		for j := 0; j < uintSize; j++ {
			if word>>j == 0 {
				break
			}
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, uint(uintSize*i+j))
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1, 144, 9, 10, 22, 34, 256, 32)
	fmt.Println(&x)

	y.Add(9, 42, 53, 23, 144, 22)
	fmt.Println(&y)

	z := *x.Copy()

	z.Remove(144)
	fmt.Println(&z)

	z.SymmetricDifference(&y)

	fmt.Println(y.Len())

	for _, e := range z.Elem() {
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}
