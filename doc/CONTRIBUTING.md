# Contributing to clix

Thank you for your interest in contributing to clix! This CLI tool is built using [Cobra](https://github.com/spf13/cobra) and written in Go. Follow these steps to get started:

## Code of Conduct

Please review and adhere to our [Code of Conduct](CODE_OF_CONDUCT.md).

## Prerequisites

- Go >= 1.20
- Git
- [Cobra CLI](https://github.com/spf13/cobra) (optional for generating commands)

## Contribution Guidelines

1. **Fork the repository**
2. **Clone your fork**
   ```bash
   git clone https://github.com/your-username/project-name.git
   cd project-name
   ```
3. **Create your feature branch.**
  ```bash
  git checkout -b my-new-feature
  ```
4. **Install dependencies**
   ```bash
   go mod tidy
   ```
5. **Build the project**
   ```bash
   go build -o mycli
   ```
6. **Run tests**
   ```bash
   go test ./...
   ```
7. **Add your changes to staging**
    ```bash
    git add .
    ```
8. **Commit your changes** 
    Write clear and concise commit messages. Use the following format: `feat(command): Add new feature` or `fix(config): Fix bug in config command`.
    ```bash
    git commit -m 'Add some feature'
    ```
9. **Push to your branch**
    ```bash
    git push origin my-new-feature
    ```
10. **Create a new pull request**

## Adding a New Command

1. Use Cobra to generate a command:
   ```bash
   cobra add mycommand
   ```
2. Implement logic in `cmd/mycommand.go`
3. Register any flags and validate input
4. Add tests for your command under `/cmd` or `/internal`

## Code Style

- Use `go fmt ./...` to format your code
- Lint with `golangci-lint run` if configured
- Follow idiomatic Go practices

## Pull Requests

- Create a feature branch:
  ```bash
  git checkout -b feature/your-feature
  ```
- Make your changes and commit them with clear messages
- Push to your fork and open a PR against `main`
- Ensure tests pass and code is reviewed

## Issues and Features

- Check existing issues before creating a new one
- Use labels if applicable (bug, enhancement, question)
- If you believe you have found a bug, please provide detailed steps for reproduction, software version and anything else you believe will be useful to help troubleshoot it (e.g. OS environment, environment variables, etc...). Also state the current behavior vs. the expected behavior.
- If you'd like to see a feature or an enhancement please open an issue with
   a clear title and description of what the feature is and why it would be
   beneficial to the project and its users.

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.
