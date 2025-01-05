module github.com/wcygan/grpc-streaming-chatroom/server

go 1.23.4

require google.golang.org/grpc v1.69.2

require github.com/wcygan/grpc-streaming-chatroom/gen v0.0.0

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)

replace github.com/wcygan/grpc-streaming-chatroom/gen => ../gen
