<div class="row justify-content-center"><div class="col-lg-5">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-pencil me-2"></i>Edit Subject</div>
  <div class="card-body">
    <form method="POST" action="/subjects/<?= $subject['id'] ?>/update">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Subject Name *</label>
        <input type="text" name="name" class="form-control"
               value="<?= htmlspecialchars($subject['name'] ?? '') ?>" required>
      </div>
      <div class="mb-4">
        <label class="form-label small fw-semibold">Subject Code *</label>
        <input type="text" name="code" class="form-control"
               value="<?= htmlspecialchars($subject['code'] ?? '') ?>" required>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary">Save Changes</button>
        <a href="/subjects" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>