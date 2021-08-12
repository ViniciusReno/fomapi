-- name: AddIngrediente :one
INSERT INTO Ingredientes (
  nome,
  ativo
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetIngreditente :one
SELECT * FROM Ingredientes
WHERE id = $1 LIMIT 1;

-- name: ListIngrediente :many
SELECT * FROM Ingredientes
ORDER BY nome
LIMIT $1
OFFSET $2;

-- name: UpdateIngrediente :one
UPDATE Ingredientes 
SET nome = $2 AND ativo = $3
WHERE id = $1
RETURNING *;

-- name: DeleteIngrediente :exec
DELETE FROM Ingredientes 
WHERE id = $1;
