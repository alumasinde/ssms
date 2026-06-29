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
    $linkedParents  = $this->api->get("/students/{$id}/parents");
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

    // 1. Upload photo to temp — don't send to API yet
    $tempPath  = '';
    $uploadDir = BASE_PATH . '/public/uploads/students/';

    if (!empty($_FILES['photo']['tmp_name'])) {
        $file    = $_FILES['photo'];
        $ext     = strtolower(pathinfo($file['name'], PATHINFO_EXTENSION));
        $allowed = ['jpg', 'jpeg', 'png', 'webp'];

        if (!in_array($ext, $allowed)) {
            Session::flash('error', 'Only JPG, PNG or WEBP images allowed.');
            $this->redirect('/students/create');
            return;
        }

        if (!is_dir($uploadDir)) {
            mkdir($uploadDir, 0775, true);
        }

        $tempPath = $uploadDir . 'tmp_' . uniqid() . '.' . $ext;
        if (!move_uploaded_file($file['tmp_name'], $tempPath)) {
            Session::flash('error', 'Photo upload failed.');
            $this->redirect('/students/create');
            return;
        }
    }

    // 2. Create student — no photo_url yet
    $res = $this->api->post('/students', [
        'admission_no' => trim($_POST['admission_no'] ?? ''),
        'first_name'   => trim($_POST['first_name'] ?? ''),
        'middle_name'  => trim($_POST['middle_name'] ?? ''),
        'last_name'    => trim($_POST['last_name'] ?? ''),
        'gender'       => $_POST['gender'] ?? '',
        'dob'          => $_POST['dob'] ?? '',
        'nationality'  => trim($_POST['nationality'] ?? ''),
        'national_id'  => trim($_POST['national_id'] ?? ''),
        'religion'     => trim($_POST['religion'] ?? ''),
        'blood_group'  => trim($_POST['blood_group'] ?? ''),
        'medical_notes' => trim($_POST['medical_notes'] ?? ''),
        'address'      => trim($_POST['address'] ?? ''),
        'class_id'     => (int)($_POST['class_id'] ?? 0),
    ]);

    if (!($res['success'] ?? false)) {
        // Clean up temp file if student creation failed
        if ($tempPath && file_exists($tempPath)) {
            unlink($tempPath);
        }
        Session::flash('error', $res['error'] ?? 'Could not enrol student.');
        $this->redirect('/students/create');
        return;
    }

    $studentID = $res['data']['id'] ?? null;

    // 3. Now we have the ID — rename and patch photo
    if ($studentID && $tempPath) {
        $ext       = pathinfo($tempPath, PATHINFO_EXTENSION);
        $finalName = 'student_' . $studentID . '.' . $ext;
        $finalPath = $uploadDir . $finalName;

        rename($tempPath, $finalPath);

        $this->api->put("/students/{$studentID}", [
            'first_name'  => trim($_POST['first_name'] ?? ''),
            'middle_name' => trim($_POST['middle_name'] ?? ''),
            'last_name'   => trim($_POST['last_name'] ?? ''),
            'gender'      => $_POST['gender'] ?? '',
            'dob'         => $_POST['dob'] ?? '',
            'nationality' => trim($_POST['nationality'] ?? ''),
            'national_id' => trim($_POST['national_id'] ?? ''),
            'religion'    => trim($_POST['religion'] ?? ''),
            'blood_group' => trim($_POST['blood_group'] ?? ''),
            'address'     => trim($_POST['address'] ?? ''),
            'medical_notes' => trim($_POST['medical_notes'] ?? ''),
            'class_id'    => (int)($_POST['class_id'] ?? 0),
            'photo_url'   => '/uploads/students/' . $finalName,
        ]);
    }

    // 4. Redirect
    if ($studentID) {
        $this->redirect("/students/{$studentID}", 'Student enrolled successfully.');
    }
    $this->redirect('/students', 'Student enrolled successfully.');
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
    $id = (int)$params['id'];

    // Handle photo upload
    $photoUrl = $_POST['existing_photo'] ?? '';  // keep existing if no new file
    if (!empty($_FILES['photo']['tmp_name'])) {
        $file     = $_FILES['photo'];
        $ext      = strtolower(pathinfo($file['name'], PATHINFO_EXTENSION));
        $allowed  = ['jpg', 'jpeg', 'png', 'webp'];

        if (!in_array($ext, $allowed)) {
            Session::flash('error', 'Only JPG, PNG or WEBP images allowed.');
            $this->redirect("/students/{$id}/edit");
            return;
        }

        $filename = 'student_' . $id . '_' . time() . '.' . $ext;
        $dest     = BASE_PATH . '/public/uploads/students/' . $filename;

        if (!is_dir(dirname($dest))) {
            mkdir(dirname($dest), 0775, true);
        }

        if (!move_uploaded_file($file['tmp_name'], $dest)) {
            Session::flash('error', 'Photo upload failed.');
            $this->redirect("/students/{$id}/edit");
            return;
        }

        $photoUrl = '/uploads/students/' . $filename;
    }

    $res = $this->api->put("/students/{$id}", [
        'first_name'    => trim($_POST['first_name'] ?? ''),
        'middle_name'   => trim($_POST['middle_name'] ?? ''),
        'last_name'     => trim($_POST['last_name'] ?? ''),
        'gender'        => $_POST['gender'] ?? '',
        'dob'           => $_POST['dob'] ?? '',
        'nationality'   => trim($_POST['nationality'] ?? ''),
        'national_id'   => trim($_POST['national_id'] ?? ''),
        'religion'      => trim($_POST['religion'] ?? ''),
        'blood_group'   => trim($_POST['blood_group'] ?? ''),
        'address'       => trim($_POST['address'] ?? ''),
        'medical_notes' => trim($_POST['medical_notes'] ?? ''),
        'class_id'      => (int)($_POST['class_id'] ?? 0),
        'photo_url'     => $photoUrl,
    ]);

    if ($res['success'] ?? false) {
        $this->redirect("/students/{$id}", 'Student updated.');
    }
    Session::flash('error', $res['error'] ?? 'Update failed.');
    $this->redirect("/students/{$id}/edit");
}
}
