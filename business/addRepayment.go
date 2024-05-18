package business

import (
	"errors"
	"fmt"
	"loan-booking/domain"
	"loan-booking/services"

	"github.com/gin-gonic/gin"
)

func AddRepayment(c *gin.Context, services *services.Services, req *domain.AddRepaymentReq) (*domain.AddRepaymentRes, error) {
	var res domain.AddRepaymentRes

	loanEntity, err := services.DB.ReadLoanInfo(c, req.UserId, req.LoanId)
	if err != nil {
		return nil, err
	} else if len(loanEntity) != 1 {
		return nil, errors.New("no loan found")
	}

	switch loanEntity[0].Status {
	case domain.STATUS_PENDING:
		return nil, errors.New("loan in PENDING state, Needs to be approved")
	case domain.STATUS_PAID:
		return nil, errors.New("loan is already PAID")
	}

	var status string
	if req.TermNo <= loanEntity[0].TermPaid {
		return nil, errors.New("termNo already paid")
	} else if req.TermNo == loanEntity[0].TermPaid+1 {
		if req.TermNo < loanEntity[0].Term {
			minAmountThatsRequired := loanEntity[0].Amount / loanEntity[0].Term
			maxAmountThatCanBePaid := loanEntity[0].Amount - loanEntity[0].AmountPaid

			if req.TermAmount == maxAmountThatCanBePaid {
				status = domain.STATUS_PAID
			} else if req.TermAmount < minAmountThatsRequired {
				return nil, errors.New("termAmount is less than the required term amount")
			} else if req.TermAmount > maxAmountThatCanBePaid {
				return nil, errors.New("termAmount is greater than the required term amount")
			}

		} else if req.TermNo == loanEntity[0].Term {
			AmountThatCanBePaid := loanEntity[0].Amount - loanEntity[0].AmountPaid

			if req.TermAmount != AmountThatCanBePaid {
				return nil, errors.New(fmt.Sprint("since this is the last term, please pay the accurate balance i.e. : ", AmountThatCanBePaid))
			} else {
				status = domain.STATUS_PAID
			}
		}
	} else {
		return nil, errors.New("invalid termNo - middle terms needs to be paid")
	}
	req.TermAmount += loanEntity[0].AmountPaid
	result, err := services.DB.UpdateRepayments(c, req, status)
	if err != nil {
		return nil, err
	} else if result != 1 {
		return nil, errors.New("no row affected")
	} else {
		res.LoanId = int64(req.LoanId)
		res.Balance = loanEntity[0].Amount - req.TermAmount
	}

	return &res, nil

}
