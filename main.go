package main

import (
	"fmt"
	"strconv"
)

// BiggestPairNumber finds the biggest pair of digits with the largest sum.
func BiggestPairNumber(numbers int) int {
	// Convert the number to string to access each digit
	numStr := strconv.Itoa(numbers)
	maxSum := 0
	maxPair := 0

	// Loop through the number string and get every pair of digits
	for i := 0; i < len(numStr)-1; i++ {
		// Ambil dua digit berurutan dari string
		pairStr := numStr[i:i+2]
		// Konversi menjadi dua digit
		digit1, _ := strconv.Atoi(string(pairStr[0]))
		digit2, _ := strconv.Atoi(string(pairStr[1]))
		sum := digit1 + digit2

		// Update maxPair jika jumlah saat ini lebih besar
		if sum > maxSum {
			maxSum = sum
			maxPair, _ = strconv.Atoi(pairStr) 
		}
	}

	return maxPair
}

func main() {
	fmt.Println(BiggestPairNumber(11223344)) // Output: 44
	fmt.Println(BiggestPairNumber(89083278)) // Output: 89
}
