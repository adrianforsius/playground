package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type bucket struct {
	items map[int]string
	mu    *sync.RWMutex
}

func (b bucket) Get(key int) string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.items[key]
}

func (b bucket) Add(key int, val string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.items[key] = string(val)
}

type Cache struct {
	buckets map[int]bucket
	mu      sync.RWMutex
}

func (c *Cache) Add(val string) int {
	h := fnv.New32a()
	h.Write([]byte(val))
	key := int(h.Sum32())
	c.buckets[key%256].Add(key, val)
	return key
}

func (c *Cache) Get(key int) string {
	return c.buckets[key%256].Get(key)
}

func New() *Cache {
	c := Cache{buckets: make(map[int]bucket)}
	for i := 0; i < 256; i++ {
		c.buckets[i] = bucket{
			mu:    &sync.RWMutex{},
			items: make(map[int]string),
		}
	}
	return &c
}

func main() {
	c := New()
	key1 := c.Add("hearaasjdas jkdasdj k asd ")
	key2 := c.Add("hearaasjdas jkdasdj k asd  ASssSDASssda ")
	key3 := c.Add("hearaasjdas jkdasdj k ")

	fmt.Println(c.Get(key1))
	fmt.Println(c.Get(key2))
	fmt.Println(c.Get(key3))
}
