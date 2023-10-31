package main

import "fmt"

func main() {
	s1 := "abcdegkal"
	s2 := "bcedaaf"

	fmt.Println(anagram(s1, s2))
}

func anagram(s1, s2 string) bool {
	var characters map[rune]bool = make(map[rune]bool)
	var set1 map[rune]int = make(map[rune]int)
	for _, c := range s1 {
		set1[c]++
		characters[c] = true
	}
	var set2 map[rune]int = make(map[rune]int)
	for _, c := range s2 {
		set2[c]++
		characters[c] = true
	}
	for char := range characters {
		if set1[char] != set2[char] {
			return false
		}
	}
	return true
}
