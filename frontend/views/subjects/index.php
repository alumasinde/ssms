<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-book me-2 text-primary"></i>Subjects</h5>
  <a href="/subjects/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Add Subject</a>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Code</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($subjects)): ?>
          <tr><td colspan="3" class="text-center text-muted py-4">No subjects yet.</td></tr>
        <?php else: foreach ($subjects as $s): ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($s['name']) ?></td>
          <td><span class="badge bg-light text-dark border"><?= htmlspecialchars($s['code']) ?></span></td>
          <td class="text-end">
            <a href="/subjects/<?= $s['id'] ?>/edit" class="btn btn-sm btn-outline-secondary">
              <i class="bi bi-pencil"></i>
            </a>
          </td>
        </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>