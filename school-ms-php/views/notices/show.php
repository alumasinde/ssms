<div class="row justify-content-center"><div class="col-lg-8">
<div class="card">
  <div class="card-header py-3 d-flex justify-content-between align-items-center">
    <span><i class="bi bi-megaphone me-2 text-warning"></i><?= htmlspecialchars($notice['title'] ?? '') ?></span>
    <span class="badge bg-secondary-subtle text-secondary"><?= $notice['audience'] ?? '' ?></span>
  </div>
  <div class="card-body">
    <p class="text-muted small mb-3">Published: <?= date('d M Y H:i', strtotime($notice['published_at'] ?? 'now')) ?></p>
    <div style="line-height:1.8"><?= nl2br(htmlspecialchars($notice['body'] ?? '')) ?></div>
  </div>
  <div class="card-footer bg-white">
    <a href="/notices" class="btn btn-outline-secondary btn-sm"><i class="bi bi-arrow-left me-1"></i>Back</a>
  </div>
</div>
</div></div>
