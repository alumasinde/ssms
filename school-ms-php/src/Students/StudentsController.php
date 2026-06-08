<?php

namespace Students;

use Core\Controller;
use Core\Session;

class StudentsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('students.view');
        $page    = max(1, (int)($_GET['page'] ?? 1));
        $perPage = 50;
        $search  = trim($_GET['search'] ?? '');
        $res     = $this->api->get('/students', ['page' => $page, 'per_page' => $perPage]);

        $this->view('students/index', [
            'title'    => 'Students',
            'students' => $res['data'] ?? [],
            'meta'     => $res['meta'] ?? [],
            'page'     => $page,
            'search'   => $search,
        ]);
    }

    public function show(array $params): void
{
    $this->requirePermission('students.view');
    $id      = (int)$params['id'];
    $res     = $this->api->get("/students/{$id}");
    if (!($res['success'] ?? false)) {
        $this->redirect('/students', 'Student not found.', 'error');
    }
    $student        = $res['data'];
    $linkedParents  = $this->api->get("/parents/student/{$id}");
    $allParents     = $this->api->get('/parents');

    $this->view('students/show', [
        'title'         => $student['first_name'] . ' ' . $student['last_name'],
        'student'       => $student,
        'parents'       => $linkedParents['data'] ?? [],
        'allParents'    => $allParents['data'] ?? [],
    ]);
}

    public function create(array $params = []): void
    {
        $this->requirePermission('students.create');
        $classes = $this->api->get('/classes');

        $this->view('students/create', [
            'title'   => 'Enrol Student',
            'classes' => $classes['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
{
    $this->requirePermission('students.create');
    $res = $this->api->post('/students', [
        'admission_no' => trim($_POST['admission_no'] ?? ''),
        'first_name'   => trim($_POST['first_name'] ?? ''),
        'middle_name'  => trim($_POST['middle_name'] ?? ''),
        'last_name'    => trim($_POST['last_name'] ?? ''),
        'gender'       => $_POST['gender'] ?? '',
        'dob'          => $_POST['dob'] ?? '',
        'class_id'     => (int)($_POST['class_id'] ?? 0),
    ]);
    if ($res['success'] ?? false) {
        $studentID = $res['data']['id'] ?? null;
        if ($studentID) {
            $this->redirect("/students/{$studentID}", 'Student enrolled. You can now link parents below.');
        }
        $this->redirect('/students', 'Student enrolled successfully.');
    }
    Session::flash('error', $res['error'] ?? 'Could not enrol student.');
    $this->redirect('/students/create');
}

    public function edit(array $params): void
    {
        $this->requirePermission('students.edit');
        $id      = (int)$params['id'];
        
        $res     = $this->api->get("/students/{$id}");
        $classes = $this->api->get('/classes');

        $this->view('students/edit', [
            'title'   => 'Edit Student',
            'student' => $res['data'] ?? [],
            'classes' => $classes['data'] ?? [],
        ]);
    }

    public function update(array $params): void
    {
        $this->requirePermission('students.edit');
        $id  = (int)$params['id'];
        $res = $this->api->put("/students/{$id}", [
            'first_name'   => trim($_POST['first_name'] ?? ''),
            'middle_name'  => trim($_POST['middle_name'] ?? ''),
            'last_name'    => trim($_POST['last_name'] ?? ''),
            'gender'   => $_POST['gender'] ?? '',
            'dob'      => $_POST['dob'] ?? '',
            'class_id' => (int)($_POST['class_id'] ?? 0),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect("/students/{$id}", 'Student updated.');
        }
        Session::flash('error', $res['error'] ?? 'Update failed.');
        $this->redirect("/students/{$id}/edit");
    }
}
