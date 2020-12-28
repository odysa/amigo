package lib

type nodeType uint8

const (
	matchAll nodeType = iota
	param
	root
	static
)

type Node struct {
	path     string
	children []*Node
	// param node or not
	isWild    bool
	maxParams uint8
	// if children are search and support, indices would be eu
	indices  string
	nodeType nodeType
	handler  HandlerFunc
}

func (n *Node) addRoute(path string, handler HandlerFunc) {
	node := n
	for {
		pathLen := len(path)
		nodePathLen := len(node.path)
		// get common prefix
		index := longestCommonPrefix(path, node.path, min(pathLen, nodePathLen))
		// spilt node, e.g. : if service and support, would make node s, ervice, and upport
		if index < nodePathLen {
			child := &Node{
				path:     node.path[index:],
				isWild:   node.isWild,
				indices:  node.indices,
				children: node.children,
				handler:  node.handler,
			}
			node.children = []*Node{child}
			node.isWild = false
			node.path = path[:index]
			node.indices = string(node.path[index])
			node.handler = nil
		}

		if index < pathLen {
			path = path[index:]

			if node.isWild {
				node := node.children[0]
				if len(path) >= len(node.path) && node.path == path[:len(node.path)] {
					if len(node.path) >= len(path) || path[len(node.path)] == '/' {
						if node.nodeType != matchAll {
							continue
						}
					}

				}
				panic("Match error")

			}
			front := path[0]
			for i,ch:=range []byte(node.indices){
				if ch == front{
					node = node.children[i]
					break
				}
			}
		}
		if node.handler != nil {
			panic("a handler is already registered")
		}
		node.handler = handler
		return
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonPrefix(s1, s2 string, l int) int {
	i := 0
	for ; i < l && s1[i] == s2[i]; {
		i++
	}
	return i
}

func countParams(path string) int {
	n := 0
	for i := range path {
		if path[i] == '*' || path[i] == ':' {
			n++
		}
	}
	return n
}
