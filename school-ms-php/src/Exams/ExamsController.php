<?php
namespace Exams;
use Core\Controller;
class ExamsController extends Controller {
    public function index(array $params = []): void {
        $this->requireAuth();
        $res = $this->api->get('/exams');
        $this->view('exams/index', ['title' => 'Exams', 'exams' => $res['data'] ?? []]);
    }
    public function create(array $params = []): void {
        $this->requireAuth();
        $this->view('exams/create', ['title' => 'Create Exam']);
    }
    public function store(array $params = []): void {
        $this->requireAuth();
        $res = $this->api->post('/exams', ['name' => $_POST['name'], 'type' => $_POST['type'], 'term_id' => (int)$_POST['term_id'], 'start_date' => $_POST['start_date'], 'end_date' => $_POST['end_date']]);
        if ($res['success'] ?? false) $this->redirect('/exams', 'Exam created.');
        $this->redirect('/exams/create', $res['error'] ?? 'Failed.', 'error');
    }
    public function results(array $params): void {
        $this->requireAuth();
        $examID = $params['id'];
        $res    = $this->api->get("/exams/{$examID}/results");
        $exam   = $this->api->get("/exams/{$examID}");
        $this->view('exams/results', ['title' => 'Exam Results', 'results' => $res['data'] ?? [], 'exam' => $exam['data'] ?? []]);
    }
}
