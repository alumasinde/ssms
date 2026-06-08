<?php
namespace Finance;
use Core\Controller;
class FinanceController extends Controller {
    public function index(array $params = []): void {
        $this->requireAuth();
        $feeTypes = $this->api->get('/finance/fee-types');
        $this->view('finance/index', ['title' => 'Finance', 'feeTypes' => $feeTypes['data'] ?? []]);
    }
    public function statement(array $params): void {
        $this->requireAuth();
        $studentID = $params['studentId'];
        $student   = $this->api->get("/students/{$studentID}");
        $stmt      = $this->api->get("/finance/statement/student/{$studentID}");
        $this->view('finance/statement', ['title' => 'Fee Statement', 'student' => $student['data'] ?? [], 'statement' => $stmt['data'] ?? []]);
    }
    public function recordPayment(array $params = []): void {
        $this->requireAuth();
        $res = $this->api->post('/finance/payments', [
            'invoice_id'  => (int)$_POST['invoice_id'],
            'amount_paid' => (float)$_POST['amount_paid'],
            'method'      => $_POST['method'],
            'ref_no'      => $_POST['ref_no'] ?? '',
        ]);
        $studentID = $_POST['student_id'] ?? '';
        if ($res['success'] ?? false) $this->redirect("/finance/statement/{$studentID}", 'Payment recorded.');
        $this->redirect("/finance/statement/{$studentID}", $res['error'] ?? 'Failed.', 'error');
    }
}
