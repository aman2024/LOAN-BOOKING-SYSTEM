package handler

import (
	"errors"
	"loan-booking/business"
	"loan-booking/domain"
	"loan-booking/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewLoanHandler(c *gin.Context, services *services.Services) {
	req, err := validateViewLoanReq(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":          false,
			"error":           "INVALID_REQ",
			"error_statement": err.Error(),
		})
		return
	}
	res, err := business.ViewLoan(c, services, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":          false,
			"error":           "VIEW_LOAN_ERROR",
			"error_statement": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   res,
	})
}

func validateViewLoanReq(c *gin.Context) (*domain.ViewLoanReq, error) {
	req := domain.ViewLoanReq{}

	if userId, ok := c.Get("userId"); ok {
		req.UserId = userId.(string)
	} else {
		return nil, errors.New("invalid userId")
	}

	return &req, nil
}
