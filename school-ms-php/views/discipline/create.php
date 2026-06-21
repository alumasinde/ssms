<?php /** @var array $students @var array $terms @var int $studentID @var array $types */ ?>
<div class="row justify-content-center"><div class="col-lg-6">
<div class="card"><div class="card-header py-3"><i class="bi bi-shield-exclamation me-2 text-danger"></i>Record Incident / Commendation</div>
<div class="card-body">
  <form method="POST" action="/discipline">
    <div class="mb-3"><label class="form-label small fw-semibold">Student *</label>
      <select name="student_id" class="form-select" required>
        <option value="">Select student...</option>
        <?php foreach ($students as $s): ?>
          <option value="<?= $s['id'] ?>" <?= $s['id']==$studentID?'selected':'' ?>>
            <?= htmlspecialchars(trim(($s['first_name']??'').' '.($s['last_name']??''))) ?> (<?= htmlspecialchars($s['admission_no']??'') ?>)
          </option>
        <?php endforeach; ?>
      </select></div>
    <div class="row g-2 mb-3">
      <div class="col"><label class="form-label small fw-semibold">Term *</label>
        <select name="term_id" class="form-select" required>
          <option value="">Select...</option>
          <?php foreach ($terms as $t): ?><option value="<?= $t['id'] ?>" <?= ($t['is_current']??false)?'selected':'' ?>><?= htmlspecialchars($t['name']) ?></option><?php endforeach; ?>
        </select></div>
      <div class="col"><label class="form-label small fw-semibold">Date *</label>
        <input type="date" name="incident_date" class="form-control" required value="<?= date('Y-m-d') ?>"></div>
    </div>
    <div class="mb-3"><label class="form-label small fw-semibold">Type *</label>
      <select name="type" class="form-select" required>
        <?php foreach ($types as $val => $label): ?><option value="<?= $val ?>"><?= $label ?></option><?php endforeach; ?>
      </select></div>
    <div class="mb-3"><label class="form-label small fw-semibold">Description *</label>
      <textarea name="description" class="form-control" rows="3" required placeholder="Describe the incident..."></textarea></div>
    <div class="mb-3"><label class="form-label small fw-semibold">Action Taken</label>
      <textarea name="action_taken" class="form-control" rows="2" placeholder="What action was taken?"></textarea></div>
    <div class="d-flex gap-2">
      <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Save Record</button>
      <a href="/discipline" class="btn btn-outline-secondary">Cancel</a>
    </div>
  </form>
</div></div></div></div>
