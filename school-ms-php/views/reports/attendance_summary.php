<h5 class="fw-bold mb-3"><i class="bi bi-pie-chart me-2 text-info"></i>Attendance Summary</h5>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Class</th><th>Total Days</th><th>Avg Attendance %</th></tr></thead>
      <tbody>
        <?php if (empty($reports)): ?>
          <tr><td colspan="3" class="text-center text-muted py-4">No data for this term.</td></tr>
        <?php else: foreach ($reports as $r): ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($r['class_name']) ?></td>
            <td><?= $r['total_days'] ?></td>
            <td>
              <div class="d-flex align-items-center gap-2">
                <div class="progress flex-grow-1" style="height:8px">
                  <div class="progress-bar <?= $r['avg_present'] >= 80 ? 'bg-success' : ($r['avg_present'] >= 60 ? 'bg-warning' : 'bg-danger') ?>" style="width:<?= $r['avg_present'] ?>%"></div>
                </div>
                <span class="small fw-bold"><?= $r['avg_present'] ?>%</span>
              </div>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>