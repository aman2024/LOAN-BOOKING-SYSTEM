package db

import (
	"loan-booking/domain"

	"github.com/gin-gonic/gin"
)

type SQLDbQuery interface {
	ReadLoanInfo(c *gin.Context, userId string, loanId int) ([]domain.ViewLoanEntity, error)
	Insert(c *gin.Context, tableName string, columns []string, values []interface{}) (int64, error)
	UpdateStatus(c *gin.Context, req *domain.ApproveLoanReq) (int64, error)
	UpdateRepayments(c *gin.Context, req *domain.AddRepaymentReq, status string) (int64, error)
}
