package main

import "fmt"

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// there is room to grow. Extend the slice
		z = x[:zlen]
	} else {
		// there is insufficient space, Allocate new array
		// grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap) //works but it's bad, if y is always bigger than x
		copy(z, x)                  //built in copy function
	}
	copy(z[len(x):], y)
	return z
}

type IntSlice struct {
	ptr      *int
	len, cap int
}

func main() {
	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)
}
