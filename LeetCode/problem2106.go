func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	result int 
	for _, fruit := range fruits {
		if fruit[0] >= startPos && fruit[0] < startPos + k - 1 {
			result += fruit[1]
		}
	}
	return result
}
