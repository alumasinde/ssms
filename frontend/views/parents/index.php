<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-people me-2 text-primary"></i>Parents</h5>
  <a href="/parents/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Add Parent</a>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Email</th><th>Phone</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($parents)): ?>
          <tr><td colspan="4" class="text-center text-muted py-4">No parents registered.</td></tr>
        <?php else: foreach ($parents as $p): ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($p['name']) ?></td>
          <td><?= htmlspecialchars($p['email']) ?></td>
          <td><?= htmlspecialchars($p['phone'] ?? '—') ?></td>
          <td class="text-end">
            <a href="/parents/<?= $p['id'] ?>" class="btn btn-sm btn-outline-primary">
              <i class="bi bi-eye"></i>
            </a>
          </td>
        </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>