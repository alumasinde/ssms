<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-bar-chart me-2 text-success"></i>Attendance Summary</h5>
</div>

<div class="card mb-3">
  <div class="card-body">
    <form method="GET" action="/attendance/summary" class="row g-2 align-items-end">
      <div class="col-md-4">
        <label class="form-label small fw-semibold">Class</label>
        <select name="class_id" class="form-select">
          <option value="">Select Class</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>"
              <?= $classID == $c['id'] ? 'selected' : '' ?>>
              <?= htmlspecialchars($c['name']) ?>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-4">
        <label class="form-label small fw-semibold">Term ID</label>
        <input type="number" name="term_id" class="form-control"
               value="<?= $termID ?>" placeholder="Enter term ID">
      </div>
      <div class="col-md-4">
        <button type="submit" class="btn btn-success w-100">
          <i class="bi bi-search me-1"></i>View Summary
        </button>
      </div>
    </form>
  </div>
</div>

<?php if (!empty($summary)): ?>
<div class="card">
  <div class="card-body p-0">
    <table class="table table-hover mb-0">
      <thead>
        <tr>
          <th>Student</th>
          <th class="text-center">Total Days</th>
          <th class="text-center">Present</th>
          <th class="text-center">Absent</th>
          <th class="text-center">Late</th>
          <th>Attendance %</th>
        </tr>
      </thead>
      <tbody>
        <?php foreach ($summary as $row): ?>
        <tr>
          <td class="fw-semibold"><?= htmlspecialchars($row['student_name']) ?></td>
          <td class="text-center"><?= $row['total'] ?></td>
          <td class="text-center text-success fw-bold"><?= $row['present'] ?></td>
          <td class="text-center text-danger"><?= $row['absent'] ?></td>
          <td class="text-center text-warning"><?= $row['late'] ?></td>
          <td>
            <div class="d-flex align-items-center gap-2">
              <div class="progress flex-grow-1" style="height:8px">
                <div class="progress-bar
                  <?= $row['percent'] >= 80 ? 'bg-success' : ($row['percent'] >= 60 ? 'bg-warning' : 'bg-danger') ?>"
                  style="width:<?= $row['percent'] ?>%"></div>
              </div>
              <span class="small fw-semibold"><?= $row['percent'] ?>%</span>
            </div>
          </td>
        </tr>
        <?php endforeach; ?>
      </tbody>
    </table>
  </div>
</div>
<?php elseif ($classID && $termID): ?>
<div class="alert alert-info">No attendance records found for the selected class and term.</div>
<?php endif; ?>