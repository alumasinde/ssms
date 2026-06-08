<?php
namespace Classes;
use Core\Controller;
use Core\Session;

class ClassesController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->get('/classes');
        $this->view('classes/index', ['title' => 'Classes', 'classes' => $res['data'] ?? []]);
    }

    public function create(array $params = []): void
    {
        $this->requireAuth();
        $this->view('classes/create', ['title' => 'Create Class']);
    }

    public function store(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->post('/classes', [
            'name'   => trim($_POST['name'] ?? ''),
            'level'  => trim($_POST['level'] ?? ''),
            'stream' => trim($_POST['stream'] ?? ''),
        ]);
        if ($res['success'] ?? false) $this->redirect('/classes', 'Class created.');
        Session::flash('error', $res['error'] ?? 'Failed to create class.');
        $this->redirect('/classes/create');
    }

    public function edit(array $params): void
    {
        $this->requireAuth();
        $res = $this->api->get("/classes/{$params['id']}");
        $this->view('classes/edit', ['title' => 'Edit Class', 'class' => $res['data'] ?? []]);
    }

    public function update(array $params): void
    {
        $this->requireAuth();
        $id  = (int)$params['id'];
        $res = $this->api->put("/classes/{$id}", [
            'name'   => trim($_POST['name'] ?? ''),
            'level'  => trim($_POST['level'] ?? ''),
            'stream' => trim($_POST['stream'] ?? ''),
        ]);
        if ($res['success'] ?? false) $this->redirect('/classes', 'Class updated.');
        Session::flash('error', $res['error'] ?? 'Update failed.');
        $this->redirect("/classes/{$id}/edit");
    }

    public function delete(array $params): void
    {
        $this->requireAuth();
        $this->api->delete("/classes/{$params['id']}");
        $this->redirect('/classes', 'Class deleted.');
    }
}