<?php
/** @var array $users */
?>

<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-person-plus me-2"></i>Add Parent/Guardian</div>
  <div class="card-body">
    <form method="POST" action="/parents">
      <div class="mb-3">
        <label class="form-label small fw-semibold">User Account *</label>
        <select name="user_id" class="form-select" required>
          <option value="">Select user account...</option>
          <?php foreach ($users as $u): ?>
            <option value="<?= $u['id'] ?>">
              <?= htmlspecialchars($u['name']) ?> (<?= htmlspecialchars($u['email']) ?>)
            </option>
          <?php endforeach; ?>
        </select>
        <div class="form-text">The parent must have a user account first.</div>
      </div>
      <div class="row g-3">
        <div class="col-md-6">
          <label class="form-label small fw-semibold">Phone</label>
          <input type="tel" name="phone" class="form-control" placeholder="+254...">
        </div>
        <div class="col-md-6">
          <label class="form-label small fw-semibold">Occupation</label>
          <input type="text" name="occupation" class="form-control">
        </div>
        <div class="col-12">
          <label class="form-label small fw-semibold">Address</label>
          <textarea name="address" class="form-control" rows="2"></textarea>
        </div>
      </div>
      <div class="mt-4 d-flex gap-2">
        <button type="submit" class="btn btn-primary">Add Parent</button>
        <a href="/parents" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>