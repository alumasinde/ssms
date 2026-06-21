<?php
namespace Assignments;
use Core\Controller;
use Core\Session;

class AssignmentsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('assignments.view');
        $termID  = (int)($_GET['term_id']  ?? 0);
        $classID = (int)($_GET['class_id'] ?? 0);
        $q       = array_filter(['term_id' => $termID ?: null, 'class_id' => $classID ?: null]);
        $res     = $this->api->get('/assignments', $q);
        $terms   = $this->api->get('/terms');
        $classes = $this->api->get('/classes');
        $this->view('assignments/index', [
            'title'       => 'Assignments',
            'assignments' => $res['data']     ?? [],
            'terms'       => $terms['data']   ?? [],
            'classes'     => $classes['data'] ?? [],
            'termID'      => $termID,
            'classID'     => $classID,
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('assignments.create');
        $terms    = $this->api->get('/terms');
        $classes  = $this->api->get('/classes');
        $subjects = $this->api->get('/subjects');
        $teachers = $this->api->get('/teachers');
        $this->view('assignments/create', [
            'title'    => 'Create Assignment',
            'terms'    => $terms['data']    ?? [],
            'classes'  => $classes['data']  ?? [],
            'subjects' => $subjects['data'] ?? [],
            'teachers' => $teachers['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('assignments.create');
        $res = $this->api->post('/assignments', [
            'class_id'    => (int)($_POST['class_id']    ?? 0),
            'subject_id'  => (int)($_POST['subject_id']  ?? 0),
            'teacher_id'  => (int)($_POST['teacher_id']  ?? 0),
            'term_id'     => (int)($_POST['term_id']     ?? 0),
            'title'       => trim($_POST['title']        ?? ''),
            'description' => trim($_POST['description']  ?? ''),
            'due_date'    => $_POST['due_date']           ?? '',
            'max_marks'   => (float)($_POST['max_marks'] ?? 100),
        ]);
        if ($res['success'] ?? false) { $this->redirect('/assignments', 'Assignment created.'); }
        Session::flash('error', $res['error'] ?? 'Failed.');
        $this->redirect('/assignments/create');
    }

    public function delete(array $params): void
    {
        $this->requirePermission('assignments.edit');
        $this->api->delete("/assignments/{$params['id']}");
        $this->redirect('/assignments', 'Assignment deleted.');
    }
}
