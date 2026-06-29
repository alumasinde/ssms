<?php
namespace Reports;

use Core\Controller;

class ReportsController extends Controller
{
    public function reportCard(array $params): void
    {
        $this->requirePermission('reports.view');
        $studentID = (int)$params['studentId'];
        $examID    = (int)($_GET['exam_id'] ?? 0);
        $res       = $this->api->get("/reports/report-card/{$studentID}", $examID ? ['exam_id' => $examID] : []);
        $exams     = $this->api->get('/exams');
        $student   = $this->api->get("/students/{$studentID}");
        $this->view('reports/report_card', [
            'title'   => 'Report Card',
            'card'    => $res['data'] ?? [],
            'exams'   => $exams['data'] ?? [],
            'student' => $student['data'] ?? [],
            'examID'  => $examID,
        ]);
    }

    public function classResults(array $params = []): void
    {
        $this->requirePermission('reports.view');
        $examID  = (int)($_GET['exam_id'] ?? 0);
        $classID = (int)($_GET['class_id'] ?? 0);
        $res     = $this->api->get('/reports/class-results', array_filter([
            'exam_id'  => $examID ?: null,
            'class_id' => $classID ?: null,
        ]));
        $exams   = $this->api->get('/exams');
        $classes = $this->api->get('/classes');
        $this->view('reports/class_results', [
            'title'   => 'Class Results',
            'results' => $res['data'] ?? [],
            'exams'   => $exams['data'] ?? [],
            'classes' => $classes['data'] ?? [],
            'examID'  => $examID,
            'classID' => $classID,
        ]);
    }

    public function feeCollection(array $params = []): void
    {
        $this->requirePermission('reports.view');
        $termID = (int)($_GET['term_id'] ?? 0);
        $res    = $this->api->get('/reports/fee-collection', $termID ? ['term_id' => $termID] : []);
        $terms  = $this->api->get('/terms');
        $this->view('reports/fee_collection', [
            'title'  => 'Fee Collection Report',
            'report' => $res['data'] ?? [],
            'terms'  => $terms['data'] ?? [],
            'termID' => $termID,
        ]);
    }

    public function attendanceSummary(array $params = []): void
    {
        $this->requirePermission('reports.view');
        $termID = (int)($_GET['term_id'] ?? 0);
        $res    = $this->api->get('/reports/attendance-summary', $termID ? ['term_id' => $termID] : []);
        $terms  = $this->api->get('/terms');
        $this->view('reports/attendance_summary', [
            'title'   => 'Attendance Summary',
            'reports' => $res['data'] ?? [],
            'terms'   => $terms['data'] ?? [],
            'termID'  => $termID,
        ]);
    }

    public function subjectPerformance(array $params = []): void
    {
        $this->requirePermission('reports.view');
        $examID = (int)($_GET['exam_id'] ?? 0);
        $res    = $this->api->get('/reports/subject-performance', $examID ? ['exam_id' => $examID] : []);
        $exams  = $this->api->get('/exams');
        $this->view('reports/subject_performance', [
            'title'  => 'Subject Performance',
            'rows'   => $res['data'] ?? [],
            'exams'  => $exams['data'] ?? [],
            'examID' => $examID,
        ]);
    }

    public function promote(array $params = []): void
    {
        $this->requirePermission('students.promote');
        $classes = $this->api->get('/classes');
        $this->view('reports/promote', [
            'title'   => 'Student Promotion',
            'classes' => $classes['data'] ?? [],
        ]);
    }

    public function storePromotion(array $params = []): void
    {
        $this->requirePermission('students.promote');
        $res = $this->api->post('/students/promote', [
            'from_class_id' => (int)($_POST['from_class_id'] ?? 0),
            'to_class_id'   => (int)($_POST['to_class_id'] ?? 0),
            'student_ids'   => [],
        ]);
        if ($res['success'] ?? false) {
            $count = $res['data']['promoted'] ?? 0;
            $this->redirect('/students', "{$count} student(s) promoted successfully.");
        }
        \Core\Session::flash('error', $res['error'] ?? 'Promotion failed.');
        $this->redirect('/reports/promote');
    }
}
