package trie

import "strings"

type Node struct {
	pattern  string
	part     string
	children []*Node
	exact    bool
}

func (n *Node) Pattern() string {
	return n.pattern
}

func (n *Node) matchChild(part string) *Node {
	for _, child := range n.children {
		if child.part == part || !child.exact {
			return child
		}
	}
	return nil
}
func (n *Node) matchChildren(part string) []*Node {
	Nodes := make([]*Node, 0)
	for _, child := range n.children {
		if child.part == part || !child.exact {
			Nodes = append(Nodes, child)
		}
	}
	return Nodes
}

func (n *Node) Find(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// pattern not found
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		if result := child.Find(parts, height+1); result != nil {
			return result
		}
	}
	return nil
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &Node{
			part:  part,
			exact: part[0] != ':' && part[0] != '*',
		}
		n.children = append(n.children, child)
	}
	child.Insert(pattern, parts, height+1)
}
