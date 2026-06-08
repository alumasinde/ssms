<?php
/** @var array $classes */
?>

<div class="row justify-content-center"><div class="col-lg-7">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-person-plus me-2 text-primary"></i>Enrol New Student</div>
  <div class="card-body">
    <form method="POST" action="/students">
      <div class="row g-3">
        <div class="col-md-6"><label class="form-label small fw-semibold">Admission No *</label><input type="text" name="admission_no" class="form-control" required></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">First Name *</label><input type="text" name="first_name" class="form-control" required></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Middle Name</label><input type="text" name="middle_name" class="form-control"></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Last Name *</label><input type="text" name="last_name" class="form-control" required></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Gender</label><select name="gender" class="form-select"><option value="">Select</option><option>male</option><option>female</option><option>other</option></select></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Date of Birth</label><input type="date" name="dob" class="form-control"></div>
        <div class="col-12"><label class="form-label small fw-semibold">Class *</label><select name="class_id" class="form-select" required><option value="">Select Class</option><?php foreach ($classes as $c): ?><option value="<?= $c['id'] ?>"><?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?></option><?php endforeach; ?></select></div>
      </div>
      <div class="mt-4 d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Enrol Student</button>
        <a href="/students" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
