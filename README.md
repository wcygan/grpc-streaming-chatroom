# gRPC Streaming Chatroom

## Quickstart

### Pull the Images

```bash
docker pull wcygan/grpc-chat-client
docker pull wcygan/grpc-chat-server
```

### Start the Server

```bash
docker run -d --name chat-server -p 50051:50051 wcygan/grpc-chat-server
```

### Start Multiple Clients

Open a new terminal for each client you want to run:

```bash
docker run -it --network host wcygan/grpc-chat-client
```

Note: We use `--network host` for the clients to allow them to connect to the server running on localhost.

### Cleanup Docker Resources

Stop and remove the server container:
```bash
docker stop chat-server
docker rm chat-server
```

Remove the images (optional):
```bash
docker rmi wcygan/grpc-chat-client
docker rmi wcygan/grpc-chat-server
```

## Things to work on

- TUI for Client (Charm Bubbletea)
- Docs on how to develop with this project (hot reloading?)
