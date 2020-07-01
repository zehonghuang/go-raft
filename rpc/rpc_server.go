package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) PreVote(ctx context.Context, req *RequestVoteRequest) (*RequestVoteResponse, error) {
	fmt.Println(req)
	granted := true
	term := int64(100)
	return &RequestVoteResponse{Granted: &granted, Term: &term}, nil
}

func (*server) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, nil
}

func StartedServer(port *int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterRaftServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
