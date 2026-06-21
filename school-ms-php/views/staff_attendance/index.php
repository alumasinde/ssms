<?php /** @var array $records @var string $date */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-person-check me-2 text-success"></i>Staff Attendance</h5>
  <div class="d-flex gap-2">
    <a href="/staff-attendance/summary" class="btn btn-outline-info btn-sm"><i class="bi bi-bar-chart me-1"></i>Summary</a>
    <a href="/staff-attendance/mark?date=<?= $date ?>" class="btn btn-primary btn-sm"><i class="bi bi-pencil-square me-1"></i>Mark Today</a>
  </div>
</div>
<div class="card mb-3"><div class="card-body py-2">
  <form method="GET" class="row g-2 align-items-end">
    <div class="col-md-4">
      <label class="form-label small fw-semibold mb-1">Date</label>
      <input type="date" name="date" class="form-control form-control-sm" value="<?= $date ?>" onchange="this.form.submit()">
    </div>
  </form>
</div></div>
<div class="card"><div class="card-body p-0">
  <table class="table table-hover mb-0">
    <thead><tr><th>Teacher</th><th>Employee No</th><th class="text-center">Status</th><th>Check In</th><th>Check Out</th><th>Remark</th></tr></thead>
    <tbody>
      <?php if (empty($records)): ?>
        <tr><td colspan="6" class="text-center text-muted py-4">No records for <?= $date ?>. <a href="/staff-attendance/mark?date=<?= $date ?>">Mark now.</a></td></tr>
      <?php else: foreach ($records as $r):
        $badge = match($r['status']) { 'present'=>'bg-success','absent'=>'bg-danger','late'=>'bg-warning text-dark','on_leave'=>'bg-info','sick_leave'=>'bg-secondary', default=>'bg-secondary' };
      ?>
      <tr>
        <td class="fw-semibold"><?= htmlspecialchars($r['teacher_name']) ?></td>
        <td class="small text-muted"><?= htmlspecialchars($r['employee_no']) ?></td>
        <td class="text-center"><span class="badge <?= $badge ?>"><?= ucwords(str_replace('_',' ',$r['status'])) ?></span></td>
        <td class="small"><?= $r['check_in']  ?: '—' ?></td>
        <td class="small"><?= $r['check_out'] ?: '—' ?></td>
        <td class="small text-muted"><?= htmlspecialchars($r['remark'] ?? '') ?></td>
      </tr>
      <?php endforeach; endif; ?>
    </tbody>
  </table>
</div></div>
