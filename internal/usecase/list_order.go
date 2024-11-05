package usecase

import (
	"github.com/psilva1982/cleanarch_challange/internal/entity"
	"github.com/psilva1982/cleanarch_challange/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderList    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	//OrderList events.EventInterface,
	//EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		//OrderList:    OrderList,
		//EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	var data []OrderOutputDTO
	orders, err := l.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, nil
	}

	for _, order := range orders {
		data = append(data, OrderOutputDTO{
			ID: order.ID,
			Price: order.Price,
			Tax: order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return data, nil
}