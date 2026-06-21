<?php
/** @var array $parent */
/** @var array $students */
$pName = htmlspecialchars($parent['name'] ?? '');
?>
<div class="row g-3">
  <div class="col-lg-4">
    <div class="card text-center p-3">
      <div class="rounded-circle bg-info mx-auto d-flex align-items-center justify-content-center text-white mb-3"
           style="width:72px;height:72px;font-size:2rem;font-weight:700">
        <?= strtoupper(substr($parent['name'] ?? 'P', 0, 1)) ?>
      </div>
      <h5 class="fw-bold mb-0"><?= $pName ?></h5>
      <div class="text-muted small"><?= htmlspecialchars($parent['email'] ?? '') ?></div>
    </div>
    <div class="card mt-3">
      <div class="card-header small fw-semibold">Details</div>
      <div class="list-group list-group-flush small">
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Phone</span><span><?= htmlspecialchars($parent['phone'] ?? '—') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Occupation</span><span><?= htmlspecialchars($parent['occupation'] ?? '—') ?></span>
        </div>
        <div class="list-group-item">
          <span class="text-muted d-block">Address</span>
          <span><?= htmlspecialchars($parent['address'] ?? '—') ?></span>
        </div>
      </div>
    </div>
    <a href="/parents" class="btn btn-sm btn-link mt-2"><i class="bi bi-arrow-left me-1"></i>Back to List</a>
  </div>

  <div class="col-lg-8">
    <div class="card">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <span><i class="bi bi-person-badge me-1"></i>Linked Students</span>
        <?php if (\Core\Session::can('parents.create')): ?>
        <button class="btn btn-sm btn-primary" data-bs-toggle="modal" data-bs-target="#linkModal">
          <i class="bi bi-link me-1"></i>Link Student
        </button>
        <?php endif; ?>
      </div>
      <div class="card-body p-0">
        <?php if (empty($students)): ?>
          <div class="p-3 text-muted small">No students linked yet.</div>
        <?php else: ?>
          <table class="table table-hover mb-0">
            <thead><tr><th>Adm No</th><th>Name</th><th>Class</th><th>Relationship</th><th></th></tr></thead>
            <tbody>
              <?php foreach ($students as $s): ?>
              <?php $sName = trim(($s['first_name'] ?? '') . ' ' . ($s['last_name'] ?? '')); ?>
              <tr>
                <td><span class="badge bg-light text-dark border"><?= htmlspecialchars($s['admission_no'] ?? '') ?></span></td>
                <td class="fw-semibold"><?= htmlspecialchars($sName) ?></td>
                <td><?= htmlspecialchars($s['class_name'] ?? ('Class #' . ($s['class_id'] ?? '—'))) ?></td>
                <td><span class="badge bg-secondary-subtle text-secondary"><?= htmlspecialchars($s['relationship'] ?? 'parent') ?></span></td>
                <td><a href="/students/<?= $s['id'] ?>" class="btn btn-sm btn-outline-primary"><i class="bi bi-eye"></i></a></td>
              </tr>
              <?php endforeach; ?>
            </tbody>
          </table>
        <?php endif; ?>
      </div>
    </div>
  </div>
</div>

<?php if (\Core\Session::can('parents.create')): ?>
<div class="modal fade" id="linkModal" tabindex="-1">
  <div class="modal-dialog modal-sm">
    <div class="modal-content">
      <div class="modal-header">
        <h6 class="modal-title">Link Student</h6>
        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
      </div>
      <div class="modal-body">
        <form method="POST" action="/parents/link-student">
          <input type="hidden" name="parent_id"  value="<?= $parent['id'] ?? '' ?>">
          <input type="hidden" name="redirect"   value="/parents/<?= $parent['id'] ?? '' ?>">
          <div class="mb-3">
            <label class="form-label small fw-semibold">Student ID</label>
            <input type="number" name="student_id" class="form-control" required>
          </div>
          <div class="mb-3">
            <label class="form-label small fw-semibold">Relationship</label>
            <select name="relationship" class="form-select">
              <option value="father">Father</option>
              <option value="mother">Mother</option>
              <option value="guardian">Guardian</option>
              <option value="other">Other</option>
            </select>
          </div>
          <button type="submit" class="btn btn-primary w-100"><i class="bi bi-link me-1"></i>Link</button>
        </form>
      </div>
    </div>
  </div>
</div>
<?php endif; ?>
