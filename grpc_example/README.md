# gRPC Examples with Go and Rust

This directory contains gRPC examples demonstrating interoperability between Go and Rust, with both languages serving as client and server.

## Structure

```
grpc_example/
├── proto/                    # Shared protobuf definitions
│   ├── greeter.proto        # Simple Greeter service
│   └── BUILD.bazel
├── go/                       # Go implementations
│   ├── server/main.go       # Go gRPC server
│   ├── client/main.go       # Go gRPC client
│   ├── go.mod
│   └── BUILD.bazel
└── rust/                     # Rust implementations
    ├── server/main.rs       # Rust gRPC server
    ├── client/main.rs       # Rust gRPC client
    ├── proto/               # Rust proto code generation
    │   ├── build.rs         # Build script for tonic
    │   ├── src/lib.rs       # Generated proto library
    │   └── BUILD.bazel
    └── BUILD.bazel
```

## Building

Build all targets:

```bash
bazel build //grpc_example/...
```

Build individual targets:

```bash
# Go
bazel build //grpc_example/go:server
bazel build //grpc_example/go:client

# Rust
bazel build //grpc_example/rust:server
bazel build //grpc_example/rust:client
```

## Running

### Example 1: Go Server with Rust Client

Terminal 1 (Start Go server on port 50051):
```bash
bazel run //grpc_example/go:server
```

Terminal 2 (Run Rust client to connect to Go server):
```bash
bazel run //grpc_example/rust:client
```

### Example 2: Rust Server with Go Client

Terminal 1 (Start Rust server on port 50052):
```bash
bazel run //grpc_example/rust:server
```

Terminal 2 (Run Go client to connect to Rust server):
```bash
bazel run //grpc_example/go:client
```

## Implementation Details

### Proto Definition

The `greeter.proto` file defines a simple unary RPC service with one method:
- `SayHello(HelloRequest) returns (HelloReply)`

### Go Implementation

- **Location**: `grpc_example/go/`
- **Server**: `server/main.go` - Listens on `[::]:50051` by default
- **Client**: `client/main.go` - Connects to `localhost:50052` by default (for Rust server)
- Uses `google.golang.org/grpc` for gRPC
- Proto code generated with `go_proto_library` using `go_grpc` compiler

### Rust Implementation

- **Location**: `grpc_example/rust/`
- **Server**: `server/main.rs` - Listens on `[::1]:50052`
- **Client**: `client/main.rs` - Connects to `http://[::1]:50051` by default (for Go server)
- Uses `tonic` v0.12 for gRPC and `prost` v0.13 for protobuf
- Proto code generated at build time using `cargo_build_script` with `tonic-build`

### Bazel Setup

- Go dependencies managed via `go.work` and Gazelle
- Rust dependencies managed via `Cargo.toml` and `crate_universe`
- Proto compilation handled by `rules_proto` for Go
- Rust proto generation uses `tonic-build` in a build script

## Key Files

- `MODULE.bazel` - Bazel module with Go, Rust, and proto dependencies
- `Cargo.toml` - Rust workspace dependencies (tonic, prost, tokio, tonic-build)
- `go.work` - Go workspace configuration
- `grpc_example/proto/greeter.proto` - Service definition

## Available Targets

**Go:**
- `//grpc_example/go:server` - Go gRPC server
- `//grpc_example/go:client` - Go gRPC client

**Rust:**
- `//grpc_example/rust:server` - Rust gRPC server
- `//grpc_example/rust:client` - Rust gRPC client

## Notes

- No direct cargo commands needed - everything is built through Bazel
- The Cargo.lock is committed to enable reproducible builds
- Proto files are automatically compiled to both Go and Rust code during build
- Both Go and Rust can act as either client or server, allowing 4 different combinations


