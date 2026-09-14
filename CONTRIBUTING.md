# Contributing to pkg-responses

Thank you for your interest in contributing!

## Prerequisites

- Go 1.25 or later
- `golangci-lint` (optional but recommended)

## Quick Start

```bash
git clone https://github.com/mawarpay/pkg-responses.git
cd pkg-responses

make test    # race + coverage
make check   # fmt-check, vet, build, test
make lint    # golangci-lint via Docker (optional)
```

## Development Workflow

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/my-feature`
3. Make your changes
4. Add or update tests for new behavior
5. Run `make check` and `make lint` (or `golangci-lint run`)
6. Commit with a descriptive message
7. Push and open a Pull Request

CI runs on every PR: tests (`-race`, Go 1.25+), lint, and security scans (`govulncheck`, gosec, CodeQL, Bearer). Tag `v*` triggers a library release via GoReleaser (changelog only).

## Code Guidelines

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Add godoc comments to all exported symbols
- Prefer table-driven tests
- Keep this package free of business logic and domain imports
- Prefer `Write*` helpers over long multi-arg wrappers for new APIs

## Reporting Issues

Use [GitHub Issues](https://github.com/mawarpay/pkg-responses/issues). Include:

- Go version (`go version`)
- OS and architecture
- Steps to reproduce
- Expected vs actual behavior
