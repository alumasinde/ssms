<?php
namespace Discipline;
use Core\Controller;
use Core\Session;

class DisciplineController extends Controller
{
    private array $types = [
        'commendation' => 'Commendation',
        'minor_offence' => 'Minor Offence',
        'major_offence' => 'Major Offence',
        'suspension'    => 'Suspension',
        'expulsion'     => 'Expulsion',
    ];

    public function index(array $params = []): void
    {
        $this->requirePermission('discipline.view');
        $termID  = (int)($_GET['term_id'] ?? 0);
        $res     = $this->api->get('/discipline', $termID ? ['term_id' => $termID] : []);
        $terms   = $this->api->get('/terms');
        $this->view('discipline/index', [
            'title'   => 'Discipline Records',
            'records' => $res['data']   ?? [],
            'terms'   => $terms['data'] ?? [],
            'termID'  => $termID,
            'types'   => $this->types,
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('discipline.create');
        $studentID = (int)($_GET['student_id'] ?? 0);
        $students  = $this->api->get('/students', ['per_page' => 200]);
        $terms     = $this->api->get('/terms');
        $this->view('discipline/create', [
            'title'      => 'Record Incident',
            'students'   => $students['data'] ?? [],
            'terms'      => $terms['data']    ?? [],
            'studentID'  => $studentID,
            'types'      => $this->types,
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('discipline.create');
        $res = $this->api->post('/discipline', [
            'student_id'    => (int)($_POST['student_id']    ?? 0),
            'term_id'       => (int)($_POST['term_id']       ?? 0),
            'incident_date' => $_POST['incident_date']        ?? date('Y-m-d'),
            'type'          => $_POST['type']                 ?? 'minor_offence',
            'description'   => trim($_POST['description']    ?? ''),
            'action_taken'  => trim($_POST['action_taken']   ?? ''),
        ]);
        if ($res['success'] ?? false) { $this->redirect('/discipline', 'Record saved.'); }
        Session::flash('error', $res['error'] ?? 'Failed to save record.');
        $this->redirect('/discipline/create');
    }

    public function delete(array $params): void
    {
        $this->requirePermission('discipline.edit');
        $this->api->delete("/discipline/{$params['id']}");
        $this->redirect('/discipline', 'Record deleted.');
    }
}
