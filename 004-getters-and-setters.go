package main

import "fmt"

/* Advantages of getters:
 * - encapsulate behavior (field validation, computing a value)
 * - hide internal representation
 * - provide a debugging point for when the property changes at runtime
 */

type Customer struct {
	balance int
}

// not GetBalance - don't need Get, public members are capitalized
func (c *Customer) Balance() int {
	fmt.Println("getting customer balance at ", c.balance)
	return c.balance
}

func (c *Customer) SetBalance(balance int) {
	fmt.Println("setting customer balance to ", balance)
	c.balance = balance
}

func GettersAndSetters() {
	c := Customer{balance: -10}
	if c.Balance() < 0 {
		c.SetBalance(0)
	}
}
