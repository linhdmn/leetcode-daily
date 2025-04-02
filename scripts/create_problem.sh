#!/bin/bash

# Helper script to create a new problem implementation
# Usage: ./scripts/create_problem.sh problem_name

if [ $# -ne 1 ]; then
    echo "Usage: $0 problem_name"
    echo "Example: $0 three_sum"
    exit 1
fi

PROBLEM_NAME=$1
PROBLEMS_DIR="problems"
TEST_CASES_DIR="test_cases"

# Create problem directory
mkdir -p "$PROBLEMS_DIR/$PROBLEM_NAME"
mkdir -p "$TEST_CASES_DIR/$PROBLEM_NAME"

# Generate problem implementation file
cat > "$PROBLEMS_DIR/$PROBLEM_NAME/${PROBLEM_NAME}.go" << EOF
package ${PROBLEM_NAME}

// TODO: Implement the solution here
// For example:
// 
// func TwoSum(nums []int, target int) []int {
//     // ...implementation...
// }
EOF

echo "Created problem template at $PROBLEMS_DIR/$PROBLEM_NAME/${PROBLEM_NAME}.go"
echo "Created test cases directory at $TEST_CASES_DIR/$PROBLEM_NAME"
echo ""
echo "Next steps:"
echo "1. Implement your solution in $PROBLEMS_DIR/$PROBLEM_NAME/${PROBLEM_NAME}.go"
echo "2. Add test cases in $TEST_CASES_DIR/$PROBLEM_NAME/ following the existing format"
echo "3. Run the tests with: go run main.go $PROBLEM_NAME" 