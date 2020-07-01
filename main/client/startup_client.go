package main

import (
	"context"
	"fmt"
	"go-raft/rpc"
	_ "go-raft/rpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:18800", grpc.WithInsecure())
	if err != nil {

	}
	defer conn.Close()

	groupId := "defult"
	serverId := "localhost:8080"
	peerId := "p"
	term := int64(100)
	prevLogTerm := int64(100)
	prevLogIndex := int64(100)
	prevote := true

	requet := rpc.RequestVoteRequest{
		GroupId:      &groupId,
		ServerId:     &serverId,
		PeerId:       &peerId,
		Term:         &term,
		PrevLogTerm:  &prevLogTerm,
		PrevLogIndex: &prevLogIndex,
		PreVote:      &prevote,
	}

	c := rpc.NewRaftServiceClient(conn)

	response, err1 := c.PreVote(context.Background(), &requet)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(response)
}
