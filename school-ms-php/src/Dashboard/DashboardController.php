<?php

namespace Dashboard;

use Core\Controller;

class DashboardController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requireAuth();

        
        // Quick stats: students, teachers, notices
        $students = $this->api->get('/students', ['per_page' => 1]);
        $teachers = $this->api->get('/teachers');
        $notices  = $this->api->get('/notices', ['audience' => 'all']);

        $totalStudents = $students['meta']['total'] ?? 0;
        $totalTeachers = count($teachers['data'] ?? []);
        $recentNotices = array_slice($notices['data'] ?? [], 0, 5);

        $this->view('dashboard/index', [
            'title'          => 'Dashboard',
            'totalStudents'  => $totalStudents,
            'totalTeachers'  => $totalTeachers,
            'recentNotices'  => $recentNotices,
        ]);
    }
}
