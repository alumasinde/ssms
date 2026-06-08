<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-cash-coin me-2 text-success"></i>Finance</h5>
</div>
<div class="card">
  <div class="card-header py-3">Fee Types</div>
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Amount (KES)</th><th>Frequency</th><th>Mandatory</th></tr></thead>
      <tbody>
        <?php if (empty($feeTypes)): ?>
          <tr><td colspan="4" class="text-center text-muted py-4">No fee types configured.</td></tr>
        <?php else: foreach ($feeTypes as $ft): ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($ft['name']) ?></td>
            <td><?= number_format($ft['amount'], 2) ?></td>
            <td><span class="badge bg-info-subtle text-info"><?= $ft['frequency'] ?></span></td>
            <td><?= $ft['is_mandatory'] ? '<span class="badge bg-danger-subtle text-danger">Yes</span>' : '<span class="badge bg-secondary-subtle text-secondary">No</span>' ?></td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>