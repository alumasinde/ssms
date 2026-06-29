<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-bar-chart-line me-2 text-primary"></i>Class Results</h5>
  <form method="GET" class="d-flex gap-2">
    <select name="exam_id" class="form-select form-select-sm" onchange="this.form.submit()">
      <option value="">Select Exam</option>
      <?php foreach ($exams as $e): ?>
        <option value="<?= $e['id'] ?>" <?= $examID == $e['id'] ? 'selected' : '' ?>><?= htmlspecialchars($e['name']) ?></option>
      <?php endforeach; ?>
    </select>
  </form>
</div>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>#</th><th>Adm No</th><th>Name</th><th>Total Marks</th><th>Average %</th></tr></thead>
      <tbody>
        <?php if (empty($results)): ?>
          <tr><td colspan="5" class="text-center text-muted py-4">Select an exam to view results.</td></tr>
        <?php else: foreach ($results as $r): ?>
          <tr>
            <td><span class="badge bg-primary"><?= $r['position'] ?></span></td>
            <td><?= htmlspecialchars($r['admission_no']) ?></td>
            <td class="fw-semibold"><?= htmlspecialchars($r['student_name']) ?></td>
            <td><?= $r['total_marks'] ?> / <?= $r['total_max'] ?></td>
            <td>
              <div class="d-flex align-items-center gap-2">
                <div class="progress flex-grow-1" style="height:8px">
                  <div class="progress-bar" style="width:<?= $r['average'] ?>%"></div>
                </div>
                <span class="small"><?= $r['average'] ?>%</span>
              </div>
            </td>
          </tr>
        <?php endforeach; endif; ?>
      </tbody>
    </table>
  </div>
</div>