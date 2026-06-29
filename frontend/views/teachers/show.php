<?php
/** @var array $teacher */
/** @var array $subjects */
/** @var array $allSubjects */
/** @var array $allClasses */
$canEdit = \Core\Session::can('teachers.edit');
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <nav aria-label="breadcrumb">
    <ol class="breadcrumb mb-0 small">
      <li class="breadcrumb-item"><a href="/teachers">Teachers</a></li>
      <li class="breadcrumb-item active"><?= htmlspecialchars($teacher['name'] ?? '') ?></li>
    </ol>
  </nav>
  <?php if ($canEdit): ?>
    <a href="/teachers/<?= $teacher['id'] ?>/edit" class="btn btn-sm btn-outline-primary">
      <i class="bi bi-pencil me-1"></i>Edit Profile
    </a>
  <?php endif; ?>
</div>

<div class="row g-3">
  <!-- Profile card -->
  <div class="col-lg-4">
    <div class="card text-center p-3">
      <div class="rounded-circle bg-primary mx-auto d-flex align-items-center justify-content-center text-white mb-3"
           style="width:80px;height:80px;font-size:2.2rem;font-weight:700">
        <?= strtoupper(substr($teacher['name'] ?? 'T', 0, 1)) ?>
      </div>
      <h5 class="fw-bold mb-0"><?= htmlspecialchars($teacher['name'] ?? '') ?></h5>
      <div class="text-muted small"><?= htmlspecialchars($teacher['email'] ?? '') ?></div>
      <div class="mt-2">
        <span class="badge bg-primary-subtle text-primary"><?= ucfirst(str_replace('_',' ',$teacher['employment_type'] ?? '')) ?></span>
      </div>
    </div>

    <div class="card mt-3">
      <div class="card-header small fw-semibold">HR Details</div>
      <div class="list-group list-group-flush small">
        <?php foreach ([
          'Employee No'    => $teacher['employee_no'] ?? '—',
          'TSC No'         => $teacher['tsc_no'] ?? '—',
          'Phone'          => $teacher['phone'] ?? '—',
          'Gender'         => ucfirst($teacher['gender'] ?? '—'),
          'Qualification'  => $teacher['qualification'] ?? '—',
          'Specialization' => $teacher['specialization'] ?? '—',
          'Hire Date'      => !empty($teacher['hire_date']) ? date('d M Y', strtotime($teacher['hire_date'])) : '—',
        ] as $label => $value): ?>
          <div class="list-group-item d-flex justify-content-between py-2">
            <span class="text-muted"><?= $label ?></span>
            <span class="text-end" style="max-width:60%"><?= htmlspecialchars($value) ?></span>
          </div>
        <?php endforeach; ?>
      </div>
    </div>
  </div>

  <!-- Assigned subjects -->
  <div class="col-lg-8">
    <div class="card">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <span><i class="bi bi-book me-1 text-primary"></i>Assigned Subjects
          <span class="badge bg-secondary-subtle text-secondary ms-1"><?= count($subjects) ?></span>
        </span>
        <?php if ($canEdit): ?>
          <button class="btn btn-sm btn-primary" data-bs-toggle="modal" data-bs-target="#assignModal">
            <i class="bi bi-plus-lg me-1"></i>Assign Subject
          </button>
        <?php endif; ?>

      </div>

      <div class="card-body p-0">
        <?php if (empty($subjects)): ?>
          <div class="p-4 text-center text-muted">
            <i class="bi bi-book fs-1 d-block mb-2 opacity-25"></i>
            No subjects assigned yet.
            <?php if ($canEdit): ?>
              <div class="mt-2 small">Click <strong>Assign Subject</strong> to get started.</div>
            <?php endif; ?>
          </div>
        <?php else: ?>
          <table class="table table-hover mb-0">
            <thead>
              <tr>
                <th>Subject</th>
                <th>Code</th>
                <th>Class</th>
                <?php if ($canEdit): ?><th></th><?php endif; ?>
              </tr>
            </thead>
            <tbody>
              <?php foreach ($subjects as $s): ?>
              <tr>
                <td class="fw-semibold"><?= htmlspecialchars($s['subject_name'] ?? '') ?></td>
                <td><span class="badge bg-primary-subtle text-primary"><?= htmlspecialchars($s['subject_code'] ?? '') ?></span></td>
                <td><?= htmlspecialchars($s['class_name'] ?? ('Class #' . ($s['class_id'] ?? '—'))) ?></td>
                <?php if ($canEdit): ?>
                <td class="text-end">
                  <form method="POST" action="/teachers/remove-subject"
                        onsubmit="return confirm('Remove <?= htmlspecialchars($s['subject_name'] ?? '') ?> from this teacher?')">
                    <input type="hidden" name="teacher_id" value="<?= $teacher['id'] ?>">
                    <input type="hidden" name="subject_id" value="<?= $s['subject_id'] ?>">
                    <input type="hidden" name="class_id"   value="<?= $s['class_id'] ?>">
                    <button class="btn btn-sm btn-outline-danger">
                      <i class="bi bi-x-lg"></i>
                    </button>
                  </form>
                </td>
                <?php endif; ?>
              </tr>
              <?php endforeach; ?>
            </tbody>
          </table>
        <?php endif; ?>
      </div>
    </div>
  </div>
</div>

<?php if ($canEdit): ?>
<!-- Assign Subject Modal -->
<div class="modal fade" id="assignModal" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h6 class="modal-title"><i class="bi bi-book me-2"></i>Assign Subject to <?= htmlspecialchars($teacher['name'] ?? '') ?></h6>
        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
      </div>
      <div class="modal-body">
        <form method="POST" action="/teachers/assign-subject">
          <input type="hidden" name="teacher_id" value="<?= $teacher['id'] ?>">

          <div class="mb-3">
            <label class="form-label small fw-semibold">Subject *</label>
            <select name="subject_id" class="form-select" required>
              <option value="">Select subject...</option>
              <?php foreach ($allSubjects as $sub): ?>
                <option value="<?= $sub['id'] ?>">
                  <?= htmlspecialchars($sub['name']) ?> (<?= htmlspecialchars($sub['code']) ?>)
                </option>
              <?php endforeach; ?>
            </select>
          </div>

          <div class="mb-3">
            <label class="form-label small fw-semibold">Class *</label>
            <select name="class_id" class="form-select" required>
              <option value="">Select class...</option>
              <?php foreach ($allClasses as $cls): ?>
                <option value="<?= $cls['id'] ?>">
                  <?= htmlspecialchars($cls['name']) ?> — <?= htmlspecialchars($cls['level'] ?? '') ?>
                </option>
              <?php endforeach; ?>
            </select>
          </div>

          <div class="alert alert-info py-2 small">
            <i class="bi bi-info-circle me-1"></i>
            The teacher will be responsible for teaching this subject to the selected class.
          </div>

          <button type="submit" class="btn btn-primary w-100">
            <i class="bi bi-check-lg me-1"></i>Assign Subject
          </button>
        </form>
      </div>
    </div>
  </div>
</div>
<?php endif; ?>
