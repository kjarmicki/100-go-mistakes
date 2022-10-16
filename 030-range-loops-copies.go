package main

import "fmt"

type account struct {
	balance int
}

func RangeLoopsCopies() {
	accounts := []account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	// in Go, every assignment is a copy
	// if a function returns a struct, the assignment will perform copy of that struct
	// if a function returns a pointer, the assignment will perform copy of the memory address
	for _, a := range accounts {
		a.balance += 1000
	}
	fmt.Println(accounts) // nothing was changed because assignment copied a structure and performed action on that copy

	// here the access is through the slice index so it goes directly to the element
	for i, _ := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts) // 1000 was added to each account's balance

	// by the same logic:
	fmt.Println(accounts[0].balance) // 1100
	acc0 := accounts[0]              // creates a copy, not a reference!
	acc0.balance += 1000
	accounts[0].balance -= 1000
	fmt.Println("in slice", accounts[0].balance) // 100
	fmt.Println("standalone", acc0.balance)      // 2100
}
