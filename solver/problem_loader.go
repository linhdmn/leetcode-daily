package solver

import (
	"fmt"
	"regexp"
	"strconv"
)

// ProblemLoader loads problem implementations
type ProblemLoader struct {
	// No need to store implementations, we'll use factory pattern
}

// NewProblemLoader creates a new problem loader
func NewProblemLoader() *ProblemLoader {
	return &ProblemLoader{}
}

// CreateSolver creates a solver for a problem type
func (l *ProblemLoader) CreateSolver(problemType ProblemType) (Problem, error) {
	// Create a solver based on the problem type
	switch problemType {
	case "merge_array":
		return &MergeSolver{
			problemType: problemType,
			loader:      l,
		}, nil
	case "two_sum":
		return &TwoSumSolver{
			problemType: problemType,
		}, nil
	case "remove_element":
		return &RemoveElementSolver{
			problemType: problemType,
		}, nil
	// Add other problem types as they're implemented
	default:
		// Try using the generic solver for unknown problem types
		return NewGenericSolver(problemType)
	}
}

// MergeSolver is a problem solver for merge array problems
type MergeSolver struct {
	problemType ProblemType
	loader      *ProblemLoader
}

// Solve implements the Problem interface
func (s *MergeSolver) Solve(params map[string]interface{}) (interface{}, error) {
	switch s.problemType {
	case "merge_array":
		return s.solveMergeArray(params)
	default:
		return nil, fmt.Errorf("unknown problem type: %s", s.problemType)
	}
}

// ParseTestCase implements the TestCaseParser interface
func (s *MergeSolver) ParseTestCase(filePath string) (TestCase, error) {
	testCase := TestCase{
		FilePath:    filePath,
		ProblemType: s.problemType,
		InputParams: make(map[string]interface{}),
	}

	// Read input and output from file
	inputLine, outputLine, err := ReadInputAndOutput(filePath)
	if err != nil {
		return testCase, err
	}

	// Parse input parameters
	nums1Regex := `nums1\s*=\s*\[(.*?)\]`
	mRegex := `m\s*=\s*(\d+)`
	nums2Regex := `nums2\s*=\s*\[(.*?)\]`
	nRegex := `n\s*=\s*(\d+)`

	nums1Str, err := extractPattern(inputLine, nums1Regex, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract nums1: %w", err)
	}

	mStr, err := extractPattern(inputLine, mRegex, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract m: %w", err)
	}

	nums2Str, err := extractPattern(inputLine, nums2Regex, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract nums2: %w", err)
	}

	nStr, err := extractPattern(inputLine, nRegex, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract n: %w", err)
	}

	nums1 := ParseIntArray(nums1Str)
	m, _ := strconv.Atoi(mStr)
	nums2 := ParseIntArray(nums2Str)
	n, _ := strconv.Atoi(nStr)

	testCase.InputParams["nums1"] = nums1
	testCase.InputParams["m"] = m
	testCase.InputParams["nums2"] = nums2
	testCase.InputParams["n"] = n

	// Parse expected output
	outputStr, err := ExtractArrayFromBrackets(outputLine)
	if err != nil {
		return testCase, fmt.Errorf("invalid output format: %w", err)
	}

	testCase.ExpectedOutput = ParseIntArray(outputStr)

	return testCase, nil
}

// solveMergeArray handles the merge array problem
func (s *MergeSolver) solveMergeArray(params map[string]interface{}) (interface{}, error) {
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

	// Implementation of merge array
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

// TwoSumSolver is a problem solver for two sum problems
type TwoSumSolver struct {
	problemType ProblemType
}

// Solve implements the Problem interface
func (s *TwoSumSolver) Solve(params map[string]interface{}) (interface{}, error) {
	// Extract parameters
	nums, ok1 := params["nums"].([]int)
	target, ok2 := params["target"].(int)

	if !ok1 || !ok2 {
		return nil, fmt.Errorf("invalid parameters for two sum problem")
	}

	// Direct implementation to avoid import dependencies
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

// ParseTestCase implements the TestCaseParser interface
func (s *TwoSumSolver) ParseTestCase(filePath string) (TestCase, error) {
	testCase := TestCase{
		FilePath:    filePath,
		ProblemType: s.problemType,
		InputParams: make(map[string]interface{}),
	}

	// Read input and output from file
	inputLine, outputLine, err := ReadInputAndOutput(filePath)
	if err != nil {
		return testCase, err
	}

	// Parse input parameters
	numsStr, err := extractPattern(inputLine, `nums\s*=\s*\[(.*?)\]`, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract nums: %w", err)
	}

	targetStr, err := extractPattern(inputLine, `target\s*=\s*(\d+)`, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract target: %w", err)
	}

	nums := ParseIntArray(numsStr)
	target, _ := strconv.Atoi(targetStr)

	testCase.InputParams["nums"] = nums
	testCase.InputParams["target"] = target

	// Parse expected output
	outputStr, err := ExtractArrayFromBrackets(outputLine)
	if err != nil {
		return testCase, fmt.Errorf("invalid output format: %w", err)
	}

	testCase.ExpectedOutput = ParseIntArray(outputStr)

	return testCase, nil
}

// RemoveElementSolver is a problem solver for remove_element problems
type RemoveElementSolver struct {
	problemType ProblemType
}

// Solve implements the Problem interface
func (s *RemoveElementSolver) Solve(params map[string]interface{}) (interface{}, error) {
	// Extract parameters
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

// ParseTestCase implements the TestCaseParser interface
func (s *RemoveElementSolver) ParseTestCase(filePath string) (TestCase, error) {
	testCase := TestCase{
		FilePath:    filePath,
		ProblemType: s.problemType,
		InputParams: make(map[string]interface{}),
	}

	// Read input and output from file
	inputLine, outputLine, err := ReadInputAndOutput(filePath)
	if err != nil {
		return testCase, err
	}

	// Parse input parameters
	// nums = [3,2,2,3], val = 3
	numsStr, err := extractPattern(inputLine, `nums\s*=\s*\[(.*?)\]`, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract nums: %w", err)
	}

	valStr, err := extractPattern(inputLine, `val\s*=\s*(\d+)`, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract val: %w", err)
	}

	nums := ParseIntArray(numsStr)
	val, _ := strconv.Atoi(valStr)

	testCase.InputParams["nums"] = nums
	testCase.InputParams["val"] = val

	// Parse expected output
	// For remove_element, the output has two parts:
	// 1. A number k (length)
	// 2. (Optional) An array of the first k elements

	// First check if there's a number in the output
	kStr, err := extractPattern(outputLine, `^\s*(\d+)`, 1)
	if err != nil {
		return testCase, fmt.Errorf("failed to extract length (k): %w", err)
	}

	k, _ := strconv.Atoi(kStr)

	// Create the expected output map
	expectedOutput := map[string]interface{}{
		"length": k,
	}

	// Check if there's also an array in the output
	arrayMatch := regexp.MustCompile(`\[(.*?)\]`).FindStringSubmatch(outputLine)
	if len(arrayMatch) >= 2 {
		expectedArray := ParseIntArray(arrayMatch[1])
		expectedOutput["array"] = expectedArray
	}

	testCase.ExpectedOutput = expectedOutput

	return testCase, nil
}

// Helper function to extract patterns from strings
func extractPattern(s, pattern string, groupIndex int) (string, error) {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(s)
	if len(match) <= groupIndex {
		return "", fmt.Errorf("pattern '%s' not found in '%s'", pattern, s)
	}
	return match[groupIndex], nil
}
