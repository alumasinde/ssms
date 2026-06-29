<?php
namespace Timetable;
use Core\Controller;
use Core\Session;

class TimetableController extends Controller
{
    private array $days = [1=>'Monday',2=>'Tuesday',3=>'Wednesday',4=>'Thursday',5=>'Friday'];

    public function index(array $params = []): void
    {
        $this->requirePermission('timetable.view');
        $termID  = (int)($_GET['term_id']  ?? 0);
        $classID = (int)($_GET['class_id'] ?? 0);

        $terms   = $this->api->get('/terms');
        $classes = $this->api->get('/classes');
        $slots   = [];

        if ($classID && $termID) {
            $res   = $this->api->get("/timetable/class/{$classID}", ['term_id' => $termID]);
            $slots = $res['data'] ?? [];
        } elseif ($termID) {
            $res   = $this->api->get('/timetable', ['term_id' => $termID]);
            $slots = $res['data'] ?? [];
        }

        // Group by class then day
        $grid = [];
        foreach ($slots as $s) {
            $grid[$s['class_id']][$s['day_of_week']][] = $s;
        }

        $this->view('timetable/index', [
            'title'   => 'Timetable',
            'terms'   => $terms['data']   ?? [],
            'classes' => $classes['data'] ?? [],
            'slots'   => $slots,
            'grid'    => $grid,
            'termID'  => $termID,
            'classID' => $classID,
            'days'    => $this->days,
        ]);
    }

    public function manage(array $params = []): void
    {
        $this->requirePermission('timetable.edit');
        $classID = (int)($params['classId'] ?? $_GET['class_id'] ?? 0);
        $termID  = (int)($_GET['term_id'] ?? 0);

        $terms    = $this->api->get('/terms');
        $classes  = $this->api->get('/classes');
        $teachers = $this->api->get('/teachers');
        $subjects = $classID
            ? $this->api->get("/classes/{$classID}/subjects")
            : $this->api->get('/subjects');
        $slots = [];
        if ($classID && $termID) {
            $res   = $this->api->get("/timetable/class/{$classID}", ['term_id' => $termID]);
            $slots = $res['data'] ?? [];
        }

        $this->view('timetable/manage', [
            'title'    => 'Manage Timetable',
            'terms'    => $terms['data']    ?? [],
            'classes'  => $classes['data']  ?? [],
            'teachers' => $teachers['data'] ?? [],
            'subjects' => $subjects['data'] ?? [],
            'slots'    => $slots,
            'classID'  => $classID,
            'termID'   => $termID,
            'days'     => $this->days,
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('timetable.edit');
        $res = $this->api->post('/timetable', [
            'class_id'   => (int)($_POST['class_id']   ?? 0),
            'subject_id' => (int)($_POST['subject_id'] ?? 0),
            'teacher_id' => (int)($_POST['teacher_id'] ?? 0),
            'term_id'    => (int)($_POST['term_id']    ?? 0),
            'day_of_week'=> (int)($_POST['day_of_week']?? 1),
            'start_time' => $_POST['start_time'] ?? '',
            'end_time'   => $_POST['end_time']   ?? '',
            'room'       => trim($_POST['room']  ?? ''),
        ]);
        $back = "/timetable/manage?class_id={$_POST['class_id']}&term_id={$_POST['term_id']}";
        if ($res['success'] ?? false) {
            $this->redirect($back, 'Slot added.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to add slot.');
        $this->redirect($back);
    }

    public function delete(array $params): void
    {
        $this->requirePermission('timetable.edit');
        $id   = (int)$params['id'];
        $back = $_POST['back'] ?? '/timetable';
        $this->api->delete("/timetable/{$id}");
        $this->redirect($back, 'Slot removed.');
    }
}
