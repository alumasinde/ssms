<div class="row justify-content-center"><div class="col-lg-5">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-pencil me-2"></i>Edit Class</div>
  <div class="card-body">
    <form method="POST" action="/classes/<?= $class['id'] ?>/update">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Class Name *</label>
        <input type="text" name="name" class="form-control"
               value="<?= htmlspecialchars($class['name'] ?? '') ?>" required>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Level *</label>
        <input type="text" name="level" class="form-control"
               value="<?= htmlspecialchars($class['level'] ?? '') ?>" required>
      </div>
      <div class="mb-4">
        <label class="form-label small fw-semibold">Stream</label>
        <input type="text" name="stream" class="form-control"
               value="<?= htmlspecialchars($class['stream'] ?? '') ?>">
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary">Save Changes</button>
        <a href="/classes" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>