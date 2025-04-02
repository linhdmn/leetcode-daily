package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"leetcodedaily/solver"
)

// Create a registry instance and auto-register problems
var registry = solver.NewRegistry()

func init() {
	// Auto-register all available problem solvers
	registry.AutoRegister()
}

func main() {
	// Check if a specific problem was requested
	problemType := ""
	if len(os.Args) > 1 {
		problemType = os.Args[1]
	}

	// Find all problem directories
	var problemDirs []string
	if problemType != "" {
		// Test only the specified problem
		problemDirs = []string{filepath.Join("test_cases", problemType)}
	} else {
		// Find all problem directories
		entries, err := os.ReadDir("test_cases")
		if err != nil {
			log.Fatalf("Failed to read test_cases directory: %v", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				problemDirs = append(problemDirs, filepath.Join("test_cases", entry.Name()))
			}
		}
	}

	if len(problemDirs) == 0 {
		log.Fatal("No problem directories found in test_cases/")
	}

	// Run tests for each problem
	totalPassed := 0
	totalFailed := 0

	for _, problemDir := range problemDirs {
		problem := filepath.Base(problemDir)
		fmt.Printf("\n=== Testing Problem: %s ===\n\n", problem)

		// Find all test files for this problem
		testFiles, err := filepath.Glob(filepath.Join(problemDir, "*.txt"))
		if err != nil {
			log.Printf("Error finding test files for %s: %v", problem, err)
			continue
		}

		if len(testFiles) == 0 {
			log.Printf("No test files found for problem %s", problem)
			continue
		}

		// Run tests for this problem
		passed, failed := runTestsForProblem(solver.ProblemType(problem), testFiles)
		totalPassed += passed
		totalFailed += failed
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total: %d tests\n", totalPassed+totalFailed)
	fmt.Printf("Passed: %d tests\n", totalPassed)
	fmt.Printf("Failed: %d tests\n", totalFailed)
}

// runTestsForProblem runs all tests for a given problem
func runTestsForProblem(problemType solver.ProblemType, testFiles []string) (int, int) {
	passed := 0
	failed := 0

	// Get the solver for this problem
	problemSolver, exists := registry.Get(problemType)
	if !exists {
		log.Printf("No solver registered for problem type: %s", problemType)
		for _, file := range testFiles {
			log.Printf("⚠️ SKIP: %s (no solver available)", filepath.Base(file))
			failed++
		}
		return passed, failed
	}

	// Check if the solver implements the TestCaseParser interface
	parserSolver, ok := problemSolver.(solver.TestCaseParser)
	if !ok {
		log.Printf("Solver for %s does not implement TestCaseParser", problemType)
		for _, file := range testFiles {
			log.Printf("⚠️ SKIP: %s (no parser available)", filepath.Base(file))
			failed++
		}
		return passed, failed
	}

	for _, testFile := range testFiles {
		// Use the solver's own parser to parse the test case
		testCase, err := parserSolver.ParseTestCase(testFile)
		if err != nil {
			log.Printf("Error parsing test file %s: %v", testFile, err)
			failed++
			continue
		}

		// Run the specific test
		result := runTest(problemSolver, testCase)
		if result {
			passed++
			fmt.Printf("✅ PASS: %s\n", filepath.Base(testFile))
		} else {
			failed++
			fmt.Printf("❌ FAIL: %s\n", filepath.Base(testFile))
		}
	}

	fmt.Printf("\nResults for %s: %d passed, %d failed\n", problemType, passed, failed)
	return passed, failed
}

// runTest runs a specific test case with the given solver
func runTest(problemSolver solver.Problem, testCase solver.TestCase) bool {
	// Call the solver
	result, err := problemSolver.Solve(testCase.InputParams)
	if err != nil {
		log.Printf("Error solving problem: %v", err)
		return false
	}

	// Compare with expected output
	expected := testCase.ExpectedOutput

	// Special handling for remove_element problem which returns a map
	if testCase.ProblemType == "remove_element" {
		// For remove_element, we need to compare the length and array separately
		resultMap, ok := result.(map[string]interface{})
		if !ok {
			log.Printf("Error: expected result to be a map, got %T", result)
			return false
		}

		expectedMap, ok := expected.(map[string]interface{})
		if !ok {
			log.Printf("Error: expected output to be a map, got %T", expected)
			return false
		}

		// Compare length
		resultLength, ok1 := resultMap["length"].(int)
		expectedLength, ok2 := expectedMap["length"].(int)

		if !ok1 || !ok2 || resultLength != expectedLength {
			fmt.Printf("   Expected length: %v\n   Got length:      %v ❌\n",
				expectedMap["length"], resultMap["length"])
			return false
		}

		// Compare array if it exists in expected output
		if expectedArray, hasArray := expectedMap["array"]; hasArray {
			resultArray, _ := resultMap["array"].([]int)
			isArrayEqual := reflect.DeepEqual(resultArray, expectedArray)

			if !isArrayEqual {
				fmt.Printf("   Expected array: %v\n   Got array:      %v ❌\n",
					expectedArray, resultMap["array"])
				return false
			}

			fmt.Printf("   Expected length: %v\n   Got length:      %v\n",
				expectedLength, resultLength)
			fmt.Printf("   Expected array: %v\n   Got array:      %v\n",
				expectedArray, resultMap["array"])
			return true
		}

		// If no array in expected, just display length
		fmt.Printf("   Expected length: %v\n   Got length:      %v\n",
			expectedLength, resultLength)
		fmt.Printf("   Got array:      %v\n", resultMap["array"])
		return true
	}

	// Default comparison for other problems
	isEqual := reflect.DeepEqual(result, expected)
	if isEqual {
		fmt.Printf("   Expected: %v\n   Got:      %v\n", expected, result)
	} else {
		fmt.Printf("   Expected: %v\n   Got:      %v ❌\n", expected, result)
	}

	return isEqual
}
