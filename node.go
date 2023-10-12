package main

type Node struct {
	children map[rune]*Node
}

type Visitor interface {
	VisitNode(*Node)
}

func NewNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

func (n *Node) Accept(v Visitor) {
	v.VisitNode(n)
}
