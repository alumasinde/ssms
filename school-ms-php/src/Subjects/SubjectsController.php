<?php
namespace Subjects;
use Core\Controller;
use Core\Session;

class SubjectsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->get('/subjects');
        $this->view('subjects/index', ['title' => 'Subjects', 'subjects' => $res['data'] ?? []]);
    }

    public function create(array $params = []): void
    {
        $this->requireAuth();
        $this->view('subjects/create', ['title' => 'Add Subject']);
    }

    public function store(array $params = []): void
    {
        $this->requireAuth();
        $res = $this->api->post('/subjects', [
            'name' => trim($_POST['name'] ?? ''),
            'code' => trim($_POST['code'] ?? ''),
        ]);
        if ($res['success'] ?? false) $this->redirect('/subjects', 'Subject added.');
        Session::flash('error', $res['error'] ?? 'Failed.');
        $this->redirect('/subjects/create');
    }

    public function edit(array $params): void
    {
        $this->requireAuth();
        $res = $this->api->get("/subjects/{$params['id']}");
        $this->view('subjects/edit', ['title' => 'Edit Subject', 'subject' => $res['data'] ?? []]);
    }

    public function update(array $params): void
    {
        $this->requireAuth();
        $id  = (int)$params['id'];
        $res = $this->api->put("/subjects/{$id}", [
            'name' => trim($_POST['name'] ?? ''),
            'code' => trim($_POST['code'] ?? ''),
        ]);
        if ($res['success'] ?? false) $this->redirect('/subjects', 'Subject updated.');
        Session::flash('error', $res['error'] ?? 'Update failed.');
        $this->redirect("/subjects/{$id}/edit");
    }
}