package handler

import (
	"errors"
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

type Category string

const (
	CategoryFood           Category = "FOOD"
	CategoryDrink          Category = "DRINK"
	CategoryCosmetic       Category = "COSMETIC"
	CategoryTransportation Category = "TRANSPORTATION"
	CategoryOthers         Category = "OTHERS"
)

type CreateTransactionHandlerRequest struct {
	Date        time.Time `json:"date" binding:"required"`
	Description string    `json:"description" binding:"required,max=50"`
	Amount      float64   `json:"amount" binding:"required,gte=1"`
	Category    Category  `json:"category" binding:"required"`
}

func ValidateCategory(category Category) error {
	switch category {
	case CategoryFood, CategoryDrink, CategoryCosmetic, CategoryTransportation, CategoryOthers:
		return nil
	default:
		return errors.New("invalid status")
	}
}

func CreateTransactionHandler(ctx *gin.Context) {
	var transactionInput CreateTransactionHandlerRequest

	err := ctx.BindJSON(&transactionInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = ValidateCategory(transactionInput.Category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	transaction := repo.Transactions{
		Date:        transactionInput.Date,
		Description: transactionInput.Description,
		Amount:      transactionInput.Amount,
		Category:    string(transactionInput.Category),
	}

	err = repo.SaveTransaction(ctx.Request.Context(), &transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transaction})
}

type DeleteTransactionHanlerRequest struct {
	ID uint `uri:"id"`
}

func DeleteTransactionHanler(ctx *gin.Context) {
	var req DeleteTransactionHanlerRequest

	err := ctx.BindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = repo.FindTransactionByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repo.DeleteTransaction(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
