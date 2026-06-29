<?php /** @var array $terms @var array $classes @var array $subjects @var array $teachers */ ?>
<div class="row justify-content-center"><div class="col-lg-7">
<div class="card"><div class="card-header py-3"><i class="bi bi-journal-plus me-2 text-warning"></i>New Assignment</div>
<div class="card-body">
  <form method="POST" action="/assignments">
    <div class="mb-3"><label class="form-label small fw-semibold">Title *</label>
      <input type="text" name="title" class="form-control" required placeholder="e.g. Chapter 5 Exercise"></div>
    <div class="mb-3"><label class="form-label small fw-semibold">Description</label>
      <textarea name="description" class="form-control" rows="3" placeholder="Instructions..."></textarea></div>
    <div class="row g-2 mb-3">
      <div class="col"><label class="form-label small fw-semibold">Term *</label>
        <select name="term_id" class="form-select" required>
          <option value="">Select...</option>
          <?php foreach ($terms as $t): ?><option value="<?= $t['id'] ?>" <?= ($t['is_current']??false)?'selected':'' ?>><?= htmlspecialchars($t['name']) ?></option><?php endforeach; ?>
        </select></div>
      <div class="col"><label class="form-label small fw-semibold">Class *</label>
        <select name="class_id" class="form-select" required>
          <option value="">Select...</option>
          <?php foreach ($classes as $c): ?><option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?></option><?php endforeach; ?>
        </select></div>
    </div>
    <div class="row g-2 mb-3">
      <div class="col"><label class="form-label small fw-semibold">Subject *</label>
        <select name="subject_id" class="form-select" required>
          <option value="">Select...</option>
          <?php foreach ($subjects as $s): ?><option value="<?= $s['id'] ?>"><?= htmlspecialchars($s['name']) ?></option><?php endforeach; ?>
        </select></div>
      <div class="col"><label class="form-label small fw-semibold">Teacher *</label>
        <select name="teacher_id" class="form-select" required>
          <option value="">Select...</option>
          <?php foreach ($teachers as $t): ?><option value="<?= $t['id'] ?>"><?= htmlspecialchars($t['name']) ?></option><?php endforeach; ?>
        </select></div>
    </div>
    <div class="row g-2 mb-3">
      <div class="col"><label class="form-label small fw-semibold">Due Date *</label>
        <input type="date" name="due_date" class="form-control" required></div>
      <div class="col"><label class="form-label small fw-semibold">Max Marks</label>
        <input type="number" name="max_marks" class="form-control" value="100" min="1"></div>
    </div>
    <div class="d-flex gap-2">
      <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Create</button>
      <a href="/assignments" class="btn btn-outline-secondary">Cancel</a>
    </div>
  </form>
</div></div></div></div>
