<?php
/** @var array $student */
/** @var array $classes */
?>

<div class="row justify-content-center"><div class="col-lg-7">
<div class="card">
  <div class="card-header py-3 d-flex justify-content-between align-items-center">
    <span><i class="bi bi-pencil me-2 text-primary"></i>Edit Student</span>
    <a href="/students/<?= $student['id'] ?>" class="btn btn-sm btn-outline-secondary">
      <i class="bi bi-arrow-left me-1"></i>Back
    </a>
  </div>
  <div class="card-body">
    <form method="POST" action="/students/<?= $student['id'] ?>/update">
      <div class="row g-3">
        <div class="col-md-6">
          <label class="form-label small fw-semibold">First Name *</label>
          <input type="text" name="first_name" class="form-control"
                 value="<?= htmlspecialchars($student['first_name'] ?? '') ?>" required>
        </div>
        <div class="col-md-6">
          <label class="form-label small fw-semibold">Middle Name</label>
          <input type="text" name="middle_name" class="form-control"
                 value="<?= htmlspecialchars($student['middle_name'] ?? '') ?>">
        </div>
        <div class="col-12">
          <label class="form-label small fw-semibold">Last Name *</label>
          <input type="text" name="last_name" class="form-control"
                 value="<?= htmlspecialchars($student['last_name'] ?? '') ?>" required>
        </div>
        <div class="col-md-6">
          <label class="form-label small fw-semibold">Admission No</label>
          <input type="text" class="form-control"
                 value="<?= htmlspecialchars($student['admission_no'] ?? '') ?>" disabled>
          <div class="form-text">Admission number cannot be changed.</div>
        </div>
        <div class="col-md-6">
          <label class="form-label small fw-semibold">Gender</label>
          <select name="gender" class="form-select">
            <option value="">Select</option>
            <?php foreach (['male','female','other'] as $g): ?>
              <option value="<?= $g ?>" <?= ($student['gender'] ?? '') === $g ? 'selected' : '' ?>>
                <?= ucfirst($g) ?>
              </option>
            <?php endforeach; ?>
          </select>
        </div>
        <div class="col-md-6">
          <label class="form-label small fw-semibold">Date of Birth</label>
          <input type="date" name="dob" class="form-control"
                 value="<?= htmlspecialchars($student['dob'] ?? '') ?>">
        </div>
        <div class="col-12">
          <label class="form-label small fw-semibold">Class *</label>
          <select name="class_id" class="form-select" required>
            <option value="">Select Class</option>
            <?php foreach ($classes as $c): ?>
              <option value="<?= $c['id'] ?>"
                <?= ($student['class_id'] ?? 0) == $c['id'] ? 'selected' : '' ?>>
                <?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?>
              </option>
            <?php endforeach; ?>
          </select>
        </div>
      </div>
      <div class="mt-4 d-flex gap-2">
        <button type="submit" class="btn btn-primary">
          <i class="bi bi-check-lg me-1"></i>Save Changes
        </button>
        <a href="/students/<?= $student['id'] ?>" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>