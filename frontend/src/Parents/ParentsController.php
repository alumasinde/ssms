<?php
namespace Parents;
use Core\Controller;
use Core\Session;

class ParentsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('parents.view');
        $res = $this->api->get('/parents');
        $this->view('parents/index', [
            'title'   => 'Parents',
            'parents' => $res['data'] ?? [],
        ]);
    }

    public function show(array $params): void
    {
        $this->requirePermission('parents.view');
        $id  = (int)$params['id'];
        $res = $this->api->get("/parents/{$id}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/parents', 'Parent not found.', 'error');
        }
        // Load linked students via the parent's students endpoint
        $studentsRes = $this->api->get("/parents/{$id}/students");
        $this->view('parents/show', [
            'title'    => $res['data']['name'] ?? 'Parent',
            'parent'   => $res['data'] ?? [],
            'students' => $studentsRes['data'] ?? [],
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('parents.create');
        $users = $this->api->get('/users');
        $this->view('parents/create', [
            'title' => 'Add Parent',
            'users' => $users['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('parents.create');
        $res = $this->api->post('/parents', [
            'user_id'    => (int)($_POST['user_id'] ?? 0),
            'phone'      => trim($_POST['phone'] ?? ''),
            'occupation' => trim($_POST['occupation'] ?? ''),
            'address'    => trim($_POST['address'] ?? ''),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/parents', 'Parent added.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to add parent.');
        $this->redirect('/parents/create');
    }

    public function linkStudent(array $params = []): void
    {
        $this->requirePermission('parents.create');
        $parentID = (int)($_POST['parent_id'] ?? 0);
        $res = $this->api->post('/parents/link-student', [
            'parent_id'    => $parentID,
            'student_id'   => (int)($_POST['student_id'] ?? 0),
            'relationship' => trim($_POST['relationship'] ?? 'parent'),
        ]);
        // Redirect back to wherever the link was triggered from
        $redirect = $_POST['redirect'] ?? "/parents/{$parentID}";
        if ($res['success'] ?? false) {
            $this->redirect($redirect, 'Student linked.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to link student.');
        $this->redirect($redirect);
    }
}
