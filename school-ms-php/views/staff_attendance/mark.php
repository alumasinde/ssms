<?php /** @var array $teachers @var string $date @var array $existingMap */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-person-check me-2 text-success"></i>Mark Staff Attendance — <?= date('d M Y', strtotime($date)) ?></h5>
  <a href="/staff-attendance" class="btn btn-sm btn-outline-secondary"><i class="bi bi-arrow-left me-1"></i>Back</a>
</div>
<form method="POST" action="/staff-attendance">
  <input type="hidden" name="date" value="<?= htmlspecialchars($date) ?>">
  <div class="card"><div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Teacher</th><th>Emp No</th><th class="text-center">Present</th><th class="text-center">Absent</th><th class="text-center">Late</th><th class="text-center">On Leave</th><th>Check In</th><th>Check Out</th><th>Remark</th></tr></thead>
      <tbody>
        <?php if (empty($teachers)): ?>
          <tr><td colspan="9" class="text-center text-muted py-4">No teachers found.</td></tr>
        <?php else: foreach ($teachers as $t):
          $ex = $existingMap[$t['id']] ?? null;
          $cs = $ex['status'] ?? 'present';
        ?>
        <tr>
          <td class="fw-semibold small"><?= htmlspecialchars($t['name']) ?></td>
          <td class="small text-muted"><?= htmlspecialchars($t['employee_no'] ?? '') ?></td>
          <?php foreach (['present','absent','late','on_leave'] as $st): ?>
          <td class="text-center"><input type="radio" name="status[<?= $t['id'] ?>]" value="<?= $st ?>" <?= $cs===$st?'checked':'' ?> class="form-check-input"></td>
          <?php endforeach; ?>
          <td><input type="time" name="check_in[<?= $t['id'] ?>]" class="form-control form-control-sm" value="<?= htmlspecialchars($ex['check_in']??'') ?>" style="width:110px"></td>
          <td><input type="time" name="check_out[<?= $t['id'] ?>]" class="form-control form-control-sm" value="<?= htmlspecialchars($ex['check_out']??'') ?>" style="width:110px"></td>
          <td><input type="text" name="remark[<?= $t['id'] ?>]" class="form-control form-control-sm" value="<?= htmlspecialchars($ex['remark']??'') ?>" placeholder="Optional"></td>
        </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
  <div class="card-footer bg-white d-flex gap-2">
    <button type="submit" class="btn btn-success"><i class="bi bi-save me-1"></i>Save</button>
    <a href="/staff-attendance" class="btn btn-outline-secondary">Cancel</a>
  </div>
</form>
