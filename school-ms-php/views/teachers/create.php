<?php /** @var array $users */ ?>
<div class="row justify-content-center"><div class="col-lg-8">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-person-plus me-2"></i>Add Teacher</div>
  <div class="card-body">
    <form method="POST" action="/teachers">
      <div class="mb-3">
        <label class="form-label small fw-semibold">User Account *</label>
        <select name="user_id" class="form-select" required>
          <option value="">Select user...</option>
          <?php foreach ($users as $u): ?>
            <option value="<?= $u['id'] ?>"><?= htmlspecialchars($u['name']) ?> (<?= htmlspecialchars($u['email']) ?>)</option>
          <?php endforeach; ?>
        </select>
        <div class="form-text">User must be created first with role = teacher.</div>
      </div>
      <div class="row g-3">
        <div class="col-md-6"><label class="form-label small fw-semibold">Employee No *</label><input type="text" name="employee_no" class="form-control" required></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">TSC No</label><input type="text" name="tsc_no" class="form-control" placeholder="e.g. TSC/12345"></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Phone</label><input type="tel" name="phone" class="form-control" placeholder="+254..."></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Gender</label>
          <select name="gender" class="form-select"><option value="">Select</option><option>male</option><option>female</option><option>other</option></select>
        </div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Date of Birth</label><input type="date" name="dob" class="form-control"></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Hire Date</label><input type="date" name="hire_date" class="form-control"></div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Employment Type</label>
          <select name="employment_type" class="form-select">
            <option value="permanent">Permanent</option>
            <option value="contract">Contract</option>
            <option value="part_time">Part Time</option>
          </select>
        </div>
        <div class="col-md-6"><label class="form-label small fw-semibold">Qualification</label><input type="text" name="qualification" class="form-control" placeholder="e.g. B.Ed Mathematics"></div>
        <div class="col-12"><label class="form-label small fw-semibold">Specialization</label><input type="text" name="specialization" class="form-control" placeholder="e.g. Mathematics & Physics"></div>
      </div>
      <div class="mt-4 d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Add Teacher</button>
        <a href="/teachers" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
