package remove_element

// RemoveElement removes all instances of val from nums and returns the new length
func RemoveElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
