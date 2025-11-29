package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	// key should be the url
	interval := 5 * time.Second
	cases := map[string]struct {
		key string
		val []byte
	}{
		"add": {
			key: "add test",
			val: []byte("test add"),
		},
		"get": {
			key: "get test",
			val: []byte("test get"),
		},
	}

	for name, c := range cases {
		t.Run(fmt.Sprintf("Test case %s", name), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find %s in -> cache: %t", c.key, ok)
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find %s in -> %s", c.val, val)
			}
		})
	}
}
