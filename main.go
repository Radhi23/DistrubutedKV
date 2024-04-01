package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// KeyValue represents a key-value pair.
type KeyValue struct {
	Key   string
	Value string
}

// ServerNode represents a server node in the distributed key-value store system.
type ServerNode struct {
	Address string          // Address of the server node
	Data    map[string]KeyValue // Data stored in the server node
}

// Put inserts a key-value pair into the server node's data store.
func (s *ServerNode) Put(kv *KeyValue, reply *bool) error {
	return nil
}

// Get retrieves the value associated with the given key from the server node's data store.
func (s *ServerNode) Get(key string, value *string) error {
	return nil
}

func main() {
	node := &ServerNode{
		Address: ":1234", 
		Data:    make(map[string]KeyValue),
	}

	// Register the ServerNode object as an RPC service
	rpc.Register(node)

	// Start the RPC server
	listener, err := net.Listen("tcp", node.Address)
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	fmt.Println("Server started at", node.Address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
