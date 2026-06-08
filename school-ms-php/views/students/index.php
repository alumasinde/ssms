<?php
/** @var array $students */
/** @var array $meta */
/** @var int $page */
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-person-badge me-2 text-primary"></i>Students</h5>
  <a href="/students/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Enrol Student</a>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Adm No</th><th>Name</th><th>Gender</th><th>Class</th><th>Status</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($students)): ?>
          <tr><td colspan="6" class="text-center text-muted py-4">No students found.</td></tr>
        <?php else: foreach ($students as $s): ?>
     
          <tr>
            <td><span class="badge bg-light text-dark border"><?= htmlspecialchars($s['admission_no']) ?></span></td>
            <td class="fw-semibold"><?= htmlspecialchars($s['first_name']) ?> <?= htmlspecialchars($s['middle_name']) ?> <?= htmlspecialchars($s['last_name']) ?></td>
            <td><?= ucfirst($s['gender'] ?? '—') ?></td>
            <td><?= $s['class_id'] ?></td>
            <td><span class="badge <?= $s['is_active'] ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>"><?= $s['is_active'] ? 'Active' : 'Inactive' ?></span></td>
            <td class="text-end">
              <a href="/students/<?= $s['id'] ?>" class="btn btn-sm btn-outline-primary"><i class="bi bi-eye"></i></a>
              <a href="/students/<?= $s['id'] ?>/edit" class="btn btn-sm btn-outline-secondary"><i class="bi bi-pencil"></i></a>
              <a href="/reports/report-card/<?= $s['id'] ?>" class="btn btn-sm btn-outline-info" title="Report Card"><i class="bi bi-file-earmark-text"></i></a>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
  <?php if (!empty($meta)): ?>
  <div class="card-footer bg-white d-flex justify-content-between align-items-center py-2">
    <small class="text-muted">Showing <?= count($students) ?> of <?= $meta['total'] ?> students</small>
    <div class="d-flex gap-1">
      <?php if ($page > 1): ?><a href="?page=<?= $page-1 ?>" class="btn btn-sm btn-outline-secondary"><i class="bi bi-chevron-left"></i></a><?php endif; ?>
      <?php if ($page < ($meta['total_pages'] ?? 1)): ?><a href="?page=<?= $page+1 ?>" class="btn btn-sm btn-outline-secondary"><i class="bi bi-chevron-right"></i></a><?php endif; ?>
    </div>
  </div>
  <?php endif; ?>
</div>
