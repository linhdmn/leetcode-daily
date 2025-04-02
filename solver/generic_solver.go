package solver

import (
	"fmt"
)

// GenericSolver is a generic implementation of Problem that adapts various problem types
type GenericSolver struct {
	problemType ProblemType
}

// NewGenericSolver creates a new generic solver for the given problem type
func NewGenericSolver(problemType ProblemType) (*GenericSolver, error) {
	return &GenericSolver{
		problemType: problemType,
	}, nil
}

// Solve implements the Problem interface
func (gs *GenericSolver) Solve(params map[string]interface{}) (interface{}, error) {
	// Call the function with the appropriate parameters
	// This will depend on the specific problem type
	switch gs.problemType {
	case "merge_array":
		return gs.solveMergeArray(params)
	case "two_sum":
		return gs.solveTwoSum(params)
	case "remove_element":
		return gs.solveRemoveElement(params)
	default:
		return nil, fmt.Errorf("unknown problem type: %s", gs.problemType)
	}
}

// ParseTestCase implements the TestCaseParser interface
func (gs *GenericSolver) ParseTestCase(filePath string) (TestCase, error) {
	// Create a temporary solver for the problem type and use its parser
	switch gs.problemType {
	case "merge_array":
		solver := &MergeSolver{problemType: gs.problemType}
		return solver.ParseTestCase(filePath)
	case "two_sum":
		solver := &TwoSumSolver{problemType: gs.problemType}
		return solver.ParseTestCase(filePath)
	case "remove_element":
		solver := &RemoveElementSolver{problemType: gs.problemType}
		return solver.ParseTestCase(filePath)
	default:
		return TestCase{}, fmt.Errorf("no parser available for problem type: %s", gs.problemType)
	}
}

// solveMergeArray solves the merge array problem
func (gs *GenericSolver) solveMergeArray(params map[string]interface{}) (interface{}, error) {
	nums1, ok1 := params["nums1"].([]int)
	m, ok2 := params["m"].(int)
	nums2, ok3 := params["nums2"].([]int)
	n, ok4 := params["n"].(int)

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("invalid parameters for merge array problem")
	}

	// Make a copy of nums1 so we don't modify the original
	result := make([]int, len(nums1))
	copy(result, nums1)

	// Use a direct implementation instead of reflection
	if n == 0 {
		return result, nil
	}

	last_num1 := m - 1
	last_num2 := n - 1
	last_merge := m + n - 1

	for last_num2 >= 0 {
		if last_num1 >= 0 && nums1[last_num1] > nums2[last_num2] {
			result[last_merge] = nums1[last_num1]
			last_num1--
		} else {
			result[last_merge] = nums2[last_num2]
			last_num2--
		}
		last_merge--
	}

	return result, nil
}

// solveTwoSum solves the two sum problem
func (gs *GenericSolver) solveTwoSum(params map[string]interface{}) (interface{}, error) {
	nums, ok1 := params["nums"].([]int)
	target, ok2 := params["target"].(int)

	if !ok1 || !ok2 {
		return nil, fmt.Errorf("invalid parameters for two sum problem")
	}

	// Implementation of two sum
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, found := numMap[complement]; found {
			return []int{j, i}, nil
		}

		numMap[num] = i
	}

	return []int{-1, -1}, nil
}

// solveRemoveElement solves the remove element problem
func (gs *GenericSolver) solveRemoveElement(params map[string]interface{}) (interface{}, error) {
	nums, ok1 := params["nums"].([]int)
	val, ok2 := params["val"].(int)

	if !ok1 || !ok2 {
		return nil, fmt.Errorf("invalid parameters for remove element problem")
	}

	// Make a copy of nums to avoid modifying the original
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	// Two-pointer approach
	slow := 0
	for fast := 0; fast < len(numsCopy); fast++ {
		if numsCopy[fast] != val {
			numsCopy[slow] = numsCopy[fast]
			slow++
		}
	}

	// Prepare the result expected by the test framework
	result := map[string]interface{}{
		"length": slow,
		"array":  numsCopy[:slow],
	}

	return result, nil
}
