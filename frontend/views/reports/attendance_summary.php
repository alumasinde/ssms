<?php /** @var array $reports @var array $terms @var int $termID */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-pie-chart me-2 text-info"></i>Attendance Summary Report</h5>
</div>
<div class="card mb-4">
  <div class="card-body py-3">
    <form method="GET" class="row g-2 align-items-end">
      <div class="col-md-5">
        <label class="form-label small fw-semibold">Term</label>
        <select name="term_id" class="form-select" onchange="this.form.submit()">
          <option value="">Select term...</option>
          <?php foreach ($terms as $t): ?>
            <option value="<?= $t['id'] ?>" <?= $termID == $t['id'] ? 'selected' : '' ?>>
              <?= htmlspecialchars($t['name']) ?><?= ($t['is_current'] ?? false) ? ' (Current)' : '' ?>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
      <?php if (!$termID): ?>
        <div class="col-md-5"><div class="alert alert-info py-2 small mb-0"><i class="bi bi-info-circle me-1"></i>Select a term to view school-wide attendance.</div></div>
      <?php endif; ?>
    </form>
  </div>
</div>

<?php if ($termID): ?>
<?php if (empty($reports)): ?>
  <div class="alert alert-warning"><i class="bi bi-exclamation-triangle me-2"></i>No attendance records found for this term.</div>
<?php else: ?>
<?php
$avgAll  = count($reports) > 0 ? round(array_sum(array_column($reports, 'avg_present')) / count($reports), 1) : 0;
$best    = array_reduce($reports, fn($c, $r) => ($c === null || $r['avg_present'] > $c['avg_present']) ? $r : $c, null);
$worst   = array_reduce($reports, fn($c, $r) => ($c === null || $r['avg_present'] < $c['avg_present']) ? $r : $c, null);
?>
<div class="row g-3 mb-4">
  <div class="col-sm-4"><div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)">
    <div class="stat-num"><?= $avgAll ?>%</div><div class="stat-label"><i class="bi bi-graph-up me-1"></i>School Average</div>
  </div></div>
  <div class="col-sm-4"><div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)">
    <div class="stat-num"><?= $best['avg_present'] ?? '—' ?>%</div>
    <div class="stat-label"><i class="bi bi-trophy me-1"></i>Best: <?= htmlspecialchars($best['class_name'] ?? '') ?></div>
  </div></div>
  <div class="col-sm-4"><div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)">
    <div class="stat-num"><?= $worst['avg_present'] ?? '—' ?>%</div>
    <div class="stat-label"><i class="bi bi-arrow-down me-1"></i>Lowest: <?= htmlspecialchars($worst['class_name'] ?? '') ?></div>
  </div></div>
</div>
<div class="card">
  <div class="card-header py-3"><i class="bi bi-table me-1"></i>Per-Class Breakdown</div>
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Class</th><th class="text-center">Days Recorded</th><th>Average Attendance</th><th class="text-center">Status</th></tr></thead>
      <tbody>
        <?php
        usort($reports, fn($a, $b) => $b['avg_present'] <=> $a['avg_present']);
        foreach ($reports as $row):
          $pct = $row['avg_present'] ?? 0;
          $color  = $pct >= 80 ? 'success' : ($pct >= 60 ? 'warning' : 'danger');
          $status = $pct >= 80 ? 'Good'    : ($pct >= 60 ? 'Fair'    : 'Poor');
        ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($row['class_name'] ?? '') ?></td>
          <td class="text-center"><?= $row['total_days'] ?? 0 ?></td>
          <td>
            <div class="d-flex align-items-center gap-2">
              <div class="progress flex-grow-1" style="height:8px;max-width:200px">
                <div class="progress-bar bg-<?= $color ?>" style="width:<?= $pct ?>%"></div>
              </div>
              <span class="small fw-bold text-<?= $color ?>"><?= $pct ?>%</span>
            </div>
          </td>
          <td class="text-center"><span class="badge bg-<?= $color ?>-subtle text-<?= $color ?>"><?= $status ?></span></td>
        </tr>
        <?php endforeach; ?>
      </tbody>
    </table>
  </div>
</div>
<?php endif; ?>
<?php endif; ?>
