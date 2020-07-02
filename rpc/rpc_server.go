package rpc

import (
	"encoding/json"
	"fmt"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (*Server) PreVote(r *RequestVoteRequest, s RaftService_PreVoteServer) error {
	return func(r *RequestVoteRequest, s RaftService_PreVoteServer) error {
		reqJson, err := json.Marshal(r)
		if err != nil {
			log.Println("what ??")
		}
		log.Println(string(reqJson))
		granted := true
		term := int64(100)
		s.Send(&RequestVoteResponse{Granted: &granted, Term: &term})
		return nil
	}(r, s)
}

func handlerPreVoteHRequest(r *RequestVoteRequest, s RaftService_PreVoteServer) {

}

func (*Server) RequestVote(*RequestVoteRequest, RaftService_RequestVoteServer) error {
	return nil
}

func StartedServer(port *int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})
	RegisterRaftServiceServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
