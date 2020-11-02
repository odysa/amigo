package lib

import "testing"

func TestNodeInsert(t *testing.T) {
	root := &Node{
		pattern: "/",
		part:    "/",
		exact:   true,
	}

	root.Insert("/hello", []string{"hello"}, 0)
	if root.matchChild("hello").pattern!="/hello"{
		t.Error("Insert failed!")
	}
}
