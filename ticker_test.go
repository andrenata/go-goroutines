package goroutines

import (
	"fmt"
	"testing"
	"time"
)

//Ticker adalah representasi kejadian yang berulang
//Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
//Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
//Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}
