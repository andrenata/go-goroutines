package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMapGoroutine(t *testing.T) {
	var data sync.Map

	var addToMap = func(value int) {
		data.Store(value, value)
	}
	for i := 0; i < 100; i++ {
		go addToMap(i)
	}

	time.Sleep(3 * time.Second)
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

//MAP AND GROUP
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}
func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
