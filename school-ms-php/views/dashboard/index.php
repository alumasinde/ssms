<?php
/** @var int        $totalStudents  */
/** @var int        $totalTeachers  */
/** @var int        $totalParents   */
/** @var array      $recentNotices  */
/** @var array|null $currentTerm    */
?>

<?php if ($currentTerm): ?>
<div class="alert alert-info py-2 small mb-3 d-flex align-items-center gap-2">
  <i class="bi bi-calendar3 fs-5"></i>
  <span>Current Term: <strong><?= htmlspecialchars($currentTerm['name']) ?></strong>
  &nbsp;·&nbsp; <?= $currentTerm['start_date'] ?> → <?= $currentTerm['end_date'] ?></span>
</div>
<?php else: ?>
<div class="alert alert-warning py-2 small mb-3">
  <i class="bi bi-exclamation-triangle me-2"></i>
  No active term set. <a href="/terms/create" class="alert-link">Create a term</a> to enable attendance and exam features.
</div>
<?php endif; ?>

<div class="row g-3 mb-4">
  <?php if (\Core\Session::can('students.view')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/students" class="text-decoration-none">
      <div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)">
        <div class="stat-num"><?= number_format($totalStudents) ?></div>
        <div class="stat-label"><i class="bi bi-person-badge me-1"></i>Students</div>
      </div>
    </a>
  </div>
  <?php endif; ?>
  <?php if (\Core\Session::can('teachers.view')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/teachers" class="text-decoration-none">
      <div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)">
        <div class="stat-num"><?= number_format($totalTeachers) ?></div>
        <div class="stat-label"><i class="bi bi-person-workspace me-1"></i>Teachers</div>
      </div>
    </a>
  </div>
  <?php endif; ?>
  <?php if (\Core\Session::can('parents.view')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/parents" class="text-decoration-none">
      <div class="stat-card" style="background:linear-gradient(135deg,#7c3aed,#a78bfa)">
        <div class="stat-num"><?= number_format($totalParents) ?></div>
        <div class="stat-label"><i class="bi bi-people me-1"></i>Parents</div>
      </div>
    </a>
  </div>
  <?php endif; ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/terms" class="text-decoration-none">
      <div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)">
        <div class="stat-num"><?= $currentTerm ? '✓' : '—' ?></div>
        <div class="stat-label"><i class="bi bi-calendar3 me-1"></i>Active Term</div>
      </div>
    </a>
  </div>
</div>

<!-- Quick actions -->
<div class="row g-3 mb-4">
  <?php if (\Core\Session::can('attendance.mark')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/attendance" class="btn btn-outline-success w-100 py-3">
      <i class="bi bi-calendar-check d-block fs-4 mb-1"></i>Mark Attendance
    </a>
  </div>
  <?php endif; ?>
  <?php if (\Core\Session::can('exams.view')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/exams" class="btn btn-outline-warning w-100 py-3">
      <i class="bi bi-pencil-square d-block fs-4 mb-1"></i>Exams
    </a>
  </div>
  <?php endif; ?>
  <?php if (\Core\Session::can('finance.view')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/finance" class="btn btn-outline-primary w-100 py-3">
      <i class="bi bi-cash-coin d-block fs-4 mb-1"></i>Finance
    </a>
  </div>
  <?php endif; ?>
  <?php if (\Core\Session::can('reports.view')): ?>
  <div class="col-sm-6 col-lg-3">
    <a href="/reports/class-results" class="btn btn-outline-info w-100 py-3">
      <i class="bi bi-bar-chart-line d-block fs-4 mb-1"></i>Reports
    </a>
  </div>
  <?php endif; ?>
</div>

<!-- Recent notices -->
<?php if (\Core\Session::can('notices.view') && !empty($recentNotices)): ?>
<div class="card">
  <div class="card-header py-3 d-flex justify-content-between align-items-center">
    <span><i class="bi bi-megaphone me-2"></i>Recent Notices</span>
    <a href="/notices" class="btn btn-sm btn-outline-secondary">View All</a>
  </div>
  <div class="list-group list-group-flush">
    <?php foreach ($recentNotices as $n): ?>
    <a href="/notices/<?= $n['id'] ?>" class="list-group-item list-group-item-action">
      <div class="d-flex justify-content-between align-items-start">
        <div>
          <div class="fw-semibold small"><?= htmlspecialchars($n['title']) ?></div>
          <div class="text-muted" style="font-size:.78rem"><?= htmlspecialchars(mb_substr($n['body'] ?? '', 0, 80)) ?>…</div>
        </div>
        <span class="badge bg-secondary-subtle text-secondary ms-2" style="font-size:.65rem;white-space:nowrap">
          <?= htmlspecialchars($n['audience']) ?>
        </span>
      </div>
    </a>
    <?php endforeach; ?>
  </div>
</div>
<?php endif; ?>
