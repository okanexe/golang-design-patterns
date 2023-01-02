package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type fiboFunc func(int) int

func wraplogger(fun fiboFunc, logger *log.Logger) fiboFunc {
	return func(n int) int {
		fn := func(n int) (result int) {
			defer func(t time.Time) {
				logger.Printf("took=%v, n=%v, result=%v", time.Since(t), n, result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

func wrapcache(fun fiboFunc, cache *sync.Map) fiboFunc {
	return func(n int) int {
		fn := func(n int) int {
			key := fmt.Sprintf("n=%d", n)
			val, ok := cache.Load(key)
			if ok {
				return val.(int)
			}

			result := fun(n)
			cache.Store(key, result)
			return result
		}
		return fn(n)
	}
}

func fibonacci(n int) int {
	ch := make(chan int)
	quit := make(chan bool)
	var result int

	go func(n int) {
		for i := 0; i < n-1; i++ {
			fmt.Println(<-ch)
		}
		result = <-ch
		quit <- false
	}(n)

	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return result
		}
	}
}

func main() {
	// f := wraplogger(fibonacci, log.New(os.Stdout, "Test ", 1))
	// f(10)

	f := wrapcache(fibonacci, &sync.Map{})
	g := wraplogger(f, log.New(os.Stdout, "Test ", 1))

	g(11)
}
