<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-list-ol me-2"></i><?= htmlspecialchars($exam['name'] ?? 'Exam') ?> — Results</h5>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Student</th><th>Subject</th><th>Marks</th><th>Max</th><th>Grade</th><th>Remarks</th></tr></thead>
      <tbody>
        <?php if (empty($results)): ?>
          <tr><td colspan="6" class="text-center text-muted py-4">No results submitted yet.</td></tr>
        <?php else: foreach ($results as $r): ?>
          <tr>
            <td><?= $r['student_id'] ?></td>
            <td><?= $r['subject_id'] ?></td>
            <td class="fw-bold"><?= $r['marks'] ?></td>
            <td class="text-muted"><?= $r['max_marks'] ?></td>
            <td><span class="badge bg-primary-subtle text-primary"><?= $r['grade'] ?></span></td>
            <td class="text-muted small"><?= htmlspecialchars($r['remarks'] ?? '') ?></td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>