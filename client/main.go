package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"buf.build/gen/go/wcygan/grpc-streaming-chatroom/grpc/go/chat/v1/chatv1grpc"
	chatv1 "buf.build/gen/go/wcygan/grpc-streaming-chatroom/protocolbuffers/go/chat/v1"
)

func main() {
	// Connect to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := chatv1grpc.NewChatServiceClient(conn)

	// Create the bidirectional stream.
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("ChatStream error: %v", err)
	}
	defer stream.CloseSend()

	// Generate a nanoid for the username
	username, err := gonanoid.New(7)
	if err != nil {
		log.Fatalf("Failed to generate nanoid: %v", err)
	}

	fmt.Printf("Your username is: %s\n", username)

	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

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
