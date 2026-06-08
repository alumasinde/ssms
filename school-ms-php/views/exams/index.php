<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-pencil-square me-2 text-warning"></i>Exams</h5>
  <a href="/exams/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Create Exam</a>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Name</th><th>Type</th><th>Start</th><th>End</th><th></th></tr></thead>
      <tbody>
        <?php if (empty($exams)): ?>
          <tr><td colspan="5" class="text-center text-muted py-4">No exams yet.</td></tr>
        <?php else: foreach ($exams as $e): ?>
          <tr>
            <td class="fw-semibold"><?= htmlspecialchars($e['name']) ?></td>
            <td><span class="badge bg-warning-subtle text-warning"><?= $e['type'] ?></span></td>
            <td><?= $e['start_date'] ?></td><td><?= $e['end_date'] ?></td>
            <td class="text-end">
              <a href="/exams/<?= $e['id'] ?>/results" class="btn btn-sm btn-outline-primary"><i class="bi bi-list-ol me-1"></i>Results</a>
              <a href="/reports/class-results?exam_id=<?= $e['id'] ?>" class="btn btn-sm btn-outline-info"><i class="bi bi-bar-chart me-1"></i>Report</a>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>