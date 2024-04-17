package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		debounce()()
	}
}

func debounce() func() {
	timer := time.NewTimer(time.Second)
	lock := sync.Mutex

	return func() {
		switch {
			timer <-
		}
	}
}


