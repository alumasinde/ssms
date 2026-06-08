<?php

declare(strict_types=1);

define('BASE_PATH', dirname(__DIR__));

// PSR-4 style autoloader
spl_autoload_register(function (string $class): void {
    $file = BASE_PATH . '/src/' . str_replace('\\', '/', $class) . '.php';
    if (file_exists($file)) {
        require $file;
    }
});

use Core\Router;
use Core\Session;
use Core\Config;

Config::load();
Session::start();

$router = new Router();

// ── Auth ──────────────────────────────────────────────────────────────────────
$router->get('/auth/login',  [\Auth\AuthController::class, 'loginForm']);
$router->post('/auth/login', [\Auth\AuthController::class, 'login']);
$router->get('/auth/logout', [\Auth\AuthController::class, 'logout']);

// ── Root redirect ─────────────────────────────────────────────────────────────
$router->get('/', function () {
    header('Location: ' . (Session::isLoggedIn() ? '/dashboard' : '/auth/login'));
    exit;
});

// ── Dashboard ─────────────────────────────────────────────────────────────────
$router->get('/dashboard', [\Dashboard\DashboardController::class, 'index']);

// ── Students ─ IMPORTANT: static routes MUST come before /{id} wildcard ──────
$router->get('/students',              [\Students\StudentsController::class, 'index']);
$router->get('/students/create',       [\Students\StudentsController::class, 'create']);
$router->post('/students',             [\Students\StudentsController::class, 'store']);
$router->get('/students/{id}/edit',    [\Students\StudentsController::class, 'edit']);    // ← moved up
$router->post('/students/{id}/update', [\Students\StudentsController::class, 'update']);  // ← moved up
$router->get('/students/{id}',         [\Students\StudentsController::class, 'show']);    // ← now last

// ── Teachers ──────────────────────────────────────────────────────────────────
$router->get('/teachers',          [\Teachers\TeachersController::class, 'index']);
$router->get('/teachers/create',   [\Teachers\TeachersController::class, 'create']);
$router->post('/teachers',         [\Teachers\TeachersController::class, 'store']);
$router->get('/teachers/{id}',     [\Teachers\TeachersController::class, 'show']);

// ── Classes ───────────────────────────────────────────────────────────────────
$router->get('/classes',              [\Classes\ClassesController::class, 'index']);
$router->get('/classes/create',       [\Classes\ClassesController::class, 'create']);
$router->post('/classes',             [\Classes\ClassesController::class, 'store']);
$router->get('/classes/{id}/edit',    [\Classes\ClassesController::class, 'edit']);
$router->post('/classes/{id}/update', [\Classes\ClassesController::class, 'update']);
$router->post('/classes/{id}/delete', [\Classes\ClassesController::class, 'delete']);

// ── Subjects ──────────────────────────────────────────────────────────────────
$router->get('/subjects',              [\Subjects\SubjectsController::class, 'index']);
$router->get('/subjects/create',       [\Subjects\SubjectsController::class, 'create']);
$router->post('/subjects',             [\Subjects\SubjectsController::class, 'store']);
$router->get('/subjects/{id}/edit',    [\Subjects\SubjectsController::class, 'edit']);
$router->post('/subjects/{id}/update', [\Subjects\SubjectsController::class, 'update']);

// ── Parents ───────────────────────────────────────────────────────────────────
$router->get('/parents',             [\Parents\ParentsController::class, 'index']);
$router->get('/parents/create',      [\Parents\ParentsController::class, 'create']);
$router->post('/parents',            [\Parents\ParentsController::class, 'store']);
$router->get('/parents/{id}',        [\Parents\ParentsController::class, 'show']);
$router->post('/parents/link-student', [\Parents\ParentsController::class, 'linkStudent']);

// ── Attendance ────────────────────────────────────────────────────────────────
$router->get('/attendance',          [\Attendance\AttendanceController::class, 'index']);
$router->get('/attendance/mark',     [\Attendance\AttendanceController::class, 'mark']);
$router->post('/attendance',         [\Attendance\AttendanceController::class, 'store']);
$router->get('/attendance/summary',  [\Attendance\AttendanceController::class, 'summary']);

// ── Exams ─────────────────────────────────────────────────────────────────────
$router->get('/exams',               [\Exams\ExamsController::class, 'index']);
$router->get('/exams/create',        [\Exams\ExamsController::class, 'create']);
$router->post('/exams',              [\Exams\ExamsController::class, 'store']);
$router->get('/exams/{id}/results',  [\Exams\ExamsController::class, 'results']);
$router->post('/exams/results',      [\Exams\ExamsController::class, 'submitResults']);

// ── Finance ───────────────────────────────────────────────────────────────────
$router->get('/finance',                       [\Finance\FinanceController::class, 'index']);
$router->get('/finance/statement/{studentId}', [\Finance\FinanceController::class, 'statement']);
$router->post('/finance/payment',              [\Finance\FinanceController::class, 'recordPayment']);

// ── Notices ───────────────────────────────────────────────────────────────────
$router->get('/notices',          [\Notices\NoticesController::class, 'index']);
$router->get('/notices/create',   [\Notices\NoticesController::class, 'create']);
$router->post('/notices',         [\Notices\NoticesController::class, 'store']);
$router->get('/notices/{id}',     [\Notices\NoticesController::class, 'show']);

// ── Reports ───────────────────────────────────────────────────────────────────
$router->get('/reports/report-card/{studentId}', [\Reports\ReportsController::class, 'reportCard']);
$router->get('/reports/class-results',            [\Reports\ReportsController::class, 'classResults']);
$router->get('/reports/fee-collection',           [\Reports\ReportsController::class, 'feeCollection']);
$router->get('/reports/attendance-summary',       [\Reports\ReportsController::class, 'attendanceSummary']);

$router->dispatch();
