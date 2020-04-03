package utils

// Sum is ...
func Sum(firstNumber float32, secondNumber float32) float32 {
	return firstNumber + secondNumber
}

// Difference is ...
func Difference(firstNumber float32, secondNumber float32) float32 {
	return firstNumber - secondNumber
}

// Product is ...
func Product(firstNumber float32, secondNumber float32) float32 {
	return firstNumber * secondNumber
}

// Quotient is ...
func Quotient(firstNumber float32, secondNumber float32) float32 {
	return firstNumber / secondNumber
}

// Remainder is ...
func Remainder(firstNumber float32, secondNumber float32) float32 {
	result := int(firstNumber) % int(secondNumber)
	return float32(result)
}

// ClearValue is ...
func ClearValue() float32 {
	return 0
}
