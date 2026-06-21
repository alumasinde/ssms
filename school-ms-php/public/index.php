<?php

declare(strict_types=1);

define('BASE_PATH', dirname(__DIR__));

spl_autoload_register(function (string $class): void {
    $file = BASE_PATH . '/src/' . str_replace('\\', '/', $class) . '.php';
    if (file_exists($file)) require $file;
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

// ── Students ──────────────────────────────────────────────────────────────────
$router->get('/students',              [\Students\StudentsController::class, 'index']);
$router->get('/students/create',       [\Students\StudentsController::class, 'create']);
$router->post('/students',             [\Students\StudentsController::class, 'store']);
$router->get('/students/{id}/edit',    [\Students\StudentsController::class, 'edit']);
$router->post('/students/{id}/update', [\Students\StudentsController::class, 'update']);
$router->get('/students/{id}',         [\Students\StudentsController::class, 'show']);

// ── Teachers ──────────────────────────────────────────────────────────────────
$router->get('/teachers/subject-admin',    [\Teachers\TeachersController::class, 'subjectAdmin']);
$router->post('/teachers/assign-subject',  [\Teachers\TeachersController::class, 'assignSubject']);
$router->post('/teachers/remove-subject',  [\Teachers\TeachersController::class, 'removeSubject']);
$router->get('/teachers/create',           [\Teachers\TeachersController::class, 'create']);
$router->post('/teachers',                 [\Teachers\TeachersController::class, 'store']);
$router->get('/teachers/{id}/edit',        [\Teachers\TeachersController::class, 'edit']);
$router->put('/teachers/{id}',             [\Teachers\TeachersController::class, 'update']);
$router->post('/teachers/{id}/update',     [\Teachers\TeachersController::class, 'update']);
$router->get('/teachers/{id}',             [\Teachers\TeachersController::class, 'show']);
$router->get('/teachers',                  [\Teachers\TeachersController::class, 'index']);

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
$router->post('/subjects/{id}/delete', [\Subjects\SubjectsController::class, 'delete']);

// ── Terms ─────────────────────────────────────────────────────────────────────
$router->get('/terms',                    [\Terms\TermsController::class, 'index']);
$router->get('/terms/create',             [\Terms\TermsController::class, 'create']);
$router->post('/terms',                   [\Terms\TermsController::class, 'store']);
$router->get('/terms/{id}/edit',          [\Terms\TermsController::class, 'edit']);
$router->post('/terms/{id}/update',       [\Terms\TermsController::class, 'update']);
$router->post('/terms/{id}/set-current',  [\Terms\TermsController::class, 'setCurrent']);
$router->post('/terms/{id}/delete',       [\Terms\TermsController::class, 'delete']);

// ── Parents ───────────────────────────────────────────────────────────────────
$router->get('/parents',               [\Parents\ParentsController::class, 'index']);
$router->get('/parents/create',        [\Parents\ParentsController::class, 'create']);
$router->post('/parents',              [\Parents\ParentsController::class, 'store']);
$router->post('/parents/link-student', [\Parents\ParentsController::class, 'linkStudent']);
$router->get('/parents/{id}',          [\Parents\ParentsController::class, 'show']);

// ── Attendance ────────────────────────────────────────────────────────────────
$router->get('/attendance',         [\Attendance\AttendanceController::class, 'index']);
$router->get('/attendance/mark',    [\Attendance\AttendanceController::class, 'mark']);
$router->post('/attendance',        [\Attendance\AttendanceController::class, 'store']);
$router->get('/attendance/summary', [\Attendance\AttendanceController::class, 'summary']);

// ── Exams ─────────────────────────────────────────────────────────────────────
$router->get('/exams',              [\Exams\ExamsController::class, 'index']);
$router->get('/exams/create',       [\Exams\ExamsController::class, 'create']);
$router->post('/exams',             [\Exams\ExamsController::class, 'store']);
$router->post('/exams/results',     [\Exams\ExamsController::class, 'submitResults']);
$router->get('/exams/{id}/results', [\Exams\ExamsController::class, 'results']);
$router->get('/exams/{id}',         [\Exams\ExamsController::class, 'show']);

// ── Finance ───────────────────────────────────────────────────────────────────
$router->get('/finance',                       [\Finance\FinanceController::class, 'index']);
$router->get('/finance/fee-types/create',      [\Finance\FinanceController::class, 'createFeeType']);
$router->post('/finance/fee-types',            [\Finance\FinanceController::class, 'storeFeeType']);
$router->get('/finance/invoices/generate',     [\Finance\FinanceController::class, 'generateInvoices']);
$router->post('/finance/invoices/generate',    [\Finance\FinanceController::class, 'storeInvoices']);
$router->get('/finance/statement/{studentId}', [\Finance\FinanceController::class, 'statement']);
$router->post('/finance/payment',              [\Finance\FinanceController::class, 'recordPayment']);
$router->get('/finance/discounts/create',      [\Finance\FinanceController::class, 'createDiscount']);
$router->post('/finance/discounts',            [\Finance\FinanceController::class, 'storeDiscount']);

// ── Notices ───────────────────────────────────────────────────────────────────
$router->get('/notices',              [\Notices\NoticesController::class, 'index']);
$router->get('/notices/create',       [\Notices\NoticesController::class, 'create']);
$router->post('/notices',             [\Notices\NoticesController::class, 'store']);
$router->post('/notices/{id}/delete', [\Notices\NoticesController::class, 'delete']);
$router->get('/notices/{id}',         [\Notices\NoticesController::class, 'show']);

// ── Reports ───────────────────────────────────────────────────────────────────
$router->get('/reports/report-card/{studentId}', [\Reports\ReportsController::class, 'reportCard']);
$router->get('/reports/class-results',            [\Reports\ReportsController::class, 'classResults']);
$router->get('/reports/fee-collection',           [\Reports\ReportsController::class, 'feeCollection']);
$router->get('/reports/attendance-summary',       [\Reports\ReportsController::class, 'attendanceSummary']);
$router->get('/reports/subject-performance',      [\Reports\ReportsController::class, 'subjectPerformance']);
$router->get('/reports/promote',                  [\Reports\ReportsController::class, 'promote']);
$router->post('/reports/promote',                 [\Reports\ReportsController::class, 'storePromotion']);

// ── Timetable ─────────────────────────────────────────────────────────────────
$router->get('/timetable',                        [\Timetable\TimetableController::class, 'index']);
$router->get('/timetable/manage',                 [\Timetable\TimetableController::class, 'manage']);
$router->post('/timetable',                       [\Timetable\TimetableController::class, 'store']);
$router->post('/timetable/{id}/delete',           [\Timetable\TimetableController::class, 'delete']);

// ── Assignments ───────────────────────────────────────────────────────────────
$router->get('/assignments',                      [\Assignments\AssignmentsController::class, 'index']);
$router->get('/assignments/create',               [\Assignments\AssignmentsController::class, 'create']);
$router->post('/assignments',                     [\Assignments\AssignmentsController::class, 'store']);
$router->post('/assignments/{id}/delete',         [\Assignments\AssignmentsController::class, 'delete']);

// ── Discipline ────────────────────────────────────────────────────────────────
$router->get('/discipline',                       [\Discipline\DisciplineController::class, 'index']);
$router->get('/discipline/create',                [\Discipline\DisciplineController::class, 'create']);
$router->post('/discipline',                      [\Discipline\DisciplineController::class, 'store']);
$router->post('/discipline/{id}/delete',          [\Discipline\DisciplineController::class, 'delete']);

// ── Staff Attendance ──────────────────────────────────────────────────────────
$router->get('/staff-attendance',                 [\StaffAttendance\StaffAttendanceController::class, 'index']);
$router->get('/staff-attendance/mark',            [\StaffAttendance\StaffAttendanceController::class, 'mark']);
$router->post('/staff-attendance',                [\StaffAttendance\StaffAttendanceController::class, 'store']);
$router->get('/staff-attendance/summary',         [\StaffAttendance\StaffAttendanceController::class, 'summary']);

// ── Grade Scales ──────────────────────────────────────────────────────────────
$router->get('/grade-scales',                     [\Exams\GradeScalesController::class, 'index']);
$router->post('/grade-scales',                    [\Exams\GradeScalesController::class, 'store']);

// ── Users ─────────────────────────────────────────────────────────────────────
$router->get('/users',                            [\Users\UsersController::class, 'index']);
$router->get('/users/create',                     [\Users\UsersController::class, 'create']);
$router->post('/users',                           [\Users\UsersController::class, 'store']);
$router->post('/users/{id}/role',                 [\Users\UsersController::class, 'updateRole']);
$router->post('/users/{id}/deactivate',           [\Users\UsersController::class, 'deactivate']);
$router->post('/users/{id}/activate',             [\Users\UsersController::class, 'activate']);

// ── Class Subjects ────────────────────────────────────────────────────────────
$router->get('/classes/{id}/subjects',                      [\Classes\ClassesController::class, 'subjects']);
$router->post('/classes/{id}/subjects',                     [\Classes\ClassesController::class, 'assignSubject']);
$router->post('/classes/{id}/subjects/{subjectId}/remove',  [\Classes\ClassesController::class, 'removeSubject']);
$router->dispatch();
