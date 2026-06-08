<div class="row justify-content-center"><div class="col-lg-6">
<div class="card">
  <div class="card-header py-3"><i class="bi bi-pencil-square me-2 text-warning"></i>Create Exam</div>
  <div class="card-body">
    <form method="POST" action="/exams">
      <div class="mb-3"><label class="form-label small fw-semibold">Exam Name *</label><input type="text" name="name" class="form-control" required></div>
      <div class="mb-3"><label class="form-label small fw-semibold">Type</label>
        <select name="type" class="form-select"><option value="endterm">End Term</option><option value="midterm">Mid Term</option><option value="cat">CAT</option><option value="mock">Mock</option><option value="opener">Opener</option></select>
      </div>
      <div class="mb-3"><label class="form-label small fw-semibold">Term ID *</label><input type="number" name="term_id" class="form-control" required></div>
      <div class="row g-2 mb-3">
        <div class="col"><label class="form-label small fw-semibold">Start Date</label><input type="date" name="start_date" class="form-control" required></div>
        <div class="col"><label class="form-label small fw-semibold">End Date</label><input type="date" name="end_date" class="form-control" required></div>
      </div>
      <div class="d-flex gap-2"><button type="submit" class="btn btn-primary">Create Exam</button><a href="/exams" class="btn btn-outline-secondary">Cancel</a></div>
    </form>
  </div>
</div>
</div></div>