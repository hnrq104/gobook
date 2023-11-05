package main

import (
	"fmt"
	"sort"
)

func isPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len(); i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type ByteSlice []byte

func (b ByteSlice) Len() int           { return len(b) }
func (b ByteSlice) Less(i, j int) bool { return b[i] < b[j] }
func (b ByteSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func main() {
	nome := "henrique"
	ela := "anna"

	teste := []int{1, 1, 3, 3, 5, 3, 3, 1, 1}

	fmt.Println(nome, isPalindrome(ByteSlice(nome)))
	fmt.Println(ela, isPalindrome(ByteSlice(ela)))
	fmt.Println(teste, isPalindrome(sort.IntSlice(teste)))

}
