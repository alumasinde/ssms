<?php
/** @var array $terms @var array $classes @var array $slots @var array $grid @var int $termID @var int $classID @var array $days */
$dayColors = [1=>'primary',2=>'success',3=>'warning',4=>'info',5=>'danger'];
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-calendar-week me-2 text-primary"></i>Timetable</h5>
  <?php if (\Core\Session::can('timetable.edit')): ?>
    <a href="/timetable/manage<?= $classID ? "?class_id={$classID}&term_id={$termID}" : '' ?>"
       class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Manage Slots</a>
  <?php endif; ?>
</div>

<div class="card mb-3">
  <div class="card-body py-2">
    <form method="GET" class="row g-2 align-items-end">
      <div class="col-md-4">
        <label class="form-label small fw-semibold mb-1">Term</label>
        <select name="term_id" class="form-select form-select-sm" onchange="this.form.submit()">
          <option value="">Select term...</option>
          <?php foreach ($terms as $t): ?>
            <option value="<?= $t['id'] ?>" <?= $termID==$t['id']?'selected':'' ?>>
              <?= htmlspecialchars($t['name']) ?><?= ($t['is_current']??false)?' (Current)':'' ?>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
      <div class="col-md-4">
        <label class="form-label small fw-semibold mb-1">Class</label>
        <select name="class_id" class="form-select form-select-sm" onchange="this.form.submit()">
          <option value="">All classes</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>" <?= $classID==$c['id']?'selected':'' ?>><?= htmlspecialchars($c['name']) ?></option>
          <?php endforeach; ?>
        </select>
      </div>
    </form>
  </div>
</div>

<?php if (!$termID): ?>
<div class="alert alert-info"><i class="bi bi-info-circle me-2"></i>Select a term to view the timetable.</div>
<?php elseif (empty($slots)): ?>
<div class="alert alert-warning"><i class="bi bi-exclamation-triangle me-2"></i>No slots configured yet.
  <?php if (\Core\Session::can('timetable.edit')): ?><a href="/timetable/manage?term_id=<?= $termID ?>">Add slots.</a><?php endif; ?>
</div>
<?php else: ?>

<?php
// Group slots by class for display
$byClass = [];
foreach ($slots as $s) { $byClass[$s['class_id']]['name'] = $s['class_name']; $byClass[$s['class_id']]['slots'][$s['day_of_week']][] = $s; }
ksort($byClass);
foreach ($byClass as $cid => $cdata):
?>
<div class="card mb-4">
  <div class="card-header py-2 fw-semibold">
    <i class="bi bi-door-open me-1 text-primary"></i><?= htmlspecialchars($cdata['name']) ?>
    <?php if (\Core\Session::can('timetable.edit')): ?>
      <a href="/timetable/manage?class_id=<?= $cid ?>&term_id=<?= $termID ?>" class="btn btn-xs btn-outline-primary ms-2" style="font-size:.7rem;padding:.15rem .5rem">Edit</a>
    <?php endif; ?>
  </div>
  <div class="card-body p-0" style="overflow-x:auto">
    <table class="table table-bordered mb-0 text-center" style="min-width:700px">
      <thead class="table-light">
        <tr><?php foreach ($days as $dn => $dl): ?><th class="text-<?= $dayColors[$dn] ?>"><?= $dl ?></th><?php endforeach; ?></tr>
      </thead>
      <tbody>
        <!-- Find max slots per day -->
        <?php
        $maxRows = 1;
        foreach ($days as $dn => $_) { $maxRows = max($maxRows, count($cdata['slots'][$dn] ?? [])); }
        for ($row = 0; $row < $maxRows; $row++):
        ?>
        <tr>
          <?php foreach ($days as $dn => $_):
            $slot = $cdata['slots'][$dn][$row] ?? null;
          ?>
          <td class="align-top p-2" style="min-width:130px">
            <?php if ($slot): ?>
              <div class="badge bg-<?= $dayColors[$dn] ?>-subtle text-<?= $dayColors[$dn] ?> border border-<?= $dayColors[$dn] ?>-subtle w-100 text-wrap p-2" style="font-size:.75rem">
                <div class="fw-bold"><?= htmlspecialchars($slot['subject_name']) ?></div>
                <div class="opacity-75"><?= $slot['start_time'] ?> – <?= $slot['end_time'] ?></div>
                <div class="opacity-75"><?= htmlspecialchars($slot['teacher_name']) ?></div>
                <?php if ($slot['room']): ?><div class="opacity-60"><?= htmlspecialchars($slot['room']) ?></div><?php endif; ?>
                <?php if (\Core\Session::can('timetable.edit')): ?>
                  <form method="POST" action="/timetable/<?= $slot['id'] ?>/delete" class="mt-1">
                    <input type="hidden" name="back" value="<?= htmlspecialchars($_SERVER['REQUEST_URI']) ?>">
                    <button type="submit" class="btn btn-xs btn-outline-danger" style="font-size:.6rem;padding:.1rem .4rem"
                            onclick="return confirm('Remove this slot?')">✕</button>
                  </form>
                <?php endif; ?>
              </div>
            <?php else: ?>
              <div class="text-muted" style="font-size:.75rem">—</div>
            <?php endif; ?>
          </td>
          <?php endforeach; ?>
        </tr>
        <?php endfor; ?>
      </tbody>
    </table>
  </div>
</div>
<?php endforeach; ?>
<?php endif; ?>
