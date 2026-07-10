# rcli - Simple Rust CRUD CLI

A lightweight command-line tool written in Rust for performing CRUD (Create, Read, Update, Delete) operations on posts via the JSONPlaceholder API. Perfect for learning async Rust, HTTP requests, and CLI development.

## Features

- **GET**: Fetch a post by ID
- **CREATE**: Create a new post with a title and body
- **UPDATE**: Update an existing post with new title and body
- **DELETE**: Delete a post by ID

All operations interact with the JSONPlaceholder public API (`https://jsonplaceholder.typicode.com/posts`).

## Prerequisites

- **Rust** 1.56 or later
- **Cargo** (comes with Rust)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Akene-Uzezi/rcli.git
cd rcli
```

2. Build the project:
```bash
cargo build --release
```

## Usage

Run commands using `cargo run` followed by the command and arguments:

### Get a Post
Fetch a post by its ID:
```bash
cargo run get 1
```

### Create a Post
Create a new post with a title and body:
```bash
cargo run create "My Title" "My post body text"
```

### Update a Post
Update an existing post with a new title and body:
```bash
cargo run update 1 "Updated Title" "Updated body text"
```

### Delete a Post
Delete a post by its ID:
```bash
cargo run delete 1
```

## Project Structure

```
rcli/
├── src/
│   └── main.rs          # Main CLI application logic
├── Cargo.toml           # Project manifest and dependencies
├── Cargo.lock           # Dependency lock file
├── .gitignore           # Git ignore file
└── README.md            # This file
```

## Dependencies

- **tokio** (v1.0): Async runtime with full feature set
- **reqwest** (v0.12): HTTP client library with JSON support
- **serde** (v1.0): Serialization/deserialization framework
- **serde_json** (v1.0): JSON support for serde

## How It Works

The application:
1. Parses command-line arguments to determine the operation
2. Creates an async HTTP client using `reqwest`
3. Constructs the appropriate HTTP request (GET, POST, PUT, DELETE)
4. Serializes/deserializes JSON data using `serde`
5. Displays formatted results to the console

## Example Output

```
$ cargo run get 1
--- READ (GET) SUCCESS ---
Post {
    id: Some(1),
    title: "sunt aut facere repellat provident...",
    body: "quia et suscipit...",
    user_id: 1,
}
```

## Notes

- The CLI uses the free **JSONPlaceholder** API for testing CRUD operations
- Post IDs in JSONPlaceholder range from 1 to 100
- The API returns mock data; changes are not persisted
- All network calls are handled asynchronously using Tokio

## Learning Resources

This project demonstrates:
- Async/await patterns in Rust with Tokio
- HTTP requests with reqwest
- JSON serialization with Serde
- Command-line argument parsing
- Error handling in async Rust
- Pattern matching for CLI commands

## License

This project is open source and available under the MIT License.

## Contributing

Feel free to fork, modify, and submit improvements!
