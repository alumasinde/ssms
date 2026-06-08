<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0">
    <i class="bi bi-calendar-check me-2 text-success"></i>
    Mark Attendance — <?= date('d M Y', strtotime($date)) ?>
  </h5>
  <a href="/attendance" class="btn btn-sm btn-outline-secondary">
    <i class="bi bi-arrow-left me-1"></i>Back
  </a>
</div>

<form method="POST" action="/attendance">
  <input type="hidden" name="class_id" value="<?= $classID ?>">
  <input type="hidden" name="date"     value="<?= htmlspecialchars($date) ?>">
  <input type="hidden" name="term_id"  value="<?= $currentTerm['id'] ?? 0 ?>">

  <?php if ($currentTerm): ?>
    <div class="alert alert-info py-2 small mb-3">
      <i class="bi bi-info-circle me-1"></i>
      Recording for: <strong><?= htmlspecialchars($currentTerm['name'] ?? 'Current Term') ?></strong>
    </div>
  <?php else: ?>
    <div class="alert alert-warning py-2 small mb-3">
      <i class="bi bi-exclamation-triangle me-1"></i>
      No current term set. Please configure a current term in Academic Years.
    </div>
  <?php endif; ?>

  <div class="card">
    <div class="card-body p-0">
      <table class="table table-hover mb-0">
        <thead>
          <tr>
            <th>Student</th>
            <th class="text-center">Present</th>
            <th class="text-center">Absent</th>
            <th class="text-center">Late</th>
            <th class="text-center">Excused</th>
            <th>Remark</th>
          </tr>
        </thead>
        <tbody>
          <?php if (empty($students)): ?>
            <tr><td colspan="6" class="text-center text-muted py-4">No students in this class.</td></tr>
          <?php else: ?>
            <?php foreach ($students as $s):
              $existing      = $existingMap[$s['id']] ?? null;
              $currentStatus = $existing['status'] ?? 'present';
            ?>
            <tr>
              <td class="fw-semibold small">
                <?= htmlspecialchars($s['name']) ?>
                <br><span class="text-muted" style="font-size:.72rem">
                  <?= htmlspecialchars($s['admission_no']) ?>
                </span>
              </td>
              <?php foreach (['present','absent','late','excused'] as $st): ?>
              <td class="text-center">
                <input type="radio"
                       name="status[<?= $s['id'] ?>]"
                       value="<?= $st ?>"
                       <?= $currentStatus === $st ? 'checked' : '' ?>
                       class="form-check-input">
              </td>
              <?php endforeach; ?>
              <td>
                <input type="text"
                       name="remark[<?= $s['id'] ?>]"
                       class="form-control form-control-sm"
                       value="<?= htmlspecialchars($existing['remark'] ?? '') ?>"
                       placeholder="Optional">
              </td>
            </tr>
            <?php endforeach; ?>
          <?php endif; ?>
        </tbody>
      </table>
    </div>
    <div class="card-footer bg-white d-flex gap-2">
      <button type="submit" class="btn btn-success">
        <i class="bi bi-save me-1"></i>Save Attendance
      </button>
      <a href="/attendance" class="btn btn-outline-secondary">Cancel</a>
    </div>
  </div>
</form>