package solver

import (
	"fmt"
)

// ReflectiveSolver uses reflection to call problem-specific functions
type ReflectiveSolver struct {
	problemType ProblemType
}

// NewReflectiveSolver creates a new reflective solver for the given problem type
func NewReflectiveSolver(problemType ProblemType) *ReflectiveSolver {
	return &ReflectiveSolver{
		problemType: problemType,
	}
}

// Solve implements the Problem interface
func (s *ReflectiveSolver) Solve(params map[string]interface{}) (interface{}, error) {
	// Based on the problem type, call the appropriate handler
	switch s.problemType {
	case "merge_array":
		return s.handleMergeArray(params)
	case "two_sum":
		return s.handleTwoSum(params)
	case "remove_element":
		return s.handleRemoveElement(params)
	default:
		return nil, fmt.Errorf("unknown problem type: %s", s.problemType)
	}
}

// handleMergeArray handles the merge array problem
func (s *ReflectiveSolver) handleMergeArray(params map[string]interface{}) (interface{}, error) {
	// Extract parameters
	nums1, ok1 := params["nums1"].([]int)
	m, ok2 := params["m"].(int)
	nums2, ok3 := params["nums2"].([]int)
	n, ok4 := params["n"].(int)

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("invalid parameters for merge array problem")
	}

	// Create a copy of nums1 to avoid modifying the original
	result := make([]int, len(nums1))
	copy(result, nums1)

	// Use direct implementation
	if n == 0 {
		return result, nil
	}

	lastNum1 := m - 1
	lastNum2 := n - 1
	lastMerge := m + n - 1

	for lastNum2 >= 0 {
		if lastNum1 >= 0 && nums1[lastNum1] > nums2[lastNum2] {
			result[lastMerge] = nums1[lastNum1]
			lastNum1--
		} else {
			result[lastMerge] = nums2[lastNum2]
			lastNum2--
		}
		lastMerge--
	}

	return result, nil
}

// handleTwoSum handles the two sum problem
func (s *ReflectiveSolver) handleTwoSum(params map[string]interface{}) (interface{}, error) {
	// We'll implement this later
	return []int{0, 0}, fmt.Errorf("two sum implementation not available")
}

// handleRemoveElement handles the remove element problem
func (s *ReflectiveSolver) handleRemoveElement(params map[string]interface{}) (interface{}, error) {
	// We'll implement this later
	result := map[string]interface{}{
		"length": 0,
		"array":  []int{},
	}

	return result, fmt.Errorf("remove element implementation not available")
}
