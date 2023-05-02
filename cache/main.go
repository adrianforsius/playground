package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"sync"
	"time"
)

const EXPIRATION = 1
const GARBAGE_COLLECTOR = 2

type item struct {
	expire time.Time
	value  string
}

type bucket struct {
	items map[int]item
	mu    *sync.RWMutex
}

func (b bucket) Get(key int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	item, ok := b.items[key]
	if !ok {
		return "", errors.New("not found")
	}
	return item.value, nil
}

func (b bucket) Add(key int, val string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.items[key] = item{time.Now().Add(time.Second * EXPIRATION), val}
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

func (c *Cache) Get(key int) (string, error) {
	return c.buckets[key%256].Get(key)
}

func New() *Cache {
	c := Cache{buckets: make(map[int]bucket)}
	for i := 0; i < 256; i++ {
		c.buckets[i] = bucket{
			mu:    &sync.RWMutex{},
			items: make(map[int]item),
		}
	}

	go func() {
		for {
			time.Sleep(time.Second * GARBAGE_COLLECTOR)
			c.mu.Lock()
			for _, bucket := range c.buckets {
				for idx, itm := range bucket.items {
					bucket.mu.Lock()
					if itm.expire.Before(time.Now()) {
						bucket.items[idx] = item{}
					}
					bucket.mu.Unlock()
				}
			}
			c.mu.Unlock()
		}
	}()

	return &c
}

func main() {
	c := New()
	key1 := c.Add("hearaasjdas jkdasdj k asd ")
	key2 := c.Add("hearaasjdas jkdasdj k asd  ASssSDASssda ")
	key3 := c.Add("hearaasjdas jkdasdj k ")

	var err error
	val1, err := c.Get(key1)
	if err != nil {
		fmt.Println("error with key", key1)
	}
	fmt.Println("val1", val1)

	time.Sleep(time.Second * 3)

	val2, err := c.Get(key2)
	if err != nil {
		fmt.Println("error with key", key2)
	}
	fmt.Println("val2", val2)

	val3, err := c.Get(key3)
	if err != nil {
		fmt.Println("error with key", key3)
	}
	fmt.Println("val3", val3)
}
