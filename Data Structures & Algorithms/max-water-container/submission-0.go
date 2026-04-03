func maxArea(heights []int) int {
	i := 0;
	j := len(heights) - 1
	maxArea := -1
	for {
		if i >= j {
			break
		}
		maxArea = max(maxArea, area(i, j, heights))
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func area(i, j int, heights []int) int {
	return (j - i) * min(heights[i], heights[j])
}
