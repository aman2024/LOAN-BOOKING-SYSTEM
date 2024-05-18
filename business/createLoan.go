package business

import (
	"loan-booking/domain"
	"loan-booking/services"

	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context, services *services.Services, req *domain.CreateLoanReq) (*domain.CreateLoanRes, error) {
	var res domain.CreateLoanRes
	result, err := services.DB.Insert(c, "loan_info", []string{"user_id", "amount", "term", "status"}, []interface{}{req.UserId, req.Amount, req.Term, domain.STATUS_PENDING})
	if err != nil {
		return nil, err
	} else {
		res.LoanId = result
	}

	return &res, err

}
