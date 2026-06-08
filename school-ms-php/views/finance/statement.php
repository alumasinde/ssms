<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-receipt me-2 text-success"></i>Fee Statement — <?= htmlspecialchars($student['name'] ?? '') ?></h5>
  <a href="/students/<?= $student['id'] ?>" class="btn btn-sm btn-outline-secondary"><i class="bi bi-arrow-left me-1"></i>Back</a>
</div>
<?php $stmt = $statement ?? []; ?>
<div class="row g-3 mb-3">
  <div class="col-md-4"><div class="stat-card" style="background:linear-gradient(135deg,#1a56db,#3b82f6)"><div class="stat-num">KES <?= number_format($stmt['total_billed'] ?? 0, 2) ?></div><div class="stat-label">Total Billed</div></div></div>
  <div class="col-md-4"><div class="stat-card" style="background:linear-gradient(135deg,#059669,#34d399)"><div class="stat-num">KES <?= number_format($stmt['total_paid'] ?? 0, 2) ?></div><div class="stat-label">Total Paid</div></div></div>
  <div class="col-md-4"><div class="stat-card" style="background:linear-gradient(135deg,#d97706,#fbbf24)"><div class="stat-num">KES <?= number_format($stmt['balance'] ?? 0, 2) ?></div><div class="stat-label">Balance</div></div></div>
</div>
<div class="card">
  <div class="card-header py-3">Invoices</div>
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Fee Type</th><th>Term</th><th>Amount</th><th>Due Date</th><th>Status</th><th></th></tr></thead>
      <tbody>
        <?php foreach ($stmt['invoices'] ?? [] as $inv): ?>
        <tr>
          <td><?= $inv['fee_type_id'] ?></td>
          <td><?= $inv['term_id'] ?></td>
          <td class="fw-semibold">KES <?= number_format($inv['amount'], 2) ?></td>
          <td><?= $inv['due_date'] ?></td>
          <td><span class="badge <?= $inv['status'] === 'paid' ? 'bg-success-subtle text-success' : ($inv['status'] === 'partial' ? 'bg-warning-subtle text-warning' : 'bg-danger-subtle text-danger') ?>"><?= ucfirst($inv['status']) ?></span></td>
          <td>
            <?php if ($inv['status'] !== 'paid'): ?>
            <button class="btn btn-sm btn-success" data-bs-toggle="modal" data-bs-target="#payModal" data-invoice="<?= $inv['id'] ?>" data-amount="<?= $inv['amount'] ?>">
              <i class="bi bi-cash me-1"></i>Pay
            </button>
            <?php endif; ?>
          </td>
        </tr>
        <?php endforeach; ?>
      </tbody>
    </table>
  </div>
</div>

<div class="modal fade" id="payModal" tabindex="-1">
  <div class="modal-dialog modal-sm">
    <div class="modal-content">
      <div class="modal-header"><h6 class="modal-title">Record Payment</h6><button type="button" class="btn-close" data-bs-dismiss="modal"></button></div>
      <div class="modal-body">
        <form method="POST" action="/finance/payment">
          <input type="hidden" name="invoice_id" id="payInvoiceID">
          <input type="hidden" name="student_id" value="<?= $student['id'] ?? '' ?>">
          <div class="mb-3"><label class="form-label small fw-semibold">Amount (KES)</label><input type="number" step="0.01" name="amount_paid" id="payAmount" class="form-control" required></div>
          <div class="mb-3"><label class="form-label small fw-semibold">Method</label><select name="method" class="form-select"><option value="cash">Cash</option><option value="mpesa">M-Pesa</option><option value="bank">Bank</option><option value="cheque">Cheque</option></select></div>
          <div class="mb-3"><label class="form-label small fw-semibold">Reference No</label><input type="text" name="ref_no" class="form-control" placeholder="e.g. M-Pesa code"></div>
          <button type="submit" class="btn btn-success w-100">Record Payment</button>
        </form>
      </div>
    </div>
  </div>
</div>
<script>
document.getElementById('payModal').addEventListener('show.bs.modal', function(e) {
  const btn = e.relatedTarget;
  document.getElementById('payInvoiceID').value = btn.dataset.invoice;
  document.getElementById('payAmount').value = btn.dataset.amount;
});
</script>