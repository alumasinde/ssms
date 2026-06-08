<div class="row g-3">
  <div class="col-lg-4">
    <div class="card text-center p-3">
      <div class="rounded-circle bg-success mx-auto d-flex align-items-center justify-content-center text-white mb-3" style="width:72px;height:72px;font-size:2rem;font-weight:700"><?= strtoupper(substr($teacher['name'] ?? 'T', 0, 1)) ?></div>
      <h5 class="fw-bold mb-0"><?= htmlspecialchars($teacher['name'] ?? '') ?></h5>
      <div class="text-muted small"><?= htmlspecialchars($teacher['email'] ?? '') ?></div>
      <div class="mt-2"><span class="badge bg-success-subtle text-success"><?= htmlspecialchars($teacher['employee_no'] ?? '') ?></span></div>
    </div>
    <div class="card mt-3">
      <div class="card-header small fw-semibold">Details</div>
      <div class="list-group list-group-flush small">
        <div class="list-group-item d-flex justify-content-between"><span class="text-muted">Phone</span><span><?= $teacher['phone'] ?? '—' ?></span></div>
        <div class="list-group-item d-flex justify-content-between"><span class="text-muted">Gender</span><span><?= ucfirst($teacher['gender'] ?? '—') ?></span></div>
        <div class="list-group-item d-flex justify-content-between"><span class="text-muted">Qualification</span><span><?= htmlspecialchars($teacher['qualification'] ?? '—') ?></span></div>
      </div>
    </div>
  </div>
  <div class="col-lg-8">
    <div class="card">
      <div class="card-header py-3"><i class="bi bi-book me-1"></i>Assigned Subjects</div>
      <div class="list-group list-group-flush">
        <?php if (empty($subjects)): ?>
          <div class="list-group-item text-muted small py-3">No subjects assigned.</div>
        <?php else: foreach ($subjects as $s): ?>
          <div class="list-group-item small">
            <div class="fw-semibold">Subject ID: <?= $s['subject_id'] ?></div>
            <div class="text-muted">Class ID: <?= $s['class_id'] ?></div>
          </div>
        <?php endforeach; endif; ?>
      </div>
    </div>
  </div>
</div>
