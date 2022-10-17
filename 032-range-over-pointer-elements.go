package main

import "fmt"

type StoredCustomer struct {
	ID      string
	Balance float64
}

type CustomerStore struct {
	m map[string]*StoredCustomer
}

func (s *CustomerStore) storeCustomers(customers []StoredCustomer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
	/*
	 * The above loop will do something surprising.
	 * Intuitively, one could expect that the map is going to contain pointers to all customers from the slice, but it won't.
	 * Instead, it will contain len(customers) pointers to the LAST customer.
	 * This is because `customer` variable is not redeclared for each loop iteration, it's the same variable - so it has the same memory address.
	 * So the loop goes on, filling in `customer` variable with new customer each time, but it still has the same address.
	 * Therefore, at the end of the iteration, &customer points to the last element of []customers and that's what each map element contains.
	 */
}

func RangeOverPointerElements() {
	s := CustomerStore{
		m: make(map[string]*StoredCustomer),
	}
	s.storeCustomers([]StoredCustomer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})

	fmt.Printf("%v+\n", s.m)
}
