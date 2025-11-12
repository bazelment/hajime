use greeter_proto::greeter::greeter_client::GreeterClient;
use greeter_proto::greeter::HelloRequest;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = std::env::args()
        .nth(1)
        .unwrap_or_else(|| "http://[::1]:50051".to_string());

    println!("Rust client connecting to Go server at {}", addr);

    let mut client = GreeterClient::connect(addr).await?;

    let request = tonic::Request::new(HelloRequest {
        name: "Rust Client".to_string(),
    });

    let response = client.say_hello(request).await?;

    println!("Response from Go server: {}", response.get_ref().message);

    Ok(())
}

