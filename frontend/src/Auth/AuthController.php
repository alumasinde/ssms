<?php

namespace Auth;

use Core\Controller;
use Core\Session;
use Core\ApiClient;

class AuthController extends Controller
{
    public function loginForm(array $params = []): void
    {
        if (Session::isLoggedIn()) {
            $this->redirect('/dashboard');
        }
        // viewPartial renders the login page which has its own full HTML
        $this->viewPartial('auth/login', ['title' => 'Sign In']);
    }

    public function login(array $params = []): void
    {
        $email    = trim($_POST['email'] ?? '');
        $password = $_POST['password'] ?? '';

        if (!$email || !$password) {
            Session::flash('error', 'Email and password are required.');
            $this->redirect('/auth/login');
        }

        // Domain-based login: only email + password.
        // The tenant is resolved by the Go backend from the X-Tenant-Slug header
        // sent automatically by ApiClient (derived from Host or .env TENANT_SLUG).
        $api = new ApiClient();
        $res = $api->post('/auth/login', [
            'email'    => $email,
            'password' => $password,
        ]);

        if (!($res['success'] ?? false)) {
            Session::flash('error', $res['error'] ?? 'Invalid credentials.');
            $this->redirect('/auth/login');
        }

        $payload = $res['data'];
        Session::set('access_token',  $payload['access_token']);
        Session::set('refresh_token', $payload['refresh_token']);
        Session::set('user',          $payload['user']);
        Session::set('tenant_id',     $payload['user']['tenant_id']);
        Session::set('school_id',     $payload['user']['school_id']);
        Session::set('permissions',   $payload['permissions'] ?? []);
        Session::set('roles',         $payload['user']['roles'] ?? []);

        $this->redirect('/dashboard');
    }

    public function logout(array $params = []): void
    {
        Session::destroy();
        $this->redirect('/auth/login');
    }
}
