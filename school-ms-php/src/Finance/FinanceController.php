<?php
namespace Finance;

use Core\Controller;
use Core\Session;

class FinanceController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('finance.view');
        $feeTypes = $this->api->get('/finance/fee-types');
        $terms    = $this->api->get('/terms');
        $this->view('finance/index', [
            'title'    => 'Finance',
            'feeTypes' => $feeTypes['data'] ?? [],
            'terms'    => $terms['data'] ?? [],
        ]);
    }

    public function createFeeType(array $params = []): void
    {
        $this->requirePermission('finance.create');
        $this->view('finance/create_fee_type', ['title' => 'Add Fee Type']);
    }

    public function storeFeeType(array $params = []): void
    {
        $this->requirePermission('finance.create');
        $res = $this->api->post('/finance/fee-types', [
            'name'         => trim($_POST['name'] ?? ''),
            'amount'       => (float)($_POST['amount'] ?? 0),
            'frequency'    => $_POST['frequency'] ?? 'termly',
            'is_mandatory' => isset($_POST['is_mandatory']) ? true : false,
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/finance', 'Fee type created.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to create fee type.');
        $this->redirect('/finance/fee-types/create');
    }

    public function statement(array $params): void
    {
        $this->requirePermission('finance.view');
        $studentID = (int)$params['studentId'];
        $student   = $this->api->get("/students/{$studentID}");
        $stmt      = $this->api->get("/finance/statement/student/{$studentID}");
        $discounts = $this->api->get("/finance/discounts/student/{$studentID}");
        $this->view('finance/statement', [
            'title'     => 'Fee Statement',
            'student'   => $student['data'] ?? [],
            'statement' => $stmt['data'] ?? [],
            'discounts' => $discounts['data'] ?? [],
        ]);
    }

    public function recordPayment(array $params = []): void
    {
        $this->requirePermission('finance.create');
        $studentID = (int)($_POST['student_id'] ?? 0);
        $res = $this->api->post('/finance/payments', [
            'invoice_id'  => (int)($_POST['invoice_id'] ?? 0),
            'amount_paid' => (float)($_POST['amount_paid'] ?? 0),
            'method'      => $_POST['method'] ?? 'cash',
            'ref_no'      => trim($_POST['ref_no'] ?? ''),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect("/finance/statement/{$studentID}", 'Payment recorded successfully.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to record payment.');
        $this->redirect("/finance/statement/{$studentID}");
    }

    public function generateInvoices(array $params = []): void
    {
        $this->requirePermission('finance.create');
        $feeTypes = $this->api->get('/finance/fee-types');
        $terms    = $this->api->get('/terms');
        $classes  = $this->api->get('/classes');
        $this->view('finance/generate_invoices', [
            'title'    => 'Generate Invoices',
            'feeTypes' => $feeTypes['data'] ?? [],
            'terms'    => $terms['data'] ?? [],
            'classes'  => $classes['data'] ?? [],
        ]);
    }

    public function storeInvoices(array $params = []): void
    {
        $this->requirePermission('finance.create');
        // API expects class_ids as array
        $classIDs = [];
        if (!empty($_POST['class_id'])) {
            $classIDs = [(int)$_POST['class_id']];
        }
        $res = $this->api->post('/finance/invoices/generate', [
            'term_id'     => (int)($_POST['term_id'] ?? 0),
            'fee_type_id' => (int)($_POST['fee_type_id'] ?? 0),
            'class_ids'   => $classIDs,
            'due_date'    => $_POST['due_date'] ?? '',
        ]);
        if ($res['success'] ?? false) {
            $count = $res['data']['invoices_created'] ?? 0;
            $this->redirect('/finance', "{$count} invoice(s) generated successfully.");
        }
        Session::flash('error', $res['error'] ?? 'Failed to generate invoices.');
        $this->redirect('/finance/invoices/generate');
    }

    public function createDiscount(array $params = []): void
    {
        $this->requirePermission('finance.discount');
        $studentID = (int)($_GET['student_id'] ?? 0);
        $feeTypes  = $this->api->get('/finance/fee-types');
        $terms     = $this->api->get('/terms');
        $this->view('finance/create_discount', [
            'title'     => 'Add Fee Discount',
            'studentID' => $studentID,
            'feeTypes'  => $feeTypes['data'] ?? [],
            'terms'     => $terms['data'] ?? [],
        ]);
    }

    public function storeDiscount(array $params = []): void
    {
        $this->requirePermission('finance.discount');
        $studentID = (int)($_POST['student_id'] ?? 0);
        $feeTypeID = (int)($_POST['fee_type_id'] ?? 0) ?: null;
        $termID    = (int)($_POST['term_id'] ?? 0) ?: null;
        $discountPct = trim($_POST['discount_pct'] ?? '');
        $discountAmt = trim($_POST['discount_amt'] ?? '');

        $res = $this->api->post('/finance/discounts', [
            'student_id'   => $studentID,
            'fee_type_id'  => $feeTypeID,
            'term_id'      => $termID,
            'label'        => trim($_POST['label'] ?? ''),
            'discount_pct' => $discountPct !== '' ? (float)$discountPct : null,
            'discount_amt' => $discountAmt !== '' ? (float)$discountAmt : null,
        ]);
        if ($res['success'] ?? false) {
            $this->redirect("/finance/statement/{$studentID}", 'Discount applied.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to apply discount.');
        $this->redirect("/finance/discounts/create?student_id={$studentID}");
    }
}
