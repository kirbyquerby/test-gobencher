package exhibit

// Counts the number of ones in an unsigned integer.
// Similar to POPCOUNT, or commonly known as Hamming distance.
func countOnes(in uint64) (count int) {
	if in == 0 {
		return 0
	}
	if in&(in-1) == 0 {
		return 1
	}
	
	fib(30)

	for {
		if in == 0 {
			return
		}
		if in&1 == 1 {
			count++
		}
		in >>= 1
	}
	return
}

func fib(n uint64) int {
	if n < 0 {
		panic(n)
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}
