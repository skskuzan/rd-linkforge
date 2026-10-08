# AGENTS.md

## General Principles
- All active Git manipulations (branching, merging, pushing) must be performed with direct human participation or explicit user confirmation.
- Pull requests and branch merges require explicit human verification and approval.
- Before any commit, run necessary code quality checks and linters (e.g., `golangci-lint`, `go fmt`, `go vet`).
- Maintain project standards and follow standard Git Flow practices.

## Git Flow Guidelines
- **Main Branch:** Protected branch, reserved for production-ready code.
- **Develop Branch:** The main integration branch for features.
- **Feature Branches:** Created from `develop` using the prefix `feature/`.
- **Bugfix Branches:** Created from `develop` using the prefix `bugfix/`.
- **Commits:** Follow conventional commit messages.
- **Workflow:**
  1. Create a feature/bugfix branch.
  2. Implement changes.
  3. Run quality checks (linters/tests).
  4. Human reviews and approves the changes.
  5. Merge to `develop` only after explicit human consent.
