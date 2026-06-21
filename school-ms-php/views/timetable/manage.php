<?php /** @var array $terms @var array $classes @var array $teachers @var array $subjects @var array $slots @var int $classID @var int $termID @var array $days */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-calendar-week me-2 text-primary"></i>Manage Timetable</h5>
  <a href="/timetable<?= $termID ? "?term_id={$termID}" : '' ?>" class="btn btn-sm btn-outline-secondary"><i class="bi bi-arrow-left me-1"></i>View Timetable</a>
</div>

<div class="row g-3">
  <!-- Filter -->
  <div class="col-12">
    <div class="card">
      <div class="card-body py-2">
        <form method="GET" class="row g-2 align-items-end">
          <div class="col-md-4">
            <label class="form-label small fw-semibold mb-1">Term *</label>
            <select name="term_id" class="form-select" required onchange="this.form.submit()">
              <option value="">Select term...</option>
              <?php foreach ($terms as $t): ?>
                <option value="<?= $t['id'] ?>" <?= $termID==$t['id']?'selected':'' ?>><?= htmlspecialchars($t['name']) ?></option>
              <?php endforeach; ?>
            </select>
          </div>
          <div class="col-md-4">
            <label class="form-label small fw-semibold mb-1">Class *</label>
            <select name="class_id" class="form-select" required onchange="this.form.submit()">
              <option value="">Select class...</option>
              <?php foreach ($classes as $c): ?>
                <option value="<?= $c['id'] ?>" <?= $classID==$c['id']?'selected':'' ?>><?= htmlspecialchars($c['name']) ?></option>
              <?php endforeach; ?>
            </select>
          </div>
        </form>
      </div>
    </div>
  </div>

  <!-- Add slot form -->
  <?php if ($classID && $termID): ?>
  <div class="col-lg-5">
    <div class="card">
      <div class="card-header py-3">Add Slot</div>
      <div class="card-body">
        <form method="POST" action="/timetable">
          <input type="hidden" name="class_id" value="<?= $classID ?>">
          <input type="hidden" name="term_id"  value="<?= $termID ?>">
          <div class="mb-2">
            <label class="form-label small fw-semibold">Day</label>
            <select name="day_of_week" class="form-select form-select-sm" required>
              <?php foreach ($days as $n => $l): ?><option value="<?= $n ?>"><?= $l ?></option><?php endforeach; ?>
            </select>
          </div>
          <div class="row g-2 mb-2">
            <div class="col"><label class="form-label small fw-semibold">Start</label>
              <input type="time" name="start_time" class="form-control form-control-sm" required></div>
            <div class="col"><label class="form-label small fw-semibold">End</label>
              <input type="time" name="end_time" class="form-control form-control-sm" required></div>
          </div>
          <div class="mb-2">
            <label class="form-label small fw-semibold">Subject</label>
            <select name="subject_id" class="form-select form-select-sm" required>
              <option value="">Select subject...</option>
              <?php foreach ($subjects as $s): ?><option value="<?= $s['id'] ?>"><?= htmlspecialchars($s['subject_name'] ?? $s['name']) ?></option><?php endforeach; ?>
            </select>
          </div>
          <div class="mb-2">
            <label class="form-label small fw-semibold">Teacher</label>
            <select name="teacher_id" class="form-select form-select-sm" required>
              <option value="">Select teacher...</option>
              <?php foreach ($teachers as $t): ?><option value="<?= $t['id'] ?>"><?= htmlspecialchars($t['name']) ?></option><?php endforeach; ?>
            </select>
          </div>
          <div class="mb-3">
            <label class="form-label small fw-semibold">Room <span class="text-muted fw-normal">(optional)</span></label>
            <input type="text" name="room" class="form-control form-control-sm" placeholder="e.g. Lab 1">
          </div>
          <button type="submit" class="btn btn-primary w-100"><i class="bi bi-plus-lg me-1"></i>Add Slot</button>
        </form>
      </div>
    </div>
  </div>

  <!-- Existing slots -->
  <div class="col-lg-7">
    <div class="card">
      <div class="card-header py-3">Current Slots — <?= count($slots) ?> total</div>
      <div class="card-body p-0">
        <table class="table table-sm table-hover mb-0">
          <thead><tr><th>Day</th><th>Time</th><th>Subject</th><th>Teacher</th><th>Room</th><th></th></tr></thead>
          <tbody>
            <?php if (empty($slots)): ?>
              <tr><td colspan="6" class="text-center text-muted py-3">No slots yet.</td></tr>
            <?php else: foreach ($slots as $s): ?>
            <tr>
              <td><span class="badge bg-primary-subtle text-primary"><?= $days[$s['day_of_week']] ?></span></td>
              <td class="small"><?= $s['start_time'] ?> – <?= $s['end_time'] ?></td>
              <td class="fw-semibold small"><?= htmlspecialchars($s['subject_name']) ?></td>
              <td class="small text-muted"><?= htmlspecialchars($s['teacher_name']) ?></td>
              <td class="small text-muted"><?= htmlspecialchars($s['room'] ?: '—') ?></td>
              <td>
                <form method="POST" action="/timetable/<?= $s['id'] ?>/delete" class="d-inline">
                  <input type="hidden" name="back" value="<?= htmlspecialchars($_SERVER['REQUEST_URI']) ?>">
                  <button type="submit" class="btn btn-xs btn-outline-danger" style="font-size:.7rem;padding:.15rem .4rem"
                          onclick="return confirm('Remove this slot?')"><i class="bi bi-trash"></i></button>
                </form>
              </td>
            </tr>
            <?php endforeach; endif; ?>
          </tbody>
        </table>
      </div>
    </div>
  </div>
  <?php else: ?>
  <div class="col-12"><div class="alert alert-info"><i class="bi bi-info-circle me-2"></i>Select a class and term to manage slots.</div></div>
  <?php endif; ?>
</div>
