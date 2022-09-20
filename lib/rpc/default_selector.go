package myrpc

import (
	"context"
	"fmt"
)

type DefaultSelector struct {
	servers []string
	cnt     int
}

func (this *DefaultSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	if len(this.servers) > 0 {
		this.cnt++
		return this.servers[this.cnt%len(this.servers)]
	}
	return ""
}

func (this *DefaultSelector) UpdateServer(servers map[string]string) {
	fmt.Println(servers)
	this.servers = this.servers[:0]
	for k, _ := range servers {
		this.servers = append(this.servers, k)
		fmt.Println("update server,", k)
	}
}
