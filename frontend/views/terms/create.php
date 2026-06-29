<?php /** @var array $years */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-calendar3 me-2 text-primary"></i>Create Term</div>
  <div class="card-body">
    <form method="POST" action="/terms">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Academic Year *</label>
        <select name="academic_year_id" class="form-select" required>
          <option value="">Select academic year...</option>
          <?php foreach ($years as $y): ?>
            <option value="<?= $y['id'] ?>" <?= ($y['is_current'] ?? false) ? 'selected' : '' ?>>
              <?= htmlspecialchars($y['name']) ?><?= ($y['is_current'] ?? false) ? ' (Current)' : '' ?>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Term Name *</label>
        <input type="text" name="name" class="form-control" required placeholder="e.g. Term 1">
      </div>
      <div class="row g-2 mb-3">
        <div class="col">
          <label class="form-label small fw-semibold">Start Date *</label>
          <input type="date" name="start_date" class="form-control" required>
        </div>
        <div class="col">
          <label class="form-label small fw-semibold">End Date *</label>
          <input type="date" name="end_date" class="form-control" required>
        </div>
      </div>
      <div class="mb-3 form-check">
        <input type="checkbox" name="is_current" class="form-check-input" id="isCurrent">
        <label class="form-check-label small" for="isCurrent">Set as current term</label>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Create Term</button>
        <a href="/terms" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
