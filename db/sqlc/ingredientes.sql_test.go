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
func createRandomIngrediente(t *testing.T) Ingrediente {
	arg := AddIngredienteParams{
		Nome:  util.RandomIngrediente(),
		Ativo: true,
	}

	Ingrediente, err := testQueries.AddIngrediente(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Ingrediente)

	require.Equal(t, arg.Ativo, Ingrediente.Ativo)
	require.Equal(t, arg.Nome, Ingrediente.Nome)

	require.NotZero(t, Ingrediente.ID)
	require.NotZero(t, Ingrediente.CriadoEm)

	return Ingrediente
}

func TestAddIngrediente(t *testing.T) {
	createRandomIngrediente(t)
}

func TestGetIngreditente(t *testing.T) {
	ing1 := createRandomIngrediente(t)
	ing2, err := testQueries.GetIngreditente(context.Background(), ing1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, ing2)
	require.Equal(t, ing1.ID, ing2.ID)
	require.Equal(t, ing1.Nome, ing2.Nome)
	require.WithinDuration(t, ing1.CriadoEm, ing2.CriadoEm, time.Second)
}

func TestUpdateIngrediente(t *testing.T) {
	ing1 := createRandomIngrediente(t)
	arg := UpdateIngredienteParams{
		ID:    ing1.ID,
		Nome:  ing1.Nome,
		Ativo: false,
	}
	ing2, err := testQueries.UpdateIngrediente(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ing2)
	require.Equal(t, ing1.ID, ing2.ID)
	require.Equal(t, ing1.Nome, ing2.Nome)
	require.NotEqual(t, ing1.Ativo, ing2.Ativo)
	require.WithinDuration(t, ing1.CriadoEm, ing2.CriadoEm, time.Second)
}

func TestDeleteIngrediente(t *testing.T) {
	ing1 := createRandomIngrediente(t)
	err := testQueries.DeleteIngrediente(context.Background(), ing1.ID)
	require.NoError(t, err)

	ing2, err := testQueries.GetIngreditente(context.Background(), ing1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ing2)
}

func TestListIngrediente(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIngrediente(t)
	}

	arg := ListIngredienteParams{
		Limit:  5,
		Offset: 5,
	}

	ings, err := testQueries.ListIngrediente(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ings, 5)

	for _, ing := range ings {
		require.NotEmpty(t, ing)
	}
}
