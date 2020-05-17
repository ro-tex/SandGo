# What is this?

This is a list of things not to forget to add to a new project.

# Things
Start with a Makefile with handy scripts to make you life easier.

## Linting and static code analysis
- go fmt and/or go imports
- spell check
- golangci-lint
- go vet
- golang.org/x/tools/go/analysis/multichecker with and without `-lockcheck`

## Testing
- run your tests
- tun them with `-race`

## Benchmarking
- write some base performance benchmarks, `testing.B` is enough
Ï
