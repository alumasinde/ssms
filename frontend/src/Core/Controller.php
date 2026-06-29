<?php

namespace Core;

class Controller
{
    protected ApiClient $api;

    public function __construct()
    {
        Session::start();
        $this->api = new ApiClient();
    }

    protected function can(string $permission): bool
{
    $permissions = \Core\Session::get('permissions', []);
    return in_array($permission, $permissions);
}

protected function requirePermission(string $permission): void
{
    if (!$this->can($permission)) {
        \Core\Session::flash('error', 'You do not have permission to do that.');
        $this->redirect('/dashboard');
    }
}
    protected function requireAuth(): void
    {
        if (!Session::isLoggedIn()) {
            header('Location: /auth/login');
            exit;
        }
    }

    protected function requireRole(string ...$roles): void
{
    $this->requireAuth();
    // Backend now returns roles[] array not role string
    $userRoles = Session::get('roles', []);
    foreach ($roles as $role) {
        if (in_array($role, $userRoles, true)) return;
    }
    http_response_code(403);
    $this->view('403');
    exit;
}

    /**
     * Render a view inside the main layout.
     * $data is extracted into variables available in the view.
     */
    protected function view(string $template, array $data = []): void
    {
        extract($data);
        $appName      = Config::appName();
        $user         = Session::user();
        $flashError   = Session::flash('error');
        $flashSuccess = Session::flash('success');
        $viewFile     = BASE_PATH . '/views/' . $template . '.php';
        if (!file_exists($viewFile)) {
            $viewFile = BASE_PATH . '/views/404.php';
        }
        include BASE_PATH . '/views/layouts/app.php';
    }

    /**
     * Render a view WITHOUT the main layout (e.g. login page, which has its own full HTML).
     */
    protected function viewPartial(string $template, array $data = []): void
    {
        extract($data);
        include BASE_PATH . '/views/' . $template . '.php';
    }

    protected function redirect(string $url, ?string $flash = null, string $flashKey = 'success'): void
    {
        if ($flash) Session::flash($flashKey, $flash);
        header('Location: ' . $url);
        exit;
    }

    protected function json(mixed $data, int $status = 200): void
    {
        http_response_code($status);
        header('Content-Type: application/json');
        echo json_encode($data);
        exit;
    }

    protected function redirectBack(?string $flash = null, string $flashKey = 'success'): void
    {
        if ($flash) Session::flash($flashKey, $flash);
        $backUrl = Session::flash('back_url') ?? '/';
        header('Location: ' . $backUrl);
        exit;
    }
}

