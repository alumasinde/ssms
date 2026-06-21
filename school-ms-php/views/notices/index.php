<?php
/** @var array $notices */
/** @var string $audience */
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-megaphone me-2 text-warning"></i>Notices</h5>
  <?php if (\Core\Session::can('notices.create')): ?>
    <a href="/notices/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Post Notice</a>
  <?php endif; ?>
</div>

<div class="mb-3">
  <?php foreach (['' => 'All', 'teachers' => 'Teachers', 'parents' => 'Parents', 'students' => 'Students'] as $val => $label): ?>
    <a href="/notices<?= $val ? '?audience=' . $val : '' ?>"
       class="btn btn-sm me-1 <?= $audience === $val ? 'btn-primary' : 'btn-outline-secondary' ?>">
      <?= $label ?>
    </a>
  <?php endforeach; ?>
</div>

<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Title</th><th>Audience</th><th>Published</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($notices)): ?>
          <tr><td colspan="4" class="text-center text-muted py-4">No notices.</td></tr>
        <?php else: foreach ($notices as $n): ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($n['title'] ?? '') ?></td>
            <td><span class="badge bg-secondary-subtle text-secondary"><?= ucfirst($n['audience'] ?? 'all') ?></span></td>
            <td class="small text-muted"><?= !empty($n['published_at']) ? date('d M Y', strtotime($n['published_at'])) : '—' ?></td>
            <td class="text-end">
              <a href="/notices/<?= $n['id'] ?>" class="btn btn-sm btn-outline-primary"><i class="bi bi-eye"></i></a>
              <?php if (\Core\Session::can('notices.delete')): ?>
                <form method="POST" action="/notices/<?= $n['id'] ?>/delete" class="d-inline"
                      onsubmit="return confirm('Delete this notice?')">
                  <button class="btn btn-sm btn-outline-danger"><i class="bi bi-trash"></i></button>
                </form>
              <?php endif; ?>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>
