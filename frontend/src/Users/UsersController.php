<?php
namespace Users;
use Core\Controller;
use Core\Session;

class UsersController extends Controller
{
    private array $roles = ['admin','teacher','parent','student'];

    public function index(array $params = []): void
    {
        $this->requireRole('admin');
        $res = $this->api->get('/users/school');
        $this->view('users/index', [
            'title' => 'Users',
            'users' => $res['data'] ?? [],
            'roles' => $this->roles,
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requireRole('admin');
        $this->view('users/create', ['title'=>'Create User','roles'=>$this->roles]);
    }

    public function store(array $params = []): void
    {
        $this->requireRole('admin');
        $res = $this->api->post('/users', [
            'name'     => trim($_POST['name']     ?? ''),
            'email'    => trim($_POST['email']    ?? ''),
            'password' => trim($_POST['password'] ?? ''),
            'role'     => $_POST['role']           ?? 'teacher',
        ]);
        if ($res['success'] ?? false) { $this->redirect('/users', 'User created.'); }
        Session::flash('error', $res['error'] ?? 'Failed to create user.');
        $this->redirect('/users/create');
    }

    public function updateRole(array $params): void
    {
        $this->requireRole('admin');
        $id  = (int)$params['id'];
        $res = $this->api->put("/users/{$id}/role", ['role' => $_POST['role'] ?? 'teacher']);
        if ($res['success'] ?? false) { $this->redirect('/users', 'Role updated.'); }
        Session::flash('error', $res['error'] ?? 'Failed to update role.');
        $this->redirect('/users');
    }

    public function deactivate(array $params): void
    {
        $this->requireRole('admin');
        $this->api->delete("/users/{$params['id']}");
        $this->redirect('/users', 'User deactivated.');
    }

    public function activate(array $params): void
    {
        $this->requireRole('admin');
        $this->api->post("/users/{$params['id']}/activate", []);
        $this->redirect('/users', 'User activated.');
    }
}
