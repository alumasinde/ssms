<?php /** 
 * @var array $exam 
 * */ 
use Core\Session;
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-pencil-square me-2 text-warning"></i><?= htmlspecialchars($exam['name'] ?? '') ?></h5>
  <a href="/exams" class="btn btn-sm btn-outline-secondary"><i class="bi bi-arrow-left me-1"></i>Back</a>
</div>
<div class="row g-3">
  <div class="col-md-4">
    <div class="card">
      <div class="card-header">Exam Details</div>
      <div class="list-group list-group-flush small">
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Type</span>
          <span class="badge bg-warning-subtle text-warning"><?= ucfirst($exam['type'] ?? '') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Term ID</span><span><?= $exam['term_id'] ?? '—' ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Class</span>
          <span><?= $exam['class_id'] ? 'Class #'.$exam['class_id'] : '<em class="text-muted">School-wide</em>' ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Start</span><span><?= Session::formatDate($exam['start_date']) ?? '—' ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">End</span><span><?= Session::formatDate($exam['end_date']) ?? '—' ?></span>
        </div>
      </div>
    </div>
  </div>
  <div class="col-md-8">
    <?php if (Session::can('exams.view')): ?>
    <div class="card">
      <div class="card-header d-flex justify-content-between align-items-center">
        <span>Results</span>
        <a href="/exams/<?= $exam['id'] ?>/results<?= $exam['class_id'] ? '?class_id='.$exam['class_id'] : '' ?>" class="btn btn-sm btn-outline-primary">
          <i class="bi bi-list-ol me-1"></i>View Results
        </a>
      </div>
      <div class="card-body text-muted small">Click "View Results" to see or manage exam results.</div>
    </div>
    <?php endif; ?>
  </div>
</div>
