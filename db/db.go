package db

import (
	"database/sql"
	"loan-booking/domain"

	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type DbClient struct {
	Client *sql.DB
}

func (dbClient *DbClient) ReadLoanInfo(c *gin.Context, userId string, loanId int) ([]domain.ViewLoanEntity, error) {
	var res []domain.ViewLoanEntity
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return res, err
	}
	defer conn.Close()

	qbuilder := squirrel.Select("id", "amount", "term", "status", "term_paid", "amount_paid", "created_at", "updated_at").
		From("loan_info")

	if userId != "" {
		qbuilder = qbuilder.Where("user_id = ?", userId)
	}
	if loanId != 0 {
		qbuilder = qbuilder.Where("id = ?", loanId)
	}

	query, qargs, err := qbuilder.ToSql()
	if err != nil {
		return res, err
	}
	rows, err := conn.QueryContext(c, query, qargs...)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		entity := domain.ViewLoanEntity{}
		err := rows.Scan(&entity.LoanId, &entity.Amount, &entity.Term, &entity.Status, &entity.TermPaid, &entity.AmountPaid, &entity.CreatedAt, &entity.UpdatedAt)
		if err != nil {
			return res, err
		}
		res = append(res, entity)
	}
	return res, nil

}

func (dbClient *DbClient) Insert(c *gin.Context, tableName string, columns []string, values []interface{}) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Insert(tableName).
		Columns(columns...).
		Values(values...)

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dbClient *DbClient) UpdateStatus(c *gin.Context, req *domain.ApproveLoanReq) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Update("loan_info").
		Set("status", domain.STATUS_APPROVED).Set("admin_approver_id", req.AdminId).
		Where(squirrel.Eq{"id": req.LoanId})

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (dbClient *DbClient) UpdateRepayments(c *gin.Context, req *domain.AddRepaymentReq, status string) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Update("loan_info").
		Set("term_paid", req.TermNo).Set("amount_paid", req.TermAmount).
		Where(squirrel.Eq{"id": req.LoanId})

	if status == domain.STATUS_PAID {
		qbuilder = qbuilder.Set("status", "PAID")
	}

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
