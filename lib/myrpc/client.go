package myrpc

import (
	"io"
	"net/rpc"
)

// Client rpc client based on net/rpc implementation
type Client struct {
	*rpc.Client
}

// NewClient Create a new rpc client
func NewClient(conn io.ReadWriteCloser) *Client {
	return &Client{rpc.NewClient(conn)}
}

// Call synchronously calls the rpc function
func (c *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.Client.Call(serviceMethod, args, reply)
}

// AsyncCall asynchronously calls the rpc function and returns a channel of *rpc.Call
func (c *Client) AsyncCall(serviceMethod string, args interface{}, reply interface{}) chan *rpc.Call {
	return c.Go(serviceMethod, args, reply, nil).Done
}
