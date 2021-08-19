CREATE TABLE IF NOT EXISTS "ingredients" (
    "id" bigserial PRIMARY KEY,
    "nome" VARCHAR NOT NULL,
    "ativo" BOOLEAN NOT NULL,
    "criado_em" timestamptz NOT NULL DEFAULT (now())
);

-- CREATE TABLE "receitas" (
--     id bigserial PRIMARY KEY,
--     nome VARCHAR(100) NOT NULL,
--     preparo VARCHAR(10485760) NOT NULL,
--     "criado_em" timestamptz NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE "receita_Ingredients" (
--     id bigserial PRIMARY KEY,
--     id_Ingredient INT NOT NULL,
--     id_receita INT NOT NULL,
--     CONSTRAINT fk_Receitas FOREIGN KEY(id_receita) REFERENCES "Receitas"(id),
--     CONSTRAINT fk_Ingredients FOREIGN KEY(id_Ingredient) REFERENCES "Ingredients"(id)
-- );
