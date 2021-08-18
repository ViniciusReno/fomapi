package api

import (
	"database/sql"
	"net/http"

	db "github.com/ViniciusReno/fomapi/db/sqlc"
	"github.com/gin-gonic/gin"
)

type addIngredientRequest struct {
	Nome  string `json:"nome" binding:"required"`
	Ativo bool   `json:"ativo" binding:"required"`
}

func (server *Server) addIngredient(ctx *gin.Context) {
	var req addIngredientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.AddIngredientParams{
		Nome:  req.Nome,
		Ativo: true,
	}

	ing, err := server.store.AddIngredient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ing)
}

type getIngredientRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getIngredient(ctx *gin.Context) {
	var req getIngredientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ing, err := server.store.GetIngreditent(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ing)
}

type listIngredientRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listIngredient(ctx *gin.Context) {
	var req listIngredientRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListIngredientsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	ings, err := server.store.ListIngredients(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ings)
}
