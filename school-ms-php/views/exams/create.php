<?php /** @var array $terms @var array $classes */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-pencil-square me-2 text-warning"></i>Create Exam</div>
  <div class="card-body">
    <form method="POST" action="/exams">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Exam Name *</label>
        <input type="text" name="name" class="form-control" required placeholder="e.g. End of Term 2 Exams">
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Type</label>
        <select name="type" class="form-select">
          <option value="endterm">End Term</option>
          <option value="midterm">Mid Term</option>
          <option value="cat">CAT</option>
          <option value="mock">Mock</option>
          <option value="opener">Opener</option>
        </select>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Term *</label>
        <?php if (!empty($terms)): ?>
          <select name="term_id" class="form-select" required>
            <option value="">Select term...</option>
            <?php foreach ($terms as $t): ?>
              <option value="<?= $t['id'] ?>" <?= ($t['is_current'] ?? false) ? 'selected' : '' ?>>
                <?= htmlspecialchars($t['name']) ?><?= ($t['is_current'] ?? false) ? ' (Current)' : '' ?>
              </option>
            <?php endforeach; ?>
          </select>
        <?php else: ?>
          <input type="number" name="term_id" class="form-control" required placeholder="Term ID">
          <div class="form-text text-warning"><i class="bi bi-exclamation-triangle me-1"></i>No terms found. <a href="/terms/create">Create a term first.</a></div>
        <?php endif; ?>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Class <span class="text-muted">(leave blank for school-wide exam)</span></label>
        <select name="class_id" class="form-select">
          <option value="">All classes (school-wide)</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="row g-2 mb-3">
        <div class="col"><label class="form-label small fw-semibold">Start Date *</label><input type="date" name="start_date" class="form-control" required></div>
        <div class="col"><label class="form-label small fw-semibold">End Date *</label><input type="date" name="end_date" class="form-control" required></div>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Create Exam</button>
        <a href="/exams" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
