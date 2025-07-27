package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/Viczdera/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createTransferRequest struct {
	FromAccount int64 `json:"from_account" binding:"required"`
	ToAccount   int64 `json:"to_account" binding:"required"`
	// Currency    string `json:"currency" binding:"required,oneof=USD EUR"`
	Currency string `json:"currency" binding:"required,currency"`
	Amount   int64  `json:"amount" binding:"required"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccount, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.ToAccount, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccount: req.FromAccount,
		ToAccount:   req.ToAccount,
		Amount:      req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountId, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return false
	}
	return true
}
