<?php

namespace Attendance;

use Core\Controller;

class AttendanceController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requireAuth();
        $classes = $this->api->get('/classes');

        $this->view('attendance/index', [
            'title'   => 'Attendance',
            'classes' => $classes['data'] ?? [],
        ]);
    }

    public function mark(array $params = []): void
    {
        $this->requireAuth();
        $classID = (int)($_GET['class_id'] ?? 0);
        $date    = $_GET['date'] ?? date('Y-m-d');

        // Fetch current term from API
        $termRes = $this->api->get('/academic-years/current');
        $currentTerm = null;
        if ($termRes['success'] ?? false) {
            // Get current term within that year
            $termListRes = $this->api->get('/terms/current');
            if ($termListRes['success'] ?? false) {
                $currentTerm = $termListRes['data'];
            }
        }
        // Fallback: fetch from session school or terms endpoint directly
        if (!$currentTerm) {
            $termListRes = $this->api->get('/terms/current');
            $currentTerm = ($termListRes['success'] ?? false) ? $termListRes['data'] : null;
        }

        $students = $this->api->get("/students/class/{$classID}");
        $existing = $this->api->get("/attendance/class/{$classID}", ['date' => $date]);

        $existingMap = [];
        foreach ($existing['data'] ?? [] as $a) {
            $existingMap[$a['student_id']] = $a;
        }

        $this->view('attendance/mark', [
            'title'       => 'Mark Attendance',
            'classID'     => $classID,
            'date'        => $date,
            'students'    => $students['data'] ?? [],
            'existingMap' => $existingMap,
            'currentTerm' => $currentTerm,
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requireAuth();
        $classID = (int)($_POST['class_id'] ?? 0);
        $termID  = (int)($_POST['term_id'] ?? 0);
        $date    = $_POST['date'] ?? date('Y-m-d');

        $records = [];
        foreach ($_POST['status'] ?? [] as $studentID => $status) {
            $records[] = [
                'student_id' => (int)$studentID,
                'status'     => $status,
                'remark'     => $_POST['remark'][$studentID] ?? '',
            ];
        }

        $res = $this->api->post('/attendance', [
            'class_id' => $classID,
            'term_id'  => $termID,
            'date'     => $date,
            'records'  => $records,
        ]);

        if ($res['success'] ?? false) {
            $this->redirect('/attendance', 'Attendance saved.');
        }
        $this->redirect('/attendance', $res['error'] ?? 'Failed to save attendance.', 'error');
    }

    public function summary(array $params = []): void
    {
        $this->requireAuth();
        $classID = (int)($_GET['class_id'] ?? 0);
        $termID  = (int)($_GET['term_id'] ?? 0);
        $classes = $this->api->get('/classes');
        $summary = [];

        if ($classID && $termID) {
            $res     = $this->api->get("/attendance/summary/class/{$classID}", ['term_id' => $termID]);
            $summary = $res['data'] ?? [];
        }

        $this->view('attendance/summary', [
            'title'   => 'Attendance Summary',
            'summary' => $summary,
            'classes' => $classes['data'] ?? [],
            'classID' => $classID,
            'termID'  => $termID,
        ]);
    }
}
