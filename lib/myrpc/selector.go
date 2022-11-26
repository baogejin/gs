package myrpc

import (
	"math/rand"
)

type Selector interface {
	Select(req interface{}) string
	UpdateServer(servers map[string]string)
}

type ServerInfo struct {
	Address string
	Info    string
}

type RandSelector struct {
	servers []*ServerInfo
}

func (this *RandSelector) Select(req interface{}) string {
	if len(this.servers) == 0 {
		return ""
	}
	idx := rand.Int31n(int32(len(this.servers)))
	return this.servers[idx].Address
}

func (this *RandSelector) UpdateServer(servers map[string]string) {
	this.servers = this.servers[:0]
	for k, v := range servers {
		this.servers = append(this.servers, &ServerInfo{Address: k, Info: v})
	}
}

type RoundSelector struct {
	servers []*ServerInfo
	round   int
}

func (this *RoundSelector) Select(req interface{}) string {
	if len(this.servers) == 0 {
		return ""
	}
	this.round++
	this.round = this.round % len(this.servers)
	return this.servers[this.round].Address
}

func (this *RoundSelector) UpdateServer(servers map[string]string) {
	this.servers = this.servers[:0]
	for k, v := range servers {
		this.servers = append(this.servers, &ServerInfo{Address: k, Info: v})
	}
}
