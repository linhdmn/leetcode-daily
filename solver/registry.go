package solver

import (
	"log"
)

// ProblemType represents the type of problem to solve
type ProblemType string

// Problem is the interface that all problem solvers must implement
type Problem interface {
	// Solve executes the problem logic with the provided input parameters
	// and returns the result and any error that occurred
	Solve(params map[string]interface{}) (interface{}, error)
}

// Registry maintains a mapping of problem types to their solvers
type Registry struct {
	solvers          map[ProblemType]Problem
	loader           *ProblemLoader
	problemDiscovery *ProblemDiscovery
}

// NewRegistry creates a new registry instance
func NewRegistry() *Registry {
	registry := &Registry{
		solvers: make(map[ProblemType]Problem),
	}

	// Create a loader
	registry.loader = NewProblemLoader()

	// Create a problem discovery instance
	registry.problemDiscovery = NewProblemDiscovery(registry, registry.loader)

	return registry
}

// Register adds a problem solver to the registry
func (r *Registry) Register(problemType ProblemType, solver Problem) {
	r.solvers[problemType] = solver
}

// Get retrieves a problem solver from the registry
func (r *Registry) Get(problemType ProblemType) (Problem, bool) {
	solver, exists := r.solvers[problemType]
	return solver, exists
}

// ListRegisteredProblems returns a list of all registered problem types
func (r *Registry) ListRegisteredProblems() []ProblemType {
	problems := make([]ProblemType, 0, len(r.solvers))
	for problemType := range r.solvers {
		problems = append(problems, problemType)
	}
	return problems
}

// AutoRegister automatically finds and registers all available problem solvers
func (r *Registry) AutoRegister() {
	// Use the problem discovery to automatically register all solvers
	log.Println("Auto-registering problem solvers...")
	r.problemDiscovery.AutoRegisterSolvers()
}
