package main

import (
	"fmt"
	"sync"
)

type DeadlockingCustomer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

func (c *DeadlockingCustomer) UpdateAge(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}
	c.age = age
	return nil
}

func (c *DeadlockingCustomer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

func StringDeadlock() {
	c := DeadlockingCustomer{}
	c.UpdateAge(-3)
	// this is going to deadlock. steps:
	// 1. lock mutex in UpdateAge
	// 2. if age is negative, use String() method to print customer
	// 3. lock mutex again in String() - deadlock
}
