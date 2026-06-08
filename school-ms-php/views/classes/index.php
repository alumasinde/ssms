<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-door-open me-2 text-primary"></i>Classes</h5>
  <a href="/classes/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>New Class</a>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Level</th><th>Stream</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($classes)): ?>
          <tr><td colspan="4" class="text-center text-muted py-4">No classes yet.</td></tr>
        <?php else: foreach ($classes as $c): ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($c['name']) ?></td>
          <td><?= htmlspecialchars($c['level']) ?></td>
          <td><?= htmlspecialchars($c['stream'] ?? '—') ?></td>
          <td class="text-end">
            <a href="/students?class_id=<?= $c['id'] ?>" class="btn btn-sm btn-outline-info" title="Students">
              <i class="bi bi-people"></i>
            </a>
            <a href="/attendance/mark?class_id=<?= $c['id'] ?>" class="btn btn-sm btn-outline-success" title="Take Attendance">
              <i class="bi bi-calendar-check"></i>
            </a>
            <a href="/classes/<?= $c['id'] ?>/edit" class="btn btn-sm btn-outline-secondary">
              <i class="bi bi-pencil"></i>
            </a>
          </td>
        </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>