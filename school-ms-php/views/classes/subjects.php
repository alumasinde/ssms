<?php /** @var array $class @var array $assigned @var array $unassigned */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <div>
    <nav aria-label="breadcrumb"><ol class="breadcrumb mb-0 small">
      <li class="breadcrumb-item"><a href="/classes">Classes</a></li>
      <li class="breadcrumb-item active"><?= htmlspecialchars($class['name']) ?> — Subjects</li>
    </ol></nav>
    <h5 class="fw-bold mb-0 mt-1"><i class="bi bi-book me-2 text-primary"></i><?= htmlspecialchars($class['name']) ?> — Subject Assignment</h5>
  </div>
  <a href="/classes" class="btn btn-sm btn-outline-secondary"><i class="bi bi-arrow-left me-1"></i>Back</a>
</div>
<div class="row g-3">
  <div class="col-md-6">
    <div class="card"><div class="card-header py-2 fw-semibold">Assigned Subjects (<?= count($assigned) ?>)</div>
    <div class="card-body p-0">
      <table class="table table-sm mb-0">
        <thead><tr><th>Subject</th><th>Code</th><th class="text-center">Compulsory</th><th></th></tr></thead>
        <tbody>
          <?php if (empty($assigned)): ?>
            <tr><td colspan="4" class="text-center text-muted py-3">No subjects assigned yet.</td></tr>
          <?php else: foreach ($assigned as $s): ?>
          <tr>
            <td class="fw-semibold small"><?= htmlspecialchars($s['subject_name']) ?></td>
            <td><span class="badge bg-primary-subtle text-primary"><?= htmlspecialchars($s['subject_code']) ?></span></td>
            <td class="text-center"><?= $s['is_compulsory'] ? '<span class="badge bg-success-subtle text-success">Yes</span>' : '<span class="badge bg-secondary-subtle text-secondary">Optional</span>' ?></td>
            <td>
              <form method="POST" action="/classes/<?= $class['id'] ?>/subjects/<?= $s['subject_id'] ?>/remove" class="d-inline">
                <button class="btn btn-xs btn-outline-danger" style="font-size:.7rem;padding:.15rem .4rem" onclick="return confirm('Remove?')"><i class="bi bi-x"></i></button>
              </form>
            </td>
          </tr>
          <?php endforeach; endif; ?>
        </tbody>
      </table>
    </div></div>
  </div>
  <div class="col-md-6">
    <div class="card"><div class="card-header py-2 fw-semibold">Add Subject</div>
    <div class="card-body">
      <?php if (empty($unassigned)): ?>
        <div class="text-muted small">All subjects are already assigned to this class.</div>
      <?php else: ?>
      <form method="POST" action="/classes/<?= $class['id'] ?>/subjects">
        <div class="mb-2">
          <label class="form-label small fw-semibold">Subject</label>
          <select name="subject_id" class="form-select" required>
            <option value="">Select subject...</option>
            <?php foreach ($unassigned as $s): ?>
              <option value="<?= $s['subject_id'] ?>"><?= htmlspecialchars($s['subject_name']) ?> (<?= htmlspecialchars($s['subject_code']) ?>)</option>
            <?php endforeach; ?>
          </select>
        </div>
        <div class="mb-3 form-check">
          <input type="checkbox" name="is_compulsory" class="form-check-input" id="comp" checked>
          <label class="form-check-label small" for="comp">Compulsory for all students in this class</label>
        </div>
        <button type="submit" class="btn btn-primary btn-sm w-100"><i class="bi bi-plus-lg me-1"></i>Assign Subject</button>
      </form>
      <?php endif; ?>
    </div></div>
  </div>
</div>
