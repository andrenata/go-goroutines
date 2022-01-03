package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//DEADLOCK (PROBLEM)
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) UnLock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user1.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{Name: "Andre Nata", Balance: 100000}

	user2 := UserBalance{Name: "Yulia Dewi", Balance: 100000}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)

	time.Sleep(3 * time.Second)

	fmt.Println("User", user1.Name, "balance", user1.Balance)
	fmt.Println("User", user2.Name, "balance", user2.Balance)

}

//DEADLOCK SOLUTION
func RunAsync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("Complate")
}
