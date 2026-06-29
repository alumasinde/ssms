package dtos

type CreateFeeTypeDTO struct {
	SchoolID    int64   `json:"school_id"`
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"`
	Frequency   string  `json:"frequency"`
	IsMandatory bool    `json:"is_mandatory"`
}

type GenerateInvoicesDTO struct {
	SchoolID  int64   `json:"school_id"`
	FeeTypeID int64   `json:"fee_type_id"`
	TermID    int64   `json:"term_id"`
	ClassIDs  []int64 `json:"class_ids"`
	DueDate   string  `json:"due_date"`
}

type RecordPaymentDTO struct {
	InvoiceID  int64   `json:"invoice_id"`
	AmountPaid float64 `json:"amount_paid"`
	Method     string  `json:"method"`
	RefNo      string  `json:"ref_no"`
}

type CreateDiscountDTO struct {
	StudentID   int64    `json:"student_id"`
	FeeTypeID   *int64   `json:"fee_type_id"`
	TermID      *int64   `json:"term_id"`
	Label       string   `json:"label"`
	DiscountPct *float64 `json:"discount_pct"`
	DiscountAmt *float64 `json:"discount_amt"`
}

type MpesaStkPushDTO struct {
	InvoiceID int64  `json:"invoice_id"`
	Phone     string `json:"phone"`
}

// MpesaCallbackDTO mirrors the Safaricom Daraja STK push callback payload.
type MpesaCallbackDTO struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  *struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}
