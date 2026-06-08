<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-megaphone me-2 text-warning"></i>Notices</h5>
  <a href="/notices/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Post Notice</a>
</div>
<div class="row g-3">
  <?php if (empty($notices)): ?>
    <div class="col"><div class="card"><div class="card-body text-center text-muted py-4">No notices yet.</div></div></div>
  <?php else: foreach ($notices as $n): ?>
  <div class="col-12">
    <div class="card">
      <div class="card-body py-3">
        <div class="d-flex justify-content-between align-items-start">
          <div>
            <h6 class="fw-bold mb-1"><?= htmlspecialchars($n['title']) ?></h6>
            <p class="text-muted small mb-0"><?= htmlspecialchars(substr($n['body'] ?? '', 0, 120)) ?>...</p>
          </div>
          <div class="text-end ms-3 flex-shrink-0">
            <span class="badge bg-secondary-subtle text-secondary"><?= $n['audience'] ?></span>
            <div class="text-muted mt-1" style="font-size:.72rem"><?= date('d M Y', strtotime($n['published_at'])) ?></div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <?php endforeach; endif; ?>
</div>