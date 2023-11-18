package api

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type currencyConversionRequest struct {
	Source string `form:"source" binding:"required,currency"`
	Target string `form:"target" binding:"required,currency"`
	Amount int    `form:"amount" binding:"required,min=1"`
}

type currencyConversionResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func (server *Server) convertCurrency(ctx *gin.Context) {
	var req currencyConversionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rates, err := server.cache.ExchangeRates()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rate := rates[req.Source][req.Target]
	targetAmount := float64(req.Amount) * rate
	formatedAmount := formatAmount(targetAmount)

	res := currencyConversionResponse{
		Msg:    "success",
		Amount: formatedAmount,
	}
	ctx.JSON(http.StatusOK, res)
}

func formatAmount(amount float64) string {
	convertedAmount := strconv.FormatFloat(amount, 'f', 2, 64)

	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != convertedAmount; {
		n = convertedAmount
		convertedAmount = re.ReplaceAllString(convertedAmount, "$1,$2")
	}
	return fmt.Sprintf("$%s", convertedAmount)
}
