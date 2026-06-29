<?php /** @var array $terms */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-calendar3 me-2 text-primary"></i>Terms</h5>
  <?php if (\Core\Session::can('academic_years.create')): ?>
    <a href="/terms/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Add Term</a>
  <?php endif; ?>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Academic Year</th><th>Start</th><th>End</th><th class="text-center">Current</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($terms)): ?>
          <tr><td colspan="6" class="text-center text-muted py-4">No terms yet. <a href="/terms/create">Create one.</a></td></tr>
        <?php else: foreach ($terms as $t): ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($t['name']) ?></td>
            <td class="text-muted small"><?= htmlspecialchars($t['academic_year_id']) ?></td>
            <td><?= $t['start_date'] ?></td>
            <td><?= $t['end_date'] ?></td>
            <td class="text-center">
              <?php if ($t['is_current'] ?? false): ?>
                <span class="badge bg-success"><i class="bi bi-check-lg me-1"></i>Current</span>
              <?php else: ?>
                <?php if (\Core\Session::can('academic_years.edit')): ?>
                  <form method="POST" action="/terms/<?= $t['id'] ?>/set-current" class="d-inline">
                    <button type="submit" class="btn btn-xs btn-outline-secondary" style="font-size:.7rem;padding:.15rem .4rem">
                      Set Current
                    </button>
                  </form>
                <?php else: ?>—<?php endif; ?>
              <?php endif; ?>
            </td>
            <td class="text-end">
              <?php if (\Core\Session::can('academic_years.edit')): ?>
                <a href="/terms/<?= $t['id'] ?>/edit" class="btn btn-sm btn-outline-secondary">
                  <i class="bi bi-pencil"></i>
                </a>
                <form method="POST" action="/terms/<?= $t['id'] ?>/delete" class="d-inline"
                      onsubmit="return confirm('Delete this term?')">
                  <button type="submit" class="btn btn-sm btn-outline-danger"><i class="bi bi-trash"></i></button>
                </form>
              <?php endif; ?>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>
