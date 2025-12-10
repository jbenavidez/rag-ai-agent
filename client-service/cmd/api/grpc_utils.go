package main

import (
	pb "client/proto/generated"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectGRPC() pb.EmbeddingServiceClient {
	var conn *grpc.ClientConn
	var err error

	for {
		conn, err = grpc.Dial(
			fmt.Sprintf("%s:%s", "embedding-service", gRpcPort),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if err == nil {
			fmt.Println("Connected to gRPC server")
			break
		}

		fmt.Println("RPC server not ready, retrying in 2s... Error: ", err)
		time.Sleep(2 * time.Second)
	}

	return pb.NewEmbeddingServiceClient(conn)
}
