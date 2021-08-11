-- nome: CriaIngrediente :one
INSERT INTO Ingredientes (nome, ativo) VALUES ($1, $2)
RETURNING *;