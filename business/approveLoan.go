package business

import (
	"errors"
	"loan-booking/domain"
	"loan-booking/services"

	"github.com/gin-gonic/gin"
)

func ApproveLoan(c *gin.Context, services *services.Services, req *domain.ApproveLoanReq) (*domain.ApproveLoanRes, error) {
	var res domain.ApproveLoanRes

	loanEntity, err := services.DB.ReadLoanInfo(c, "", req.LoanId)
	if err != nil {
		return nil, err
	} else if len(loanEntity) != 1 {
		return nil, errors.New("no loan found")
	}

	if loanEntity[0].Status != "PENDING" {
		return nil, errors.New("loan not in PENDING state")
	}

	result, err := services.DB.UpdateStatus(c, req)
	if err != nil {
		return nil, err
	} else if result != 1 {
		return nil, errors.New("no row affected")
	} else {
		res.LoanId = int64(req.LoanId)
		res.Status = domain.STATUS_APPROVED
	}

	return &res, nil

}
