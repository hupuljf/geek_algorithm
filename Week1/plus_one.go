package Week1

func plusOne(digits []int) []int {
	//只有全为9才会有进一的情况
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	//如果首位为0 代表全为9
	if digits[0] == 0 {
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	}
	return digits
}
