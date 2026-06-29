<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-plus-circle me-2 text-success"></i>Add Fee Type</div>
  <div class="card-body">
    <form method="POST" action="/finance/fee-types">
      <div class="mb-3"><label class="form-label small fw-semibold">Fee Name *</label>
        <input type="text" name="name" class="form-control" required placeholder="e.g. Tuition Fee"></div>
      <div class="mb-3"><label class="form-label small fw-semibold">Amount (KES) *</label>
        <input type="number" step="0.01" name="amount" class="form-control" required placeholder="0.00"></div>
      <div class="mb-3"><label class="form-label small fw-semibold">Frequency</label>
        <select name="frequency" class="form-select">
          <option value="termly">Termly</option>
          <option value="monthly">Monthly</option>
          <option value="annual">Annual</option>
          <option value="once">Once</option>
        </select>
      </div>
      <div class="mb-3 form-check">
        <input type="checkbox" name="is_mandatory" class="form-check-input" id="isMandatory" value="1" checked>
        <label class="form-check-label small" for="isMandatory">Mandatory for all students</label>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Save Fee Type</button>
        <a href="/finance" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
