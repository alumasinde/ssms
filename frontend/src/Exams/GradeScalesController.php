<?php
namespace Exams;

use Core\Controller;
use Core\Session;

class GradeScalesController extends Controller
{
    private array $presets = [
        'kcse' => [
            ['A',  '80',  '100', 'Excellent'],
            ['A-', '75',  '79',  'Very Good'],
            ['B+', '70',  '74',  'Good'],
            ['B',  '65',  '69',  'Good'],
            ['B-', '60',  '64',  'Above Average'],
            ['C+', '55',  '59',  'Average'],
            ['C',  '50',  '54',  'Average'],
            ['C-', '45',  '49',  'Below Average'],
            ['D+', '40',  '44',  'Poor'],
            ['D',  '35',  '39',  'Poor'],
            ['D-', '30',  '34',  'Very Poor'],
            ['E',  '0',   '29',  'Fail'],
        ],
        'cbc' => [
            ['EE', '80', '100', 'Exceeds Expectation'],
            ['ME', '50', '79',  'Meets Expectation'],
            ['AE', '30', '49',  'Approaches Expectation'],
            ['BE', '0',  '29',  'Below Expectation'],
        ],
        'pct' => [
            ['A', '75', '100', 'Distinction'],
            ['B', '60', '74',  'Credit'],
            ['C', '50', '59',  'Pass'],
            ['D', '40', '49',  'Marginal'],
            ['E', '0',  '39',  'Fail'],
        ],
    ];

    public function index(array $params = []): void
    {
        $this->requirePermission('exams.view');
        $res = $this->api->get('/grade-scales');
        $this->view('grade_scales/index', [
            'title'   => 'Grade Scales',
            'scales'  => $res['data'] ?? [],
            'presets' => array_keys($this->presets),
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('exams.create');

        // ── Seed a preset ─────────────────────────────────────────────────
        if (!empty($_POST['preset']) && isset($this->presets[$_POST['preset']])) {
            $preset = $_POST['preset'];

            // Clear existing before seeding so there are no duplicates
            $this->clearAll();

            $errors = 0;
            foreach ($this->presets[$preset] as [$grade, $min, $max, $remark]) {
                $res = $this->api->post('/grade-scales', [
                    'grade'     => $grade,
                    'min_score' => (float)$min,
                    'max_score' => (float)$max,
                    'remark'    => $remark,
                ]);
                if (!($res['success'] ?? false)) {
                    $errors++;
                }
            }

            if ($errors === 0) {
                $this->redirect('/grade-scales', strtoupper($preset) . ' grade scale applied successfully.');
            }
            Session::flash('error', "Preset applied with {$errors} error(s). Check the scale below.");
            $this->redirect('/grade-scales');
            return;
        }

        // ── Single manual entry ───────────────────────────────────────────
        $grade    = trim($_POST['grade']    ?? '');
        $minScore = (float)($_POST['min_score'] ?? 0);
        $maxScore = (float)($_POST['max_score'] ?? 100);
        $remark   = trim($_POST['remark']   ?? '');

        if ($grade === '') {
            Session::flash('error', 'Grade is required.');
            $this->redirect('/grade-scales');
            return;
        }
        if ($minScore >= $maxScore) {
            Session::flash('error', 'Min score must be less than max score.');
            $this->redirect('/grade-scales');
            return;
        }

        $res = $this->api->post('/grade-scales', [
            'grade'     => $grade,
            'min_score' => $minScore,
            'max_score' => $maxScore,
            'remark'    => $remark,
        ]);

        if ($res['success'] ?? false) {
            $this->redirect('/grade-scales', "Grade '{$grade}' added.");
        }
        Session::flash('error', $res['error'] ?? 'Failed to add grade entry.');
        $this->redirect('/grade-scales');
    }

public function update(array $params): void
{
    $this->requirePermission('settings.manage');
    $id   = (int)$params['id'];
    $body = json_decode(file_get_contents('php://input'), true) ?? [];

    if (empty($body['grade'])) {
        $this->json(['success' => false, 'error' => 'Grade is required.'], 400);
    }
    if ((float)($body['min_score'] ?? 0) >= (float)($body['max_score'] ?? 0)) {
        $this->json(['success' => false, 'error' => 'Min must be less than Max.'], 400);
    }

    $res = $this->api->put("/grade-scales/{$id}", [
        'grade'     => trim($body['grade']),
        'min_score' => (float)$body['min_score'],
        'max_score' => (float)$body['max_score'],
        'remark'    => trim($body['remark'] ?? ''),
    ]);

    $this->json($res);
}
    public function delete(array $params): void
    {
        $this->requirePermission('exams.delete');
        $id  = (int)$params['id'];
        $res = $this->api->delete("/grade-scales/{$id}");

        if ($res['success'] ?? false) {
            $this->redirect('/grade-scales', 'Grade entry deleted.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to delete.');
        $this->redirect('/grade-scales');
    }

    public function clear(array $params = []): void
    {
        $this->requirePermission('exams.delete');
        $this->clearAll();
        $this->redirect('/grade-scales', 'All grade scale entries cleared.');
    }

    // ── Private helper ────────────────────────────────────────────────────────

    private function clearAll(): void
    {
        $res    = $this->api->get('/grade-scales');
        $scales = $res['data'] ?? [];
        foreach ($scales as $s) {
            $this->api->delete('/grade-scales/' . $s['id']);
        }
    }
}