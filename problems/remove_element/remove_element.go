package remove_element

// RemoveElement removes all instances of val from nums and returns the new length
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// Two-pointer approach
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		// If the current element is not the value to remove
		if nums[fast] != val {
			// Copy it to the slow pointer position
			nums[slow] = nums[fast]
			// Move slow pointer forward
			slow++
		}
		// If it is the value to remove, just move the fast pointer forward
	}

	return slow // new length of array
}

// TODO: Implement the solution here
// For example:
//
// func TwoSum(nums []int, target int) []int {
//     // ...implementation...
// }
