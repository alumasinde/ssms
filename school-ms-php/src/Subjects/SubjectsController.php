<?php
namespace Subjects;
use Core\Controller;
use Core\Session;

class SubjectsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('subjects.view');
        $res = $this->api->get('/subjects');
        $this->view('subjects/index', [
            'title'    => 'Subjects',
            'subjects' => $res['data'] ?? [],
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('subjects.create');
        $this->view('subjects/create', ['title' => 'Add Subject']);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('subjects.create');
        $res = $this->api->post('/subjects', [
            'name' => trim($_POST['name'] ?? ''),
            'code' => strtoupper(trim($_POST['code'] ?? '')),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/subjects', 'Subject added.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to add subject.');
        $this->redirect('/subjects/create');
    }

    public function edit(array $params): void
    {
        $this->requirePermission('subjects.edit');
        $res = $this->api->get("/subjects/{$params['id']}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/subjects', 'Subject not found.', 'error');
        }
        $this->view('subjects/edit', [
            'title'   => 'Edit Subject',
            'subject' => $res['data'] ?? [],
        ]);
    }

    public function update(array $params): void
    {
        $this->requirePermission('subjects.edit');
        $id  = (int)$params['id'];
        $res = $this->api->put("/subjects/{$id}", [
            'name' => trim($_POST['name'] ?? ''),
            'code' => strtoupper(trim($_POST['code'] ?? '')),
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/subjects', 'Subject updated.');
        }
        Session::flash('error', $res['error'] ?? 'Update failed.');
        $this->redirect("/subjects/{$id}/edit");
    }

    public function delete(array $params): void
    {
        $this->requirePermission('subjects.delete');
        $this->api->delete("/subjects/{$params['id']}");
        $this->redirect('/subjects', 'Subject deleted.');
    }
}
