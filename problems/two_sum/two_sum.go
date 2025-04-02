package two_sum

// TwoSum finds two numbers in the array that add up to the target
// and returns their indices
func TwoSum(nums []int, target int) []int {
	// Create a map to store values and their indices
	numMap := make(map[int]int)

	// Go through the array once
	for i, num := range nums {
		// Calculate the complement needed
		complement := target - num

		// Check if the complement exists in our map
		if j, found := numMap[complement]; found {
			// Return the indices
			return []int{j, i}
		}

		// Store current number and its index
		numMap[num] = i
	}

	// No solution found
	return []int{-1, -1}
}
