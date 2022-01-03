package goroutines

import (
	"fmt"
	"testing"
	"time"
)

//BASIC GOROUTINE
func RunHelloWorld() {
	fmt.Println("Hello World")
}
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

//RUN MANY
func DisplayNumber(number int) {
	fmt.Println("Display:", number)
}
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(11 * time.Second)
}
