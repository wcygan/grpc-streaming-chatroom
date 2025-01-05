# gRPC Streaming Chatroom

## Quickstart

Generate the latest protobuf code:

```bash
buf generate
```

Install dependencies:

```bash
cd server

go get google.golang.org/grpc/credentials@v1.69.2 \
    google.golang.org/grpc/internal/pretty@v1.69.2 \
    google.golang.org/grpc/encoding/proto@v1.69.2 \
    github.com/wcygan/grpc-streaming-chatroom/gen/chat/v1@v0.0.0 \
    google.golang.org/grpc/internal/binarylog@v1.69.2 \
    google.golang.org/grpc/internal/status@v1.69.2 \
    google.golang.org/grpc/reflection@v1.69.2
    
cd ../client

go get google.golang.org/grpc/credentials@v1.69.2 \
    google.golang.org/grpc/internal/pretty@v1.69.2 \
    google.golang.org/grpc/encoding/proto@v1.69.2 \
    github.com/wcygan/grpc-streaming-chatroom/gen/chat/v1@v0.0.0 \
    google.golang.org/grpc/internal/binarylog@v1.69.2 \
    google.golang.org/grpc/internal/status@v1.69.2 \
    google.golang.org/grpc/reflection@v1.69.2
```

## Things to work on

- Push Protos to BSR (use Github Actions)
- Build and Push Docker Images (use Github Actions)
- TUI for Client (Charm Bubbletea)