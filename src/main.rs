use clap::Parser;

/// A simple http client cli tool built with clap
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Cli {
    /// The target URL for the request
    #[arg(short, long)]
    url: String,

    /// The HTTP method to use(defaults to GET)
    #[arg(short, long, default_value = "GET")]
    method: String,

    /// Turn on verbose logging output
    #[arg(short, long)]
    verbose: bool,
}

fn main() {
    // This one line automatically parses args, validates types,
    // and exits with a clean error/help menu if anything is wrong.
    let args = Cli::parse();

    println!("--- Parsed Rust flags ---");
    println!("URL: {}", args.url);
    println!("Method: {}", args.method);
    println!("Verbose: {}", args.verbose);
}
