<?php
/** @var array $student   */
/** @var array $statement */
/** @var array $discounts */
$s    = $statement ?? [];
$name = trim(($student['first_name'] ?? '') . ' ' . ($student['last_name'] ?? ''));
$sid  = $student['id'] ?? '';
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-receipt me-2 text-success"></i>Fee Statement — <?= htmlspecialchars($name) ?></h5>
  <div class="d-flex gap-2">
    <?php if (\Core\Session::can('finance.discount')): ?>
      <a href="/finance/discounts/create?student_id=<?= $sid ?>" class="btn btn-sm btn-outline-warning">
        <i class="bi bi-tag me-1"></i>Add Discount
      </a>
    <?php endif; ?>
    <a href="/students/<?= $sid ?>" class="btn btn-sm btn-outline-secondary"><i class="bi bi-arrow-left me-1"></i>Back</a>
  </div>
</div>

<div class="row g-3 mb-3">
  <div class="col-md-4">
    <div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)">
      <div class="stat-num">KES <?= number_format($s['total_billed'] ?? 0, 2) ?></div>
      <div class="stat-label">Total Billed</div>
    </div>
  </div>
  <div class="col-md-4">
    <div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)">
      <div class="stat-num">KES <?= number_format($s['total_paid'] ?? 0, 2) ?></div>
      <div class="stat-label">Total Paid</div>
    </div>
  </div>
  <div class="col-md-4">
    <div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)">
      <div class="stat-num">KES <?= number_format($s['balance'] ?? 0, 2) ?></div>
      <div class="stat-label">Balance</div>
    </div>
  </div>
</div>

<!-- Discounts -->
<?php if (!empty($discounts)): ?>
<div class="card mb-3">
  <div class="card-header py-2 d-flex justify-content-between align-items-center">
    <span><i class="bi bi-tag me-1 text-warning"></i>Active Discounts / Bursaries</span>
  </div>
  <div class="card-body p-0">
    <table class="table table-sm mb-0">
      <thead><tr><th>Label</th><th>Fee Type</th><th>Term</th><th>Discount</th></tr></thead>
      <tbody>
        <?php foreach ($discounts as $d): ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($d['label']) ?></td>
          <td class="text-muted small"><?= $d['fee_type_id'] ? '#'.$d['fee_type_id'] : 'All' ?></td>
          <td class="text-muted small"><?= $d['term_id'] ? '#'.$d['term_id'] : 'Recurring' ?></td>
          <td>
            <?php if ($d['discount_pct']): ?>
              <span class="badge bg-warning-subtle text-warning"><?= $d['discount_pct'] ?>%</span>
            <?php elseif ($d['discount_amt']): ?>
              <span class="badge bg-warning-subtle text-warning">KES <?= number_format($d['discount_amt'],2) ?></span>
            <?php endif; ?>
          </td>
        </tr>
        <?php endforeach; ?>
      </tbody>
    </table>
  </div>
</div>
<?php endif; ?>

<div class="card">
  <div class="card-header py-3 d-flex justify-content-between align-items-center">
    <span>Invoices</span>
    <?php if (\Core\Session::can('finance.create')): ?>
      <a href="/finance/invoices/generate" class="btn btn-sm btn-outline-primary"><i class="bi bi-plus-lg me-1"></i>Generate Invoices</a>
    <?php endif; ?>
  </div>
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Fee Type</th><th>Term</th><th>Amount (KES)</th><th>Due Date</th><th>Status</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($s['invoices'])): ?>
          <tr><td colspan="6" class="text-center text-muted py-4">No invoices found.</td></tr>
        <?php else: foreach ($s['invoices'] as $inv): ?>
          <tr>
            <td><?= htmlspecialchars($inv['fee_type_name'] ?? 'Fee #'.$inv['fee_type_id']) ?></td>
            <td class="text-muted small">Term #<?= $inv['term_id'] ?></td>
            <td class="fw-semibold"><?= number_format($inv['amount'], 2) ?></td>
            <td><?= $inv['due_date'] ?? '—' ?></td>
            <td>
              <?php $st = $inv['status'] ?? 'unpaid'; ?>
              <span class="badge <?= $st === 'paid' ? 'bg-success-subtle text-success' : ($st === 'partial' ? 'bg-warning-subtle text-warning' : 'bg-danger-subtle text-danger') ?>">
                <?= ucfirst($st) ?>
              </span>
            </td>
            <td>
              <?php if (($inv['status'] ?? '') !== 'paid' && \Core\Session::can('finance.create')): ?>
                <button class="btn btn-sm btn-success"
                        data-bs-toggle="modal" data-bs-target="#payModal"
                        data-invoice="<?= $inv['id'] ?>"
                        data-amount="<?= $inv['amount'] ?>">
                  <i class="bi bi-cash me-1"></i>Pay
                </button>
              <?php endif; ?>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>

<?php if (\Core\Session::can('finance.create')): ?>
<div class="modal fade" id="payModal" tabindex="-1">
  <div class="modal-dialog modal-sm">
    <div class="modal-content">
      <div class="modal-header">
        <h6 class="modal-title">Record Payment</h6>
        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
      </div>
      <div class="modal-body">
        <form method="POST" action="/finance/payment">
          <input type="hidden" name="invoice_id" id="payInvoiceID">
          <input type="hidden" name="student_id" value="<?= $sid ?>">
          <div class="mb-2">
            <label class="form-label small fw-semibold">Amount (KES)</label>
            <input type="number" step="0.01" name="amount_paid" id="payAmount" class="form-control" required>
          </div>
          <div class="mb-2">
            <label class="form-label small fw-semibold">Method</label>
            <select name="method" class="form-select" id="payMethod">
              <option value="cash">Cash</option>
              <option value="mpesa">M-Pesa</option>
              <option value="bank">Bank Transfer</option>
              <option value="cheque">Cheque</option>
            </select>
          </div>
          <div class="mb-2">
            <label class="form-label small fw-semibold">Reference / M-Pesa Code</label>
            <input type="text" name="ref_no" class="form-control" placeholder="Optional reference or M-Pesa code">
          </div>
          <button type="submit" class="btn btn-success w-100"><i class="bi bi-check-circle me-1"></i>Record Payment</button>
        </form>
      </div>
    </div>
  </div>
</div>
<script>
document.getElementById('payModal').addEventListener('show.bs.modal', function(e) {
  const btn = e.relatedTarget;
  document.getElementById('payInvoiceID').value = btn.dataset.invoice;
  document.getElementById('payAmount').value    = btn.dataset.amount;
});
</script>
<?php endif; ?>
