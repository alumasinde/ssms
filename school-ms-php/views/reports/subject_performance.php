<?php /** @var array $rows @var array $exams @var int $examID */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-graph-up me-2 text-primary"></i>Subject Performance Analysis</h5>
</div>
<div class="card mb-3">
  <div class="card-body py-2">
    <form method="GET" class="row g-2 align-items-end">
      <div class="col-md-5">
        <label class="form-label small fw-semibold mb-1">Exam</label>
        <select name="exam_id" class="form-select" onchange="this.form.submit()">
          <option value="">Select Exam...</option>
          <?php foreach ($exams as $e): ?>
            <option value="<?= $e['id'] ?>" <?= $examID == $e['id'] ? 'selected' : '' ?>><?= htmlspecialchars($e['name']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
    </form>
  </div>
</div>
<?php if ($examID && empty($rows)): ?>
  <div class="alert alert-info">No results found for this exam.</div>
<?php elseif (!empty($rows)): ?>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead><tr><th>Subject</th><th class="text-center">Code</th><th class="text-center">Students</th><th class="text-center">Avg %</th><th class="text-center">Min %</th><th class="text-center">Max %</th><th class="text-center">Pass Rate</th></tr></thead>
      <tbody>
        <?php foreach ($rows as $r):
          $passRate = $r['entry_count'] > 0 ? round($r['pass_count']/$r['entry_count']*100,1) : 0;
        ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($r['subject_name']) ?></td>
          <td class="text-center"><span class="badge bg-primary-subtle text-primary"><?= htmlspecialchars($r['subject_code']) ?></span></td>
          <td class="text-center"><?= $r['entry_count'] ?></td>
          <td class="text-center">
            <span class="badge <?= $r['avg_score'] >= 50 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>">
              <?= $r['avg_score'] ?>%
            </span>
          </td>
          <td class="text-center text-muted"><?= $r['min_score'] ?>%</td>
          <td class="text-center text-muted"><?= $r['max_score'] ?>%</td>
          <td class="text-center">
            <div class="d-flex align-items-center gap-2 justify-content-center">
              <div class="progress" style="width:80px;height:6px">
                <div class="progress-bar <?= $passRate >= 50 ? 'bg-success' : 'bg-danger' ?>" style="width:<?= $passRate ?>%"></div>
              </div>
              <span class="small"><?= $passRate ?>%</span>
            </div>
          </td>
        </tr>
        <?php endforeach; ?>
      </tbody>
    </table>
  </div>
</div>
<?php endif; ?>
