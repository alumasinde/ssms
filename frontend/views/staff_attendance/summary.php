<?php /** @var array $records @var string $month */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-bar-chart me-2 text-info"></i>Staff Attendance Summary</h5>
</div>
<div class="card mb-3"><div class="card-body py-2">
  <form method="GET" class="row g-2">
    <div class="col-md-4">
      <label class="form-label small fw-semibold mb-1">Month</label>
      <input type="month" name="month" class="form-control form-control-sm" value="<?= $month ?>" onchange="this.form.submit()">
    </div>
  </form>
</div></div>
<div class="card"><div class="card-body p-0">
  <table class="table table-hover mb-0">
    <thead><tr><th>Teacher</th><th>Employee No</th><th class="text-center">Days Recorded</th><th class="text-center">Present</th><th>Attendance %</th></tr></thead>
    <tbody>
      <?php if (empty($records)): ?>
        <tr><td colspan="5" class="text-center text-muted py-4">No records for <?= $month ?>.</td></tr>
      <?php else: foreach ($records as $r):
        $pct = isset($r['total_days']) && $r['total_days'] > 0 ? round($r['present_days']/$r['total_days']*100,1) : 0;
      ?>
      <tr>
        <td class="fw-semibold"><?= htmlspecialchars($r['teacher_name']) ?></td>
        <td class="small text-muted"><?= htmlspecialchars($r['employee_no']) ?></td>
        <td class="text-center"><?= $r['total_days'] ?? 0 ?></td>
        <td class="text-center text-success fw-bold"><?= $r['present_days'] ?? 0 ?></td>
        <td>
          <div class="d-flex align-items-center gap-2">
            <div class="progress flex-grow-1" style="height:8px;max-width:120px">
              <div class="progress-bar <?= $pct>=80?'bg-success':($pct>=60?'bg-warning':'bg-danger') ?>" style="width:<?= $pct ?>%"></div>
            </div>
            <span class="small fw-bold"><?= $pct ?>%</span>
          </div>
        </td>
      </tr>
      <?php endforeach; endif; ?>
    </tbody>
  </table>
</div></div>
