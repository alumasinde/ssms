<?php
namespace Exams;

use Core\Controller;
use Core\Session;

class ExamsController extends Controller
{
    public function index(array $params = []): void
    {
        $this->requirePermission('exams.view');
        $termID  = (int)($_GET['term_id'] ?? 0);
        $classID = (int)($_GET['class_id'] ?? 0);

        $query = array_filter(['term_id' => $termID ?: null, 'class_id' => $classID ?: null]);
        $res   = $this->api->get('/exams', $query);

        $terms   = $this->api->get('/terms');
        $classes = $this->api->get('/classes');

        $this->view('exams/index', [
            'title'   => 'Exams',
            'exams'   => $res['data'] ?? [],
            'terms'   => $terms['data'] ?? [],
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
            'terms'   => $terms['data'] ?? [],
            'classes' => $classes['data'] ?? [],
        ]);
    }

    public function store(array $params = []): void
    {
        $this->requirePermission('exams.create');
        $classID = (int)($_POST['class_id'] ?? 0);
        $res = $this->api->post('/exams', [
            'name'       => trim($_POST['name'] ?? ''),
            'type'       => $_POST['type'] ?? 'endterm',
            'term_id'    => (int)($_POST['term_id'] ?? 0),
            'class_id'   => $classID ?: null,
            'start_date' => $_POST['start_date'] ?? '',
            'end_date'   => $_POST['end_date'] ?? '',
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
            'exam'  => $res['data'] ?? [],
        ]);
    }

    public function results(array $params): void
    {
        $this->requirePermission('exams.view');
        $examID  = (int)$params['id'];
        $classID = (int)($_GET['class_id'] ?? 0);

        $exam    = $this->api->get("/exams/{$examID}");
        $examData = $exam['data'] ?? [];

        // Use class_id from exam itself if not passed as filter
        if (!$classID && !empty($examData['class_id'])) {
            $classID = (int)$examData['class_id'];
        }

        $query   = $classID ? ['class_id' => $classID] : [];
        $res     = $this->api->get("/exams/{$examID}/results", $query);

        // Load class students for the grade entry form
        $students = [];
        if ($classID) {
            $sRes     = $this->api->get("/students/class/{$classID}");
            $students = $sRes['data'] ?? [];
        } else {
            $sRes     = $this->api->get('/students', ['per_page' => 200]);
            $students = $sRes['data'] ?? [];
        }

        $subjects = $this->api->get('/subjects');
        $classes  = $this->api->get('/classes');

        $this->view('exams/results', [
            'title'    => 'Results — ' . ($examData['name'] ?? ''),
            'results'  => $res['data'] ?? [],
            'exam'     => $examData,
            'students' => $students,
            'subjects' => $subjects['data'] ?? [],
            'classes'  => $classes['data'] ?? [],
            'classID'  => $classID,
        ]);
    }

    /**
     * Submit bulk results for one exam (all subjects for a student, or class-wide).
     * The new backend expects: {exam_id, class_id, results:[{student_id, subject_id, marks, max_marks, remarks}]}
     */
    public function submitResults(array $params = []): void
    {
        $this->requirePermission('exams.grade');
        $examID  = (int)($_POST['exam_id'] ?? 0);
        $classID = (int)($_POST['class_id'] ?? 0);

        // Build results array from POST fields
        $results = [];
        foreach ($_POST['student_id'] ?? [] as $idx => $studentID) {
            $subjectID = (int)($_POST['subject_id'][$idx] ?? 0);
            $marks     = (float)($_POST['marks'][$idx] ?? 0);
            $maxMarks  = (float)($_POST['max_marks'][$idx] ?? 100);
            if (!$studentID || !$subjectID) continue;
            $results[] = [
                'student_id' => (int)$studentID,
                'subject_id' => $subjectID,
                'marks'      => $marks,
                'max_marks'  => $maxMarks,
                'remarks'    => trim($_POST['remarks'][$idx] ?? ''),
            ];
        }

        // Fallback: single-entry form (old-style)
        if (empty($results) && isset($_POST['student_id']) && !is_array($_POST['student_id'])) {
            $results[] = [
                'student_id' => (int)($_POST['student_id'] ?? 0),
                'subject_id' => (int)($_POST['subject_id'] ?? 0),
                'marks'      => (float)($_POST['marks'] ?? 0),
                'max_marks'  => (float)($_POST['max_marks'] ?? 100),
                'remarks'    => trim($_POST['remarks'] ?? ''),
            ];
        }

        $res = $this->api->post('/exams/results', [
            'exam_id'  => $examID,
            'class_id' => $classID,
            'results'  => $results,
        ]);

        if ($res['success'] ?? false) {
            $this->redirect("/exams/{$examID}/results" . ($classID ? "?class_id={$classID}" : ''), 'Results submitted.');
        }
        Session::flash('error', $res['error'] ?? 'Failed to submit results.');
        $this->redirect("/exams/{$examID}/results" . ($classID ? "?class_id={$classID}" : ''));
    }
}
