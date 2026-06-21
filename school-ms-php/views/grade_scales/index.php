<?php /** @var array $scales @var array $presets */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-award me-2 text-warning"></i>Grade Scales</h5>
</div>
<div class="row g-3">
  <div class="col-lg-7">
    <div class="card"><div class="card-header py-3">Current Grade Scale</div><div class="card-body p-0">
      <table class="table table-hover mb-0">
        <thead><tr><th>Grade</th><th>Min %</th><th>Max %</th><th>Remark</th></tr></thead>
        <tbody>
          <?php if (empty($scales)): ?>
            <tr><td colspan="4" class="text-center text-muted py-4">No grade scale configured. Seed one from the right.</td></tr>
          <?php else: foreach ($scales as $s): ?>
          <tr>
            <td><span class="badge bg-primary-subtle text-primary fw-bold"><?= htmlspecialchars($s['grade']) ?></span></td>
            <td><?= $s['min_score'] ?>%</td>
            <td><?= $s['max_score'] ?>%</td>
            <td class="text-muted small"><?= htmlspecialchars($s['remark']) ?></td>
          </tr>
          <?php endforeach; endif; ?>
        </tbody>
      </table>
    </div></div>
  </div>
  <div class="col-lg-5">
    <div class="card mb-3"><div class="card-header py-3">Seed Preset</div><div class="card-body">
      <p class="text-muted small">Quickly load a standard Kenyan grading system.</p>
      <form method="POST" action="/grade-scales">
        <div class="mb-3">
          <select name="preset" class="form-select" required>
            <option value="">Choose preset...</option>
            <option value="kcse">KCSE (A to E — 12 grades)</option>
            <option value="cbc">CBC (EE / ME / AE / BE)</option>
            <option value="pct">Simple % (A–E, 5 grades)</option>
          </select>
        </div>
        <button type="submit" class="btn btn-warning w-100"><i class="bi bi-lightning me-1"></i>Load Preset</button>
      </form>
    </div></div>
    <div class="card"><div class="card-header py-3">Add Single Entry</div><div class="card-body">
      <form method="POST" action="/grade-scales">
        <div class="row g-2 mb-2">
          <div class="col-4"><label class="form-label small fw-semibold">Grade</label>
            <input type="text" name="grade" class="form-control form-control-sm" required placeholder="A"></div>
          <div class="col-4"><label class="form-label small fw-semibold">Min %</label>
            <input type="number" name="min_score" class="form-control form-control-sm" required min="0" max="100"></div>
          <div class="col-4"><label class="form-label small fw-semibold">Max %</label>
            <input type="number" name="max_score" class="form-control form-control-sm" required min="0" max="100"></div>
        </div>
        <div class="mb-2"><label class="form-label small fw-semibold">Remark</label>
          <input type="text" name="remark" class="form-control form-control-sm" placeholder="e.g. Excellent"></div>
        <button type="submit" class="btn btn-outline-primary btn-sm w-100"><i class="bi bi-plus-lg me-1"></i>Add</button>
      </form>
    </div></div>
  </div>
</div>
