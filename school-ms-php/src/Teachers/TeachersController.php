<?php
namespace Teachers;

use Core\Controller;
use Core\Session;

class TeachersController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('teachers.view');
        $res = $this->api->get('/teachers');
        $this->view('teachers/index', [
            'title'    => 'Teachers',
            'teachers' => $res['data'] ?? [],
        ]);
    }

    public function show(array $params): void
    {
        $this->requirePermission('teachers.view');
        $id          = (int)$params['id'];
        $res         = $this->api->get("/teachers/{$id}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/teachers', 'Teacher not found.', 'error');
        }
        $subjects    = $this->api->get("/teachers/{$id}/subjects");
        $allSubjects = $this->api->get('/subjects');
        $allClasses  = $this->api->get('/classes');

        $this->view('teachers/show', [
            'title'       => $res['data']['name'] ?? 'Teacher',
            'teacher'     => $res['data'] ?? [],
            'subjects'    => $subjects['data'] ?? [],
            'allSubjects' => $allSubjects['data'] ?? [],
            'allClasses'  => $allClasses['data'] ?? [],
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('teachers.create');
        $users = $this->api->get('/users');
        $this->view('teachers/create', [
            'title' => 'Add Teacher',
            'users' => $users['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('teachers.create');
        $phone         = trim($_POST['phone'] ?? '');
        $gender        = trim($_POST['gender'] ?? '');
        $dob           = trim($_POST['dob'] ?? '');
        $qualification = trim($_POST['qualification'] ?? '');
        $tscNo         = trim($_POST['tsc_no'] ?? '');
        $hireDate      = trim($_POST['hire_date'] ?? '');

        $res = $this->api->post('/teachers', [
            'user_id'       => (int)($_POST['user_id'] ?? 0),
            'employee_no'   => trim($_POST['employee_no'] ?? ''),
            'phone'         => $phone !== '' ? $phone : null,
            'gender'        => $gender !== '' ? $gender : null,
            'dob'           => $dob !== '' ? $dob : null,
            'qualification' => $qualification !== '' ? $qualification : null,
            'tsc_no'        => $tscNo !== '' ? $tscNo : null,
            'hire_date'     => $hireDate !== '' ? $hireDate : null,
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/teachers', 'Teacher added successfully.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to add teacher.');
        $this->redirect('/teachers/create');
    }

    public function edit(array $params = []): void
    {
        $this->requirePermission('teachers.edit');
        $id  = (int)$params['id'];
        $res = $this->api->get("/teachers/{$id}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/teachers', 'Teacher not found.', 'error');
        }
        $this->view('teachers/edit', [
            'title'   => 'Edit Teacher',
            'teacher' => $res['data'] ?? [],
        ]);
    }

    public function update(array $params = []): void
    {
        $this->requirePermission('teachers.edit');
        $teacherID = (int)($params['id'] ?? 0);
        $phone         = trim($_POST['phone'] ?? '');
        $gender        = trim($_POST['gender'] ?? '');
        $dob           = trim($_POST['dob'] ?? '');
        $qualification = trim($_POST['qualification'] ?? '');
        $tscNo         = trim($_POST['tsc_no'] ?? '');
        $hireDate      = trim($_POST['hire_date'] ?? '');

        $res = $this->api->put("/teachers/{$teacherID}", [
            'id'            => $teacherID,
            'phone'         => $phone !== '' ? $phone : null,
            'gender'        => $gender !== '' ? $gender : null,
            'dob'           => $dob !== '' ? $dob : null,
            'qualification' => $qualification !== '' ? $qualification : null,
            'tsc_no'        => $tscNo !== '' ? $tscNo : null,
            'hire_date'     => $hireDate !== '' ? $hireDate : null,
        ]);

        if ($res['success'] ?? false) {
            $this->redirect("/teachers/{$teacherID}", 'Teacher updated successfully.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to update teacher.');
        $this->redirect("/teachers/{$teacherID}/edit");
    }

    public function assignSubject(array $params = []): void
    {
        $this->requirePermission('teachers.edit');
        $teacherID = (int)($_POST['teacher_id'] ?? 0);
        $redirect  = $_POST['redirect'] ?? "/teachers/{$teacherID}";

        $res = $this->api->post('/teachers/assign-subject', [
            'teacher_id' => $teacherID,
            'subject_id' => (int)($_POST['subject_id'] ?? 0),
            'class_id'   => (int)($_POST['class_id'] ?? 0),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect($redirect, 'Subject assigned.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to assign subject.');
        $this->redirect($redirect);
    }

    public function removeSubject(array $params = []): void
    {
        $this->requirePermission('teachers.edit');
        $teacherID = (int)($_POST['teacher_id'] ?? 0);
        $redirect  = $_POST['redirect'] ?? "/teachers/{$teacherID}";

        $res = $this->api->post('/teachers/remove-subject', [
            'teacher_id' => $teacherID,
            'subject_id' => (int)($_POST['subject_id'] ?? 0),
            'class_id'   => (int)($_POST['class_id'] ?? 0),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect($redirect, 'Subject removed.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to remove subject.');
        $this->redirect($redirect);
    }

    public function subjectAdmin(array $params = []): void
    {
        $this->requirePermission('teachers.edit');
        $teachers = $this->api->get('/teachers');
        $subjects = $this->api->get('/subjects');
        $classes  = $this->api->get('/classes');

        $matrix = [];
        foreach ($teachers['data'] ?? [] as $t) {
            $sRes = $this->api->get("/teachers/{$t['id']}/subjects");
            $matrix[$t['id']] = [
                'teacher'  => $t,
                'subjects' => $sRes['data'] ?? [],
            ];
        }

        $this->view('teachers/subject_admin', [
            'title'    => 'Subject Assignment',
            'matrix'   => $matrix,
            'subjects' => $subjects['data'] ?? [],
            'classes'  => $classes['data'] ?? [],
        ]);
    }
}
