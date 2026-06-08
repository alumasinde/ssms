<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-person-workspace me-2 text-primary"></i>Teachers</h5>
  <a href="/teachers/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Add Teacher</a>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Employee No</th><th>Name</th><th>Email</th><th>Phone</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($teachers)): ?>
          <tr><td colspan="5" class="text-center text-muted py-4">No teachers found.</td></tr>
        <?php else: foreach ($teachers as $t): ?>
          <tr>
            <td><span class="badge bg-light text-dark border"><?= htmlspecialchars($t['employee_no']) ?></span></td>
            <td class="fw-semibold"><?= htmlspecialchars($t['name']) ?></td>
            <td><?= htmlspecialchars($t['email']) ?></td>
            <td><?= htmlspecialchars($t['phone'] ?? '—') ?></td>
            <td class="text-end"><a href="/teachers/<?= $t['id'] ?>" class="btn btn-sm btn-outline-primary"><i class="bi bi-eye"></i></a></td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>