<?php
namespace Exams;

use Core\Controller;
use Core\Session;

class ExamsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('exams.view');
        $termID  = (int)($_GET['term_id']  ?? 0);
        $classID = (int)($_GET['class_id'] ?? 0);

        $query   = array_filter(['term_id' => $termID ?: null, 'class_id' => $classID ?: null]);
        $res     = $this->api->get('/exams', $query);
        $terms   = $this->api->get('/terms');
        $classes = $this->api->get('/classes');

        $this->view('exams/index', [
            'title'   => 'Exams',
            'exams'   => $res['data']     ?? [],
            'terms'   => $terms['data']   ?? [],
            'classes' => $classes['data'] ?? [],
            'termID'  => $termID,
            'classID' => $classID,
        ]);
    }

    public function create(array $params = []): void
    {
        $this->requirePermission('exams.create');
        $terms   = $this->api->get('/terms');
        $classes = $this->api->get('/classes');
        $this->view('exams/create', [
            'title'   => 'Create Exam',
            'terms'   => $terms['data']   ?? [],
            'classes' => $classes['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('exams.create');
        $classID = (int)($_POST['class_id'] ?? 0);
        $res = $this->api->post('/exams', [
            'name'       => trim($_POST['name'] ?? ''),
            'type'       => $_POST['type']       ?? 'endterm',
            'term_id'    => (int)($_POST['term_id'] ?? 0),
            'class_id'   => $classID ?: null,
            'start_date' => $_POST['start_date'] ?? '',
            'end_date'   => $_POST['end_date']   ?? '',
        ]);
        if ($res['success'] ?? false) {
            $this->redirect('/exams', 'Exam created.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to create exam.');
        $this->redirect('/exams/create');
    }

    public function show(array $params): void
    {
        $this->requirePermission('exams.view');
        $id  = (int)$params['id'];
        $res = $this->api->get("/exams/{$id}");
        if (!($res['success'] ?? false)) {
            $this->redirect('/exams', 'Exam not found.', 'error');
        }
        $this->view('exams/show', [
            'title' => $res['data']['name'] ?? 'Exam',
            'exam'  => $res['data']         ?? [],
        ]);
    }

    public function results(array $params): void
    {
        $this->requirePermission('exams.view');
        $examID  = (int)$params['id'];
        $classID = (int)($_GET['class_id'] ?? 0);

        $exam     = $this->api->get("/exams/{$examID}");
        $examData = $exam['data'] ?? [];

        // Use class_id from exam itself if not passed as GET filter
        if (!$classID && !empty($examData['class_id'])) {
            $classID = (int)$examData['class_id'];
        }

        $query   = $classID ? ['class_id' => $classID] : [];
        $res     = $this->api->get("/exams/{$examID}/results", $query);

        // Students — scoped to class if known
        if ($classID) {
            $sRes     = $this->api->get("/students/class/{$classID}");
            $students = $sRes['data'] ?? [];
        } else {
            $sRes     = $this->api->get('/students', ['per_page' => 200]);
            $students = $sRes['data'] ?? [];
        }

        // Subjects — admin sees all, teacher sees only their assigned subjects
        if (Session::hasRole('teacher') && $classID) {
            $subRes  = $this->api->get('/teacher/my-subjects', ['class_id' => $classID]);
            $subjects = $subRes['data'] ?? [];

            // Fallback — if endpoint not available yet, load all
            if (empty($subjects)) {
                $subRes   = $this->api->get('/subjects');
                $subjects = $subRes['data'] ?? [];
            }
        } else {
            $subRes   = $this->api->get('/subjects');
            $subjects = $subRes['data'] ?? [];
        }

        $classes = $this->api->get('/classes');

        $this->view('exams/results', [
            'title'    => 'Results — ' . ($examData['name'] ?? ''),
            'results'  => $res['data']      ?? [],
            'exam'     => $examData,
            'students' => $students,
            'subjects' => $subjects,
            'classes'  => $classes['data']  ?? [],
            'classID'  => $classID,
        ]);
    }

    public function submitResults(array $params = []): void
    {
        // FIXED: permission matches what's in DB
        $this->requirePermission('exams.results');

        $examID  = (int)($_POST['exam_id']  ?? 0);
        $classID = (int)($_POST['class_id'] ?? 0);

        // FIXED: read new rows[N][field] structure from view
        $results = [];
        foreach ($_POST['rows'] ?? [] as $row) {
            $studentID = (int)($row['student_id'] ?? 0);
            $subjectID = (int)($row['subject_id'] ?? 0);
            $marksRaw  = $row['marks'] ?? '';

            // Skip blank cells — teacher left them empty intentionally
            if ($marksRaw === '' || $marksRaw === null) continue;
            if (!$studentID || !$subjectID) continue;

            $results[] = [
                'student_id' => $studentID,
                'subject_id' => $subjectID,
                'marks'      => (float)$marksRaw,
                'max_marks'  => (float)($row['max_marks'] ?? 100),
                'remarks'    => trim($row['remarks'] ?? ''),
            ];
        }

        if (empty($results)) {
            Session::flash('error', 'No marks were entered.');
            $this->redirect("/exams/{$examID}/results" . ($classID ? "?class_id={$classID}" : ''));
            return;
        }

        $res = $this->api->post('/exams/results', [
            'exam_id'  => $examID,
            'class_id' => $classID,
            'results'  => $results,
        ]);

        if ($res['success'] ?? false) {
            $count = count($results);
            $this->redirect(
                "/exams/{$examID}/results" . ($classID ? "?class_id={$classID}" : ''),
                "{$count} result(s) saved successfully."
            );
        }

        Session::flash('error', $res['error'] ?? 'Failed to submit results.');
        $this->redirect("/exams/{$examID}/results" . ($classID ? "?class_id={$classID}" : ''));
    }
}