<?php /** @var array $assignments @var array $terms @var array $classes @var int $termID @var int $classID */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-journal-check me-2 text-warning"></i>Assignments</h5>
  <?php if (\Core\Session::can('assignments.create')): ?>
    <a href="/assignments/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>New Assignment</a>
  <?php endif; ?>
</div>
<div class="card mb-3"><div class="card-body py-2">
  <form method="GET" class="row g-2 align-items-end">
    <div class="col-md-4">
      <label class="form-label small fw-semibold mb-1">Term</label>
      <select name="term_id" class="form-select form-select-sm" onchange="this.form.submit()">
        <option value="">All terms</option>
        <?php foreach ($terms as $t): ?><option value="<?= $t['id'] ?>" <?= $termID==$t['id']?'selected':'' ?>><?= htmlspecialchars($t['name']) ?></option><?php endforeach; ?>
      </select>
    </div>
    <div class="col-md-4">
      <label class="form-label small fw-semibold mb-1">Class</label>
      <select name="class_id" class="form-select form-select-sm" onchange="this.form.submit()">
        <option value="">All classes</option>
        <?php foreach ($classes as $c): ?><option value="<?= $c['id'] ?>" <?= $classID==$c['id']?'selected':'' ?>><?= htmlspecialchars($c['name']) ?></option><?php endforeach; ?>
      </select>
    </div>
  </form>
</div></div>
<div class="card"><div class="card-body p-0">
  <table class="table table-hover mb-0">
    <thead><tr><th>Title</th><th>Class</th><th>Subject</th><th>Teacher</th><th>Due</th><th>Max</th><th></th></tr></thead>
    <tbody>
      <?php if (empty($assignments)): ?>
        <tr><td colspan="7" class="text-center text-muted py-4">No assignments found.</td></tr>
      <?php else: foreach ($assignments as $a): ?>
      <tr>
        <td class="fw-semibold"><?= htmlspecialchars($a['title']) ?></td>
        <td class="small text-muted"><?= htmlspecialchars($a['class_name']) ?></td>
        <td class="small"><?= htmlspecialchars($a['subject_name']) ?></td>
        <td class="small text-muted"><?= htmlspecialchars($a['teacher_name']) ?></td>
        <td class="small"><?= $a['due_date'] ?></td>
        <td class="small"><?= $a['max_marks'] ?></td>
        <td class="text-end">
          <?php if (\Core\Session::can('assignments.edit')): ?>
            <form method="POST" action="/assignments/<?= $a['id'] ?>/delete" class="d-inline" onsubmit="return confirm('Delete?')">
              <button class="btn btn-sm btn-outline-danger"><i class="bi bi-trash"></i></button>
            </form>
          <?php endif; ?>
        </td>
      </tr>
      <?php endforeach; endif; ?>
    </tbody>
  </table>
</div></div>
