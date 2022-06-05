package Week1

//dp解法
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	heights := make([]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, getMaximal(heights))
	}
	return ans
}

func getMaximal(heights []int) int {
	n := len(heights)
	if n < 1 {
		return 0
	}
	if n < 2 {
		return heights[0]
	}
	minIdx := 0
	for i := 1; i < n; i++ {
		if heights[i] < heights[minIdx] {
			minIdx = i
		}
	}
	return max(heights[minIdx]*n, max(getMaximal(heights[:minIdx]), getMaximal(heights[minIdx+1:])))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
