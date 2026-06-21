<?php
namespace Classes;
use Core\Controller;
use Core\Session;

class ClassesController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('classes.view');
        $res = $this->api->get('/classes');
        $this->view('classes/index', [
            'title'   => 'Classes',
            'classes' => $res['data'] ?? [],
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('classes.create');
        $this->view('classes/create', ['title' => 'Create Class']);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('classes.create');
        $res = $this->api->post('/classes', [
            'name'   => trim($_POST['name']   ?? ''),
            'level'  => trim($_POST['level']  ?? ''),
            'stream' => trim($_POST['stream'] ?? '') ?: null,
        ]);
        if ($res['success'] ?? false) { $this->redirect('/classes', 'Class created.'); }
        Session::flash('error', $res['error'] ?? 'Failed to create class.');
        $this->redirect('/classes/create');
    }

    public function edit(array $params): void
    {
        $this->requirePermission('classes.edit');
        $res = $this->api->get("/classes/{$params['id']}");
        if (!($res['success'] ?? false)) { $this->redirect('/classes','Class not found.','error'); }
        $this->view('classes/edit', ['title'=>'Edit Class','class'=>$res['data']??[]]);
    }

    public function update(array $params): void
    {
        $this->requirePermission('classes.edit');
        $id  = (int)$params['id'];
        $res = $this->api->put("/classes/{$id}", [
            'name'   => trim($_POST['name']   ?? ''),
            'level'  => trim($_POST['level']  ?? ''),
            'stream' => trim($_POST['stream'] ?? '') ?: null,
        ]);
        if ($res['success'] ?? false) { $this->redirect('/classes','Class updated.'); }
        Session::flash('error', $res['error'] ?? 'Update failed.');
        $this->redirect("/classes/{$id}/edit");
    }

    public function delete(array $params): void
    {
        $this->requirePermission('classes.delete');
        $res = $this->api->delete("/classes/{$params['id']}");
        if ($res['success'] ?? false) { $this->redirect('/classes','Class deleted.'); }
        Session::flash('error', $res['error'] ?? 'Could not delete class.');
        $this->redirect('/classes');
    }

    // ── Class Subjects ─────────────────────────────────────────────────────
    public function subjects(array $params): void
    {
        $this->requirePermission('classes.view');
        $id        = (int)$params['id'];
        $classRes  = $this->api->get("/classes/{$id}");
        $assigned  = $this->api->get("/classes/{$id}/subjects");
        $unassigned = $this->api->get("/classes/{$id}/subjects/unassigned");
        $this->view('classes/subjects', [
            'title'      => 'Class Subjects',
            'class'      => $classRes['data'] ?? [],
            'assigned'   => $assigned['data']   ?? [],
            'unassigned' => $unassigned['data']  ?? [],
        ]);
    }

    public function assignSubject(array $params): void
    {
        $this->requirePermission('classes.edit');
        $id = (int)$params['id'];
        $this->api->post("/classes/{$id}/subjects", [
            'subject_id'   => (int)($_POST['subject_id'] ?? 0),
            'is_compulsory' => isset($_POST['is_compulsory']),
        ]);
        $this->redirect("/classes/{$id}/subjects", 'Subject assigned.');
    }

    public function removeSubject(array $params): void
    {
        $this->requirePermission('classes.edit');
        $id        = (int)$params['id'];
        $subjectID = (int)$params['subjectId'];
        $this->api->delete("/classes/{$id}/subjects/{$subjectID}");
        $this->redirect("/classes/{$id}/subjects", 'Subject removed.');
    }
}
