package tax

import (
	"fmt"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		fmt.Println("Olá")
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
