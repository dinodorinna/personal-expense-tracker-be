package main

import (
	"github.com/dinodorinna/personal-expense-tracker-be/internal/handler"
	"github.com/dinodorinna/personal-expense-tracker-be/internal/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	repo.SetupDatabase()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/transactions", handler.GetTransactionsHandler)
	r.POST("/transaction", handler.CreateTransactionHandler)
	r.DELETE("/delete/:id", handler.DeleteTransactionHanler)
	r.Run()
}
