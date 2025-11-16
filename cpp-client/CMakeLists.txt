cmake_minimum_required(VERSION 3.15)
project(grpc-chat-client CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Help CMake find Homebrew packages on macOS
if(APPLE)
    execute_process(
        COMMAND brew --prefix protobuf
        OUTPUT_VARIABLE PROTOBUF_PREFIX
        OUTPUT_STRIP_TRAILING_WHITESPACE
    )
    execute_process(
        COMMAND brew --prefix grpc
        OUTPUT_VARIABLE GRPC_PREFIX
        OUTPUT_STRIP_TRAILING_WHITESPACE
    )
    list(APPEND CMAKE_PREFIX_PATH "${PROTOBUF_PREFIX}" "${GRPC_PREFIX}")
endif()

# Find Protobuf and gRPC
find_package(Protobuf REQUIRED)
find_package(gRPC CONFIG REQUIRED)

# Include directories
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/generated)

# Generated protobuf/gRPC sources
set(PROTO_SRCS
    generated/chat/v1/chat.pb.cc
    generated/chat/v1/chat.grpc.pb.cc
)

set(PROTO_HDRS
    generated/chat/v1/chat.pb.h
    generated/chat/v1/chat.grpc.pb.h
)

# Main client executable
add_executable(chat_client
    main.cpp
    ${PROTO_SRCS}
)

# Link libraries
target_link_libraries(chat_client
    gRPC::grpc++
    protobuf::libprotobuf
)
