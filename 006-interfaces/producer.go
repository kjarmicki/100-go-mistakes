package interfaces

type Customer struct{}

// Anitpattern. Don't expose interface from the producer package, let the client decide what abstraction it needs.
type ConsumerStore interface {
	StoreCustomer(customer Customer) error
	UpdateCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	// (...) etc. - a full blown store description
}

type MySQLConsumerStore struct{}

func (s *MySQLConsumerStore) StoreCustomer(customer Customer) error {
	return nil
}

func (s *MySQLConsumerStore) UpdateCustomer(customer Customer) error {
	return nil
}

func (s *MySQLConsumerStore) GetCustomer(id string) (Customer, error) {
	return Customer{}, nil
}
