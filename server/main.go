package main

import (
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chatv1 "github.com/wcygan/grpc-streaming-chatroom/gen/chat/v1"
)

// chatServer implements chatv1.ChatServiceServer.
type chatServer struct {
	chatv1.UnimplementedChatServiceServer

	// We store a list of active client streams. Each client
	// receives messages via its stream.
	mu      sync.Mutex
	clients map[chatv1.ChatService_ChatStreamServer]struct{}
}

func newChatServer() *chatServer {
	return &chatServer{
		clients: make(map[chatv1.ChatService_ChatStreamServer]struct{}),
	}
}

// ChatStream handles a bidirectional streaming RPC.
func (s *chatServer) ChatStream(stream chatv1.ChatService_ChatStreamServer) error {
	// Register this new stream (client connection).
	s.addClient(stream)
	defer s.removeClient(stream)

	// Start a loop to continuously read incoming messages from this stream.
	for {
		msg, err := stream.Recv()
		if err != nil {
			// This typically indicates the client disconnected or an error occurred.
			log.Printf("stream.Recv() error: %v", err)
			return err
		}

		// Broadcast the received message to all connected clients.
		log.Printf("Received message from %s: %s", msg.User, msg.Text)
		s.broadcastMessage(msg)
	}
}

func (s *chatServer) addClient(stream chatv1.ChatService_ChatStreamServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[stream] = struct{}{}
	log.Printf("Client connected. Total clients: %d", len(s.clients))
}

func (s *chatServer) removeClient(stream chatv1.ChatService_ChatStreamServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, stream)
	log.Printf("Client disconnected. Total clients: %d", len(s.clients))
}

func (s *chatServer) broadcastMessage(msg *chatv1.ChatMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for clientStream := range s.clients {
		// Send message in a separate goroutine to avoid blocking others
		go func(cs chatv1.ChatService_ChatStreamServer) {
			if err := cs.Send(msg); err != nil {
				log.Printf("Failed to send message to client: %v", err)
			}
		}(clientStream)
	}
}

func main() {
	// Create a TCP listener.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Println("Server listening on port 50051...")

	// Create a gRPC server and register our chatServer.
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer) // Allows tools like grpcui or grpcurl to inspect the service

	chatv1.RegisterChatServiceServer(grpcServer, newChatServer())

	// Start serving.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
