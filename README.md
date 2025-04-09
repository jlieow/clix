# CliX - Command Line Extender

## Overview

CliX is a command-line tool designed to enhance and extend the functionality of existing command-line interfaces (CLIs). It allows users to define, customize, and chain commands with ease, empowering developers, system administrators, and power users to streamline their workflows.

## Features

-   **Command Aliasing:** Create custom aliases for frequently used commands.
-   **Pre and Post Hooks:** Define commands or functions to run before and after executing an alias.
-   **Extensibility:** Extend CLI functionality with custom functions written in Go.
-   **Configuration:** Store aliases and hooks in a configuration file for easy management.

## Installation

1.  **Install Go:** Ensure you have Go installed on your system.
2.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd clix
    ```
3.  **Build the application:**

    ```bash
    go build -o clix main.go
    ```
4.  **Move the binary to your desired location (e.g., /usr/local/bin):**

    ```bash
    mv clix /usr/local/bin/
    ```

## Usage

### Defining Aliases

Aliases are defined in the `config.json` file. The file is automatically created when the `clix` command is first run.

### Running Aliases

To run an alias, simply type the alias name in the command line. CliX will execute the corresponding command and any associated pre or post hooks.

### Hooks

Hooks are commands or functions that are executed before or after an alias. They can be used to perform tasks such as setting environment variables, loading configuration files, or printing messages.

#### Available Functions

-   `set_env`: Sets environment variables.
-   `load_env_file`: Loads environment variables from a file.
-   `print_all_env`: Prints all environment variables.
-   `hello_world`: Prints "Hello, world!".
-   `run_python`: Runs a Python script.
-   `run_javascript`: Runs a JavaScript script.

## Command Structure

The `clix` command has the following structure:

-   `clix`: The main command.
-   `cmd/config`: Contains configuration related commands.
-   `cmd/symlink`: Contains symlink related commands.
-   `Execute()`: Adds all child commands to the root command and sets flags appropriately. This is called by `main.main()`. It also creates a directory with config files.

## Configuration

The configuration file (`config.json`) is stored in the `$HOME/.clix` directory. It contains the definitions for aliases and hooks.

## Contributing

Contributions are welcome! Please submit a pull request with your changes.

## License

This project is licensed under the MIT License.