<?php
/** @var array      $card    */
/** @var array      $exams   */
/** @var array      $student */
/** @var int        $examID  */
$c        = $card ?? [];
$hasCard  = !empty($c['results']);
$studentID = $student['id'] ?? 0;
$studentName = trim(($student['first_name'] ?? '') . ' ' . ($student['last_name'] ?? ''));
?>

<div class="d-flex justify-content-between align-items-center mb-3">
  <div>
    <nav aria-label="breadcrumb">
      <ol class="breadcrumb mb-0 small">
        <li class="breadcrumb-item"><a href="/students">Students</a></li>
        <?php if ($studentID): ?>
          <li class="breadcrumb-item"><a href="/students/<?= $studentID ?>"><?= htmlspecialchars($studentName) ?></a></li>
        <?php endif; ?>
        <li class="breadcrumb-item active">Report Card</li>
      </ol>
    </nav>
    <h5 class="fw-bold mb-0 mt-1"><i class="bi bi-file-earmark-text me-2 text-info"></i>Report Card</h5>
  </div>
  <?php if ($hasCard): ?>
    <button onclick="window.print()" class="btn btn-outline-secondary btn-sm">
      <i class="bi bi-printer me-1"></i>Print
    </button>
  <?php endif; ?>
</div>

<!-- Exam selector -->
<div class="card mb-4">
  <div class="card-body py-3">
    <form method="GET" class="row g-2 align-items-end">
      <div class="col-md-5">
        <label class="form-label small fw-semibold">Select Exam</label>
        <select name="exam_id" class="form-select" onchange="this.form.submit()">
          <option value="">— Choose an exam —</option>
          <?php foreach ($exams as $e): ?>
            <option value="<?= $e['id'] ?>" <?= $examID == $e['id'] ? 'selected' : '' ?>>
              <?= htmlspecialchars($e['name']) ?>
              <span>(<?= ucfirst($e['type'] ?? '') ?>)</span>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
      <?php if (!$examID): ?>
        <div class="col-md-7">
          <div class="alert alert-info py-2 small mb-0">
            <i class="bi bi-arrow-left me-1"></i>Select an exam above to generate the report card.
          </div>
        </div>
      <?php endif; ?>
    </form>
  </div>
</div>

<?php if ($examID && !$hasCard): ?>
  <div class="alert alert-warning">
    <i class="bi bi-exclamation-triangle me-2"></i>
    No results found for <strong><?= htmlspecialchars($studentName) ?></strong> in this exam.
    Results may not have been entered yet.
  </div>
<?php endif; ?>

<?php if ($hasCard): ?>
<!-- Printable report card -->
<div id="printArea">
  <!-- Header -->
  <div class="card mb-3">
    <div class="card-body">
      <div class="row align-items-center">
        <div class="col-md-8">
          <h4 class="fw-bold mb-1"><?= htmlspecialchars($c['student_name'] ?? $studentName) ?></h4>
          <div class="text-muted small">
            Adm No: <strong><?= htmlspecialchars($c['admission_no'] ?? '') ?></strong>
            &nbsp;·&nbsp;
            Class: <strong><?= htmlspecialchars($c['class_name'] ?? '') ?></strong>
          </div>
          <div class="text-muted small mt-1">
            Exam: <strong><?= htmlspecialchars($c['exam_name'] ?? '') ?></strong>
            &nbsp;·&nbsp;
            Term: <strong><?= htmlspecialchars($c['term_name'] ?? '') ?></strong>
          </div>
        </div>
        <div class="col-md-4 text-md-end mt-3 mt-md-0">
          <div class="d-flex flex-column gap-1">
            <div>
              <span class="text-muted small">Position:</span>
              <strong class="fs-5 ms-1"><?= $c['position'] ?? '—' ?></strong>
              <span class="text-muted small">/ <?= $c['class_size'] ?? '—' ?></span>
            </div>
            <div>
              <span class="text-muted small">Average:</span>
              <?php $avg = round($c['average'] ?? 0, 1); ?>
              <span class="badge fs-6 <?= $avg >= 50 ? 'bg-success' : 'bg-danger' ?> ms-1">
                <?= $avg ?>%
              </span>
            </div>
            <div>
              <span class="text-muted small">Attendance:</span>
              <strong class="ms-1"><?= $c['attendance_pct'] ?? 0 ?>%</strong>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Subject results -->
  <div class="card mb-3">
    <div class="card-body p-0">
      <table class="table table-bordered mb-0">
        <thead class="table-light">
          <tr>
            <th>Subject</th>
            <th class="text-center" style="width:80px">Marks</th>
            <th class="text-center" style="width:60px">Max</th>
            <th class="text-center" style="width:70px">%</th>
            <th class="text-center" style="width:60px">Grade</th>
            <th>Remarks</th>
          </tr>
        </thead>
        <tbody>
          <?php foreach ($c['results'] ?? [] as $r): ?>
          <?php $pct = round($r['percentage'] ?? 0, 1); ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($r['subject_name']) ?></td>
            <td class="text-center fw-bold"><?= $r['marks'] ?></td>
            <td class="text-center text-muted"><?= $r['max_marks'] ?></td>
            <td class="text-center">
              <div class="progress" style="height:6px;margin-bottom:3px">
                <div class="progress-bar <?= $pct >= 50 ? 'bg-success' : 'bg-danger' ?>" style="width:<?= min($pct,100) ?>%"></div>
              </div>
              <span style="font-size:.75rem"><?= $pct ?>%</span>
            </td>
            <td class="text-center">
              <span class="badge <?= $pct >= 50 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>">
                <?= htmlspecialchars($r['grade']) ?>
              </span>
            </td>
            <td class="text-muted small"><?= htmlspecialchars($r['remarks'] ?? '') ?></td>
          </tr>
          <?php endforeach; ?>
        </tbody>
        <tfoot class="table-light fw-bold">
          <tr>
            <td>TOTAL</td>
            <td class="text-center"><?= $c['total_marks'] ?? 0 ?></td>
            <td class="text-center"><?= $c['total_max'] ?? 0 ?></td>
            <td class="text-center"><?= round($c['average'] ?? 0, 1) ?>%</td>
            <td class="text-center">
              <span class="badge bg-primary"><?= htmlspecialchars($c['overall_grade'] ?? '—') ?></span>
            </td>
            <td></td>
          </tr>
        </tfoot>
      </table>
    </div>
  </div>

  <!-- Signatures row -->
  <div class="card">
    <div class="card-body">
      <div class="row text-center">
        <div class="col-4">
          <div style="border-top:1px solid #dee2e6;margin-top:2rem;padding-top:.5rem">
            <div class="small text-muted">Class Teacher</div>
          </div>
        </div>
        <div class="col-4">
          <div style="border-top:1px solid #dee2e6;margin-top:2rem;padding-top:.5rem">
            <div class="small text-muted">Principal</div>
          </div>
        </div>
        <div class="col-4">
          <div style="border-top:1px solid #dee2e6;margin-top:2rem;padding-top:.5rem">
            <div class="small text-muted">Parent / Guardian</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
@media print {
  .sidebar, .topbar, .breadcrumb, .card:first-child, button, .btn, nav { display: none !important; }
  .main-wrap { margin-left: 0 !important; }
  .page-content { padding: 0 !important; }
  #printArea .card { box-shadow: none !important; border: 1px solid #dee2e6 !important; }
}
</style>
<?php endif; ?>
