package tests

import (
	"loan-booking/business"
	"loan-booking/domain"
	"loan-booking/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddRepayment(t *testing.T) {

	t.Run("TestAddRepayment", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Term: 2, Amount: 500, Status: domain.STATUS_APPROVED}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)
		dbClient.On("UpdateRepayments", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.AddRepaymentReq{
			LoanId:     1,
			UserId:     "1",
			TermNo:     1,
			TermAmount: 300,
		}
		res, err := business.AddRepayment(&gin.Context{}, &service, &req)

		expectedRes := domain.AddRepaymentRes{LoanId: 1, Balance: 200}
		assert.NoError(t, err)
		assert.Equal(t, &expectedRes, res)
	})

	t.Run("TestAddRepayment", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Term: 2, Amount: 500, Status: domain.STATUS_APPROVED}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.AddRepaymentReq{
			LoanId:     1,
			UserId:     "1",
			TermNo:     2,
			TermAmount: 100,
		}
		_, err := business.AddRepayment(&gin.Context{}, &service, &req)

		assert.Error(t, err)
	})

	t.Run("TestAddRepayment", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Term: 2, Amount: 500, Status: domain.STATUS_APPROVED, TermPaid: 1, AmountPaid: 300}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)
		dbClient.On("UpdateRepayments", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.AddRepaymentReq{
			LoanId:     1,
			UserId:     "1",
			TermNo:     2,
			TermAmount: 200,
		}
		res, err := business.AddRepayment(&gin.Context{}, &service, &req)

		expectedRes := domain.AddRepaymentRes{LoanId: 1, Balance: 0}
		assert.NoError(t, err)
		assert.Equal(t, &expectedRes, res)
	})

}
