package mathx

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	g := GCD(a, b)
	return abs(a*b) / g
}

// LCMMultiple returns the least common multiple of multiple numbers
func LCMMultiple(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return abs(numbers[0])
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = LCM(result, numbers[i])
		if result == 0 {
			return 0
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
