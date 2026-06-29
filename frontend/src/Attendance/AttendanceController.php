<?php
namespace Attendance;

use Core\Controller;
use Core\Session;

class AttendanceController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('attendance.view');
        $classes = $this->api->get('/classes');
        $terms   = $this->api->get('/terms');
        $this->view('attendance/index', [
            'title'   => 'Attendance',
            'classes' => $classes['data'] ?? [],
            'terms'   => $terms['data'] ?? [],
        ]);
    }

    public function mark(array $params = []): void
    {
        $this->requirePermission('attendance.mark');
        $classID = (int)($_GET['class_id'] ?? 0);
        $date    = $_GET['date'] ?? date('Y-m-d');

        if (!$classID) {
            Session::flash('error', 'Please select a class first.');
            $this->redirect('/attendance');
        }

        // Use new /terms/current endpoint
        $tRes        = $this->api->get('/terms/current');
        $currentTerm = ($tRes['success'] ?? false) ? ($tRes['data'] ?? null) : null;

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
        $this->requirePermission('attendance.mark');
        $classID = (int)($_POST['class_id'] ?? 0);
        $termID  = (int)($_POST['term_id'] ?? 0);
        $date    = $_POST['date'] ?? date('Y-m-d');

        if (!$termID) {
            Session::flash('error', 'No current term is set. Please configure the current term first.');
            $this->redirect('/attendance');
        }

        $records = [];
        foreach ($_POST['status'] ?? [] as $studentID => $status) {
            $records[] = [
                'student_id' => (int)$studentID,
                'status'     => $status,
                'remark'     => trim($_POST['remark'][$studentID] ?? ''),
            ];
        }

        if (empty($records)) {
            Session::flash('error', 'No students to mark.');
            $this->redirect('/attendance');
        }

        $res = $this->api->post('/attendance', [
            'class_id' => $classID,
            'term_id'  => $termID,
            'date'     => $date,
            'records'  => $records,
        ]);

        if ($res['success'] ?? false) {
            $this->redirect('/attendance', 'Attendance saved successfully.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to save attendance.');
        $this->redirect("/attendance/mark?class_id={$classID}&date={$date}");
    }

    public function summary(array $params = []): void
    {
        $this->requirePermission('attendance.view');
        $classID = (int)($_GET['class_id'] ?? 0);
        $termID  = (int)($_GET['term_id'] ?? 0);
        $classes = $this->api->get('/classes');
        $terms   = $this->api->get('/terms');
        $summary = [];

        if ($classID && $termID) {
            $res     = $this->api->get("/attendance/summary/class/{$classID}", ['term_id' => $termID]);
            $summary = $res['data'] ?? [];
        }

        $this->view('attendance/summary', [
            'title'   => 'Attendance Summary',
            'summary' => $summary,
            'classes' => $classes['data'] ?? [],
            'terms'   => $terms['data'] ?? [],
            'classID' => $classID,
            'termID'  => $termID,
        ]);
    }
}
