package main

import (
	"fmt"
	"hash/fnv"
	"sync"
	"time"
)

const EXPIRATION = 60 * 3
const GARBAGE_COLLECTOR = 30

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
		return "", fmt.Errorf("not found (%d), %+v", key, b.items)
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
	// Add time to trigger weight
	time.Sleep(time.Millisecond)

	h := fnv.New32a()
	h.Write([]byte(val))
	key := int(h.Sum32())
	c.buckets[key%256].Add(key, val)
	return key
}

func (c *Cache) Get(key int) (string, error) {
	// Add time to trigger weight
	time.Sleep(time.Millisecond)
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

// Weight is to balance read and write on a 80/20 basis
func NewWeight(cache Cacher) Cacher {
	addChan := make(chan int)
	getChan := make(chan int)
	go func() {
		read := 0
		for {
			if read < 4 {
				select {
				case <-getChan:
					getChan <- 1
					read++
				default:
				}
			}

			read = 0
			select {
			case <-addChan:
				addChan <- 1
			default:
			}

			select {
			case <-addChan:
				addChan <- 1
			case <-getChan:
				getChan <- 1
			default:
			}

		}

	}()

	return &Weight{cache: cache, get: getChan, add: addChan}
}

type Weight struct {
	cache Cacher
	add   chan int
	get   chan int
}

func (w *Weight) Add(val string) int {
	w.add <- 1
	<-w.add
	fmt.Println("adding value")
	return w.cache.Add(val)
}

func (w *Weight) Get(key int) (string, error) {
	w.get <- 1
	<-w.get
	fmt.Println("getting value")
	return w.cache.Get(key)
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

	var err error
	val1, err := c.Get(key1)
	if err != nil {
		fmt.Println("error with key", key1)
	}
	fmt.Println("val1", val1, "key1", key1)

	val2, err := c.Get(key2)
	if err != nil {
		fmt.Println("error with key", key2)
	}
	fmt.Println("val2", val2, "key2", key2)

	val3, err := c.Get(key3)
	if err != nil {
		fmt.Println("error with key", key3)
	}
	fmt.Println("val3", val3, "key3", key3)

	var keys []int
	for i := 0; i < 500; i++ {
		v := fmt.Sprintf("%d", i)
		key := c.Add(string(v))
		keys = append(keys, key)
	}

	for i, k := range keys {
		go func(key1, i int) {
			// fmt.Println("getting val", key1)
			val, err := c.Get(key1)
			if err != nil {
				fmt.Println("error with key init", key1, "val", val, "err", err)
			}

			if i%3 == 0 {
				v := fmt.Sprintf("%d", key1)
				k := c.Add(v)
				time.Sleep(time.Millisecond * 5)
				value, err := c.Get(k)
				if err != nil {
					fmt.Println("error with key after add", k, "val", value, "error", err)
				}
				if value != v {
					fmt.Println("expected", v, "got", value, "error", err, "key1", k)
				}
			}
			fmt.Println("added value!")
		}(k, i)
	}
	time.Sleep(time.Second * 10)

}
