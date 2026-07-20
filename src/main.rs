use clap::{Parser, Subcommand};
use reqwest::blocking::Client;
use reqwest::header::{HeaderMap, HeaderName, HeaderValue};
use std::io::{self, Read};
use std::process;

// Define root cli command structure
#[derive(Parser)]
#[command(name = "httpcli", about = "httpcli is a minimal HTTP client cli")]
struct Cli {
    #[command(subcommand)]
    command: Option<Commands>,

    // url for default get request
    url: Option<String>,
}

// define the available subcommand
#[derive(Subcommand)]
enum Commands {
    //make a post request
    Post {
        url: String,

        #[arg(short = 'H', long = "headers")]
        headers: Vec<String>,

        #[arg(short, long)]
        data: Option<String>,
    },
    Del {
        url: String,

        #[arg(short = 'H', long = "headers")]
        headers: Vec<String>,

        #[arg(short, long)]
        data: Option<String>,
    },
}

fn main() {
    let cli = Cli::parse();
    let client = Client::new();

    // match against subcommand provided by the user
    match cli.command {
        Some(Commands::Post { url, headers, data }) => {
            // passing borrowed reference &headers
            let header_map = parse_headers(&headers);

            let body_content = data.unwrap_or_default();

            let response = client
                .post(&url)
                .headers(header_map)
                .body(body_content)
                .send();

            handle_response(response);
        }
        Some(Commands::Del { url, headers, data }) => {
            let header_map = parse_headers(&headers);
            let body_content = data.unwrap_or_default();

            let response = client
                .delete(&url)
                .headers(header_map)
                .body(body_content)
                .send();
            handle_response(response);
        }
        None => {
            if let Some(url) = cli.url {
                let response = client.get(&url).send()
                    handle_response(response);
            } else {
                eprintln!("error: exact positional argument (url) required");
                process::exit(1);
            }
        }
    }
}
