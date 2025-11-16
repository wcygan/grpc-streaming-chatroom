#include <grpcpp/grpcpp.h>
#include <iostream>
#include <memory>
#include <string>
#include <thread>
#include <random>
#include <chrono>
#include <iomanip>
#include <sstream>

#include "chat/v1/chat.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReaderWriter;
using grpc::Status;

using chat::v1::ChatMessage;
using chat::v1::ChatService;

// Generate a random alphanumeric username similar to nanoid
std::string generate_username(int length = 7) {
    static const char charset[] =
        "0123456789"
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz";

    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<> dis(0, sizeof(charset) - 2);

    std::string result;
    result.reserve(length);
    for (int i = 0; i < length; ++i) {
        result += charset[dis(gen)];
    }
    return result;
}

// Get current timestamp in RFC3339 format
std::string get_timestamp() {
    auto now = std::chrono::system_clock::now();
    auto in_time_t = std::chrono::system_clock::to_time_t(now);

    std::stringstream ss;
    ss << std::put_time(std::gmtime(&in_time_t), "%Y-%m-%dT%H:%M:%S");

    auto ms = std::chrono::duration_cast<std::chrono::milliseconds>(
        now.time_since_epoch()) % 1000;
    ss << '.' << std::setfill('0') << std::setw(3) << ms.count() << 'Z';

    return ss.str();
}

int main() {
    // Connect to the gRPC server
    auto channel = grpc::CreateChannel(
        "localhost:50051",
        grpc::InsecureChannelCredentials()
    );

    auto stub = ChatService::NewStub(channel);

    // Create the bidirectional stream
    ClientContext context;
    auto stream = stub->ChatStream(&context);

    // Generate username
    std::string username = generate_username();
    std::cout << "Your username is: " << username << std::endl;

    // Start a thread to receive messages from the server
    std::thread receiver([&stream]() {
        ChatMessage msg;
        while (stream->Read(&msg)) {
            std::cout << "\r" << msg.user() << ": " << msg.text() << std::endl;
            std::cout << "> " << std::flush;
        }
    });

    // Main loop: read user input and send to server
    std::string text;
    while (true) {
        std::cout << "> " << std::flush;
        std::getline(std::cin, text);

        // Trim whitespace
        text.erase(0, text.find_first_not_of(" \t\n\r"));
        text.erase(text.find_last_not_of(" \t\n\r") + 1);

        if (text.empty()) {
            continue;
        }

        // Build and send ChatMessage
        ChatMessage chat_msg;
        chat_msg.set_user(username);
        chat_msg.set_text(text);
        chat_msg.set_timestamp(get_timestamp());

        if (!stream->Write(chat_msg)) {
            std::cerr << "Failed to send message" << std::endl;
            break;
        }
    }

    // Close the stream and wait for receiver thread
    stream->WritesDone();
    receiver.join();

    Status status = stream->Finish();
    if (!status.ok()) {
        std::cerr << "Stream error: " << status.error_message() << std::endl;
        return 1;
    }

    return 0;
}
