<div class="row g-3">
  <div class="col-lg-4">
    <div class="card text-center p-3">
      <div class="rounded-circle bg-info mx-auto d-flex align-items-center justify-content-center text-white mb-3"
           style="width:72px;height:72px;font-size:2rem;font-weight:700">
        <?= strtoupper(substr($parent['name'] ?? 'P', 0, 1)) ?>
      </div>
      <h5 class="fw-bold mb-0"><?= htmlspecialchars($parent['name'] ?? '') ?></h5>
      <div class="text-muted small"><?= htmlspecialchars($parent['email'] ?? '') ?></div>
    </div>
    <div class="card mt-3">
      <div class="card-header small fw-semibold">Details</div>
      <div class="list-group list-group-flush small">
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Phone</span><span><?= htmlspecialchars($parent['phone'] ?? '—') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Occupation</span><span><?= htmlspecialchars($parent['occupation'] ?? '—') ?></span>
        </div>
        <div class="list-group-item d-flex justify-content-between">
          <span class="text-muted">Address</span><span><?= htmlspecialchars($parent['address'] ?? '—') ?></span>
        </div>
      </div>
    </div>
  </div>
  <div class="col-lg-8">
    <div class="card">
      <div class="card-header py-3 d-flex justify-content-between align-items-center">
        <span><i class="bi bi-person-badge me-1"></i>Linked Students</span>
        <button class="btn btn-sm btn-outline-primary" data-bs-toggle="modal" data-bs-target="#linkModal">
          <i class="bi bi-link me-1"></i>Link Student
        </button>
      </div>
      <div class="card-body text-muted small">
        Link students to this parent using the button above.
      </div>
    </div>
  </div>
</div>

<div class="modal fade" id="linkModal" tabindex="-1">
  <div class="modal-dialog modal-sm">
    <div class="modal-content">
      <div class="modal-header"><h6 class="modal-title">Link Student</h6>
        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
      </div>
      <div class="modal-body">
        <form method="POST" action="/parents/link-student">
          <input type="hidden" name="parent_id" value="<?= $parent['id'] ?>">
          <div class="mb-3">
            <label class="form-label small fw-semibold">Student ID</label>
            <input type="number" name="student_id" class="form-control" required>
          </div>
          <div class="mb-3">
            <label class="form-label small fw-semibold">Relationship</label>
            <select name="relationship" class="form-select">
              <option value="father">Father</option>
              <option value="mother">Mother</option>
              <option value="guardian">Guardian</option>
              <option value="other">Other</option>
            </select>
          </div>
          <button type="submit" class="btn btn-primary w-100">Link</button>
        </form>
      </div>
    </div>
  </div>
</div>