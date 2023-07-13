package main

type Order struct {
}

func (o Order) Validate() error {
	return nil
}

type OrderRepository interface {
	Save(order Order) error
	GetByID(id int) (Order, error)
}

type OrderService struct {
	orderRepository OrderRepository
}

func (o *OrderService) CreateOrder(order Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	if err := o.orderRepository.Save(order); err != nil {
		return err
	}

	return nil
}
