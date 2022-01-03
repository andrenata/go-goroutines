package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoolPattern(t *testing.T) {
	//var pool sync.Pool
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Wayan")
	pool.Put("Andre")
	pool.Put("Nata")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
}
