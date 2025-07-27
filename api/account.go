package api

import (
	"database/sql"
	"net/http"

	db "github.com/Viczdera/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	//get reqBody
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	//account = db.Account{} to check get test

	ctx.JSON(http.StatusOK, account)
}

type getAccountsListRequest struct {
	CurrentPage int32 `form:"current_page" binding:"required,min=1"`
	PageSize    int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getAccountsList(ctx *gin.Context) {
	var req getAccountsListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	args := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.CurrentPage - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, args)

	// log.Printf("Accounts: %+v", accounts)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
