# rcli

Minimal HTTP client CLI in Go and Rust.

## Go

```bash
cd withgo && go build -o httpcli
```

```
./httpcli "https://example.com"                           # GET
./httpcli post "https://example.com" -d '{"a":1}'         # POST
./httpcli put "https://example.com" -d '{"a":1}'          # PUT
./httpcli del "https://example.com"                       # DELETE (stub)
```

Flags: `-H, --headers` (repeatable), `-d, --data` (string)

## Rust

```bash
cargo build --release
```

```bash
cargo run -- --url "https://example.com" --method GET -v
```

Flags: `-u, --url` (required), `-m, --method` (default: GET), `-v, --verbose`

## Dependencies

| Language | Dependencies |
|----------|-------------|
| Go | `cobra`, `pflag` |
| Rust | `clap`, `tokio`, `reqwest`, `serde`, `serde_json` |

## License

MIT
