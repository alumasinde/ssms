<h5 class="fw-bold mb-3"><i class="bi bi-receipt me-2 text-success"></i>Fee Collection Report</h5>
<?php $r = $report ?? []; ?>
<div class="row g-3 mb-4">
  <div class="col-md-4"><div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)"><div class="stat-num">KES <?= number_format($r['total_billed'] ?? 0, 2) ?></div><div class="stat-label">Total Billed</div></div></div>
  <div class="col-md-4"><div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)"><div class="stat-num">KES <?= number_format($r['total_paid'] ?? 0, 2) ?></div><div class="stat-label">Collected</div></div></div>
  <div class="col-md-4"><div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)"><div class="stat-num">KES <?= number_format($r['balance'] ?? 0, 2) ?></div><div class="stat-label">Outstanding</div></div></div>
</div>
<div class="card"><div class="card-body">
  <div class="row text-center">
    <div class="col"><span class="display-6 text-success fw-bold"><?= $r['paid_count'] ?? 0 ?></span><div class="small text-muted mt-1">Fully Paid</div></div>
    <div class="col"><span class="display-6 text-warning fw-bold"><?= $r['partial_count'] ?? 0 ?></span><div class="small text-muted mt-1">Partial</div></div>
    <div class="col"><span class="display-6 text-danger fw-bold"><?= $r['unpaid_count'] ?? 0 ?></span><div class="small text-muted mt-1">Unpaid</div></div>
  </div>
</div></div>