package main

import (
	r "go-raft/rpc"
)

func main() {
	port := 18800
	r.StartedServer(&port)
}
