# C++ gRPC Chat Client

Standard gRPC C++ client for the streaming chatroom.

## Build

**Install dependencies:**
```bash
# macOS
brew install grpc protobuf abseil

# Ubuntu
apt-get install libgrpc++-dev libprotobuf-dev libabsl-dev protobuf-compiler-grpc
```

**Build with Makefile:**
```bash
make
```

## Run

Start the server first (see main README), then:
```bash
./chat_client
```

## Generated Files

The `generated/` directory contains protobuf-compiled C++ files committed to git:
- `chat/v1/chat.pb.{h,cc}` - Protobuf message definitions
- `chat/v1/chat.grpc.pb.{h,cc}` - gRPC service stubs

To regenerate (requires `protoc` and `grpc_cpp_plugin`):
```bash
protoc --cpp_out=generated --grpc_out=generated \
  --plugin=protoc-gen-grpc=$(which grpc_cpp_plugin) \
  -I ../proto ../proto/chat/v1/chat.proto
```
