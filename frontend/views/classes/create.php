<div class="row justify-content-center"><div class="col-lg-5">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-plus-lg me-2 text-primary"></i>Create Class</div>
  <div class="card-body">
    <form method="POST" action="/classes">
      <div class="mb-3">
        <label class="form-label small fw-semibold">Class Name *</label>
        <input type="text" name="name" class="form-control" placeholder="e.g. Form 1 East" required>
      </div>
      <div class="mb-3">
        <label class="form-label small fw-semibold">Level *</label>
        <input type="text" name="level" class="form-control" placeholder="e.g. Form 1" required>
      </div>
      <div class="mb-4">
        <label class="form-label small fw-semibold">Stream</label>
        <input type="text" name="stream" class="form-control" placeholder="e.g. East">
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary">Create Class</button>
        <a href="/classes" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>