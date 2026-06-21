package services

import (
	"fmt"

	dtos   "school-ms/internal/Modules/Finance/DTOs"
	models "school-ms/internal/Modules/Finance/Models"
	repos  "school-ms/internal/Modules/Finance/Repositories"
)

type FinanceService struct{ repo *repos.FinanceRepository }

func NewFinanceService(repo *repos.FinanceRepository) *FinanceService { return &FinanceService{repo: repo} }

func (s *FinanceService) CreateFeeType(dto dtos.CreateFeeTypeDTO) (*models.FeeType, error) {
	ft := &models.FeeType{SchoolID: dto.SchoolID, Name: dto.Name, Amount: dto.Amount, Frequency: dto.Frequency, IsMandatory: dto.IsMandatory}
	return ft, s.repo.CreateFeeType(ft)
}

func (s *FinanceService) ListFeeTypes(schoolID int64) ([]models.FeeType, error) {
	return s.repo.ListFeeTypes(schoolID)
}

func (s *FinanceService) GenerateInvoices(dto dtos.GenerateInvoicesDTO) (int, error) {
	ft, err := s.repo.ListFeeTypes(dto.SchoolID)
	if err != nil { return 0, err }
	var amount float64
	for _, f := range ft {
		if f.ID == dto.FeeTypeID { amount = f.Amount; break }
	}
	count := 0
	for _, classID := range dto.ClassIDs {
		studentIDs, err := s.repo.ListStudentsByClass(classID)
		if err != nil { return count, err }
		for _, sid := range studentIDs {
			inv := &models.FeeInvoice{StudentID: sid, FeeTypeID: dto.FeeTypeID, TermID: dto.TermID, Amount: amount, DueDate: dto.DueDate}
			if err := s.repo.CreateInvoice(inv); err == nil { count++ }
		}
	}
	return count, nil
}

func (s *FinanceService) GetStudentStatement(studentID int64) (*models.StudentFeeStatement, error) {
	invoices, err := s.repo.ListStudentInvoices(studentID)
	if err != nil { return nil, err }
	var totalBilled, totalPaid float64
	for _, inv := range invoices {
		totalBilled += inv.Amount
		totalPaid += s.repo.TotalPaidForInvoice(inv.ID)
	}
	return &models.StudentFeeStatement{Invoices: invoices, TotalBilled: totalBilled, TotalPaid: totalPaid, Balance: totalBilled - totalPaid}, nil
}

func (s *FinanceService) RecordPayment(dto dtos.RecordPaymentDTO) error {
	inv, err := s.repo.GetInvoiceByID(dto.InvoiceID)
	if err != nil { return err }
	p := &models.FeePayment{InvoiceID: dto.InvoiceID, AmountPaid: dto.AmountPaid, Method: dto.Method, RefNo: dto.RefNo}
	if err := s.repo.RecordPayment(p); err != nil { return err }
	totalPaid := s.repo.TotalPaidForInvoice(dto.InvoiceID)
	status := "partial"
	if totalPaid >= inv.Amount { status = "paid" }
	return s.repo.UpdateInvoiceStatus(dto.InvoiceID, status)
}

func (s *FinanceService) GetInvoicePayments(invoiceID int64) ([]models.FeePayment, error) {
	return s.repo.GetPaymentsByInvoice(invoiceID)
}

func (s *FinanceService) CreateDiscount(dto dtos.CreateDiscountDTO, schoolID, approvedBy int64) (*models.FeeDiscount, error) {
	d := &models.FeeDiscount{
		SchoolID: schoolID, StudentID: dto.StudentID,
		FeeTypeID: dto.FeeTypeID, TermID: dto.TermID,
		Label: dto.Label, DiscountPct: dto.DiscountPct, DiscountAmt: dto.DiscountAmt,
		ApprovedBy: approvedBy,
	}
	return d, s.repo.CreateDiscount(d)
}

func (s *FinanceService) ListDiscounts(studentID int64) ([]models.FeeDiscount, error) {
	return s.repo.ListDiscountsByStudent(studentID)
}

func (s *FinanceService) InitiateStkPush(dto dtos.MpesaStkPushDTO, schoolID int64) (map[string]interface{}, error) {
	_, balance, err := s.repo.GetInvoiceForMpesa(dto.InvoiceID)
	if err != nil { return nil, err }
	if balance <= 0 { return nil, fmt.Errorf("invoice already fully paid") }
	return map[string]interface{}{
		"invoice_id": dto.InvoiceID,
		"phone":      dto.Phone,
		"amount":     balance,
		"status":     "pending",
		"message":    "Configure MPESA_* env vars and implement Daraja HTTP call to go live",
	}, nil
}

func (s *FinanceService) HandleMpesaCallback(cb dtos.MpesaCallbackDTO) {
	stk := cb.Body.StkCallback
	if stk.ResultCode != 0 { return }
	var amount float64
	var mpesaCode string
	for _, item := range stk.CallbackMetadata.Item {
		switch item.Name {
		case "Amount":
			if v, ok := item.Value.(float64); ok { amount = v }
		case "MpesaReceiptNumber":
			if v, ok := item.Value.(string); ok { mpesaCode = v }
		}
	}
	// TODO: look up pending_mpesa_pushes by CheckoutRequestID, record payment
	_ = amount
	_ = mpesaCode
}
