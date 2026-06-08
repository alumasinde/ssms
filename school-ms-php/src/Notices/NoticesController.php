<?php
namespace Notices;
use Core\Controller;
class NoticesController extends Controller {
    public function index(array $params = []): void {
        $this->requireAuth();
        $audience = $_GET['audience'] ?? '';
        $res = $this->api->get('/notices', ['audience' => $audience]);
        $this->view('notices/index', ['title' => 'Notices', 'notices' => $res['data'] ?? []]);
    }
    public function create(array $params = []): void {
        $this->requireAuth();
        $this->view('notices/create', ['title' => 'Post Notice']);
    }
    public function store(array $params = []): void {
        $this->requireAuth();
        $res = $this->api->post('/notices', [
            'title'    => trim($_POST['title'] ?? ''),
            'body'     => trim($_POST['body'] ?? ''),
            'audience' => $_POST['audience'] ?? 'all',
        ]);
        if ($res['success'] ?? false) $this->redirect('/notices', 'Notice published.');
        $this->redirect('/notices/create', $res['error'] ?? 'Failed.', 'error');
    }
    public function show(array $params): void {
        $this->requireAuth();
        $res = $this->api->get("/notices/{$params['id']}");
        $this->view('notices/show', ['title' => $res['data']['title'] ?? 'Notice', 'notice' => $res['data'] ?? []]);
    }
}
