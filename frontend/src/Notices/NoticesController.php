<?php
namespace Notices;
use Core\Controller;
use Core\Session;

class NoticesController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('notices.view');
        $audience = $_GET['audience'] ?? '';
        $res = $this->api->get('/notices', ['audience' => $audience ?: null]);
        $this->view('notices/index', [
            'title'    => 'Notices',
            'notices'  => $res['data'] ?? [],
            'audience' => $audience,
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('notices.create');
        $this->view('notices/create', ['title' => 'Post Notice']);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('notices.create');
        $res = $this->api->post('/notices', [
            'title'    => trim($_POST['title'] ?? ''),
            'body'     => trim($_POST['body'] ?? ''),
            'audience' => $_POST['audience'] ?? 'all',
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/notices', 'Notice published.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to publish notice.');
        $this->redirect('/notices/create');
    }

    public function show(array $params): void
    {
        $this->requirePermission('notices.view');
        $res = $this->api->get("/notices/{$params['id']}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/notices', 'Notice not found.', 'error');
        }
        $this->view('notices/show', [
            'title'  => $res['data']['title'] ?? 'Notice',
            'notice' => $res['data'] ?? [],
        ]);
    }

    public function delete(array $params): void
    {
        $this->requirePermission('notices.delete');
        $this->api->delete("/notices/{$params['id']}");
        $this->redirect('/notices', 'Notice deleted.');
    }
}
