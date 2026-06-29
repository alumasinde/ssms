<?php
/** @var array $report @var array $terms @var int $termID */
$r = $report ?? [];
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-receipt me-2 text-success"></i>Fee Collection Report</h5>
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
        <div class="col-md-5"><div class="alert alert-info py-2 small mb-0"><i class="bi bi-info-circle me-1"></i>Select a term to view the fee collection summary.</div></div>
      <?php endif; ?>
    </form>
  </div>
</div>

<?php if ($termID): ?>
<div class="row g-3 mb-4">
  <div class="col-sm-6 col-lg-4">
    <div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)">
      <div class="stat-num">KES <?= number_format($r['total_billed'] ?? 0, 2) ?></div>
      <div class="stat-label"><i class="bi bi-receipt me-1"></i>Total Billed</div>
    </div>
  </div>
  <div class="col-sm-6 col-lg-4">
    <div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)">
      <div class="stat-num">KES <?= number_format($r['total_paid'] ?? 0, 2) ?></div>
      <div class="stat-label"><i class="bi bi-cash-coin me-1"></i>Collected</div>
    </div>
  </div>
  <div class="col-sm-6 col-lg-4">
    <div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)">
      <div class="stat-num">KES <?= number_format($r['balance'] ?? 0, 2) ?></div>
      <div class="stat-label"><i class="bi bi-exclamation-circle me-1"></i>Outstanding</div>
    </div>
  </div>
</div>
<?php
$billed = $r['total_billed'] ?? 0;
$paid   = $r['total_paid'] ?? 0;
$rate   = $billed > 0 ? round($paid / $billed * 100, 1) : 0;
?>
<div class="card mb-3">
  <div class="card-body">
    <div class="d-flex justify-content-between mb-2">
      <span class="small fw-semibold">Collection Rate</span>
      <span class="small fw-bold <?= $rate >= 80 ? 'text-success' : ($rate >= 50 ? 'text-warning' : 'text-danger') ?>"><?= $rate ?>%</span>
    </div>
    <div class="progress" style="height:10px">
      <div class="progress-bar <?= $rate >= 80 ? 'bg-success' : ($rate >= 50 ? 'bg-warning' : 'bg-danger') ?>" style="width:<?= $rate ?>%"></div>
    </div>
  </div>
</div>
<div class="card">
  <div class="card-header py-3">Invoice Status Breakdown</div>
  <div class="card-body">
    <?php $total = ($r['paid_count'] ?? 0) + ($r['partial_count'] ?? 0) + ($r['unpaid_count'] ?? 0); ?>
    <div class="row g-3 text-center mb-4">
      <?php foreach ([['Fully Paid',$r['paid_count']??0,'success','bi-check-circle-fill'],['Partial',$r['partial_count']??0,'warning','bi-dash-circle-fill'],['Unpaid',$r['unpaid_count']??0,'danger','bi-x-circle-fill']] as [$label,$count,$color,$icon]): ?>
        <div class="col-4">
          <div class="p-3 rounded-3 bg-light">
            <i class="bi <?= $icon ?> fs-2 text-<?= $color ?> d-block mb-1"></i>
            <div class="display-6 fw-bold text-<?= $color ?>"><?= number_format($count) ?></div>
            <div class="small text-muted"><?= $label ?></div>
            <?php if ($total > 0): ?><div class="small text-<?= $color ?>"><?= round($count/$total*100,1) ?>%</div><?php endif; ?>
          </div>
        </div>
      <?php endforeach; ?>
    </div>
  </div>
</div>
<?php endif; ?>
