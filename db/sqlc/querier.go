// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	AddIngrediente(ctx context.Context, arg AddIngredienteParams) (Ingrediente, error)
	DeleteIngrediente(ctx context.Context, id int64) error
	GetIngreditente(ctx context.Context, id int64) (Ingrediente, error)
	ListIngrediente(ctx context.Context, arg ListIngredienteParams) ([]Ingrediente, error)
	UpdateIngrediente(ctx context.Context, arg UpdateIngredienteParams) (Ingrediente, error)
}

var _ Querier = (*Queries)(nil)