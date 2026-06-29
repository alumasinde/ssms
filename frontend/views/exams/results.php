<?php
/** @var array $exam     */
/** @var array $results  */
/** @var array $students */
/** @var array $subjects */
/** @var array $classes  */
/** @var int   $classID  */

use Core\Session;


$canGrade = Session::can('exams.results');
$examID   = $exam['id'] ?? 0;
$examClassID = $exam['class_id'] ?? null;
$userRole = Session::role(); // 'admin' or 'teacher'

// ── Index existing results keyed [student_id][subject_id] ─────────────────────
$existing = [];
foreach ($results as $r) {
    $existing[$r['student_id']][$r['subject_id']] = $r;
}

// ── Build student list ────────────────────────────────────────────────────────
$entryStudents = [];
foreach ($students as $s) {
    $entryStudents[$s['id']] = [
        'id'           => $s['id'],
        'name'         => trim(($s['first_name'] ?? '') . ' ' . ($s['last_name'] ?? '')),
        'admission_no' => $s['admission_no'] ?? '',
    ];
}
// Merge any extra students found only in results
foreach ($results as $r) {
    if (!isset($entryStudents[$r['student_id']])) {
        $entryStudents[$r['student_id']] = [
            'id'           => $r['student_id'],
            'name'         => $r['student_name'] ?? 'Unknown',
            'admission_no' => $r['admission_no'] ?? '',
        ];
    }
}
uasort($entryStudents, fn($a, $b) => strcmp($a['name'], $b['name']));

// ── Build subject list ────────────────────────────────────────────────────────
$entrySubjects = [];
foreach ($subjects as $s) {
    $entrySubjects[$s['id']] = [
        'id'   => $s['id'],
        'name' => $s['name'],
        'code' => $s['code'],
    ];
}

// ── Ranking for view mode ─────────────────────────────────────────────────────
$ranked = [];
foreach ($entryStudents as $st) {
    $total = 0; $max = 0;
    foreach ($entrySubjects as $sub) {
        $r = $existing[$st['id']][$sub['id']] ?? null;
        if ($r) { $total += $r['marks']; $max += $r['max_marks']; }
    }
    $ranked[] = [
        'st'    => $st,
        'total' => $total,
        'max'   => $max,
        'avg'   => $max > 0 ? round($total / $max * 100, 1) : null,
    ];
}
usort($ranked, fn($a, $b) => ($b['avg'] ?? -1) <=> ($a['avg'] ?? -1));

$mode = $_GET['mode'] ?? ($canGrade && empty($results) ? 'entry' : 'view');
?>

