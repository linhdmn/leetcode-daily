# LeetCode Daily

A modular framework for solving and testing LeetCode problems with automatic test case verification and problem discovery.

## Features

- **Automatic Problem Discovery** - New problem implementations are automatically discovered and registered
- **Modular Design** - Each problem is isolated in its own package
- **Extensible Architecture** - Easy to add new problem types and test parsers
- **Test Framework** - Run tests for specific problems or all problems at once
- **Interface-Based** - Clean separation of concerns through well-defined interfaces

## Getting Started

1. Create a new problem implementation using the helper script:

```bash
./scripts/create_problem.sh problem_name
```

For example:

```bash
./scripts/create_problem.sh two_sum
```

2. Implement the solution in the generated file. For example, for `two_sum`:

```go
package two_sum

// TwoSum finds two numbers in the array that add up to the target
func TwoSum(nums []int, target int) []int {
    // Your implementation here
    numMap := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        if j, found := numMap[complement]; found {
            return []int{j, i}
        }
        numMap[num] = i
    }
    
    return []int{-1, -1}
}
```

3. Create test cases in the `test_cases/problem_name/` directory. Each test case should be in a separate `.txt` file with the following format:

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
```

4. Run the tests for your problem:

```bash
go run main.go problem_name
```

For example:

```bash
go run main.go two_sum
```

5. To run tests for all problems:

```bash
go run main.go
```

## Project Structure

```
leetcode_daily/
├── main.go                  # Main test runner
├── go.mod                   # Go module file
├── README.md                # This file
├── scripts/                 # Helper scripts
│   └── create_problem.sh    # Script to create new problems
├── problems/                # Problem implementations
│   ├── merge_array/         # Example problem implementation
│   │   └── merge_array.go   # Implementation file
│   ├── two_sum/             # Example problem implementation
│   │   └── two_sum.go       # Implementation file
│   └── ...
├── solver/                  # Solver framework
│   ├── registry.go          # Registry of problem solvers
│   ├── problem_loader.go    # Problem implementation loader
│   ├── problem_discovery.go # Automatic problem discovery
│   ├── generic_solver.go    # Generic problem solver
│   ├── testcase.go          # Test case parsing utilities
│   └── ...
└── test_cases/              # Test cases for each problem
    ├── merge_array/         # Test cases for merge_array problem
    │   ├── test1.txt        # Example test case
    │   └── ...
    ├── two_sum/             # Test cases for two_sum problem
    │   ├── test1.txt        # Example test case
    │   └── ...
    └── ...
```

## Architecture

The framework uses a modular architecture with the following components:

1. **Problem Interface**: All problem solvers implement the `Problem` interface with a `Solve` method
2. **TestCaseParser Interface**: Problem solvers implement the `TestCaseParser` interface to parse test cases
3. **Registry**: Keeps track of all available problem solvers
4. **ProblemDiscovery**: Automatically discovers problem implementations in the problems directory
5. **ProblemLoader**: Creates appropriate solvers for each problem type
6. **Main Runner**: Finds test cases and runs them against the appropriate solver

## Adding a New Problem

1. Use the script to create a new problem:
   ```bash
   ./scripts/create_problem.sh new_problem_name
   ```

2. Implement the solution in the generated file:
   ```go
   package new_problem_name
   
   // Implement your solution function with an appropriate name
   func SolutionFunction(params...) returnType {
       // Your implementation here
   }
   ```

3. Create test cases in the `test_cases/new_problem_name/` directory following the LeetCode format:
   ```
   Input: param1 = value1, param2 = value2
   Output: expected_result
   ```

4. Run the tests to verify your solution:
   ```bash
   go run main.go new_problem_name
   ```

## Supported Problem Types

Currently, the framework includes examples for the following problem types:

- `merge_array`: Merge two sorted arrays
- `two_sum`: Find two numbers that add up to a target
- `remove_element`: Remove elements from an array

All new problems will be automatically detected and registered as long as:

1. They follow the directory structure: `problems/problem_name/problem_name.go`
2. Test cases follow the format: `test_cases/problem_name/*.txt`

## Extending the Framework

For special problem types that need custom test case parsing:

1. Update the `solver/problem_loader.go` file to add a specific solver implementation
2. Implement the `Problem` and `TestCaseParser` interfaces for your new problem type
3. The framework will automatically register and use your new solver 