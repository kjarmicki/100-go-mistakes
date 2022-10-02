package main

import (
	"github.com/kjarmicki/100-go-mistakes/interfaces"
)

/*
 * In Go, iterfaces are satisfied implicitly. This means that the consuming side can define a right
 * level of abstraction and there's no need to expose a best-guess interface from the producer.
 * Some clients may want the whole thing, some may want only one method. Let them decide.
 *
 * There are exceptions though: throughout standard library there are interfaces declared in the producer
 * packages. These interfaces are known to be useful and were discovered (not foreseen) when developing
 * a standard library. If you know the needs of the clients, you may provide useful interfaces.
 */

type customerGetter interface {
	GetCustomer(id string) (interfaces.Customer, error)
}

func InterfaceOnTheProducerSide() {
	doSomeStuffWithCustomer := func(getter customerGetter) error {
		customer, err := getter.GetCustomer("abc")
		if err != nil {
			return err
		}
		_ = customer
		return nil
	}

	doSomeStuffWithCustomer(&interfaces.MySQLConsumerStore{})
}
