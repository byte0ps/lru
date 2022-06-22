package lru

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	c := New()
	if c.Set("key", "123456") {
		if v, ok := c.Get("key"); ok {
			fmt.Println(v.(string))
		}
		c.Del("key")
		_, ok := c.Get("key")
		if ok {
			t.Error("output does not contain expected string")
		}
	}
}
