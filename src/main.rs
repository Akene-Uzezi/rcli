use serde::{Deserialize, Serialize};
use std::env;
use std::error::Error;

const BASE_URL: &str = "https://jsonplaceholder.typicode.com/posts";

#[derive(Serialize, Deserialize, Debug)]
struct Post {
    #[serde(skip_serializing_if = "Option::is_none")]
    id: Option<u32>,
    title: String,
    body: String,
    #[serde(rename = "userId")]
    user_id: u32,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    //collect command line arguments
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        print_usage();
        return Ok(());
    }

    let command = args[1].as_str();
    let client = reqwest::Client::new();

    match command {
        "get" => {
            if args.len() < 3 {
                println!("Error: missing post id. Usage: cargo run get <id>");
                return Ok(());
            }
            let id = &args[2];
            let url = format!("{}/{}", BASE_URL, id);

            let response = client.get(&url).send().await?.json::<Post>().await?;
            println!("--- READ (GET) SUCCESS ---");
            println!("{:#?}", response);
        }

        "create" => {
            if args.len() < 4 {
                println!("Error: Missing arguments. Usage: cargo run create <title> <body>");
                return Ok(());
            }
            let new_post = Post {
                id: None,
                title: args[2].clone(),
                body: args[3].clone(),
                user_id: 1,
            };

            let response = client
                .post(BASE_URL)
                .json(&new_post)
                .send()
                .await?
                .json::<Post>()
                .await?;
            println!("--- CREATE (POST) SUCCESS ---");
            println!("{:#?}", response);
        }

        "update" => {
            if args.len() < 5 {
                println!("Error: Missing arguments. Usage: cargo run update <id> <title> <body>");
                return Ok(());
            }
            let id = &args[2];
            let url = format!("{}/{}", BASE_URL, id);

            let updated_post = Post {
                id: Some(id.parse()?),
                title: args[3].clone(),
                body: args[4].clone(),
                user_id: 1,
            };

            let response = client
                .put(&url)
                .json(&updated_post)
                .send()
                .await?
                .json::<Post>()
                .await?;
            println!("--- UPDATE (PUT) SUCCESS ---");
            println!("{:#?}", response);
        }

        "delete" => {
            if args.len() < 3 {
                println!("Error: Missing post ID. Usage: cargo run delete <id>");
                return Ok(());
            }
            let id = &args[2];
            let url = format!("{}/{}", BASE_URL, id);

            let response = client.delete(&url).send().await?;
            if response.status().is_success() {
                println!("--- DELETE SUCCESS ---");
                println!(
                    "Post {} deleted successfully (Status: {}).",
                    id,
                    response.status()
                );
            }
        }

        _ => {
            println!("unknown command: {}", command);
            print_usage();
        }
    }

    Ok(())
}

fn print_usage() {
    println!("\n🚀 Simple Rust CRUD CLI");
    println!("Usage:");
    println!("  cargo run get <id>                      | Fetch a post");
    println!("  cargo run create <title> <body>         | Create a new post");
    println!("  cargo run update <id> <title> <body>    | Update an existing post");
    println!("  cargo run delete <id>                   | Delete a post\n");
}
