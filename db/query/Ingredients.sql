-- name: AddIngredient :one
INSERT INTO Ingredients (
  nome,
  ativo
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetIngreditent :one
SELECT * FROM Ingredients
WHERE id = $1 LIMIT 1;

-- name: ListIngredients :many
SELECT * FROM Ingredients
ORDER BY nome
LIMIT $1
OFFSET $2;

-- name: UpdateIngredient :one
UPDATE Ingredients 
SET nome = $2 AND ativo = $3
WHERE id = $1
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM Ingredients 
WHERE id = $1;
