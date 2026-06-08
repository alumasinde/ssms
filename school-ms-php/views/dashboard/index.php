<?php
/** @var int $totalStudents */
/** @var int $totalTeachers */
/** @var array $recentNotices */
?>
<div class="row g-3 mb-4">
  <div class="col-sm-6 col-xl-3">
    <div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)">
      <div class="stat-num"><?= number_format($totalStudents) ?></div>
      <div class="stat-label"><i class="bi bi-person-badge me-1"></i>Total Students</div>
    </div>
  </div>
  <div class="col-sm-6 col-xl-3">
    <div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)">
      <div class="stat-num"><?= number_format($totalTeachers) ?></div>
      <div class="stat-label"><i class="bi bi-person-workspace me-1"></i>Teachers</div>
    </div>
  </div>
  <div class="col-sm-6 col-xl-3">
    <div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)">
      <div class="stat-num"><?= count($recentNotices) ?></div>
      <div class="stat-label"><i class="bi bi-megaphone me-1"></i>Active Notices</div>
    </div>
  </div>
  <div class="col-sm-6 col-xl-3">
    <div class="stat-card" style="background:linear-gradient(135deg,#7c3aed,#a78bfa)">
      <div class="stat-num">—</div>
      <div class="stat-label"><i class="bi bi-calendar-check me-1"></i>Today's Attendance</div>
    </div>
  </div>
</div>

<div class="row g-3">
  <div class="col-lg-7">
    <div class="card">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <span><i class="bi bi-lightning-charge me-1 text-primary"></i>Quick Actions</span>
      </div>
      <div class="card-body">
        <div class="row g-2">
          <?php $actions = [
            ['/students/create','bi-person-plus','Enrol Student','primary'],
            ['/attendance/mark','bi-calendar-check','Mark Attendance','success'],
            ['/exams/create','bi-pencil-square','Create Exam','warning'],
            ['/notices/create','bi-megaphone','Post Notice','info'],
            ['/reports/class-results','bi-bar-chart-line','Class Results','secondary'],
            ['/finance','bi-cash-coin','Finance','danger'],
          ]; ?>
          <?php foreach ($actions as [$url, $icon, $label, $color]): ?>
          <div class="col-6 col-md-4">
            <a href="<?= $url ?>" class="btn btn-<?= $color ?>-subtle text-<?= $color ?> w-100 py-2 border-0 rounded-3 d-flex flex-column align-items-center gap-1 text-decoration-none" style="font-size:.82rem;font-weight:600">
              <i class="bi <?= $icon ?> fs-4"></i><?= $label ?>
            </a>
          </div>
          <?php endforeach; ?>
        </div>
      </div>
    </div>
  </div>

  <div class="col-lg-5">
    <div class="card h-100">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <span><i class="bi bi-megaphone me-1 text-warning"></i>Recent Notices</span>
        <a href="/notices" class="btn btn-sm btn-outline-secondary">View All</a>
      </div>
      <div class="list-group list-group-flush">
        <?php if (empty($recentNotices)): ?>
          <div class="list-group-item text-muted small py-3">No notices yet.</div>
        <?php else: ?>
          <?php foreach ($recentNotices as $n): ?>
          <a href="/notices/<?= $n['id'] ?? '' ?>" class="list-group-item list-group-item-action px-3 py-2">
            <div class="fw-semibold small"><?= htmlspecialchars($n['title'] ?? 'Untitled') ?></div>
            <div class="text-muted" style="font-size:.73rem">
              <span class="badge bg-secondary-subtle text-secondary"><?= $n['audience'] ?? 'All' ?></span>
              · <?= date('d M Y', strtotime($n['published_at'] ?? '')) ?>
            </div>
          </a>
          <?php endforeach; ?>
        <?php endif; ?>
      </div>
    </div>
  </div>
</div>
