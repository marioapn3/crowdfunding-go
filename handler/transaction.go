package handler

import (
	"crowdfunding/helper"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignTractions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.transactionService.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedTransactions := transaction.FormatCampaignTransactions(transactions)

	response := helper.ApiResponse("Success Get List Of Campaign's Transactions", http.StatusOK, "success", formattedTransactions)
	c.JSON(http.StatusOK, response)

}
