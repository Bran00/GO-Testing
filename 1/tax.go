// Package tax provides tax calculation utilities.
package tax

import "fmt"

func CalculateTax(amount float64) float64 {
	fmt.Println("Hello")
	if amount >= 1000 {
		return 10.0
	}

	return 5.0
}