<style>
.marks-tbl th, .marks-tbl td { vertical-align: middle; white-space: nowrap; }
.marks-tbl thead th { background: #f1f5f9; font-size: .72rem; text-transform: uppercase; letter-spacing: .04em; }
.marks-tbl td.sticky-col { position: sticky; left: 0; background: #fff; z-index: 2; }
.marks-tbl thead th.sticky-col { background: #f1f5f9; z-index: 3; }
.marks-input { width: 64px; text-align: center; font-weight: 600; }
.marks-input:focus { border-color: #6366f1; box-shadow: 0 0 0 2px rgba(99,102,241,.15); }
.marks-input.over { border-color: #ef4444 !important; }
.grade-pill { font-size: .62rem; padding: 1px 5px; border-radius: 20px; }
</style>

<!-- ── Header ─────────────────────────────────────────────────────────────── -->
<div class="d-flex justify-content-between align-items-start mb-3 flex-wrap gap-2">
  <div>
    <nav aria-label="breadcrumb"><ol class="breadcrumb mb-0 small">
      <li class="breadcrumb-item"><a href="/exams">Exams</a></li>
      <li class="breadcrumb-item active"><?= htmlspecialchars($exam['name'] ?? '') ?></li>
    </ol></nav>
    <h5 class="fw-bold mb-0 mt-1">
      <i class="bi bi-list-ol me-2 text-warning"></i><?= htmlspecialchars($exam['name'] ?? '') ?>
      <span class="badge bg-warning-subtle text-warning border border-warning-subtle ms-1" style="font-size:.7rem">
        <?= ucfirst($exam['type'] ?? '') ?>
      </span>
    </h5>
    <div class="text-muted small mt-1">
      <?= $exam['start_date'] ?? '' ?> → <?= $exam['end_date'] ?? '' ?>
      &nbsp;·&nbsp; <?= count($entryStudents) ?> students
      &nbsp;·&nbsp; <?= count($entrySubjects) ?> subjects
    </div>
  </div>

  <div class="d-flex gap-2 flex-wrap align-items-start">
    <?php if ($canGrade): ?>
      <?php if ($mode === 'view'): ?>
        <a href="?<?= http_build_query(array_merge($_GET, ['mode'=>'entry'])) ?>"
           class="btn btn-primary btn-sm">
          <i class="bi bi-pencil-square me-1"></i><?= empty($results) ? 'Enter Marks' : 'Edit Marks' ?>
        </a>
      <?php else: ?>
        <a href="?<?= http_build_query(array_merge($_GET, ['mode'=>'view'])) ?>"
           class="btn btn-outline-secondary btn-sm">
          <i class="bi bi-eye me-1"></i>View Results
        </a>
      <?php endif; ?>
    <?php endif; ?>
    <a href="/exams" class="btn btn-outline-secondary btn-sm">
      <i class="bi bi-arrow-left me-1"></i>Back
    </a>
  </div>
</div>

<!-- ── Class filter ───────────────────────────────────────────────────────── -->
<?php if (!$examClassID && !empty($classes)): ?>
<div class="d-flex align-items-center gap-2 mb-3">
  <label class="form-label small fw-semibold mb-0 text-nowrap">Class:</label>
  <form method="GET" class="d-flex gap-2 align-items-center">
    <input type="hidden" name="mode" value="<?= htmlspecialchars($mode) ?>">
    <select name="class_id" class="form-select form-select-sm" style="max-width:220px"
            onchange="this.form.submit()">
      <option value="">All classes</option>
      <?php foreach ($classes as $c): ?>
        <option value="<?= $c['id'] ?>" <?= $classID == $c['id'] ? 'selected' : '' ?>>
          <?= htmlspecialchars($c['name']) ?> — <?= htmlspecialchars($c['level']) ?>
        </option>
      <?php endforeach; ?>
    </select>
  </form>
</div>
<?php endif; ?>

<?php if ($mode === 'entry' && $canGrade): ?>
<!-- ══════════════════════════════════════════════════════════
     MARK ENTRY
══════════════════════════════════════════════════════════ -->

<?php if (empty($entryStudents)): ?>
  <div class="alert alert-warning">
    <i class="bi bi-exclamation-triangle me-2"></i>
    No students in this class. <a href="/students">Enrol students</a> first.
  </div>
<?php elseif (empty($entrySubjects)): ?>
  <div class="alert alert-warning">
    <i class="bi bi-exclamation-triangle me-2"></i>
    No subjects assigned. <a href="/subjects">Set up subjects</a> first.
  </div>
<?php else: ?>

<form method="POST" action="/exams/results" id="marksForm">
  <input type="hidden" name="exam_id"  value="<?= $examID ?>">
  <input type="hidden" name="class_id" value="<?= $classID ?: ($examClassID ?? '') ?>">

  <div class="card">
    <div class="card-header py-2 d-flex justify-content-between align-items-center">
      <div class="d-flex align-items-center gap-3">
        <span class="fw-semibold small"><i class="bi bi-pencil-square me-1 text-primary"></i>Enter Marks</span>
        <!-- Global max setter -->
        <div class="d-flex align-items-center gap-1">
          <span class="text-muted small">Default max:</span>
          <input type="number" id="globalMax" value="100" min="1" max="1000"
                 class="form-control form-control-sm" style="width:65px">
          <button type="button" class="btn btn-outline-secondary btn-sm" onclick="applyGlobalMax()">
            Apply
          </button>
        </div>
      </div>
      <button type="submit" class="btn btn-success btn-sm">
        <i class="bi bi-check-lg me-1"></i>Save All
      </button>
    </div>

    <div style="overflow-x:auto">
      <table class="table table-bordered table-sm mb-0 marks-tbl"
             style="min-width:<?= 260 + count($entrySubjects) * 110 ?>px">
        <thead>
          <tr>
            <th class="sticky-col" style="min-width:160px">Student</th>
            <th style="min-width:80px">Adm No</th>
            <?php foreach ($entrySubjects as $sub): ?>
            <th class="text-center" style="min-width:110px">
              <?= htmlspecialchars($sub['code']) ?>
              <div class="text-muted fw-normal" style="font-size:.62rem;text-transform:none">
                <?= htmlspecialchars($sub['name']) ?>
              </div>
              <div class="mt-1 d-flex align-items-center justify-content-center gap-1">
                <span class="text-muted" style="font-size:.6rem">max</span>
                <input type="number" class="col-max form-control form-control-sm text-center p-0"
                       data-subid="<?= $sub['id'] ?>"
                       value="100" min="1" max="1000" style="width:42px;font-size:.7rem;height:20px"
                       onchange="applyColMax(this)">
              </div>
            </th>
            <?php endforeach; ?>
            <th class="text-center" style="min-width:70px">Total</th>
            <th class="text-center" style="min-width:60px">Avg%</th>
          </tr>
        </thead>
        <tbody>
          <?php $ri = 0; foreach ($entryStudents as $st): ?>
          <tr>
            <td class="sticky-col fw-semibold small"><?= htmlspecialchars($st['name']) ?></td>
            <td class="text-muted small"><?= htmlspecialchars($st['admission_no']) ?></td>

            <?php foreach ($entrySubjects as $sub):
              $prev    = $existing[$st['id']][$sub['id']] ?? null;
              $prevM   = $prev['marks']     ?? '';
              $prevMax = $prev['max_marks'] ?? 100;
              $prevRem = $prev['remarks']   ?? '';
            ?>
            <td class="text-center p-1">
              <input type="hidden" name="rows[<?= $ri ?>][student_id]" value="<?= $st['id'] ?>">
              <input type="hidden" name="rows[<?= $ri ?>][subject_id]" value="<?= $sub['id'] ?>">
              <input type="hidden" name="rows[<?= $ri ?>][remarks]"    value="<?= htmlspecialchars($prevRem) ?>">
              <input type="hidden" name="rows[<?= $ri ?>][max_marks]"
                     class="max-h" data-row="<?= $ri ?>" data-subid="<?= $sub['id'] ?>"
                     value="<?= $prevMax ?>">

              <input type="number"
                     name="rows[<?= $ri ?>][marks]"
                     class="form-control form-control-sm marks-input"
                     data-row="<?= $ri ?>" data-subid="<?= $sub['id'] ?>"
                     min="0" step="0.5" placeholder="—"
                     value="<?= $prevM ?>"
                     oninput="onInput(this)">

              <?php if ($prev && $prev['grade']): ?>
              <span class="grade-pill mt-1 d-inline-block
                <?= ($prevM / max($prevMax,1)) >= 0.5 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>">
                <?= htmlspecialchars($prev['grade']) ?>
              </span>
              <?php endif; ?>
            </td>
            <?php $ri++; endforeach; ?>

            <td class="text-center fw-bold small row-total" data-student="<?= $st['id'] ?>">—</td>
            <td class="text-center small row-avg"           data-student="<?= $st['id'] ?>">—</td>
          </tr>
          <?php endforeach; ?>
        </tbody>

        <tfoot class="table-light">
          <tr>
            <td colspan="2" class="fw-semibold small sticky-col" style="background:#f1f5f9">Class Avg</td>
            <?php foreach ($entrySubjects as $sub): ?>
            <td class="text-center small fw-semibold col-avg" data-subid="<?= $sub['id'] ?>">—</td>
            <?php endforeach; ?>
            <td colspan="2"></td>
          </tr>
        </tfoot>
      </table>
    </div>

    <div class="card-footer d-flex justify-content-end gap-2 py-2">
      <a href="?<?= http_build_query(array_merge($_GET, ['mode'=>'view'])) ?>"
         class="btn btn-outline-secondary btn-sm">Cancel</a>
      <button type="submit" class="btn btn-success btn-sm">
        <i class="bi bi-check-lg me-1"></i>Save All Marks
      </button>
    </div>
  </div>
</form>

<?php endif; ?>

<?php else: ?>
<!-- ══════════════════════════════════════════════════════════
     VIEW / RESULTS GRID
══════════════════════════════════════════════════════════ -->
<?php if (empty($results)): ?>
<div class="card">
  <div class="card-body text-center py-5 text-muted">
    <i class="bi bi-clipboard-x fs-1 d-block mb-2 opacity-25"></i>
    No results yet.
    <?php if ($canGrade): ?>
    <div class="mt-3">
      <a href="?<?= http_build_query(array_merge($_GET, ['mode'=>'entry'])) ?>"
         class="btn btn-primary btn-sm">
        <i class="bi bi-pencil-square me-1"></i>Enter Marks
      </a>
    </div>
    <?php endif; ?>
  </div>
</div>
<?php else: ?>
<div class="card">
  <div class="card-header py-2 d-flex justify-content-between align-items-center">
    <span class="fw-semibold small"><i class="bi bi-table me-1"></i>Results Grid</span>
    <button class="btn btn-sm btn-outline-secondary" onclick="window.print()">
      <i class="bi bi-printer me-1"></i>Print
    </button>
  </div>
  <div class="card-body p-0" style="overflow-x:auto">
    <table class="table table-bordered table-hover table-sm mb-0 marks-tbl"
           style="min-width:<?= 250 + count($entrySubjects) * 90 ?>px">
      <thead>
        <tr>
          <th class="sticky-col" style="min-width:160px"># Student</th>
          <th style="min-width:75px">Adm No</th>
          <?php foreach ($entrySubjects as $sub): ?>
          <th class="text-center" style="min-width:88px;font-size:.72rem">
            <span class="badge bg-primary-subtle text-primary"><?= htmlspecialchars($sub['code']) ?></span>
          </th>
          <?php endforeach; ?>
          <th class="text-center" style="min-width:70px">Total</th>
          <th class="text-center" style="min-width:60px">Avg%</th>
          <th class="text-center" style="min-width:46px">Pos</th>
        </tr>
      </thead>
      <tbody>
        <?php foreach ($ranked as $pos => $item): $st = $item['st']; ?>
        <tr>
          <td class="sticky-col fw-semibold small"><?= htmlspecialchars($st['name']) ?></td>
          <td class="small text-muted"><?= htmlspecialchars($st['admission_no']) ?></td>
          <?php foreach ($entrySubjects as $sub):
            $r = $existing[$st['id']][$sub['id']] ?? null;
            $pct = $r ? $r['marks'] / max($r['max_marks'],1) : null;
          ?>
          <td class="text-center">
            <?php if ($r): ?>
              <span class="fw-bold"><?= $r['marks'] ?></span>
              <div>
                <span class="grade-pill <?= $pct >= 0.5 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>">
                  <?= $r['grade'] ?>
                </span>
              </div>
            <?php else: ?><span class="text-muted">—</span><?php endif; ?>
          </td>
          <?php endforeach; ?>
          <td class="text-center fw-bold small">
            <?= $item['max'] > 0 ? "{$item['total']}/{$item['max']}" : '—' ?>
          </td>
          <td class="text-center">
            <?php if ($item['avg'] !== null): ?>
              <span class="badge <?= $item['avg'] >= 50 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>">
                <?= $item['avg'] ?>%
              </span>
            <?php else: ?><span class="text-muted">—</span><?php endif; ?>
          </td>
          <td class="text-center">
            <span class="badge bg-secondary-subtle text-secondary"><?= $pos + 1 ?></span>
          </td>
        </tr>
        <?php endforeach; ?>
      </tbody>
      <tfoot class="table-light">
        <tr>
          <td colspan="2" class="fw-semibold small sticky-col" style="background:#f1f5f9">Class Average</td>
          <?php foreach ($entrySubjects as $sub):
            $vals = [];
            foreach ($existing as $stRes) {
                if (isset($stRes[$sub['id']])) {
                    $r = $stRes[$sub['id']];
                    $vals[] = $r['marks'] / max($r['max_marks'],1) * 100;
                }
            }
            $avg = count($vals) ? round(array_sum($vals)/count($vals),1) : null;
          ?>
          <td class="text-center small fw-semibold">
            <?php if ($avg !== null): ?>
              <span class="<?= $avg >= 50 ? 'text-success' : 'text-danger' ?>"><?= $avg ?>%</span>
            <?php else: ?>—<?php endif; ?>
          </td>
          <?php endforeach; ?>
          <td colspan="3"></td>
        </tr>
      </tfoot>
    </table>
  </div>
</div>
<?php endif; ?>
<?php endif; ?>

<?php if ($mode === 'entry' && $canGrade): ?>
<script>
(function () {
  const subIDs = <?= json_encode(array_keys($entrySubjects)) ?>;

  // Get all marks inputs for a student ID
  function studentInputs(sid) {
    return [...document.querySelectorAll(`.marks-input[data-row]`)].filter(inp => {
      // find hidden student_id in same row
      const tr = inp.closest('tr');
      return tr && tr.querySelector(`input[name*="[student_id]"][value="${sid}"]`);
    });
  }

  // Recompute row total + avg for a given student_id
  function recomputeRow(studentID) {
    const tr = [...document.querySelectorAll('tbody tr')].find(r =>
      r.querySelector(`input[name*="[student_id]"][value="${studentID}"]`)
    );
    if (!tr) return;

    let total = 0, maxTotal = 0, filled = 0;
    tr.querySelectorAll('.marks-input').forEach(inp => {
      if (inp.value === '') return;
      const row   = inp.dataset.row;
      const subid = inp.dataset.subid;
      const maxH  = document.querySelector(`.max-h[data-row="${row}"][data-subid="${subid}"]`);
      const mx    = parseFloat(maxH?.value || 100);
      total    += parseFloat(inp.value) || 0;
      maxTotal += mx;
      filled++;
    });

    const totEl = tr.querySelector('.row-total');
    const avgEl = tr.querySelector('.row-avg');
    if (totEl) totEl.textContent = filled ? `${Math.round(total*10)/10}/${maxTotal}` : '—';
    if (avgEl) {
      if (filled) {
        const pct = Math.round(total / maxTotal * 1000) / 10;
        avgEl.innerHTML = `<span class="badge ${pct >= 50 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger'}">${pct}%</span>`;
      } else {
        avgEl.textContent = '—';
      }
    }
  }

  // Recompute column average for a subject
  function recomputeColAvg(subid) {
    const inputs = document.querySelectorAll(`.marks-input[data-subid="${subid}"]`);
    let sum = 0, maxSum = 0, count = 0;
    inputs.forEach(inp => {
      if (inp.value === '') return;
      const row  = inp.dataset.row;
      const maxH = document.querySelector(`.max-h[data-row="${row}"][data-subid="${subid}"]`);
      const mx   = parseFloat(maxH?.value || 100);
      sum    += parseFloat(inp.value) || 0;
      maxSum += mx;
      count++;
    });
    const cell = document.querySelector(`.col-avg[data-subid="${subid}"]`);
    if (!cell) return;
    if (count) {
      const pct = Math.round(sum / maxSum * 1000) / 10;
      cell.innerHTML = `<span class="${pct >= 50 ? 'text-success' : 'text-danger'}">${pct}%</span>`;
    } else {
      cell.textContent = '—';
    }
  }

  // Called on every marks input change
  window.onInput = function (inp) {
    const row   = inp.dataset.row;
    const subid = inp.dataset.subid;
    const maxH  = document.querySelector(`.max-h[data-row="${row}"][data-subid="${subid}"]`);
    const max   = parseFloat(maxH?.value || 100);

    if (parseFloat(inp.value) > max) {
      inp.classList.add('over');
      inp.value = max;
      setTimeout(() => inp.classList.remove('over'), 900);
    }

    // find student id from same row
    const tr  = inp.closest('tr');
    const sid = tr?.querySelector('input[name*="[student_id]"]')?.value;
    if (sid) recomputeRow(sid);
    recomputeColAvg(subid);
  };

  // Apply column max from header input
  window.applyColMax = function (el) {
    const subid = el.dataset.subid;
    const val   = el.value;
    document.querySelectorAll(`.max-h[data-subid="${subid}"]`).forEach(h => {
      h.value = val;
      const marksInp = document.querySelector(`.marks-input[data-row="${h.dataset.row}"][data-subid="${subid}"]`);
      if (marksInp) {
        const tr  = marksInp.closest('tr');
        const sid = tr?.querySelector('input[name*="[student_id]"]')?.value;
        if (sid) recomputeRow(sid);
      }
    });
    recomputeColAvg(subid);
  };

  // Apply global max to all columns
  window.applyGlobalMax = function () {
    const val = document.getElementById('globalMax').value;
    document.querySelectorAll('.col-max').forEach(el => {
      el.value = val;
      applyColMax(el);
    });
  };

  // Validation on submit
  document.getElementById('marksForm').addEventListener('submit', function (e) {
    let hasValue = false, hasError = false;
    document.querySelectorAll('.marks-input').forEach(inp => {
      if (inp.value !== '') {
        hasValue = true;
        const row   = inp.dataset.row;
        const subid = inp.dataset.subid;
        const maxH  = document.querySelector(`.max-h[data-row="${row}"][data-subid="${subid}"]`);
        const max   = parseFloat(maxH?.value || 100);
        if (parseFloat(inp.value) > max) { hasError = true; inp.classList.add('over'); }
      }
    });
    if (!hasValue) {
      e.preventDefault();
      alert('Enter at least one mark before saving.');
      return;
    }
    if (hasError) {
      e.preventDefault();
      alert('Some marks exceed the maximum. Please fix before saving.');
    }
  });

  // Initial compute for pre-filled values
  document.querySelectorAll('tbody tr').forEach(tr => {
    const sid = tr.querySelector('input[name*="[student_id]"]')?.value;
    if (sid) recomputeRow(sid);
  });
  subIDs.forEach(s => recomputeColAvg(s));
})();
</script>
<?php endif; ?>