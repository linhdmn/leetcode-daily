package solver

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// TestCase represents a generic test case
type TestCase struct {
	FilePath       string
	ProblemType    ProblemType
	InputParams    map[string]interface{}
	ExpectedOutput interface{}
}

// TestCaseParser is the interface for problem-specific test case parsers
type TestCaseParser interface {
	// ParseTestCase parses a test file into a TestCase structure
	ParseTestCase(filePath string) (TestCase, error)
}

// ProblemWithParser is an extended interface that includes test case parsing
type ProblemWithParser interface {
	Problem
	TestCaseParser
}

// ReadInputAndOutput reads input and output lines from a test file
func ReadInputAndOutput(filePath string) (string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var inputLine, outputLine string
	scanner := bufio.NewScanner(file)

	// Read input line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Input:") {
			inputLine = strings.TrimPrefix(line, "Input:")
			break
		}
	}

	// Read output line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Output:") {
			outputLine = strings.TrimPrefix(line, "Output:")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("error reading file: %w", err)
	}

	return inputLine, outputLine, nil
}

// ParseIntArray converts a comma-separated string of integers to an int slice
func ParseIntArray(s string) []int {
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Printf("Warning: could not parse '%s' as integer, skipping", part)
			continue
		}
		result = append(result, num)
	}

	return result
}

// ExtractArrayFromBrackets extracts an array from a string with the format [...].
// It returns the array as a string, without the brackets.
func ExtractArrayFromBrackets(s string) (string, error) {
	re := regexp.MustCompile(`\[(.*?)\]`)
	match := re.FindStringSubmatch(s)
	if len(match) < 2 {
		return "", fmt.Errorf("no array found in string: %s", s)
	}
	return match[1], nil
}

// ExtractIntValue extracts an integer value from a string parameter (e.g., "target = 9")
func ExtractIntValue(s, paramName string) (int, error) {
	re := regexp.MustCompile(paramName + `\s*=\s*(\d+)`)
	match := re.FindStringSubmatch(s)
	if len(match) < 2 {
		return 0, fmt.Errorf("no %s value found in: %s", paramName, s)
	}
	return strconv.Atoi(match[1])
}
