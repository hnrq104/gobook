package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three", "", "one", "henrqie", "naju", "four"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	fmt.Printf("%q\n", nodups(data))
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, v := range strings {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

//exercise 4.5

func nodups(strings []string) []string {
	m := make(map[string]bool)
	i := 0
	for _, s := range strings {
		if !m[s] {
			strings[i] = s
			m[s] = true
			i++
		}
	}
	return strings[:i]
}
