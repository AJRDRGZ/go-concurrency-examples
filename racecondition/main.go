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
	return fmt.Sprintf("%s: Balance: %d\t", a.name, a.balance)
}

func main() {
	alexys := account{name: "Alexys", balance: 500}
	beto := account{name: "Beto", balance: 900}

	for i := 0; i <= 5; i++ {
		go transfer(100, &alexys, &beto) // bad
	}

	time.Sleep(time.Second)
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("%s: transfer denied the current balance is: %d\n",
			source.name,
			source.balance)
		return
	}

	dest.balance += amount
	source.balance -= amount

	fmt.Println(source, dest)
}
