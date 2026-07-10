# rcli - Minimal HTTP Client CLI

A lightweight command-line HTTP client tool demonstrating CLI development in both **Rust** and **Go**.

## Overview

This repository contains two implementations of a simple HTTP client CLI:
- **Rust version** (`src/main.rs`): Built with `clap` for argument parsing
- **Go version** (`withgo/main.go`): Built with `cobra` for command handling

Both versions provide a minimal interface for making HTTP requests and inspecting responses.

## Project Structure

```
rcli/
├── src/
│   └── main.rs           # Rust implementation using clap
├── withgo/
│   ├── main.go           # Go implementation using cobra
│   ├── go.mod            # Go module definition
│   └── go.sum            # Go dependencies lock file
├── Cargo.toml            # Rust project manifest and dependencies
├── Cargo.lock            # Rust dependency lock file
├── .gitignore            # Git ignore file
└── README.md             # This file
```

## Language Composition

- **Rust**: 59%
- **Go**: 41%

## Prerequisites

### For Rust Version
- **Rust** 1.56 or later
- **Cargo** (comes with Rust)

### For Go Version
- **Go** 1.26.5 or later

## Installation & Setup

### Rust Version

1. Clone the repository:
```bash
git clone https://github.com/Akene-Uzezi/rcli.git
cd rcli
```

2. Build the project:
```bash
cargo build --release
```

3. Run:
```bash
cargo run -- --url "https://example.com" --method GET
```

### Go Version

1. Navigate to the Go directory:
```bash
cd withgo
```

2. Build the project:
```bash
go build -o httpcli
```

3. Run:
```bash
./httpcli get "https://example.com"
```

## Usage

### Rust Implementation

```bash
cargo run -- --url <URL> [--method <METHOD>] [--verbose]
```

**Options:**
- `--url, -u <URL>`: The target URL for the request (required)
- `--method, -m <METHOD>`: HTTP method to use (default: GET)
- `--verbose, -v`: Enable verbose logging output

**Example:**
```bash
cargo run -- --url "https://jsonplaceholder.typicode.com/posts/1" --method GET
```

### Go Implementation

```bash
./httpcli get <URL>
```

**Subcommands:**
- `get [url]`: Fetch content from a URL

**Example:**
```bash
./httpcli get "https://jsonplaceholder.typicode.com/posts/1"
```

## Dependencies

### Rust
- **clap** (v4.4): Command-line argument parser with derive macros
- **tokio** (v1.0): Async runtime (available with full feature set)
- **reqwest** (v0.12): HTTP client library with JSON support
- **serde** (v1.0): Serialization/deserialization framework
- **serde_json** (v1.0): JSON support for serde

### Go
- **cobra** (v1.10.2): Command CLI framework
- **pflag** (v1.0.9): POSIX/GNU-style flags library

## Features

- **Multi-language implementation**: Compare Rust and Go approaches to CLI development
- **HTTP requests**: Make GET requests to any URL
- **Error handling**: Graceful error reporting
- **Command-line parsing**: Learn from two different CLI frameworks (clap vs cobra)
- **Response inspection**: View HTTP status and response body

## Learning Resources

This project demonstrates:
- **Rust**: Command-line argument parsing with clap, structured programming patterns
- **Go**: Command framework with cobra, idiomatic Go patterns
- HTTP request handling in both languages
- CLI design patterns and best practices
- Cross-language project structure and organization

## License

This project is open source and available under the MIT License.

## Contributing

Feel free to fork, modify, and submit improvements! Contributions that enhance either implementation or add new features are welcome.
