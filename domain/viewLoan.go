package domain

type ViewLoanReq struct {
	UserId string `json:"userId"`
}

type ViewLoanRes struct {
	LoanInfo []ViewLoanEntity `json:"loanInfo"`
}
type ViewLoanEntity struct {
	LoanId     int64  `json:"loanId"`
	Amount     int    `json:"amount"`
	Term       int    `json:"term"`
	Status     string `json:"status"`
	TermPaid   int    `json:"termPaid"`
	AmountPaid int    `json:"amountPaid"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
