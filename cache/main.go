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
	// add buckets to reduce lock contention
	buckets map[int]bucket
	mu      sync.RWMutex
}

func (c *Cache) Add(val string) int {
	h := fnv.New32a()
	time.Sleep(time.Millisecond * 100)
	h.Write([]byte(val))
	key := int(h.Sum32())
	c.buckets[key%256].Add(key, val)
	return key
}

func (c *Cache) Get(key int) (string, error) {
	time.Sleep(time.Millisecond * 100)
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

type add struct {
	key   int
	value string
}

type get struct {
	key   int
	value string
	err   error
}

// Weight is to balance read and write on a 80/20 basis
func NewWeight(cache Cacher) *Weight {
	addChan := make(chan add)
	getChan := make(chan get)
	go func() {
		read := 0
		for {
			if read < 4 {
				select {
				case r := <-getChan:
					val, err := cache.Get(r.key)
					getChan <- get{value: val, err: err}
					read++
				default:
				}
			}

			read = 0
			select {
			case r := <-addChan:
				addChan <- add{key: cache.Add(r.value)}
			default:
			}

			select {
			case r := <-addChan:
				addChan <- add{key: cache.Add(r.value)}
			case r := <-getChan:
				val, err := cache.Get(r.key)
				getChan <- get{value: val, err: err}
			default:
			}

		}

	}()

	return &Weight{cache: cache, get: getChan, add: addChan}
}

type Weight struct {
	cache Cacher
	add   chan add
	get   chan get
}

func (w *Weight) Add(val string) int {
	w.add <- add{value: val}
	resp := <-w.add
	return resp.key
}

func (w *Weight) Get(key int) (string, error) {
	w.get <- get{key: key}
	resp := <-w.get
	return resp.value, resp.err
}

type Cacher interface {
	Add(string) int
	Get(int) (string, error)
}

func main() {
	c := NewWeight(New())
	fmt.Println("adding values")
	key1 := c.Add("hearaasjdas jkdasdj k asd ")
	key2 := c.Add("hearaasjdas jkdasdj k asd  ASssSDASssda ")
	key3 := c.Add("hearaasjdas jkdasdj k ")

	for i := 0; i < 5_000; i++ {
		go func(i int) {
			val, err := c.Get(key1)
			if err != nil {
				fmt.Println("error with key", key1, val, err)
			}
			fmt.Println("got value!")
			if i%3 == 0 {
				c.Add(fmt.Sprintf("hearaasjdas jkdasdj k  %d", i))
			}
			fmt.Println("added value!")
		}(i)
	}

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
