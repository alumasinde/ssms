<?php
namespace Terms;

use Core\Controller;
use Core\Session;

class TermsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('academic_years.view');
        $res = $this->api->get('/terms');
        
        $this->view('terms/index', [
            'title' => 'Terms',
            'terms' => $res['data'] ?? [],
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('academic_years.create');
        $years = $this->api->get('/academic-years');
        $this->view('terms/create', [
            'title' => 'Create Term',
            'years' => $years['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('academic_years.create');
        $res = $this->api->post('/terms', [
            'academic_year_id' => (int)($_POST['academic_year_id'] ?? 0),
            'name'             => trim($_POST['name'] ?? ''),
            'start_date'       => $_POST['start_date'] ?? '',
            'end_date'         => $_POST['end_date'] ?? '',
            'is_current'       => isset($_POST['is_current']) ? true : false,
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/terms', 'Term created.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to create term.');
        $this->redirect('/terms/create');
    }

    public function edit(array $params): void
    {
        $this->requirePermission('academic_years.edit');
        $id  = (int)$params['id'];
        $res = $this->api->get("/terms/{$id}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/terms', 'Term not found.', 'error');
        }
        $years = $this->api->get('/academic-years');
        $this->view('terms/edit', [
            'title' => 'Edit Term',
            'term'  => $res['data'] ?? [],
            'years' => $years['data'] ?? [],
        ]);
    }

    public function update(array $params): void
    {
        $this->requirePermission('academic_years.edit');
        $id  = (int)$params['id'];
        $res = $this->api->put("/terms/{$id}", [
            'name'       => trim($_POST['name'] ?? ''),
            'start_date' => $_POST['start_date'] ?? '',
            'end_date'   => $_POST['end_date'] ?? '',
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/terms', 'Term updated.');
        }
        Session::flash('error', $res['error'] ?? 'Update failed.');
        $this->redirect("/terms/{$id}/edit");
    }

    public function setCurrent(array $params): void
    {
        $this->requirePermission('academic_years.edit');
        $id  = (int)$params['id'];
        $res = $this->api->post("/terms/{$id}/set-current", []);
        if ($res['success'] ?? false) {
            $this->redirect('/terms', 'Current term updated.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to set current term.');
        $this->redirect('/terms');
    }

    public function delete(array $params): void
    {
        $this->requirePermission('academic_years.edit');
        $id  = (int)$params['id'];
        $this->api->delete("/terms/{$id}");
        $this->redirect('/terms', 'Term deleted.');
    }
}
