-- name: AddIngredient :one
INSERT INTO ingredients (
  nome,
  ativo
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetIngreditent :one
SELECT * FROM ingredients
WHERE id = $1 LIMIT 1;

-- name: ListIngredients :many
SELECT * FROM ingredients
ORDER BY nome
LIMIT $1
OFFSET $2;

-- name: UpdateIngredient :one
UPDATE ingredients 
SET nome = $2, ativo = $3
WHERE id = $1
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM ingredients 
WHERE id = $1;
