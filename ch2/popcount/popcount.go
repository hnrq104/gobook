package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the per-
formance of the two versions. (Section 11.4 shows how to compare the performance of differ-
ent implementations systematically.) */

func PopCountLoop(x uint64) int {
	var pop int
	for i := 0; i < 8; i++ {
		pop += int(pc[byte(x>>(i*8))])
	}
	return pop
}

/*Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
bit positions, testing the rightmost bit each time. Compare its performance to the table-
lookup version.*/

func PopCountShift(x uint64) int {
	var pop int
	// shifted := x
	for i := 0; i < 64; i++ {
		pop += int(x & 1)
		x >>= 1
	}
	return pop
}

/*Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
of PopCount that counts bits by using this fact, and assess its performance.*/

func PopCountClears(x uint64) int {
	var pop int
	// cleared := x
	for x != 0 {
		pop++
		x = x & (x - 1)
	}
	return pop
}
