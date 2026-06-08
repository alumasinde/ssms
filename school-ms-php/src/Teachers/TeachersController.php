<?php
namespace Teachers;
use Core\Controller;
class TeachersController extends Controller {
    public function index(array $params = []): void {
        $this->requireAuth();
        $res = $this->api->get('/teachers');
        $this->view('teachers/index', ['title' => 'Teachers', 'teachers' => $res['data'] ?? []]);
    }
    public function show(array $params): void {
        $this->requireAuth();
        $res      = $this->api->get("/teachers/{$params['id']}");
        $subjects = $this->api->get("/teachers/{$params['id']}/subjects");
        $this->view('teachers/show', ['title' => $res['data']['name'] ?? 'Teacher', 'teacher' => $res['data'] ?? [], 'subjects' => $subjects['data'] ?? []]);
    }
    public function create(array $params = []): void {
        $this->requireAuth();
        $users = $this->api->get('/users');
        $this->view('teachers/create', ['title' => 'Add Teacher', 'users' => $users['data'] ?? []]);
    }
    public function store(array $params = []): void {
        $this->requireAuth();
        $res = $this->api->post('/teachers', [
            'user_id' => (int)$_POST['user_id'],
            'employee_no' => $_POST['employee_no'],
            'phone' => $_POST['phone'],
            'gender' => $_POST['gender'],
            'dob' => $_POST['dob'],
            'qualification' => $_POST['qualification']
        ]);
        if ($res['success'] ?? false) $this->redirect('/teachers', 'Teacher added.');
        $this->redirect('/teachers/create', $res['error'] ?? 'Failed.', 'error');
    }
}
