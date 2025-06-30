package handler

import (
	"net/http"
	"time"

	"github.com/dinodorinna/personal-expense-tracker-be/internal/repo"
	"github.com/gin-gonic/gin"
)

type GetTransactionsHandlerRequest struct {
	Date  time.Time `form:"date" binding:"required" time_format:"2006-01-02"`
	Limit int       `form:"limit" binding:"required,gte=1,lte=100"`
	Page  int       `form:"page" binding:"required,gte=1"`
}

func GetTransactionsHandler(ctx *gin.Context) {

	var input GetTransactionsHandlerRequest

	err := ctx.ShouldBindQuery(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := repo.FindAllTransactionByDate(ctx.Request.Context(), input.Date, input.Limit, input.Page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})

}

func CreateTransactionHandler(ctx *gin.Context) {

	transaction := repo.Transactions{
		Date:        time.Unix(1751294006, 0),
		Description: "lunch",
		Amount:      200.56,
		Category:    "Food",
	}

	err := repo.SaveTransaction(ctx.Request.Context(), &transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": transaction})

}
