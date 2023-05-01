package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type bucket struct {
	items map[uint32]string
	mu    *sync.RWMutex
}

func (b bucket) Get(key uint32) string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.items[key]
}

func (b bucket) Add(key uint32, val []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.items[key] = string(val)
}

type Cache struct {
	buckets map[uint32]bucket
	mu      sync.RWMutex
}

func (c *Cache) Add(val []byte) uint32 {
	h := fnv.New32a()
	h.Write(val)
	key := h.Sum32()
	c.buckets[key%256].Add(key, val)
	return key
}

func (c *Cache) Get(key uint32) string {
	return c.buckets[key%256].Get(key)
}

func New() *Cache {
	c := Cache{buckets: make(map[uint32]bucket)}
	for i := 0; i < 256; i++ {
		c.buckets[uint32(i)] = bucket{
			mu:    &sync.RWMutex{},
			items: make(map[uint32]string),
		}
	}
	return &c
}

func main() {
	c := New()
	key1 := c.Add([]byte("hearaasjdas jkdasdj k asd "))
	key2 := c.Add([]byte("hearaasjdas afcvassdf k asd "))
	key3 := c.Add([]byte("hearaasjdas xsada k asd "))

	fmt.Println(c.Get(key1))
	fmt.Println(c.Get(key2))
	fmt.Println(c.Get(key3))
}
