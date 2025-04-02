package merge_array

// Merge merges nums2 into nums1 in-place
func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	lastNum1 := m - 1
	lastNum2 := n - 1
	lastMerge := m + n - 1

	for lastNum2 >= 0 {
		if lastNum1 >= 0 && nums1[lastNum1] > nums2[lastNum2] {
			nums1[lastMerge] = nums1[lastNum1]
			lastNum1--
		} else {
			nums1[lastMerge] = nums2[lastNum2]
			lastNum2--
		}
		lastMerge--
	}
}
