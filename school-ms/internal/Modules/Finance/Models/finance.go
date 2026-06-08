package models

import "time"

type FeeType struct {
	ID          int64   `db:"id"`
	SchoolID    int64   `db:"school_id"`
	Name        string  `db:"name"`
	Amount      float64 `db:"amount"`
	Frequency   string  `db:"frequency"` // termly, monthly, annual, once
	IsMandatory bool    `db:"is_mandatory"`
}

type FeeInvoice struct {
	ID        int64     `db:"id"`
	StudentID int64     `db:"student_id"`
	FeeTypeID int64     `db:"fee_type_id"`
	TermID    int64     `db:"term_id"`
	Amount    float64   `db:"amount"`
	Status    string    `db:"status"` // unpaid, partial, paid
	DueDate   string    `db:"due_date"`
	CreatedAt time.Time `db:"created_at"`
}

type FeePayment struct {
	ID         int64     `db:"id"`
	InvoiceID  int64     `db:"invoice_id"`
	AmountPaid float64   `db:"amount_paid"`
	Method     string    `db:"method"` // cash, mpesa, bank, cheque
	RefNo      string    `db:"ref_no"`
	PaidAt     time.Time `db:"paid_at"`
}

type StudentFeeStatement struct {
	StudentName string       `json:"student_name"`
	AdmissionNo string       `json:"admission_no"`
	Invoices    []FeeInvoice `json:"invoices"`
	TotalBilled float64      `json:"total_billed"`
	TotalPaid   float64      `json:"total_paid"`
	Balance     float64      `json:"balance"`
}
