// Code generated by sqlc. DO NOT EDIT.
// source: Ingredients.sql

package db

import (
	"context"
)

const addIngredient = `-- name: AddIngredient :one
INSERT INTO ingredients (
  nome,
  ativo
) VALUES (
  $1, $2
) RETURNING id, nome, ativo, criado_em
`

type AddIngredientParams struct {
	Nome  string `json:"nome"`
	Ativo bool   `json:"ativo"`
}

func (q *Queries) AddIngredient(ctx context.Context, arg AddIngredientParams) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, addIngredient, arg.Nome, arg.Ativo)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Ativo,
		&i.CriadoEm,
	)
	return i, err
}

const deleteIngredient = `-- name: DeleteIngredient :exec
DELETE FROM ingredients 
WHERE id = $1
`

func (q *Queries) DeleteIngredient(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteIngredient, id)
	return err
}

const getIngreditent = `-- name: GetIngreditent :one
SELECT id, nome, ativo, criado_em FROM ingredients
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetIngreditent(ctx context.Context, id int64) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, getIngreditent, id)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Ativo,
		&i.CriadoEm,
	)
	return i, err
}

const listIngredients = `-- name: ListIngredients :many
SELECT id, nome, ativo, criado_em FROM ingredients
ORDER BY nome
LIMIT $1
OFFSET $2
`

type ListIngredientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListIngredients(ctx context.Context, arg ListIngredientsParams) ([]Ingredient, error) {
	rows, err := q.db.QueryContext(ctx, listIngredients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ingredient{}
	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Ativo,
			&i.CriadoEm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateIngredient = `-- name: UpdateIngredient :one
UPDATE ingredients 
SET nome = $2, ativo = $3
WHERE id = $1
RETURNING id, nome, ativo, criado_em
`

type UpdateIngredientParams struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Ativo bool   `json:"ativo"`
}

func (q *Queries) UpdateIngredient(ctx context.Context, arg UpdateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, updateIngredient, arg.ID, arg.Nome, arg.Ativo)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Ativo,
		&i.CriadoEm,
	)
	return i, err
}
