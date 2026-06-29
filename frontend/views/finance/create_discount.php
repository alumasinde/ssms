<?php /** @var int $studentID @var array $feeTypes @var array $terms */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-tag me-2 text-warning"></i>Apply Fee Discount</div>
  <div class="card-body">
    <form method="POST" action="/finance/discounts">
      <input type="hidden" name="student_id" value="<?= $studentID ?>">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Discount Label *</label>
        <input type="text" name="label" class="form-control" required placeholder="e.g. Bursary, Scholarship, Staff Child">
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Fee Type <span class="text-muted">(leave blank = applies to all)</span></label>
        <select name="fee_type_id" class="form-select">
          <option value="">All fee types</option>
          <?php foreach ($feeTypes as $ft): ?>
            <option value="<?= $ft['id'] ?>"><?= htmlspecialchars($ft['name']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Term <span class="text-muted">(leave blank = recurring)</span></label>
        <select name="term_id" class="form-select">
          <option value="">All terms (recurring)</option>
          <?php foreach ($terms as $t): ?>
            <option value="<?= $t['id'] ?>"><?= htmlspecialchars($t['name']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="row g-2 mb-3">
        <div class="col">
          <label class="form-label small fw-semibold">Discount % <span class="text-muted">(or)</span></label>
          <div class="input-group">
            <input type="number" name="discount_pct" class="form-control" min="0" max="100" step="0.5" placeholder="0">
            <span class="input-group-text">%</span>
          </div>
        </div>
        <div class="col">
          <label class="form-label small fw-semibold">Fixed Amount KES</label>
          <div class="input-group">
            <span class="input-group-text">KES</span>
            <input type="number" name="discount_amt" class="form-control" min="0" step="0.01" placeholder="0.00">
          </div>
        </div>
      </div>
      <div class="alert alert-info py-2 small"><i class="bi bi-info-circle me-1"></i>Enter either a percentage or a fixed amount, not both.</div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-warning"><i class="bi bi-check-lg me-1"></i>Apply Discount</button>
        <a href="/finance/statement/<?= $studentID ?>" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
