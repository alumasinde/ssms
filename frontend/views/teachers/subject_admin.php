<?php
/** @var array $matrix   - [teacher_id => ['teacher' => [...], 'subjects' => [...]]] */
/** @var array $subjects */
/** @var array $classes  */
?>

<div class="d-flex justify-content-between align-items-center mb-3">
  <div>
    <h5 class="fw-bold mb-0"><i class="bi bi-grid-3x3-gap me-2 text-primary"></i>Subject Assignment</h5>
    <div class="text-muted small mt-1">Manage which teachers teach which subjects in which classes</div>
  </div>
  <div class="d-flex gap-2">
    <button class="btn btn-sm btn-outline-secondary" id="toggleMatrixBtn">
      <i class="bi bi-table me-1"></i>Matrix View
    </button>
    <a href="/teachers" class="btn btn-sm btn-outline-secondary">
      <i class="bi bi-arrow-left me-1"></i>Back to Teachers
    </a>
  </div>
</div>

<!-- Quick Assign Card -->
<div class="card mb-4" id="quickAssignCard">
  <div class="card-header py-3">
    <i class="bi bi-plus-circle me-1 text-primary"></i>Quick Assign
  </div>
  <div class="card-body">
    <form method="POST" action="/teachers/assign-subject" class="row g-3 align-items-end">
      <div class="col-md-4">
        <label class="form-label small fw-semibold">Teacher *</label>
        <select name="teacher_id" class="form-select" required id="qaTeacher">
          <option value="">Select teacher...</option>
          <?php foreach ($matrix as $tid => $row): ?>
            <option value="<?= $tid ?>"><?= htmlspecialchars($row['teacher']['name']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-3">
        <label class="form-label small fw-semibold">Subject *</label>
        <select name="subject_id" class="form-select" required>
          <option value="">Select subject...</option>
          <?php foreach ($subjects as $s): ?>
            <option value="<?= $s['id'] ?>"><?= htmlspecialchars($s['name']) ?> (<?= htmlspecialchars($s['code']) ?>)</option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-3">
        <label class="form-label small fw-semibold">Class *</label>
        <select name="class_id" class="form-select" required>
          <option value="">Select class...</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level'] ?? '') ?></option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-2">
        <button type="submit" class="btn btn-primary w-100">
          <i class="bi bi-check-lg me-1"></i>Assign
        </button>
      </div>
    </form>
  </div>
</div>

<!-- List View (default) -->
<div id="listView">
  <?php if (empty($matrix)): ?>
    <div class="card"><div class="card-body text-muted text-center py-5">No teachers found.</div></div>
  <?php else: ?>
    <?php foreach ($matrix as $tid => $row):
      $teacher  = $row['teacher'];
      $assigned = $row['subjects'];
    ?>
    <div class="card mb-3">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <div class="d-flex align-items-center gap-2">
          <div class="rounded-circle bg-primary text-white d-flex align-items-center justify-content-center fw-bold"
               style="width:36px;height:36px;font-size:.9rem;flex-shrink:0">
            <?= strtoupper(substr($teacher['name'] ?? 'T', 0, 1)) ?>
          </div>
          <div>
            <div class="fw-semibold"><?= htmlspecialchars($teacher['name'] ?? '') ?></div>
            <div class="text-muted" style="font-size:.75rem">
              <?= htmlspecialchars($teacher['employee_no'] ?? '') ?>
              <?php if (!empty($teacher['specialization'])): ?>
                · <?= htmlspecialchars($teacher['specialization']) ?>
              <?php endif; ?>
            </div>
          </div>
        </div>
        <div class="d-flex align-items-center gap-2">
          <span class="badge bg-<?= count($assigned) > 0 ? 'success' : 'secondary' ?>-subtle text-<?= count($assigned) > 0 ? 'success' : 'secondary' ?>">
            <?= count($assigned) ?> subject<?= count($assigned) !== 1 ? 's' : '' ?>
          </span>
          <a href="/teachers/<?= $tid ?>" class="btn btn-sm btn-outline-secondary">
            <i class="bi bi-eye me-1"></i>Profile
          </a>
        </div>
      </div>

      <div class="card-body p-0">
        <?php if (empty($assigned)): ?>
          <div class="p-3 text-muted small fst-italic">No subjects assigned — use Quick Assign above.</div>
        <?php else: ?>
          <table class="table table-sm table-hover mb-0">
            <thead>
              <tr>
                <th style="width:35%">Subject</th>
                <th style="width:15%">Code</th>
                <th style="width:35%">Class</th>
                <th style="width:15%"></th>
              </tr>
            </thead>
            <tbody>
              <?php foreach ($assigned as $a): ?>
              <tr>
                <td class="fw-semibold small"><?= htmlspecialchars($a['subject_name'] ?? '') ?></td>
                <td><span class="badge bg-primary-subtle text-primary"><?= htmlspecialchars($a['subject_code'] ?? '') ?></span></td>
                <td class="small"><?= htmlspecialchars($a['class_name'] ?? '') ?></td>
                <td class="text-end">
                  <form method="POST" action="/teachers/remove-subject"
                        onsubmit="return confirm('Remove this assignment?')" class="d-inline">
                    <input type="hidden" name="teacher_id" value="<?= $tid ?>">
                    <input type="hidden" name="subject_id" value="<?= $a['subject_id'] ?>">
                    <input type="hidden" name="class_id"   value="<?= $a['class_id'] ?>">
                    <button class="btn btn-xs btn-outline-danger py-0 px-1" style="font-size:.75rem">
                      <i class="bi bi-x-lg"></i>
                    </button>
                  </form>
                </td>
              </tr>
              <?php endforeach; ?>
            </tbody>
          </table>
        <?php endif; ?>
      </div>
    </div>
    <?php endforeach; ?>
  <?php endif; ?>
</div>

<!-- Matrix View (toggled) -->
<div id="matrixView" style="display:none">
  <?php
  // Build axis: unique subjects across all assignments
  $subjectAxis = [];
  foreach ($matrix as $row) {
    foreach ($row['subjects'] as $s) {
      $key = $s['subject_id'] . '_' . $s['class_id'];
      $subjectAxis[$key] = [
        'subject_id'   => $s['subject_id'],
        'class_id'     => $s['class_id'],
        'subject_name' => $s['subject_name'],
        'subject_code' => $s['subject_code'],
        'class_name'   => $s['class_name'],
      ];
    }
  }
  // If no assignments yet, seed axis from all subjects × all classes
  if (empty($subjectAxis)) {
    foreach ($subjects as $s) {
      foreach ($classes as $c) {
        $key = $s['id'] . '_' . $c['id'];
        $subjectAxis[$key] = [
          'subject_id'   => $s['id'],
          'class_id'     => $c['id'],
          'subject_name' => $s['name'],
          'subject_code' => $s['code'],
          'class_name'   => $c['name'],
        ];
      }
    }
  }
  // Build lookup: [teacher_id][subject_id_class_id] = true
  $lookup = [];
  foreach ($matrix as $tid => $row) {
    foreach ($row['subjects'] as $s) {
      $lookup[$tid][$s['subject_id'] . '_' . $s['class_id']] = true;
    }
  }
  ?>
  <div class="card">
    <div class="card-header py-3 small fw-semibold">
      <i class="bi bi-table me-1"></i>Assignment Matrix
      <span class="text-muted fw-normal ms-2">✓ = assigned · click to toggle</span>
    </div>
    <div class="card-body p-0" style="overflow-x:auto">
      <table class="table table-bordered table-sm mb-0" style="min-width:600px">
        <thead class="table-light">
          <tr>
            <th style="min-width:160px;position:sticky;left:0;background:#f9fafb;z-index:1">Teacher</th>
            <?php foreach ($subjectAxis as $key => $col): ?>
              <th class="text-center" style="min-width:90px;font-size:.7rem;font-weight:500">
                <span class="badge bg-primary-subtle text-primary d-block mb-1"><?= htmlspecialchars($col['subject_code']) ?></span>
                <span class="text-muted d-block" style="font-size:.65rem"><?= htmlspecialchars($col['class_name']) ?></span>
              </th>
            <?php endforeach; ?>
          </tr>
        </thead>
        <tbody>
          <?php foreach ($matrix as $tid => $row): ?>
          <tr>
            <td style="position:sticky;left:0;background:#fff;z-index:1;font-size:.82rem" class="fw-semibold">
              <?= htmlspecialchars($row['teacher']['name']) ?>
            </td>
            <?php foreach ($subjectAxis as $key => $col): ?>
              <?php $assigned = !empty($lookup[$tid][$key]); ?>
              <td class="text-center p-1">
                <form method="POST"
                      action="<?= $assigned ? '/teachers/remove-subject' : '/teachers/assign-subject' ?>"
                      onsubmit="return <?= $assigned ? "confirm('Remove this assignment?')" : 'true' ?>">
                  <input type="hidden" name="teacher_id" value="<?= $tid ?>">
                  <input type="hidden" name="subject_id" value="<?= $col['subject_id'] ?>">
                  <input type="hidden" name="class_id"   value="<?= $col['class_id'] ?>">
                  <input type="hidden" name="redirect"   value="/teachers/subject-admin">
                  <button type="submit"
                          class="btn btn-sm border-0 w-100 <?= $assigned ? 'btn-success' : 'btn-light text-muted' ?>"
                          title="<?= $assigned ? 'Assigned — click to remove' : 'Not assigned — click to assign' ?>"
                          style="font-size:.85rem;padding:.2rem">
                    <?= $assigned ? '<i class="bi bi-check-lg"></i>' : '<i class="bi bi-dash"></i>' ?>
                  </button>
                </form>
              </td>
            <?php endforeach; ?>
          </tr>
          <?php endforeach; ?>
        </tbody>
      </table>
    </div>
  </div>
</div>

<script>
const toggleBtn  = document.getElementById('toggleMatrixBtn');
const listView   = document.getElementById('listView');
const matrixView = document.getElementById('matrixView');
let showingMatrix = false;

toggleBtn.addEventListener('click', function () {
  showingMatrix = !showingMatrix;
  listView.style.display   = showingMatrix ? 'none' : '';
  matrixView.style.display = showingMatrix ? '' : 'none';
  toggleBtn.innerHTML = showingMatrix
    ? '<i class="bi bi-list-ul me-1"></i>List View'
    : '<i class="bi bi-table me-1"></i>Matrix View';
});

// Pre-filter list by teacher name
document.getElementById('qaTeacher').addEventListener('change', function () {
  const name = this.options[this.selectedIndex]?.text?.toLowerCase() ?? '';
  document.querySelectorAll('#listView .card').forEach(card => {
    const n = card.querySelector('.fw-semibold')?.textContent?.toLowerCase() ?? '';
    card.style.display = (!name || n.includes(name)) ? '' : 'none';
  });
});
</script>
