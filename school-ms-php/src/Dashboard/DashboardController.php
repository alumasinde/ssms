<?php
namespace Dashboard;

use Core\Controller;
use Core\Session;

class DashboardController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requireAuth();

        $totalStudents = 0;
        $totalTeachers = 0;
        $totalParents  = 0;
        $recentNotices = [];
        $currentTerm   = null;

        if ($this->can('students.view')) {
            $s = $this->api->get('/students', ['per_page' => 1]);
            $totalStudents = $s['meta']['total'] ?? 0;
        }
        if ($this->can('teachers.view')) {
            $t = $this->api->get('/teachers');
            $totalTeachers = count($t['data'] ?? []);
        }
        if ($this->can('parents.view')) {
            $p = $this->api->get('/parents');
            $totalParents = count($p['data'] ?? []);
        }
        if ($this->can('notices.view')) {
            $n = $this->api->get('/notices');
            $recentNotices = array_slice($n['data'] ?? [], 0, 5);
        }
        // Always try to load current term for context strip
        $tRes = $this->api->get('/terms/current');
        if ($tRes['success'] ?? false) {
            $currentTerm = $tRes['data'];
        }

        $this->view('dashboard/index', [
            'title'         => 'Dashboard',
            'totalStudents' => $totalStudents,
            'totalTeachers' => $totalTeachers,
            'totalParents'  => $totalParents,
            'recentNotices' => $recentNotices,
            'currentTerm'   => $currentTerm,
        ]);
    }
}
