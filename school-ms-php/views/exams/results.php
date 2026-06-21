<?php
/** @var array $exam     */
/** @var array $results  */
/** @var array $students */
/** @var array $subjects */
/** @var array $classes  */
/** @var int   $classID  */

$canGrade    = \Core\Session::can('exams.grade');
$examID      = $exam['id'] ?? 0;
$examClassID = $exam['class_id'] ?? null;

// ── Index existing results keyed by [student_id][subject_id] ──────────────────
$existing     = [];
$studentIndex = [];
$subjectIndex = [];
foreach ($results as $r) {
    $existing[$r['student_id']][$r['subject_id']] = $r;
    $studentIndex[$r['student_id']] = [
        'id'           => $r['student_id'],
        'name'         => $r['student_name'],
        'admission_no' => $r['admission_no'],
    ];
    $subjectIndex[$r['subject_id']] = [
        'id'   => $r['subject_id'],
        'name' => $r['subject_name'],
        'code' => $r['subject_code'],
    ];
}

// ── Build the student list for bulk entry ─────────────────────────────────────
// Merge students from results + the loaded $students list (class roster)
$entryStudents = $studentIndex; // already indexed by id
foreach ($students as $s) {
    if (!isset($entryStudents[$s['id']])) {
        $entryStudents[$s['id']] = [
            'id'           => $s['id'],
            'name'         => trim(($s['first_name'] ?? '') . ' ' . ($s['last_name'] ?? '')),
            'admission_no' => $s['admission_no'] ?? '',
        ];
    }
}
// Sort by name
uasort($entryStudents, fn($a, $b) => strcmp($a['name'], $b['name']));

// ── Build subject list for bulk entry ────────────────────────────────────────
$entrySubjects = $subjectIndex;
foreach ($subjects as $sub) {
    if (!isset($entrySubjects[$sub['id']])) {
        $entrySubjects[$sub['id']] = [
            'id'   => $sub['id'],
            'name' => $sub['name'],
            'code' => $sub['code'],
        ];
    }
}

