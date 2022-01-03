package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //close channel wajib

	//data to channel
	//channel <- "Andre"

	//channel to variable
	//data := <-channel

	//fmt.Println(<-channel)

	//annonimious func
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Andre Nata"
		fmt.Println("Selesai mengirim data")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

//CHANNEL AS PARAMETER
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Andre Nata"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

//CHANNEL IN AND OUT
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Andre Nata"
}
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
func TestOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

//BUFFERED CHANNEL
func TestBufferedChannel(t *testing.T) {
	//default channel 1
	//channel := make(chan string)

	//buffered channel custom
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Andre Nata"
		channel <- "Yulia Dewi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println(cap(channel)) // melihat panjang channel
	fmt.Println(len(channel)) // melihat jumlah data di channel

	time.Sleep(2 * time.Second)
}

//RANGE CHANNEL
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i <= 100; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

//SELECT MULTIPLE CHANNEL
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go func() {
		for i := 0; i <= 20; i++ {
			channel1 <- "Data1"
			channel2 <- "Data2"
		}
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		}
		if counter == 10 {
			break
		}
	}
}

//DEFAULT SELECT
func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
