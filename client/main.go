package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	chatv1 "github.com/wcygan/grpc-streaming-chatroom/gen/chat/v1"
)

func main() {
	// Connect to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := chatv1.NewChatServiceClient(conn)

	// Create the bidirectional stream.
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("ChatStream error: %v", err)
	}
	defer stream.CloseSend()

	// Prompt user for a name.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if username == "" {
		username = "Anonymous"
	}

	// Start a goroutine to receive messages from the server.
	go func() {
		for {
			// Use stream.Recv() to get incoming messages from the server.
			msg, err := stream.Recv()
			if err != nil {
				log.Printf("stream.Recv() error: %v", err)
				return
			}
			fmt.Printf("\r%s: %s\n", msg.User, msg.Text)
			fmt.Print("> ") // re-print prompt
		}
	}()

	// Main loop: read user input from stdin, send to the server.
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		// Build a ChatMessage and send it over the stream.
		chatMsg := &chatv1.ChatMessage{
			User:      username,
			Text:      text,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		if err := stream.Send(chatMsg); err != nil {
			log.Printf("Failed to send message: %v", err)
			return
		}
	}
}
