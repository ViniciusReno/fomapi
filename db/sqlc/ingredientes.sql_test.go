package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ViniciusReno/fomapi/util"
	"github.com/stretchr/testify/require"
)

// TODO MOCK
func createRandomIngredient(t *testing.T) Ingredient {
	arg := AddIngredientParams{
		Nome:  util.RandomIngredient(),
		Ativo: true,
	}

	ingredient, err := testQueries.AddIngredient(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ingredient)

	require.Equal(t, arg.Ativo, ingredient.Ativo)
	require.Equal(t, arg.Nome, ingredient.Nome)

	require.NotZero(t, ingredient.ID)
	require.NotZero(t, ingredient.CriadoEm)

	return ingredient
}

func TestAddIngredient(t *testing.T) {
	createRandomIngredient(t)
}

func TestGetIngreditente(t *testing.T) {
	ing1 := createRandomIngredient(t)
	ing2, err := testQueries.GetIngreditent(context.Background(), ing1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, ing2)
	require.Equal(t, ing1.ID, ing2.ID)
	require.Equal(t, ing1.Nome, ing2.Nome)
	require.WithinDuration(t, ing1.CriadoEm, ing2.CriadoEm, time.Second)
}

func TestUpdateIngredient(t *testing.T) {
	ing1 := createRandomIngredient(t)
	arg := UpdateIngredientParams{
		ID:    ing1.ID,
		Nome:  ing1.Nome,
		Ativo: false,
	}
	ing2, err := testQueries.UpdateIngredient(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ing2)
	require.Equal(t, ing1.ID, ing2.ID)
	require.Equal(t, ing1.Nome, ing2.Nome)
	require.NotEqual(t, ing1.Ativo, ing2.Ativo)
	require.WithinDuration(t, ing1.CriadoEm, ing2.CriadoEm, time.Second)
}

func TestDeleteIngredient(t *testing.T) {
	ing1 := createRandomIngredient(t)
	err := testQueries.DeleteIngredient(context.Background(), ing1.ID)
	require.NoError(t, err)

	ing2, err := testQueries.GetIngreditent(context.Background(), ing1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ing2)
}

func TestListIngredient(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIngredient(t)
	}

	arg := ListIngredientsParams{
		Limit:  5,
		Offset: 5,
	}

	ings, err := testQueries.ListIngredients(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ings, 5)

	for _, ing := range ings {
		require.NotEmpty(t, ing)
	}
}
