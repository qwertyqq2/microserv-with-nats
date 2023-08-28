package store

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewTTLCache[string, int](10*time.Second, 5*time.Second)

	cache.Add("first", 1, 0)
	cache.Add("second", 2, 0)

	val, ok := cache.Get("first")
	if !ok {
		t.Fatal("not found")
	}
	if val != 1 {
		t.Fatal("neq")
	}

}

func TestCachePurge(t *testing.T) {
	cache := NewTTLCache[string, int](10*time.Second, 5*time.Second)

	cache.Add("first", 1, 3*time.Second)

	<-time.After(3 * time.Second)

	_, ok := cache.Get("first")
	if ok {
		t.Fatal("found")
	}

	cache.Add("second", 2, 0)

	<-time.After(cache.defaultExpiration)

	_, ok = cache.Get("second")
	if ok {
		t.Fatal("found")
	}
}
