# rcli - Minimal HTTP Client CLI

A lightweight command-line HTTP client tool demonstrating CLI development in both **Rust** and **Go**.

## Overview

This repository contains two implementations of a simple HTTP client CLI:
- **Go version** (`withgo/main.go`): Built with `cobra` for command handling
- **Rust version** (`src/main.rs`): Built with `clap` for argument parsing

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

- **Go**: 77.7%
- **Rust**: 22.3%

## Prerequisites

### For Go Version
- **Go** 1.26.5 or later

### For Rust Version
- **Rust** 1.56 or later
- **Cargo** (comes with Rust)

## Installation & Setup

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

## Usage

### Go Implementation

The Go CLI tool supports multiple HTTP methods through subcommands:

#### Default GET Command (Root)

Make a GET request by passing a URL directly:

```bash
./httpcli [url]
```

**Example:**
```bash
./httpcli "https://jsonplaceholder.typicode.com/posts/1"
```

**Output:**
```
status: 200 OK
{response body}
```

#### POST Command

Make a POST request with optional headers and data payload:

```bash
./httpcli post [url] [flags]
```

**Flags:**
- `-H, --headers`: HTTP headers to pass (can be used multiple times)
- `-d, --data`: Raw string data payload for the POST body

**Examples:**
```bash
./httpcli post "https://jsonplaceholder.typicode.com/posts" \
  -d '{"title":"Test","body":"Content"}' \
  -H "Content-Type: application/json"
```

```bash
./httpcli post "https://example.com/api" \
  -d "key=value&foo=bar" \
  -H "Content-Type: application/x-www-form-urlencoded"
```

#### DELETE Command

Make a DELETE request with optional headers and data:

```bash
./httpcli del [url] [flags]
```

**Flags:**
- `-H, --headers`: HTTP headers to pass (can be used multiple times)
- `-d, --data`: Raw string data payload for the request body

**Example:**
```bash
./httpcli del "https://jsonplaceholder.typicode.com/posts/1" \
  -H "Authorization: Bearer token123"
```

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

## Dependencies

### Go
- **cobra** (v1.10.2): Command CLI framework for building CLI applications
- **pflag** (v1.0.9): POSIX/GNU-style flags library (used by cobra)
- **mousetrap** (v1.1.0): Windows console event handling (indirect dependency)

### Rust
- **clap** (v4.4): Command-line argument parser with derive macros
- **tokio** (v1.0): Async runtime (available with full feature set)
- **reqwest** (v0.12): HTTP client library with JSON support
- **serde** (v1.0): Serialization/deserialization framework
- **serde_json** (v1.0): JSON support for serde

## Features

- **Multi-language implementation**: Compare Go and Rust approaches to CLI development
- **Multiple HTTP methods**: Support for GET, POST, DELETE requests
- **Custom headers**: Pass custom HTTP headers to requests
- **Request payloads**: Send raw data in request bodies
- **Error handling**: Graceful error reporting
- **Command-line parsing**: Learn from two different CLI frameworks (cobra vs clap)
- **Response inspection**: View HTTP status and response body

## Learning Resources

This project demonstrates:
- **Go**: Command framework with cobra, subcommands, flags, idiomatic Go patterns
- **Rust**: Command-line argument parsing with clap, structured programming patterns
- HTTP request handling in both languages
- CLI design patterns and best practices
- Cross-language project structure and organization

## License

This project is open source and available under the MIT License.

## Contributing

Feel free to fork, modify, and submit improvements! Contributions that enhance either implementation or add new features are welcome.
