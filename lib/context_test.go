package lib

import (
	"net/http"
	"net/url"
	"testing"
)

func TestContext_GetParam(t *testing.T) {
	c := Context{
		Params: map[string]string{
			"abc": "123",
		},
	}
	if c.GetParam("abc") != "123" {
		t.Errorf("value should be %s, but got %s", "123", c.GetParam("abc"))
	}
}
func TestContext_GetQuery(t *testing.T) {
	c := Context{
		R: &http.Request{
			URL: &url.URL{
				RawQuery: "a=1&b=2",
			},
		},
	}
	if c.GetQuery("a") != "1" || c.GetQuery("b") != "2" {
		t.Errorf("invalid query %s %s", c.GetQuery("a"), c.GetQuery("b"))
	}
}
