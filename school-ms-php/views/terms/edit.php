<?php /** @var array $term */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-pencil me-2 text-primary"></i>Edit Term</div>
  <div class="card-body">
    <form method="POST" action="/terms/<?= $term['id'] ?>/update">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Term Name *</label>
        <input type="text" name="name" class="form-control" required
               value="<?= htmlspecialchars($term['name'] ?? '') ?>">
      </div>
      <div class="row g-2 mb-3">
        <div class="col">
          <label class="form-label small fw-semibold">Start Date *</label>
          <input type="date" name="start_date" class="form-control" required
                 value="<?= htmlspecialchars($term['start_date'] ?? '') ?>">
        </div>
        <div class="col">
          <label class="form-label small fw-semibold">End Date *</label>
          <input type="date" name="end_date" class="form-control" required
                 value="<?= htmlspecialchars($term['end_date'] ?? '') ?>">
        </div>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Update Term</button>
        <a href="/terms" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
