package main

import (
	"fmt"
	"math"
	"time"
)

func divSum(n int) int {
	sum := 0
	if n == 1 {
		return 0
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == (n / i) {
				sum += i
			} else {
				sum += i + n/i
			}
		}
	}

	return sum + 1
}

func isAbundant(n int) bool {
	return divSum(n) > n
}

func main() {
	start := time.Now()

	upper := 28123
	abundantNums := []int{}
	set := make(map[int]bool)

	// fill abundant numbers
	for n := 1; n < upper; n++ {
		if isAbundant(n) {
			abundantNums = append(abundantNums, n)
		}
	}

	// get sums less than upper bound
	for _, i := range abundantNums {
		for _, j := range abundantNums {
			sum := i + j
			if sum > upper {
				break
			}
			if !set[sum] {
				set[sum] = true
			}
		}
	}

	sum := 0

	// aggregate non sums
	for n := 1; n < upper; n++ {
		if !set[n] {
			sum += n
		}
	}

	fmt.Println("The result is", sum)

	elapsed := time.Since(start)
	fmt.Printf("Executed in %s", elapsed)
}
