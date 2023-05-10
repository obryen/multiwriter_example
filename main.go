package main

import (
	"bytes"
	"fmt"
	"io"
)

type Conn struct {
	io.Writer
}

func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}

func (c *Conn) Write(b []byte) (int, error) {
	fmt.Println("Writing to underlyng conn", string(b))
	return c.Writer.Write(b)
}

type Server struct {
	peers map[*Conn]bool
}

func NewServer() *Server {
	s := &Server{
		peers: make(map[*Conn]bool),
	}

	for i := 0; i < 10; i++ {
		s.peers[NewConn()] = true
	}
	return s
}
func (s *Server) broadCast(msg []byte) error {
	peers := []io.Writer{}
	for peer := range s.peers {
		peers = append(peers, peer)
		// if _, err := peer.Write(msg); err != nil {
		// 	log.Fatal(err)
		// }
	}
	mw := io.MultiWriter(peers...)
	_, err := mw.Write(msg)
	return err
}
func main() {
	s := NewServer()
	for b := range s.peers {
		fmt.Println("emprty init conns:", b)
	}
	s.broadCast([]byte("hello kenya"))

}
