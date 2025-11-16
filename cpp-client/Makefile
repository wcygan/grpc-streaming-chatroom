CXX = g++
CXXFLAGS = -std=c++17 -Wall -Wextra

# Homebrew paths (macOS)
BREW_PREFIX_PROTOBUF = $(shell brew --prefix protobuf 2>/dev/null || echo /usr)
BREW_PREFIX_GRPC = $(shell brew --prefix grpc 2>/dev/null || echo /usr)
BREW_PREFIX_ABSEIL = $(shell brew --prefix abseil 2>/dev/null || echo /usr)

# Include paths
INCLUDES = -I. -Igenerated \
           -I$(BREW_PREFIX_PROTOBUF)/include \
           -I$(BREW_PREFIX_GRPC)/include \
           -I$(BREW_PREFIX_ABSEIL)/include

# Library paths
LDFLAGS = -L$(BREW_PREFIX_PROTOBUF)/lib \
          -L$(BREW_PREFIX_GRPC)/lib \
          -L$(BREW_PREFIX_ABSEIL)/lib

# Libraries (order matters for linking)
LIBS = -lgrpc++ -lgrpc -lgpr -lprotobuf \
       -labsl_log_internal_message \
       -labsl_log_internal_check_op \
       -labsl_log_internal_nullguard \
       -labsl_status \
       -labsl_cord \
       -labsl_cordz_info \
       -labsl_cordz_handle \
       -labsl_cordz_functions \
       -labsl_statusor \
       -labsl_strings \
       -labsl_str_format_internal \
       -labsl_base \
       -labsl_raw_logging_internal \
       -labsl_throw_delegate \
       -labsl_int128 \
       -labsl_synchronization \
       -labsl_time \
       -labsl_time_zone \
       -lpthread

# Source files
PROTO_SRCS = generated/chat/v1/chat.pb.cc \
             generated/chat/v1/chat.grpc.pb.cc

SRCS = main.cpp $(PROTO_SRCS)

OBJS = $(SRCS:.cpp=.o)
OBJS := $(OBJS:.cc=.o)

TARGET = chat_client

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) $(LDFLAGS) -o $@ $^ $(LIBS)

%.o: %.cpp
	$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

%.o: %.cc
	$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

clean:
	rm -f $(OBJS) $(TARGET)

.PHONY: all clean
