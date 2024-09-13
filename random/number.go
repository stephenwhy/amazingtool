package random

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateLuhn generates a Luhn valid number of specified length
func GenerateLuhn(length int) string {
	if length < 2 {
		return ""
	}

	rand.NewSource(time.Now().UnixNano())
	digits := make([]int, length-1)

	// Make sure the first digit is not 0
	digits[0] = rand.Intn(9) + 1

	for i := 1; i < length-1; i++ {
		digits[i] = rand.Intn(10)
	}

	// Calculate the check digit
	checkDigit := calculateLuhn(digits)
	digits = append(digits, checkDigit)

	// Convert digits to a string
	var result string
	for _, digit := range digits {
		result += fmt.Sprintf("%d", digit)
	}

	return result
}

// calculateLuhn calculates the Luhn check digit for a list of integers
func calculateLuhn(digits []int) int {
	sum := 0
	double := len(digits)%2 == 0

	for _, digit := range digits {
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return (10 - (sum % 10)) % 10
}
