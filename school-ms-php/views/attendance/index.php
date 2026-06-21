<?php /** @var array $classes @var array $terms */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-calendar-check me-2 text-success"></i>Attendance</h5>
  <a href="/attendance/summary" class="btn btn-outline-info btn-sm"><i class="bi bi-bar-chart me-1"></i>Summary</a>
</div>
<div class="card">
  <div class="card-header py-3">Mark Attendance for a Class</div>
  <div class="card-body">
    <form method="GET" action="/attendance/mark" class="row g-3">
      <div class="col-md-5">
        <label class="form-label small fw-semibold">Class</label>
        <select name="class_id" class="form-select" required>
          <option value="">Select class...</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-4">
        <label class="form-label small fw-semibold">Date</label>
        <input type="date" name="date" class="form-control" value="<?= date('Y-m-d') ?>">
      </div>
      <div class="col-md-3 d-flex align-items-end">
        <button class="btn btn-success w-100"><i class="bi bi-arrow-right me-1"></i>Go</button>
      </div>
    </form>
  </div>
</div>
<?php if (empty($terms)): ?>
<div class="alert alert-warning mt-3">
  <i class="bi bi-exclamation-triangle me-2"></i>No terms configured.
  <a href="/terms/create" class="alert-link">Create a term</a> before marking attendance.
</div>
<?php endif; ?>
