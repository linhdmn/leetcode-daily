package solver

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ProblemDiscovery discovers problems in the codebase
type ProblemDiscovery struct {
	registry *Registry
	loader   *ProblemLoader
}

// NewProblemDiscovery creates a new problem discovery instance
func NewProblemDiscovery(registry *Registry, loader *ProblemLoader) *ProblemDiscovery {
	return &ProblemDiscovery{
		registry: registry,
		loader:   loader,
	}
}

// DiscoverProblems finds all problems in the problems directory and registers them
func (pd *ProblemDiscovery) DiscoverProblems() {
	// Check if problems directory exists
	if _, err := os.Stat("problems"); os.IsNotExist(err) {
		log.Println("Problems directory not found")
		return
	}

	// Walk through all problem directories
	err := filepath.WalkDir("problems", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root problems directory and non-directories
		if path == "problems" || !d.IsDir() {
			return nil
		}

		// Skip subdirectories of problem directories
		if strings.Count(path, string(os.PathSeparator)) > 1 {
			return filepath.SkipDir
		}

		// Get the problem type from the directory name
		problemType := ProblemType(filepath.Base(path))

		// Check if this is a valid problem directory
		mainFile := filepath.Join(path, fmt.Sprintf("%s.go", problemType))
		if _, err := os.Stat(mainFile); os.IsNotExist(err) {
			log.Printf("Main file not found for problem %s at %s", problemType, mainFile)
			return nil
		}

		// Create and register a solver for this problem
		solver, err := pd.loader.CreateSolver(problemType)
		if err != nil {
			log.Printf("Warning: Could not create solver for %s: %v", problemType, err)
			return nil
		}

		// Register the solver
		pd.registry.Register(problemType, solver)
		log.Printf("Registered solver for problem: %s", problemType)

		return nil
	})

	if err != nil {
		log.Printf("Error walking through problems directory: %v", err)
	}
}

// AutoRegisterSolvers is a helper method to register all known solvers
func (pd *ProblemDiscovery) AutoRegisterSolvers() {
	// First, discover problems in the problems directory
	pd.DiscoverProblems()

	// You could add additional registration methods here, e.g., for built-in solvers
}
