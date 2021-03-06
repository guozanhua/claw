// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master

import (
	"fmt"
)

type Handler interface {
	Handle(node *Node)
}

const (
	LOGIN = iota
	BROADCAST
)

type Login struct {
	Name string
}

func (m *Login) Handle(node *Node) {
	node.Name = m.Name
	if ret := nodes.AddPeer(node.Name, node); !ret {
		fmt.Println("Master.proto add peer fail, repeated, node.Name:", node.Name)
	}
}

var (
	handlers map[uint16]func() Handler
)

func init() {
	handlers = make(map[uint16]func() Handler)

	handlers[LOGIN] = func() Handler { return &Login{} }
}
