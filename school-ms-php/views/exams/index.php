<?php /** 
 * @var array $exams 
 * @var array $terms 
 * @var array $classes 
 * @var int $termID 
 * @var int $classID */ 

 use Core\Session;
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-pencil-square me-2 text-warning"></i>Exams</h5>
  <?php if (Session::can('exams.create')): ?>
    <a href="/exams/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Create Exam</a>
  <?php endif; ?>
</div>

<!-- Filters -->
<div class="card mb-3">
  <div class="card-body py-2">
    <form method="GET" class="row g-2 align-items-end">
      <div class="col-md-4">
        <label class="form-label small fw-semibold mb-1">Filter by Term</label>
        <select name="term_id" class="form-select form-select-sm" onchange="this.form.submit()">
          <option value="">All terms</option>
          <?php foreach ($terms as $t): ?>
            <option value="<?= $t['id'] ?>" <?= $termID == $t['id'] ? 'selected' : '' ?>>
              <?= htmlspecialchars($t['name']) ?><?= ($t['is_current'] ?? false) ? ' ★' : '' ?>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-4">
        <label class="form-label small fw-semibold mb-1">Filter by Class</label>
        <select name="class_id" class="form-select form-select-sm" onchange="this.form.submit()">
          <option value="">All classes</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>" <?= $classID == $c['id'] ? 'selected' : '' ?>><?= htmlspecialchars($c['name']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-4">
        <?php if ($termID || $classID): ?>
          <a href="/exams" class="btn btn-sm btn-outline-secondary">
            <i class="bi bi-x me-1"></i>Clear Filters
          </a>
        <?php endif; ?>
      </div>
    </form>
  </div>
</div>

<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Type</th><th>Class</th><th>Start</th><th>End</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($exams)): ?>
          <tr><td colspan="6" class="text-center text-muted py-4">No exams found.</td></tr>
        <?php else: foreach ($exams as $e): ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($e['name']) ?></td>
            <td><span class="badge bg-warning-subtle text-warning"><?= $e['type'] ?></span></td>
            <td class="text-muted small"><?= $e['class_id'] ? '#'.$e['class_id'] : '<span class="text-muted">School-wide</span>' ?></td>
            <td><?= Session::formatDate($e['start_date']) ?></td>
            <td><?= Session::formatDate($e['end_date']) ?></td>
            <td class="text-end">
              <a href="/exams/<?= $e['id'] ?>/results<?= $e['class_id'] ? '?class_id='.$e['class_id'] : '' ?>" class="btn btn-sm btn-outline-primary">
                <i class="bi bi-list-ol me-1"></i>Results
              </a>
              <a href="/reports/class-results?exam_id=<?= $e['id'] ?>" class="btn btn-sm btn-outline-info">
                <i class="bi bi-bar-chart me-1"></i>Report
              </a>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>
