package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	cases := []struct {
		key	string
		val	[]byte
	}{
		{
		key:	"https://example.com",
		val:	[]byte("justsometestdata"),
		},
		{
		key:	"https://example.com/path",
		val:	[]byte("alittlemoretestdata"),
		},
	}
	
	for i, c := range cases {
		t.Run(
			fmt.Sprintf("Test Case %v", i),
			func (t *testing.T) {
				cache.Add(c.key, c.val)
				val, ok := cache.Get(c.key)
				if !ok {
					t.Errorf("Unexpected cache miss")
					return
				}
				if string(val) != string(c.val) {
					t.Errorf("Unexpected val return")
					return
				}
			},
		)
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const block = interval + 5 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("Unexpected cache miss")
		return
	}
	time.Sleep(block)
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("Unexpected cache hit after Reap")
		return
	}
}



