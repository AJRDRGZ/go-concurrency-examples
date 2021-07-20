package main

import (
	"fmt"
	"time"
)

type account struct {
	name    string
	balance int
}

func (a *account) String() string {
	return fmt.Sprintf("(%s: Balance: %d)", a.name, a.balance)
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("❌: %s\n", fmt.Sprintf("%s %s", source, dest))
		return
	}
	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount

	fmt.Printf("✅: %s\n", fmt.Sprintf("%s %s", source, dest))
}

type bankOperation struct {
	amount int
	done   chan struct{}
}

func main() {
	signal := make(chan struct{})
	transaction := make(chan *bankOperation)

	alexys := account{name: "Alexys", balance: 500}
	beto := account{name: "Beto", balance: 900}

	// cajero
	go func() {
		for {
			request := <-transaction
			transfer(request.amount, &alexys, &beto)
			request.done <- struct{}{}
		}
	}()

	for _, value := range []int{300, 300} {
		go func(amount int) {
			requestTransaction := bankOperation{amount: amount, done: make(chan struct{})}
			transaction <- &requestTransaction

			signal <- <-requestTransaction.done
		}(value)
	}

	<-signal
	<-signal
}
