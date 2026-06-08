package repositories

import (
	models "school-ms/internal/Modules/Finance/Models"
	"github.com/jmoiron/sqlx"
)

type FinanceRepository struct{ db *sqlx.DB }

func NewFinanceRepository(db *sqlx.DB) *FinanceRepository { return &FinanceRepository{db: db} }

func (r *FinanceRepository) CreateFeeType(ft *models.FeeType) error {
	res, err := r.db.Exec(`INSERT INTO fee_types (school_id,name,amount,frequency,is_mandatory) VALUES (?,?,?,?,?)`,
		ft.SchoolID, ft.Name, ft.Amount, ft.Frequency, ft.IsMandatory)
	if err != nil { return err }
	id, _ := res.LastInsertId(); ft.ID = id; return nil
}

func (r *FinanceRepository) ListFeeTypes(schoolID int64) ([]models.FeeType, error) {
	var list []models.FeeType
	return list, r.db.Select(&list, `SELECT * FROM fee_types WHERE school_id=?`, schoolID)
}

func (r *FinanceRepository) CreateInvoice(inv *models.FeeInvoice) error {
	res, err := r.db.Exec(`INSERT INTO fee_invoices (student_id,fee_type_id,term_id,amount,status,due_date) VALUES (?,?,?,?,?,?)`,
		inv.StudentID, inv.FeeTypeID, inv.TermID, inv.Amount, "unpaid", inv.DueDate)
	if err != nil { return err }
	id, _ := res.LastInsertId(); inv.ID = id; return nil
}

func (r *FinanceRepository) ListStudentInvoices(studentID int64) ([]models.FeeInvoice, error) {
	var list []models.FeeInvoice
	return list, r.db.Select(&list, `SELECT * FROM fee_invoices WHERE student_id=? ORDER BY due_date DESC`, studentID)
}

func (r *FinanceRepository) GetInvoiceByID(id int64) (*models.FeeInvoice, error) {
	var inv models.FeeInvoice; return &inv, r.db.Get(&inv, `SELECT * FROM fee_invoices WHERE id=?`, id)
}

func (r *FinanceRepository) UpdateInvoiceStatus(id int64, status string) error {
	_, err := r.db.Exec(`UPDATE fee_invoices SET status=? WHERE id=?`, status, id); return err
}

func (r *FinanceRepository) RecordPayment(p *models.FeePayment) error {
	res, err := r.db.Exec(`INSERT INTO fee_payments (invoice_id,amount_paid,method,ref_no,paid_at) VALUES (?,?,?,?,NOW())`,
		p.InvoiceID, p.AmountPaid, p.Method, p.RefNo)
	if err != nil { return err }
	id, _ := res.LastInsertId(); p.ID = id; return nil
}

func (r *FinanceRepository) GetPaymentsByInvoice(invoiceID int64) ([]models.FeePayment, error) {
	var list []models.FeePayment
	return list, r.db.Select(&list, `SELECT * FROM fee_payments WHERE invoice_id=? ORDER BY paid_at`, invoiceID)
}

func (r *FinanceRepository) TotalPaidForInvoice(invoiceID int64) float64 {
	var total float64
	r.db.Get(&total, `SELECT COALESCE(SUM(amount_paid),0) FROM fee_payments WHERE invoice_id=?`, invoiceID)
	return total
}

func (r *FinanceRepository) ListStudentsByClass(classID int64) ([]int64, error) {
	var ids []int64
	return ids, r.db.Select(&ids, `SELECT id FROM students WHERE class_id=? AND is_active=1`, classID)
}
