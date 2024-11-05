package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/psilva1982/cleanarch_challange/internal/infra/graph/model"
	"github.com/psilva1982/cleanarch_challange/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

func (r *queryResolver) ListOrder(ctx context.Context) ([]*model.Order, error) {

	data, err := r.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*model.Order
	for _, order := range data {
		var o = model.Order{
			ID:         order.ID,
			Price:      order.Price,
			FinalPrice: order.FinalPrice,
			Tax:        order.Tax,
		}
		orders = append(orders, &o)
	}

	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
