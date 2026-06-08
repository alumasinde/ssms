<?php $c = $card ?? []; ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-file-earmark-text me-2 text-info"></i>Report Card</h5>
  <button onclick="window.print()" class="btn btn-outline-secondary btn-sm"><i class="bi bi-printer me-1"></i>Print</button>
</div>
<div class="card mb-3">
  <div class="card-body">
    <div class="row">
      <div class="col-md-6">
        <p class="mb-1"><strong>Student:</strong> <?= htmlspecialchars($c['student_name'] ?? '') ?></p>
        <p class="mb-1"><strong>Admission No:</strong> <?= htmlspecialchars($c['admission_no'] ?? '') ?></p>
        <p class="mb-1"><strong>Class:</strong> <?= htmlspecialchars($c['class_name'] ?? '') ?></p>
      </div>
      <div class="col-md-6">
        <p class="mb-1"><strong>Exam:</strong> <?= htmlspecialchars($c['exam_name'] ?? '') ?></p>
        <p class="mb-1"><strong>Term:</strong> <?= htmlspecialchars($c['term_name'] ?? '') ?></p>
        <p class="mb-1"><strong>Position:</strong> <?= $c['position'] ?? '—' ?> / <?= $c['class_size'] ?? '—' ?></p>
      </div>
    </div>
  </div>
</div>
<div class="card mb-3">
  <div class="card-body p-0">
    <table class="table mb-0">
      <thead><tr><th>Subject</th><th>Marks</th><th>Max</th><th>%</th><th>Grade</th><th>Remarks</th></tr></thead>
      <tbody>
        <?php foreach ($c['results'] ?? [] as $r): ?>
        <tr>
          <td><?= htmlspecialchars($r['subject_name']) ?></td>
          <td class="fw-bold"><?= $r['marks'] ?></td>
          <td class="text-muted"><?= $r['max_marks'] ?></td>
          <td><?= round($r['percentage'] ?? 0, 1) ?>%</td>
          <td><span class="badge bg-primary-subtle text-primary"><?= $r['grade'] ?></span></td>
          <td class="text-muted small"><?= htmlspecialchars($r['remarks'] ?? '') ?></td>
        </tr>
        <?php endforeach; ?>
        <tr class="table-light fw-bold">
          <td>TOTAL</td>
          <td><?= $c['total_marks'] ?? 0 ?></td>
          <td><?= $c['total_max'] ?? 0 ?></td>
          <td><?= round($c['average'] ?? 0, 1) ?>%</td>
          <td><span class="badge bg-success"><?= $c['overall_grade'] ?? '—' ?></span></td>
          <td>Attendance: <?= $c['attendance_pct'] ?? 0 ?>%</td>
        </tr>
      </tbody>
    </table>
  </div>
</div>