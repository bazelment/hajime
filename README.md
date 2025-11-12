# Hajime

## Purpose

**Hajime** (はじめ - "beginning" in Japanese) is a polyglot monorepo starter template that demonstrates modern multi-language development using Bazel. It provides working examples of:

- **Multi-language build system**: Build Go and Rust applications in the same repository with consistent tooling
- **gRPC microservices**: Complete client/server implementations in both Go and Rust, showcasing language interoperability
- **Protocol Buffers**: Automated protobuf code generation for both languages from shared schema definitions
- **Modern Bazel**: Uses MODULE.bazel (bzlmod) for dependency management, replacing the legacy WORKSPACE approach

## What's Included

- **grpc_example/**: Working gRPC service implementations
  - Go client and server
  - Rust client and server
  - Shared protobuf definitions with automatic code generation
- **Bazel configuration**: Pre-configured rules for Go (`rules_go`), Rust (`rules_rust`), and protobuf code generation

## Use Cases

- Starting a polyglot microservices project
- Learning how to structure multi-language Bazel workspaces
- Understanding gRPC service implementations across different languages
- Evaluating Bazel for multi-language projects