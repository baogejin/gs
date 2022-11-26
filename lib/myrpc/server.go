package myrpc

import (
	"gs/lib/mylog"
	"net"
	"net/rpc"
)

// Server rpc server based on net/rpc implementation
type Server struct {
	*rpc.Server
}

// NewServer Create a new rpc server
func NewServer() *Server {
	return &Server{&rpc.Server{}}
}

// Register register rpc function
func (s *Server) Register(rcvr interface{}) error {
	return s.Server.Register(rcvr)
}

// RegisterName register the rpc function with the specified name
func (s *Server) RegisterName(name string, rcvr interface{}) error {
	return s.Server.RegisterName(name, rcvr)
}

// Serve start service
func (s *Server) Serve(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		addr := conn.RemoteAddr()
		mylog.Debug("new conn ", addr)
		go func() {
			s.Server.ServeConn(conn)
			mylog.Debug("conn close ", addr)
		}()
	}
}
