package main

func fullAdder(a, b int) (int, int) {
	sum := a + b
	return sum % 10, sum / 10
}

func plusOne(digits []int) []int {
	sum := make([]int, len(digits))
	res, carry := 0, 1
	for i := len(digits) - 1; i >= 0; i-- {
		res, carry = fullAdder(digits[i], carry)
		sum[i] = res
	}
	if carry != 0 {
		sum = append([]int{carry}, sum...)
	}
	return sum
}
