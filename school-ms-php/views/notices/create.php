<div class="row justify-content-center"><div class="col-lg-7">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-megaphone me-2 text-warning"></i>Post Notice</div>
  <div class="card-body">
    <form method="POST" action="/notices">
      <div class="mb-3"><label class="form-label small fw-semibold">Title *</label><input type="text" name="title" class="form-control" required></div>
      <div class="mb-3"><label class="form-label small fw-semibold">Body *</label><textarea name="body" class="form-control" rows="5" required></textarea></div>
      <div class="mb-4"><label class="form-label small fw-semibold">Audience</label>
        <select name="audience" class="form-select">
          <option value="all">Everyone</option>
          <option value="teachers">Teachers</option>
          <option value="parents">Parents</option>
          <option value="students">Students</option>
        </select>
      </div>
      <div class="d-flex gap-2">
        <button type="submit" class="btn btn-primary"><i class="bi bi-send me-1"></i>Publish</button>
        <a href="/notices" class="btn btn-outline-secondary">Cancel</a>
      </div>
    </form>
  </div>
</div>
</div></div>
