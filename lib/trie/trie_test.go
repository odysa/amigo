package trie

import (
	"testing"
)

func TestNodeInsert(t *testing.T) {
	root := &Node{
		pattern: "/",
		part:    "/",
		exact:   true,
	}

	root.Insert("/hello", []string{"hello"}, 0)
	if root.matchChild("hello").pattern != "/hello" {
		t.Errorf("Insert failed!")
	}
}

func TestNode_Find(t *testing.T) {
	root := &Node{
		pattern: "/",
		part:    "/",
		exact:   true,
	}

	root.Insert("/hello", []string{"hello"}, 0)

	res := root.Find([]string{"hello"}, 0)
	if res.pattern != "/hello" {
		t.Errorf("failed to find node, pattern should be %s, but got %s", "/hello", res.pattern)
	}
}
