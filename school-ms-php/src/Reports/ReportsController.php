<?php
namespace Reports;
use Core\Controller;
class ReportsController extends Controller {
    public function reportCard(array $params): void {
        $this->requireAuth();
        $studentID = $params['studentId'];
        $examID    = (int)($_GET['exam_id'] ?? 0);
        $res       = $this->api->get("/reports/report-card/{$studentID}", ['exam_id' => $examID]);
        $exams     = $this->api->get('/exams');
        $this->view('reports/report_card', ['title' => 'Report Card', 'card' => $res['data'] ?? [], 'exams' => $exams['data'] ?? []]);
    }
    public function classResults(array $params = []): void {
        $this->requireAuth();
        $examID = (int)($_GET['exam_id'] ?? 0);
        $res    = $this->api->get('/reports/class-results', ['exam_id' => $examID]);
        $exams  = $this->api->get('/exams');
        $this->view('reports/class_results', ['title' => 'Class Results', 'results' => $res['data'] ?? [], 'exams' => $exams['data'] ?? [], 'examID' => $examID]);
    }
    public function feeCollection(array $params = []): void {
        $this->requireAuth();
        $termID = (int)($_GET['term_id'] ?? 0);
        $res    = $this->api->get('/reports/fee-collection', ['term_id' => $termID]);
        $this->view('reports/fee_collection', ['title' => 'Fee Collection Report', 'report' => $res['data'] ?? []]);
    }
    public function attendanceSummary(array $params = []): void {
        $this->requireAuth();
        $termID = (int)($_GET['term_id'] ?? 0);
        $res    = $this->api->get('/reports/attendance-summary', ['term_id' => $termID]);
        $this->view('reports/attendance_summary', ['title' => 'Attendance Summary', 'reports' => $res['data'] ?? []]);
    }
}
