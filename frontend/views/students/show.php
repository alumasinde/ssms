<?php
/** @var array $student */
/** @var array $parents */
/** @var array $allParents */
$fullName = trim(($student['first_name'] ?? '') . ' ' . ($student['middle_name'] ?? '') . ' ' . ($student['last_name'] ?? ''));
?>
<div class="row g-3">
  <div class="col-lg-4">
    <div class="card text-center p-3">
      <?php if (!empty($student['photo_url'])): ?>
  <img src="<?= htmlspecialchars($student['photo_url']) ?>"
       alt="<?= htmlspecialchars($student['first_name']) ?>"
       class="rounded-circle mx-auto d-block mb-3"
       style="width:72px;height:72px;object-fit:cover;">
<?php else: ?>
  <div class="rounded-circle bg-primary mx-auto d-flex align-items-center justify-content-center text-white mb-3"
       style="width:72px;height:72px;font-size:2rem;font-weight:700">
    <?= strtoupper(substr($student['first_name'] ?? 'S', 0, 1)) ?>
  </div>
<?php endif; ?>
      <h5 class="fw-bold mb-0"><?= htmlspecialchars($fullName) ?></h5>
      <div class="text-muted small"><?= htmlspecialchars($student['admission_no'] ?? '') ?></div>
    </div>

    <div class="card mt-3">
      <div class="card-header small fw-semibold">Details</div>
      <div class="list-group list-group-flush small">
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Gender</span>
          <span><?= ucfirst($student['gender'] ?? '—') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">DOB</span>
          <span><?= !empty($student['dob']) ? date('d M Y', strtotime($student['dob'])) : '—' ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Class</span>
          <span><?= htmlspecialchars($student['class_name'] ?? ('Class #' . ($student['class_id'] ?? '—'))) ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Nationality</span>
          <span><?= htmlspecialchars($student['nationality'] ?? 'Kenyan') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Blood Group</span>
          <span><?= htmlspecialchars($student['blood_group'] ?? '—') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Status</span>
          <span class="badge <?= ($student['is_active'] ?? false) ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>">
            <?= ($student['is_active'] ?? false) ? 'Active' : 'Inactive' ?>
          </span>
        </div>
        <?php if (!empty($student['medical_notes'])): ?>
        <div class="list-group-item">
          <span class="text-muted d-block">Medical Notes</span>
          <span class="text-danger small"><?= htmlspecialchars($student['medical_notes']) ?></span>
        </div>
        <?php endif; ?>
      </div>
    </div>

    <div class="d-flex gap-2 mt-3 flex-wrap">
      <?php if (\Core\Session::can('students.edit')): ?>
        <a href="/students/<?= $student['id'] ?>/edit" class="btn btn-outline-secondary btn-sm">
          <i class="bi bi-pencil me-1"></i>Edit
        </a>
      <?php endif; ?>
      <?php if (\Core\Session::can('reports.view')): ?>
        <a href="/reports/report-card/<?= $student['id'] ?>" class="btn btn-outline-info btn-sm">
          <i class="bi bi-file-earmark-text me-1"></i>Report Card
        </a>
      <?php endif; ?>
      <?php if (\Core\Session::can('finance.view')): ?>
        <a href="/finance/statement/<?= $student['id'] ?>" class="btn btn-outline-success btn-sm">
          <i class="bi bi-cash me-1"></i>Fee Statement
        </a>
      <?php endif; ?>
    </div>
    <a href="/students" class="btn btn-sm btn-link mt-2"><i class="bi bi-arrow-left me-1"></i>Back to List</a>
  </div>

  <div class="col-lg-8">
    <div class="card">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <span><i class="bi bi-people me-1"></i>Parents / Guardians</span>
        <div class="d-flex gap-2">
          <?php if (\Core\Session::can('parents.create')): ?>
            <a href="/parents/create" class="btn btn-sm btn-outline-secondary">
              <i class="bi bi-person-plus me-1"></i>New Parent
            </a>
            <button class="btn btn-sm btn-primary" data-bs-toggle="modal" data-bs-target="#linkParentModal">
              <i class="bi bi-link me-1"></i>Link Parent
            </button>
          <?php endif; ?>
        </div>
      </div>
      <div class="list-group list-group-flush">
        <?php if (empty($parents)): ?>
          <div class="list-group-item text-muted small py-3">No parents linked yet.</div>
        <?php else: foreach ($parents as $p): ?>
          <div class="list-group-item d-flex justify-content-between align-items-center">
            <div>
              <div class="fw-semibold small"><?= htmlspecialchars($p['name'] ?? '') ?></div>
              <div class="text-muted" style="font-size:.75rem">
                <?= htmlspecialchars($p['email'] ?? '') ?>
                <?php if (!empty($p['phone'])): ?> · <?= htmlspecialchars($p['phone']) ?><?php endif; ?>
              </div>
            </div>
            <a href="/parents/<?= $p['id'] ?? '' ?>" class="btn btn-sm btn-outline-primary"><i class="bi bi-eye"></i></a>
          </div>
        <?php endforeach; endif; ?>
      </div>
    </div>
  </div>
</div>

<?php if (\Core\Session::can('parents.create')): ?>
<div class="modal fade" id="linkParentModal" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h6 class="modal-title"><i class="bi bi-link me-2"></i>Link Parent to Student</h6>
        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
      </div>
      <div class="modal-body">
        <form method="POST" action="/parents/link-student">
          <input type="hidden" name="student_id" value="<?= $student['id'] ?>">
          <input type="hidden" name="redirect"   value="/students/<?= $student['id'] ?>">
          <div class="mb-3">
            <label class="form-label small fw-semibold">Select Parent</label>
            <select name="parent_id" class="form-select" required>
              <option value="">Choose parent...</option>
              <?php foreach ($allParents as $p): ?>
                <option value="<?= $p['id'] ?? '' ?>">
                  <?= htmlspecialchars($p['name'] ?? '') ?>
                  <?php if (!empty($p['phone'])): ?> — <?= htmlspecialchars($p['phone']) ?><?php endif; ?>
                </option>
              <?php endforeach; ?>
            </select>
            <div class="form-text">Don't see the parent? <a href="/parents/create" target="_blank">Create one first</a>.</div>
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
          <button type="submit" class="btn btn-primary w-100">
            <i class="bi bi-link me-1"></i>Link Parent
          </button>
        </form>
      </div>
    </div>
  </div>
</div>
<?php endif; ?>
