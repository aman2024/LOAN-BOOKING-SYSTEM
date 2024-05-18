package domain

type AddRepaymentReq struct {
	UserId     string `json:"userId"`
	LoanId     int    `json:"loanId"`
	TermAmount int    `json:"termAmount"`
	TermNo     int    `json:"termNo"`
}

type AddRepaymentRes struct {
	LoanId  int64 `json:"loanId"`
	Balance int   `json:"balance"`
}
