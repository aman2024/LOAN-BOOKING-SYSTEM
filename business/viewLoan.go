package business

import (
	"loan-booking/domain"
	"loan-booking/services"

	"github.com/gin-gonic/gin"
)

func ViewLoan(c *gin.Context, services *services.Services, req *domain.ViewLoanReq) (*domain.ViewLoanRes, error) {
	var res domain.ViewLoanRes
	result, err := services.DB.ReadLoanInfo(c, req.UserId, 0)
	if err != nil {
		return nil, err
	} else {
		res.LoanInfo = result
	}

	return &res, err

}
