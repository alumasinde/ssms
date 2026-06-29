<?php /** @var array $classes */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-arrow-up-circle me-2 text-success"></i>Student Class Promotion</div>
  <div class="card-body">
    <div class="alert alert-warning py-2 small mb-3">
      <i class="bi bi-exclamation-triangle me-1"></i><strong>This action moves all active students from one class to another.</strong>
      Run this at the end of each academic year. It is logged in the audit trail.
    </div>
    <form method="POST" action="/reports/promote"
          onsubmit="return confirm('Are you sure? This will move ALL active students from the selected class.')">
      <div class="mb-3">
        <label class="form-label small fw-semibold">From Class *</label>
        <select name="from_class_id" class="form-select" required>
          <option value="">Select current class...</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">To Class *</label>
        <select name="to_class_id" class="form-select" required>
          <option value="">Select target class...</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-success"><i class="bi bi-arrow-up-circle me-1"></i>Promote Students</button>
        <a href="/students" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
