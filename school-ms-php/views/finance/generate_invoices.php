<?php /** @var array $feeTypes @var array $terms @var array $classes */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-file-earmark-plus me-2 text-primary"></i>Generate Invoices</div>
  <div class="card-body">
    <p class="text-muted small mb-3">Generate fee invoices for all students in a class for a given term and fee type.</p>
    <form method="POST" action="/finance/invoices/generate">
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
        <?php endif; ?>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Fee Type *</label>
        <?php if (!empty($feeTypes)): ?>
          <select name="fee_type_id" class="form-select" required>
            <option value="">Select fee type...</option>
            <?php foreach ($feeTypes as $ft): ?>
              <option value="<?= $ft['id'] ?>"><?= htmlspecialchars($ft['name']) ?> — KES <?= number_format($ft['amount'],2) ?></option>
            <?php endforeach; ?>
          </select>
        <?php else: ?>
          <input type="number" name="fee_type_id" class="form-control" required>
          <div class="form-text text-warning">No fee types found. <a href="/finance/fee-types/create">Add one first.</a></div>
        <?php endif; ?>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Class <span class="text-muted">(leave blank = all classes)</span></label>
        <select name="class_id" class="form-select">
          <option value="">All classes</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Due Date *</label>
        <input type="date" name="due_date" class="form-control" required>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Generate</button>
        <a href="/finance" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