// ── Ranking for the display grid ─────────────────────────────────────────────
$ranked = [];
foreach ($studentIndex as $st) {
    $total = 0; $max = 0;
    foreach ($subjectIndex as $sub) {
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

// Decide which mode we're in: 'entry' shows the bulk form, 'view' shows results grid
$mode = $_GET['mode'] ?? ($canGrade && empty($results) ? 'entry' : 'view');
?>

<!-- ── Breadcrumb + page header ──────────────────────────────────────────── -->
<div class="d-flex justify-content-between align-items-start mb-3 flex-wrap gap-2">
  <div>
    <nav aria-label="breadcrumb">
      <ol class="breadcrumb mb-0 small">
        <li class="breadcrumb-item"><a href="/exams">Exams</a></li>
        <li class="breadcrumb-item active"><?= htmlspecialchars($exam['name'] ?? '') ?></li>
      </ol>
    </nav>
    <h5 class="fw-bold mb-0 mt-1">
      <i class="bi bi-list-ol me-2 text-warning"></i><?= htmlspecialchars($exam['name'] ?? '') ?> — Results
    </h5>
  </div>

  <div class="d-flex gap-2 flex-wrap">
    <?php if ($canGrade): ?>
      <?php if ($mode === 'view'): ?>
        <a href="?<?= http_build_query(array_merge($_GET, ['mode' => 'entry'])) ?>"
           class="btn btn-primary btn-sm">
          <i class="bi bi-pencil-square me-1"></i>
          <?= empty($results) ? 'Enter Results' : 'Edit / Add Results' ?>
        </a>
      <?php else: ?>
        <a href="?<?= http_build_query(array_merge($_GET, ['mode' => 'view'])) ?>"
           class="btn btn-outline-secondary btn-sm">
          <i class="bi bi-eye me-1"></i>View Grid
        </a>
      <?php endif; ?>
    <?php endif; ?>

    <?php if (!empty($results)): ?>
      <a href="/reports/class-results?exam_id=<?= $examID ?><?= $classID ? '&class_id=' . $classID : '' ?>"
         class="btn btn-outline-info btn-sm">
        <i class="bi bi-bar-chart me-1"></i>Full Report
      </a>
      <a href="/reports/subject-performance?exam_id=<?= $examID ?>"
         class="btn btn-outline-secondary btn-sm">
        <i class="bi bi-graph-up me-1"></i>Subject Analysis
      </a>
    <?php endif; ?>

    <a href="/exams" class="btn btn-outline-secondary btn-sm">
      <i class="bi bi-arrow-left me-1"></i>Back
    </a>
  </div>
</div>

<!-- ── Class filter (school-wide exams only) ─────────────────────────────── -->
<?php if (!$examClassID && !empty($classes)): ?>
<div class="card mb-3">
  <div class="card-body py-2">
    <form method="GET" class="row g-2 align-items-end">
      <input type="hidden" name="mode" value="<?= htmlspecialchars($mode) ?>">
      <div class="col-md-5">
        <label class="form-label small fw-semibold mb-1">Filter by Class</label>
        <select name="class_id" class="form-select form-select-sm" onchange="this.form.submit()">
          <option value="">All classes</option>
          <?php foreach ($classes as $c): ?>
            <option value="<?= $c['id'] ?>" <?= $classID == $c['id'] ? 'selected' : '' ?>>
              <?= htmlspecialchars($c['name']) ?>
            </option>
          <?php endforeach; ?>
        </select>
      </div>
    </form>
  </div>
</div>
<?php endif; ?>

<!-- ── Exam meta badges ───────────────────────────────────────────────────── -->
<div class="d-flex gap-2 mb-3 flex-wrap">
  <span class="badge bg-warning-subtle text-warning border border-warning-subtle">
    <i class="bi bi-tag me-1"></i><?= ucfirst($exam['type'] ?? '') ?>
  </span>
  <span class="badge bg-secondary-subtle text-secondary border">
    <i class="bi bi-calendar me-1"></i><?= $exam['start_date'] ?? '' ?> → <?= $exam['end_date'] ?? '' ?>
  </span>
  <span class="badge bg-primary-subtle text-primary border">
    <i class="bi bi-people me-1"></i><?= count($entryStudents) ?> student(s)
  </span>
  <span class="badge bg-info-subtle text-info border">
    <i class="bi bi-book me-1"></i><?= count($entrySubjects) ?> subject(s)
  </span>
</div>


<?php if ($mode === 'entry' && $canGrade): ?>
<!-- ══════════════════════════════════════════════════════════════════════════
     BULK ENTRY GRID
     One row per student, one column per subject.
     All marks submitted in a single POST.
════════════════════════════════════════════════════════════════════════════ -->

<?php if (empty($entryStudents)): ?>
  <div class="alert alert-warning">
    <i class="bi bi-exclamation-triangle me-2"></i>
    No students found for this class. <a href="/students">Enrol students</a> first.
  </div>
<?php elseif (empty($entrySubjects)): ?>
  <div class="alert alert-warning">
    <i class="bi bi-exclamation-triangle me-2"></i>
    No subjects found. <a href="/subjects">Add subjects</a> first.
  </div>
<?php else: ?>

<form method="POST" action="/exams/results" id="bulkForm">
  <input type="hidden" name="exam_id"  value="<?= $examID ?>">
  <input type="hidden" name="class_id" value="<?= $classID ?: ($examClassID ?? '') ?>">

  <div class="card">
    <div class="card-header py-3 d-flex justify-content-between align-items-center">
      <span><i class="bi bi-pencil-square me-2 text-primary"></i>
        Bulk Mark Entry
        <span class="text-muted small ms-1">— fill marks then click Save All</span>
      </span>
      <div class="d-flex gap-2 align-items-center">
        <!-- Global max marks setter -->
        <label class="form-label small fw-semibold mb-0 text-muted">Max marks:</label>
        <input type="number" id="globalMax" value="100" min="1" max="1000"
               class="form-control form-control-sm" style="width:75px"
               title="Set max marks for all cells">
        <button type="button" class="btn btn-sm btn-outline-secondary" onclick="applyGlobalMax()">
          Apply All
        </button>
        <button type="submit" class="btn btn-success btn-sm">
          <i class="bi bi-check-lg me-1"></i>Save All Results
        </button>
      </div>
    </div>

    <!-- Sticky header + scrollable body -->
    <div style="overflow-x:auto;max-height:75vh;overflow-y:auto">
      <table class="table table-bordered table-sm mb-0" id="bulkTable"
             style="min-width:<?= 220 + count($entrySubjects) * 130 ?>px">

        <!-- ── Column headers ── -->
        <thead style="position:sticky;top:0;z-index:10;background:#f8f9fa">
          <tr>
            <th class="align-middle" style="min-width:180px;position:sticky;left:0;z-index:11;background:#f8f9fa">
              Student
            </th>
            <th class="text-center align-middle" style="min-width:70px;position:sticky;left:180px;z-index:11;background:#f8f9fa">
              Adm No
            </th>
            <?php foreach ($entrySubjects as $sub): ?>
            <th class="text-center" style="min-width:130px">
              <div class="fw-semibold" style="font-size:.75rem">
                <?= htmlspecialchars($sub['name']) ?>
              </div>
              <div class="text-muted" style="font-size:.65rem">
                <?= htmlspecialchars($sub['code']) ?>
                — <span class="text-muted">out of</span>
                <input type="number"
                       class="col-max-input form-control form-control-sm d-inline-block p-0 text-center border-0 bg-transparent fw-semibold"
                       data-subid="<?= $sub['id'] ?>"
                       value="100" min="1" max="1000"
                       style="width:40px;font-size:.7rem"
                       title="Max marks for <?= htmlspecialchars($sub['name']) ?>"
                       onchange="applyColMax(this)">
              </div>
            </th>
            <?php endforeach; ?>
            <th class="text-center align-middle" style="min-width:70px;font-size:.75rem">Total</th>
            <th class="text-center align-middle" style="min-width:60px;font-size:.75rem">Avg%</th>
          </tr>
        </thead>

        <tbody>
          <?php $rowIdx = 0; foreach ($entryStudents as $st): ?>
          <tr data-row="<?= $rowIdx ?>">

            <!-- Student name (sticky) -->
            <td class="fw-semibold small align-middle"
                style="position:sticky;left:0;background:#fff;z-index:2">
              <?= htmlspecialchars($st['name']) ?>
            </td>

            <!-- Admission no (sticky) -->
            <td class="text-muted small align-middle text-center"
                style="position:sticky;left:180px;background:#fff;z-index:2">
              <?= htmlspecialchars($st['admission_no']) ?>
            </td>

            <?php foreach ($entrySubjects as $sub):
              $prev = $existing[$st['id']][$sub['id']] ?? null;
              $prevMarks = $prev ? $prev['marks']    : '';
              $prevMax   = $prev ? $prev['max_marks'] : 100;
              $prevRem   = $prev ? $prev['remarks']   : '';
            ?>
            <td class="p-1 align-middle" style="min-width:130px">

              <!-- Hidden: student_id, subject_id, remarks -->
              <input type="hidden" name="student_id[<?= $rowIdx ?>]" value="<?= $st['id'] ?>">
              <input type="hidden" name="subject_id[<?= $rowIdx ?>]" value="<?= $sub['id'] ?>">
              <input type="hidden" name="remarks[<?= $rowIdx ?>]"    value="<?= htmlspecialchars($prevRem) ?>">
              <input type="hidden" name="max_marks[<?= $rowIdx ?>]"
                     class="max-hidden" data-subid="<?= $sub['id'] ?>"
                     value="<?= $prevMax ?>">

              <div class="d-flex align-items-center gap-1">
                <!-- Marks input -->
                <input type="number"
                       name="marks[<?= $rowIdx ?>]"
                       class="form-control form-control-sm marks-input text-center p-1"
                       min="0" step="0.5"
                       value="<?= $prevMarks ?>"
                       placeholder="—"
                       data-row="<?= $rowIdx ?>"
                       data-subid="<?= $sub['id'] ?>"
                       data-max="<?= $prevMax ?>"
                       style="width:58px"
                       oninput="onMarksInput(this)">
                <span class="text-muted small">/</span>
                <!-- Max marks (per cell, editable) -->
                <input type="number"
                       class="form-control form-control-sm text-center p-1 cell-max"
                       min="1" step="1"
                       value="<?= $prevMax ?>"
                       data-row="<?= $rowIdx ?>"
                       data-subid="<?= $sub['id'] ?>"
                       style="width:48px"
                       title="Max marks"
                       onchange="onCellMaxChange(this)">
              </div>

              <!-- Grade badge (live-computed) -->
              <?php if ($prev): ?>
              <div class="mt-1 text-center">
                <span class="badge grade-badge
                  <?= ($prev['marks'] / max($prev['max_marks'], 1)) >= 0.5
                        ? 'bg-success-subtle text-success'
                        : 'bg-danger-subtle text-danger' ?>"
                  style="font-size:.6rem">
                  <?= htmlspecialchars($prev['grade']) ?>
                </span>
              </div>
              <?php endif; ?>
            </td>
            <?php $rowIdx++; endforeach; // end subjects ?>

            <!-- Live total -->
            <td class="text-center align-middle fw-bold small row-total" data-row="<?= $rowIdx - count($entrySubjects) ?>">—</td>
            <!-- Live avg -->
            <td class="text-center align-middle row-avg" data-row="<?= $rowIdx - count($entrySubjects) ?>">—</td>
          </tr>
          <?php endforeach; // end students ?>
        </tbody>

        <!-- ── Column averages footer ── -->
        <tfoot style="position:sticky;bottom:0;background:#f8f9fa;z-index:9">
          <tr>
            <td class="fw-semibold small" style="position:sticky;left:0;background:#f8f9fa"
                colspan="2">Class Average</td>
            <?php foreach ($entrySubjects as $sub): ?>
            <td class="text-center small fw-semibold col-avg" data-subid="<?= $sub['id'] ?>">—</td>
            <?php endforeach; ?>
            <td class="text-center small fw-semibold overall-total">—</td>
            <td class="text-center small fw-semibold overall-avg">—</td>
          </tr>
        </tfoot>
      </table>
    </div><!-- /overflow wrapper -->

    <div class="card-footer d-flex justify-content-between align-items-center py-2">
      <span class="small text-muted">
        <i class="bi bi-info-circle me-1"></i>
        Grades are auto-assigned from the school grade scale on save.
        Leaving a cell blank skips that student–subject pair.
      </span>
      <div class="d-flex gap-2">
        <a href="?<?= http_build_query(array_merge($_GET, ['mode' => 'view'])) ?>"
           class="btn btn-outline-secondary btn-sm">Cancel</a>
        <button type="submit" class="btn btn-success">
          <i class="bi bi-check-lg me-1"></i>Save All Results
        </button>
      </div>
    </div>
  </div>
</form>

<?php endif; // entryStudents / entrySubjects ?>


<?php else: ?>
<!-- ══════════════════════════════════════════════════════════════════════════
     READ-ONLY RESULTS GRID
════════════════════════════════════════════════════════════════════════════ -->

<?php if (empty($results)): ?>
<div class="card">
  <div class="card-body text-center py-5 text-muted">
    <i class="bi bi-clipboard-x fs-1 d-block mb-2 opacity-25"></i>
    No results submitted yet.
    <?php if ($canGrade): ?>
    <div class="mt-3">
      <a href="?<?= http_build_query(array_merge($_GET, ['mode' => 'entry'])) ?>"
         class="btn btn-primary">
        <i class="bi bi-pencil-square me-1"></i>Enter Results Now
      </a>
    </div>
    <?php endif; ?>
  </div>
</div>

<?php else: ?>
<div class="card">
  <div class="card-header py-3 d-flex justify-content-between align-items-center">
    <span><i class="bi bi-table me-1"></i>Results Grid</span>
    <button class="btn btn-sm btn-outline-secondary" onclick="window.print()">
      <i class="bi bi-printer me-1"></i>Print
    </button>
  </div>
  <div class="card-body p-0" style="overflow-x:auto">
    <table class="table table-bordered table-hover table-sm mb-0"
           style="min-width:<?= 240 + count($subjectIndex) * 90 ?>px">
      <thead class="table-light">
        <tr>
          <th style="min-width:160px;position:sticky;left:0;background:#f9fafb">#&nbsp;Student</th>
          <th style="min-width:70px;position:sticky;left:160px;background:#f9fafb">Adm No</th>
          <?php foreach ($subjectIndex as $sub): ?>
          <th class="text-center" style="min-width:90px;font-size:.75rem">
            <span class="badge bg-primary-subtle text-primary"><?= htmlspecialchars($sub['code']) ?></span>
            <div style="font-size:.65rem;color:#6b7280;font-weight:400"><?= htmlspecialchars($sub['name']) ?></div>
          </th>
          <?php endforeach; ?>
          <th class="text-center" style="min-width:70px">Total</th>
          <th class="text-center" style="min-width:60px">Avg%</th>
          <th class="text-center" style="min-width:50px">Pos</th>
        </tr>
      </thead>
      <tbody>
        <?php foreach ($ranked as $pos => $item):
          $st = $item['st'];
        ?>
        <tr>
          <td class="fw-semibold small" style="position:sticky;left:0;background:#fff">
            <?= htmlspecialchars($st['name']) ?>
          </td>
          <td class="small text-muted" style="position:sticky;left:160px;background:#fff">
            <?= htmlspecialchars($st['admission_no']) ?>
          </td>
          <?php foreach ($subjectIndex as $sub):
            $r = $existing[$st['id']][$sub['id']] ?? null;
            $pct = $r ? ($r['marks'] / max($r['max_marks'], 1)) : null;
          ?>
          <td class="text-center">
            <?php if ($r): ?>
              <span class="fw-bold"><?= $r['marks'] ?></span>
              <span class="text-muted small">/<?= $r['max_marks'] ?></span>
              <div>
                <span class="badge <?= $pct >= 0.5 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger' ?>"
                      style="font-size:.6rem"><?= $r['grade'] ?></span>
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
            <span class="badge bg-secondary"><?= $pos + 1 ?></span>
          </td>
        </tr>
        <?php endforeach; ?>
      </tbody>
      <!-- Column averages -->
      <tfoot class="table-light fw-semibold small">
        <tr>
          <td colspan="2" style="position:sticky;left:0;background:#f8f9fa">Class Average</td>
          <?php foreach ($subjectIndex as $sub):
            $vals = array_filter(array_map(
              fn($r) => isset($r[$sub['id']]) ? ($r[$sub['id']]['marks'] / max($r[$sub['id']]['max_marks'],1) * 100) : null,
              $existing
            ));
            $avg = count($vals) > 0 ? round(array_sum($vals) / count($vals), 1) : null;
          ?>
          <td class="text-center">
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
<?php endif; // empty results ?>
<?php endif; // mode ?>


<?php if ($mode === 'entry' && $canGrade): ?>
<!-- ── Live totals + validation JS ────────────────────────────────────────── -->
<script>
(function () {
  const subjectIDs  = <?= json_encode(array_keys($entrySubjects)) ?>;
  const numSubjects = subjectIDs.length;

  /* ── Helpers ─────────────────────────────────────────────────────────── */
  function getRowInputs(row) {
    return document.querySelectorAll(`.marks-input[data-row="${row}"]`);
  }
  function getMaxHidden(row, subID) {
    return document.querySelector(`input.max-hidden[data-subid="${subID}"][name="max_marks[${row}]"]`);
  }

  /* ── Recompute a single row's total + avg ───────────────────────────── */
  function recomputeRow(row) {
    let total = 0, maxTotal = 0, filled = 0;
    subjectIDs.forEach(sid => {
      const inp = document.querySelector(`.marks-input[data-row="${row}"][data-subid="${sid}"]`);
      const mxH = getMaxHidden(row, sid);
      if (!inp || inp.value === '') return;
      const marks = parseFloat(inp.value) || 0;
      const mx    = parseFloat(mxH ? mxH.value : inp.dataset.max) || 100;
      total   += marks;
      maxTotal += mx;
      filled++;
    });
    const avgEl   = document.querySelector(`.row-avg[data-row="${row}"]`);
    const totalEl = document.querySelector(`.row-total[data-row="${row}"]`);

    if (totalEl) totalEl.textContent = filled ? `${total}/${maxTotal}` : '—';
    if (avgEl) {
      if (filled) {
        const pct = Math.round(total / maxTotal * 1000) / 10;
        avgEl.innerHTML = `<span class="badge ${pct >= 50 ? 'bg-success-subtle text-success' : 'bg-danger-subtle text-danger'}">${pct}%</span>`;
      } else {
        avgEl.textContent = '—';
      }
    }
  }

  /* ── Recompute column averages ──────────────────────────────────────── */
  function recomputeColAvg(subID) {
    const inputs = document.querySelectorAll(`.marks-input[data-subid="${subID}"]`);
    let sum = 0, maxSum = 0, count = 0;
    inputs.forEach(inp => {
      if (inp.value === '') return;
      const marks = parseFloat(inp.value) || 0;
      const mxH   = getMaxHidden(inp.dataset.row, subID);
      const mx     = parseFloat(mxH ? mxH.value : inp.dataset.max) || 100;
      sum    += marks;
      maxSum += mx;
      count++;
    });
    const cell = document.querySelector(`.col-avg[data-subid="${subID}"]`);
    if (!cell) return;
    if (count) {
      const pct = Math.round(sum / maxSum * 1000) / 10;
      cell.innerHTML = `<span class="${pct >= 50 ? 'text-success' : 'text-danger'}">${pct}%</span>`;
    } else {
      cell.textContent = '—';
    }
  }

  /* ── Recompute overall footer ───────────────────────────────────────── */
  function recomputeOverall() {
    const allMarks = document.querySelectorAll('.marks-input');
    let sum = 0, maxSum = 0, count = 0;
    allMarks.forEach(inp => {
      if (inp.value === '') return;
      const marks = parseFloat(inp.value) || 0;
      const mxH   = getMaxHidden(inp.dataset.row, inp.dataset.subid);
      const mx     = parseFloat(mxH ? mxH.value : inp.dataset.max) || 100;
      sum    += marks;
      maxSum += mx;
      count++;
    });
    const tot = document.querySelector('.overall-total');
    const avg = document.querySelector('.overall-avg');
    if (tot) tot.textContent = count ? `${Math.round(sum*10)/10}/${maxSum}` : '—';
    if (avg) {
      if (count) {
        const pct = Math.round(sum / maxSum * 1000) / 10;
        avg.innerHTML = `<span class="${pct >= 50 ? 'text-success' : 'text-danger'}">${pct}%</span>`;
      } else {
        avg.textContent = '—';
      }
    }
  }

  /* ── Called when a marks cell changes ──────────────────────────────── */
  window.onMarksInput = function (inp) {
    const row   = inp.dataset.row;
    const subID = inp.dataset.subid;
    // Clamp to max
    const mxH = getMaxHidden(row, subID);
    const max  = parseFloat(mxH ? mxH.value : inp.dataset.max) || 100;
    if (parseFloat(inp.value) > max) {
      inp.value = max;
      inp.classList.add('is-invalid');
      setTimeout(() => inp.classList.remove('is-invalid'), 1200);
    }
    recomputeRow(parseInt(row));
    recomputeColAvg(parseInt(subID));
    recomputeOverall();
  };

  /* ── Per-cell max changed ───────────────────────────────────────────── */
  window.onCellMaxChange = function (el) {
    const row   = el.dataset.row;
    const subID = el.dataset.subid;
    const mxH   = getMaxHidden(row, subID);
    if (mxH) mxH.value = el.value;
    // Update the marks input's data-max too
    const marksInp = document.querySelector(`.marks-input[data-row="${row}"][data-subid="${subID}"]`);
    if (marksInp) marksInp.dataset.max = el.value;
    recomputeRow(parseInt(row));
    recomputeColAvg(parseInt(subID));
    recomputeOverall();
  };

  /* ── "Apply All" global max ─────────────────────────────────────────── */
  window.applyGlobalMax = function () {
    const val = document.getElementById('globalMax').value;
    document.querySelectorAll('.cell-max').forEach(el => {
      el.value = val;
      onCellMaxChange(el);
    });
    document.querySelectorAll('.col-max-input').forEach(el => {
      el.value = val;
    });
  };

  /* ── Per-column max (header input) ─────────────────────────────────── */
  window.applyColMax = function (el) {
    const subID = el.dataset.subid;
    const val   = el.value;
    document.querySelectorAll(`.cell-max[data-subid="${subID}"]`).forEach(cell => {
      cell.value = val;
      onCellMaxChange(cell);
    });
  };

  /* ── Tab-key navigation across the grid ────────────────────────────── */
  document.addEventListener('keydown', function (e) {
    if (e.key !== 'Tab') return;
    const active = document.activeElement;
    if (!active || !active.classList.contains('marks-input')) return;
    // default tab already moves to next input; just ensure we don't skip to footer
  });

  /* ── Validation before submit ───────────────────────────────────────── */
  document.getElementById('bulkForm').addEventListener('submit', function (e) {
    let hasValue = false;
    let hasError = false;
    document.querySelectorAll('.marks-input').forEach(inp => {
      if (inp.value !== '') {
        hasValue = true;
        const mxH = getMaxHidden(inp.dataset.row, inp.dataset.subid);
        const max  = parseFloat(mxH ? mxH.value : inp.dataset.max) || 100;
        if (parseFloat(inp.value) > max) {
          hasError = true;
          inp.classList.add('is-invalid');
        }
      }
    });
    if (!hasValue) {
      e.preventDefault();
      alert('Please enter at least one mark before saving.');
      return;
    }
    if (hasError) {
      e.preventDefault();
      alert('Some marks exceed the maximum allowed. Please correct them before saving.');
    }
  });

  /* ── Initial compute (pre-filled from existing results) ────────────── */
  const rows = new Set([...document.querySelectorAll('.marks-input')].map(i => i.dataset.row));
  rows.forEach(r => recomputeRow(parseInt(r)));
  subjectIDs.forEach(s => recomputeColAvg(s));
  recomputeOverall();
})();
</script>
<?php endif; ?>