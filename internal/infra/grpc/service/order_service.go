package service

import (
	"context"

	"github.com/psilva1982/cleanarch_challange/internal/infra/grpc/pb"
	"github.com/psilva1982/cleanarch_challange/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.Blank) (*pb.OrderListResponse, error) {

	data, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.OrderResponse
	for _, order := range data {
		var o = pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			FinalPrice: float32(order.FinalPrice),
			Tax:        float32(order.Tax),
		}
		orders = append(orders, &o)
	}

	return &pb.OrderListResponse{Orders: orders}, nil
}
