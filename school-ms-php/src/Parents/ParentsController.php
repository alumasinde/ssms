<?php
namespace Parents;
use Core\Controller;
use Core\Session;

class ParentsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->get('/parents');
        $this->view('parents/index', ['title' => 'Parents', 'parents' => $res['data'] ?? []]);
    }

    public function show(array $params): void
    {
        $this->requireAuth();
        $id       = (int)$params['id'];
        $res      = $this->api->get("/parents/{$id}");
        $this->view('parents/show', [
            'title'  => $res['data']['name'] ?? 'Parent',
            'parent' => $res['data'] ?? [],
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requireAuth();
        $users = $this->api->get('/users');
        $this->view('parents/create', ['title' => 'Add Parent', 'users' => $users['data'] ?? []]);
    }

    public function store(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->post('/parents', [
            'user_id'    => (int)($_POST['user_id'] ?? 0),
            'phone'      => trim($_POST['phone'] ?? ''),
            'occupation' => trim($_POST['occupation'] ?? ''),
            'address'    => trim($_POST['address'] ?? ''),
        ]);
        if ($res['success'] ?? false) $this->redirect('/parents', 'Parent added.');
        Session::flash('error', $res['error'] ?? 'Failed.');
        $this->redirect('/parents/create');
    }

    public function linkStudent(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->post('/parents/link-student', [
            'parent_id'    => (int)($_POST['parent_id'] ?? 0),
            'student_id'   => (int)($_POST['student_id'] ?? 0),
            'relationship' => trim($_POST['relationship'] ?? 'parent'),
        ]);
        $parentID = $_POST['parent_id'] ?? '';
        if ($res['success'] ?? false) $this->redirect("/parents/{$parentID}", 'Student linked.');
        Session::flash('error', $res['error'] ?? 'Failed.');
        $this->redirect("/parents/{$parentID}");
    }
}