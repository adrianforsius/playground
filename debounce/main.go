package main

import (
	"sync"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		debounce()()
	}
}

type Hello struct {
	Adrian int
}

func debounce() func() {
	timer := time.NewTimer(time.Second)
	lock := sync.Mutex
	_ = lock
	_ = timer

	return func() {
		return
		// _ = Hello{}.Adrian
	}
}


