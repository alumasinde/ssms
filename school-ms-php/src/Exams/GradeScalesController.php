<?php
namespace Exams;
use Core\Controller;
use Core\Session;

class GradeScalesController extends Controller
{
    // Preset grading systems for quick seeding
    private array $presets = [
        'kcse'  => [['A','80','100','Excellent'],['A-','75','79','Very Good'],['B+','70','74','Good'],['B','65','69','Good'],['B-','60','64','Above Average'],['C+','55','59','Average'],['C','50','54','Average'],['C-','45','49','Below Average'],['D+','40','44','Poor'],['D','35','39','Poor'],['D-','30','34','Very Poor'],['E','0','29','Fail']],
        'cbc'   => [['EE','80','100','Exceeds Expectation'],['ME','50','79','Meets Expectation'],['AE','30','49','Approaches Expectation'],['BE','0','29','Below Expectation']],
        'pct'   => [['A','75','100','Distinction'],['B','60','74','Credit'],['C','50','59','Pass'],['D','40','49','Marginal'],['E','0','39','Fail']],
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
        // Seed a preset
        if (!empty($_POST['preset']) && isset($this->presets[$_POST['preset']])) {
            foreach ($this->presets[$_POST['preset']] as [$grade,$min,$max,$remark]) {
                $this->api->post('/grade-scales', ['grade'=>$grade,'min_score'=>(float)$min,'max_score'=>(float)$max,'remark'=>$remark]);
            }
            $this->redirect('/grade-scales', ucfirst($_POST['preset']).' grade scale seeded.');
        }
        // Single entry
        $res = $this->api->post('/grade-scales', [
            'grade'     => trim($_POST['grade']   ?? ''),
            'min_score' => (float)($_POST['min_score'] ?? 0),
            'max_score' => (float)($_POST['max_score'] ?? 100),
            'remark'    => trim($_POST['remark']  ?? ''),
        ]);
        if ($res['success'] ?? false) { $this->redirect('/grade-scales', 'Grade scale entry added.'); }
        Session::flash('error', $res['error'] ?? 'Failed.');
        $this->redirect('/grade-scales');
    }
}
