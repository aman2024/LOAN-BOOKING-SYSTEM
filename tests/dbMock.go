package tests

import (
	"loan-booking/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type DbClientMock struct {
	mock.Mock
}

// ReadLoanInfo mocks the ReadLoanInfo method
func (m *DbClientMock) ReadLoanInfo(c *gin.Context, userId string, loanId int) ([]domain.ViewLoanEntity, error) {
	args := m.Called(c, userId, loanId)
	return args.Get(0).([]domain.ViewLoanEntity), args.Error(1)
}

// Insert mocks the Insert method
func (m *DbClientMock) Insert(c *gin.Context, tableName string, columns []string, values []interface{}) (int64, error) {
	args := m.Called(c, tableName, columns, values)
	return args.Get(0).(int64), args.Error(1)
}

// UpdateStatus mocks the UpdateStatus method
func (m *DbClientMock) UpdateStatus(c *gin.Context, req *domain.ApproveLoanReq) (int64, error) {
	args := m.Called(c, req)
	return args.Get(0).(int64), args.Error(1)
}

// UpdateRepayments mocks the UpdateRepayments method
func (m *DbClientMock) UpdateRepayments(c *gin.Context, req *domain.AddRepaymentReq, status string) (int64, error) {
	args := m.Called(c, req, status)
	return args.Get(0).(int64), args.Error(1)
}
