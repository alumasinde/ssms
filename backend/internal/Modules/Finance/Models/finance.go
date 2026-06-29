package models

import "time"

type FeeType struct {
	ID          int64   `db:"id"           json:"id"`
	SchoolID    int64   `db:"school_id"    json:"school_id"`
	Name        string  `db:"name"         json:"name"`
	Amount      float64 `db:"amount"       json:"amount"`
	Frequency   string  `db:"frequency"    json:"frequency"`
	IsMandatory bool    `db:"is_mandatory" json:"is_mandatory"`
}

type FeeInvoice struct {
	ID        int64     `db:"id"          json:"id"`
	StudentID int64     `db:"student_id"  json:"student_id"`
	FeeTypeID int64     `db:"fee_type_id" json:"fee_type_id"`
	TermID    int64     `db:"term_id"     json:"term_id"`
	Amount    float64   `db:"amount"      json:"amount"`
	Status    string    `db:"status"      json:"status"`
	DueDate   string    `db:"due_date"    json:"due_date"`
	CreatedAt time.Time `db:"created_at"  json:"created_at"`
}

type FeePayment struct {
	ID         int64     `db:"id"          json:"id"`
	InvoiceID  int64     `db:"invoice_id"  json:"invoice_id"`
	AmountPaid float64   `db:"amount_paid" json:"amount_paid"`
	Method     string    `db:"method"      json:"method"`
	RefNo      string    `db:"ref_no"      json:"ref_no"`
	PaidAt     time.Time `db:"paid_at"     json:"paid_at"`
}

type StudentFeeStatement struct {
	StudentName string       `json:"student_name"`
	AdmissionNo string       `json:"admission_no"`
	Invoices    []FeeInvoice `json:"invoices"`
	TotalBilled float64      `json:"total_billed"`
	TotalPaid   float64      `json:"total_paid"`
	Balance     float64      `json:"balance"`
}

type FeeDiscount struct {
	ID          int64    `db:"id"           json:"id"`
	SchoolID    int64    `db:"school_id"    json:"school_id"`
	StudentID   int64    `db:"student_id"   json:"student_id"`
	FeeTypeID   *int64   `db:"fee_type_id"  json:"fee_type_id"`
	TermID      *int64   `db:"term_id"      json:"term_id"`
	Label       string   `db:"label"        json:"label"`
	DiscountPct *float64 `db:"discount_pct" json:"discount_pct"`
	DiscountAmt *float64 `db:"discount_amt" json:"discount_amt"`
	ApprovedBy  int64    `db:"approved_by"  json:"approved_by"`
	IsActive    bool     `db:"is_active"    json:"is_active"`
}
