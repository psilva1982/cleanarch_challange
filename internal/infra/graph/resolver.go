package graph

import (
	"github.com/psilva1982/cleanarch_challange/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase usecase.ListOrderUseCase
}
