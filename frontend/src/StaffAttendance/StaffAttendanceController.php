<?php
namespace StaffAttendance;
use Core\Controller;
use Core\Session;

class StaffAttendanceController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('staff_attendance.view');
        $date = $_GET['date'] ?? date('Y-m-d');
        $res  = $this->api->get('/staff-attendance', ['date' => $date]);
        $this->view('staff_attendance/index', [
            'title'    => 'Staff Attendance',
            'records'  => $res['data'] ?? [],
            'date'     => $date,
        ]);
    }

    public function mark(array $params = []): void
    {
        $this->requirePermission('staff_attendance.mark');
        $date     = $_GET['date'] ?? date('Y-m-d');
        $teachers = $this->api->get('/teachers');
        $existing = $this->api->get('/staff-attendance', ['date' => $date]);
        $existingMap = [];
        foreach ($existing['data'] ?? [] as $r) { $existingMap[$r['teacher_id']] = $r; }
        $this->view('staff_attendance/mark', [
            'title'       => 'Mark Staff Attendance',
            'teachers'    => $teachers['data'] ?? [],
            'date'        => $date,
            'existingMap' => $existingMap,
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('staff_attendance.mark');
        $date    = $_POST['date'] ?? date('Y-m-d');
        $records = [];
        foreach ($_POST['status'] ?? [] as $teacherID => $status) {
            $records[] = [
                'teacher_id' => (int)$teacherID,
                'status'     => $status,
                'check_in'   => $_POST['check_in'][$teacherID]  ?? '',
                'check_out'  => $_POST['check_out'][$teacherID] ?? '',
                'remark'     => trim($_POST['remark'][$teacherID] ?? ''),
            ];
        }
        $res = $this->api->post('/staff-attendance', ['date' => $date, 'records' => $records]);
        if ($res['success'] ?? false) { $this->redirect('/staff-attendance', 'Staff attendance saved.'); }
        Session::flash('error', $res['error'] ?? 'Failed.');
        $this->redirect('/staff-attendance/mark?date='.$date);
    }

    public function summary(array $params = []): void
    {
        $this->requirePermission('staff_attendance.view');
        $month = $_GET['month'] ?? date('Y-m');
        $res   = $this->api->get('/staff-attendance/summary', ['month' => $month]);
        $this->view('staff_attendance/summary', [
            'title'   => 'Staff Attendance Summary',
            'records' => $res['data'] ?? [],
            'month'   => $month,
        ]);
    }
}
