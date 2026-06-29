<?php
/** @var array  $scales  — current grade scale rows from DB  */
/** @var array  $presets — preset names: ['kcse','cbc','pct'] */

$presetData = [
  'kcse' => [['A','80','100','Excellent'],['A-','75','79','Very Good'],['B+','70','74','Good'],['B','65','69','Good'],['B-','60','64','Above Average'],['C+','55','59','Average'],['C','50','54','Average'],['C-','45','49','Below Average'],['D+','40','44','Poor'],['D','35','39','Poor'],['D-','30','34','Very Poor'],['E','0','29','Fail']],
  'cbc'  => [['EE','80','100','Exceeds Expectation'],['ME','50','79','Meets Expectation'],['AE','30','49','Approaches Expectation'],['BE','0','29','Below Expectation']],
  'pct'  => [['A','75','100','Distinction'],['B','60','74','Credit'],['C','50','59','Pass'],['D','40','49','Marginal'],['E','0','39','Fail']],
];

function gradeBadgeClass(float $min): string {
  return match(true) {
    $min >= 75 => 'bg-success',
    $min >= 60 => 'bg-primary',
    $min >= 50 => 'bg-info text-dark',
    $min >= 35 => 'bg-warning text-dark',
    default    => 'bg-danger',
  };
}
?>

<div class="d-flex justify-content-between align-items-center mb-3 flex-wrap gap-2">
  <div>
    <h5 class="fw-bold mb-0"><i class="bi bi-sliders me-2 text-primary"></i>Grade Scales</h5>
    <div class="text-muted small">Define how marks map to grades for this school.</div>
  </div>
</div>

<?php if (empty($scales)): ?>
<!-- ── Empty state ───────────────────────────────────────────────────────── -->
<div class="card mb-4">
  <div class="card-body text-center py-5">
    <i class="bi bi-patch-exclamation fs-1 text-warning d-block mb-2 opacity-50"></i>
    <p class="fw-semibold mb-1">No grade scale set up yet.</p>
    <p class="text-muted small mb-3">Seed a preset to get started, or add entries manually below.</p>
    <div class="d-flex justify-content-center gap-2 flex-wrap">
      <?php foreach ($presets as $p): ?>
      <form method="POST" action="/grade-scales/store">
        <input type="hidden" name="preset" value="<?= $p ?>">
        <button class="btn btn-outline-primary btn-sm">
          <i class="bi bi-lightning me-1"></i>Seed <?= strtoupper($p) ?>
        </button>
      </form>
      <?php endforeach; ?>
    </div>
  </div>
</div>

<?php else: ?>
<!-- ── Current scale table ───────────────────────────────────────────────── -->
<div class="card mb-4">
  <div class="card-header py-2 d-flex justify-content-between align-items-center">
    <span class="fw-semibold small">
      <i class="bi bi-table me-1"></i>Current Grade Scale
      <span class="badge bg-secondary-subtle text-secondary ms-1"><?= count($scales) ?> entries</span>
    </span>
    <div class="d-flex gap-2">
      <div class="dropdown">
        <button class="btn btn-outline-warning btn-sm dropdown-toggle" data-bs-toggle="dropdown">
          <i class="bi bi-lightning me-1"></i>Replace with Preset
        </button>
        <ul class="dropdown-menu dropdown-menu-end">
          <?php foreach ($presets as $p): ?>
          <li>
            <form method="POST" action="/grade-scales/store"
                  onsubmit="return confirm('Replace current scale with <?= strtoupper($p) ?>? All existing entries will be deleted.')">
              <input type="hidden" name="preset" value="<?= $p ?>">
              <button class="dropdown-item small"><?= strtoupper($p) ?></button>
            </form>
          </li>
          <?php endforeach; ?>
        </ul>
      </div>
      <form method="POST" action="/grade-scales/clear"
            onsubmit="return confirm('Delete ALL grade scale entries? This cannot be undone.')">
        <button class="btn btn-outline-danger btn-sm">
          <i class="bi bi-trash3 me-1"></i>Clear All
        </button>
      </form>
    </div>
  </div>

  <div class="table-responsive">
    <table class="table table-sm table-hover align-middle mb-0">
      <thead class="table-light">
        <tr>
          <th style="width:90px">Grade</th>
          <th style="width:110px">Min %</th>
          <th style="width:110px">Max %</th>
          <th>Remark</th>
          <th style="width:110px" class="text-center">Actions</th>
        </tr>
      </thead>
      <tbody>
        <?php foreach ($scales as $s): ?>
        <tr id="row-<?= $s['id'] ?>">

          <!-- VIEW cells -->
          <td class="vm">
            <span class="badge fw-bold <?= gradeBadgeClass((float)$s['min_score']) ?>">
              <?= htmlspecialchars($s['grade']) ?>
            </span>
          </td>
          <td class="vm fw-semibold"><?= $s['min_score'] ?>%</td>
          <td class="vm fw-semibold"><?= $s['max_score'] ?>%</td>
          <td class="vm text-muted small"><?= htmlspecialchars($s['remark'] ?? '') ?></td>
          <td class="vm text-center">
            <div class="d-flex gap-1 justify-content-center">
              <button class="btn btn-outline-secondary btn-sm py-0 px-2"
                      onclick="editRow(<?= $s['id'] ?>)" title="Edit">
                <i class="bi bi-pencil"></i>
              </button>
              <form method="POST" action="/grade-scales/<?= $s['id'] ?>/delete"
                    onsubmit="return confirm('Delete grade <?= addslashes($s['grade']) ?>?')"
                    class="d-inline">
                <button class="btn btn-outline-danger btn-sm py-0 px-2" title="Delete">
                  <i class="bi bi-trash"></i>
                </button>
              </form>
            </div>
          </td>

          <!-- EDIT cells (hidden by default) -->
          <td class="em d-none">
            <input type="text" id="eg-<?= $s['id'] ?>"
                   class="form-control form-control-sm text-center fw-bold"
                   value="<?= htmlspecialchars($s['grade']) ?>"
                   maxlength="5" style="width:58px">
          </td>
          <td class="em d-none">
            <div class="input-group input-group-sm" style="width:95px">
              <input type="number" id="emin-<?= $s['id'] ?>"
                     class="form-control form-control-sm text-center"
                     value="<?= $s['min_score'] ?>" min="0" max="100" step="0.5">
              <span class="input-group-text px-1">%</span>
            </div>
          </td>
          <td class="em d-none">
            <div class="input-group input-group-sm" style="width:95px">
              <input type="number" id="emax-<?= $s['id'] ?>"
                     class="form-control form-control-sm text-center"
                     value="<?= $s['max_score'] ?>" min="0" max="100" step="0.5">
              <span class="input-group-text px-1">%</span>
            </div>
          </td>
          <td class="em d-none">
            <input type="text" id="erm-<?= $s['id'] ?>"
                   class="form-control form-control-sm"
                   value="<?= htmlspecialchars($s['remark'] ?? '') ?>"
                   maxlength="100" placeholder="Remark">
          </td>
          <td class="em d-none text-center">
            <div class="d-flex gap-1 justify-content-center">
              <button class="btn btn-success btn-sm py-0 px-2"
                      id="save-btn-<?= $s['id'] ?>"
                      onclick="saveRow(<?= $s['id'] ?>)" title="Save (Enter)">
                <i class="bi bi-check-lg"></i>
              </button>
              <button class="btn btn-outline-secondary btn-sm py-0 px-2"
                      onclick="cancelEdit(<?= $s['id'] ?>)" title="Cancel (Esc)">
                <i class="bi bi-x-lg"></i>
              </button>
            </div>
          </td>

        </tr>
        <?php endforeach; ?>
      </tbody>
    </table>
  </div>
</div>
<?php endif; ?>

<!-- ── Add single entry ───────────────────────────────────────────────────── -->
<div class="card mb-4">
  <div class="card-header py-2">
    <span class="fw-semibold small">
      <i class="bi bi-plus-circle me-1 text-success"></i>Add Grade Entry
    </span>
  </div>
  <div class="card-body">
    <form method="POST" action="/grade-scales/store" class="row g-3 align-items-end">
      <div class="col-6 col-sm-2">
        <label class="form-label small fw-semibold">Grade <span class="text-danger">*</span></label>
        <input type="text" name="grade"
               class="form-control form-control-sm text-center fw-bold"
               placeholder="A" maxlength="5" required>
      </div>
      <div class="col-6 col-sm-2">
        <label class="form-label small fw-semibold">Min % <span class="text-danger">*</span></label>
        <input type="number" name="min_score"
               class="form-control form-control-sm"
               min="0" max="100" step="0.5" placeholder="75" required>
      </div>
      <div class="col-6 col-sm-2">
        <label class="form-label small fw-semibold">Max % <span class="text-danger">*</span></label>
        <input type="number" name="max_score"
               class="form-control form-control-sm"
               min="0" max="100" step="0.5" placeholder="100" required>
      </div>
      <div class="col-6 col-sm-4">
        <label class="form-label small fw-semibold">Remark</label>
        <input type="text" name="remark"
               class="form-control form-control-sm"
               placeholder="e.g. Excellent" maxlength="100">
      </div>
      <div class="col-12 col-sm-2">
        <button type="submit" class="btn btn-success btn-sm w-100">
          <i class="bi bi-plus-lg me-1"></i>Add
        </button>
      </div>
    </form>
  </div>
</div>

<!-- ── Preset reference ───────────────────────────────────────────────────── -->
<p class="text-muted small fw-semibold mb-2">
  <i class="bi bi-info-circle me-1"></i>Preset Reference
</p>
<div class="row g-3 mb-4">
  <?php foreach ($presetData as $name => $rows): ?>
  <div class="col-md-4">
    <div class="card border-0 bg-light h-100">
      <div class="card-header py-1 bg-transparent border-bottom d-flex justify-content-between align-items-center">
        <span class="fw-bold small text-uppercase"><?= $name ?></span>
        <form method="POST" action="/grade-scales/store" class="d-inline"
              onsubmit="return confirm('Apply <?= strtoupper($name) ?> preset? Current scale will be replaced.')">
          <input type="hidden" name="preset" value="<?= $name ?>">
          <button class="btn btn-outline-primary btn-sm py-0 px-2" style="font-size:.65rem">
            <i class="bi bi-lightning me-1"></i>Apply
          </button>
        </form>
      </div>
      <div class="card-body p-0">
        <table class="table table-sm mb-0" style="font-size:.74rem">
          <thead class="table-light">
            <tr><th>Grade</th><th>Min</th><th>Max</th><th>Remark</th></tr>
          </thead>
          <tbody>
            <?php foreach ($rows as [$g, $min, $max, $rem]): ?>
            <tr>
              <td><span class="fw-bold badge <?= gradeBadgeClass((float)$min) ?>"><?= $g ?></span></td>
              <td><?= $min ?>%</td>
              <td><?= $max ?>%</td>
              <td class="text-muted"><?= $rem ?></td>
            </tr>
            <?php endforeach; ?>
          </tbody>
        </table>
      </div>
    </div>
  </div>
  <?php endforeach; ?>
</div>

<script>
function editRow(id) {
  const row = document.getElementById('row-' + id);
  row.querySelectorAll('.vm').forEach(el => el.classList.add('d-none'));
  row.querySelectorAll('.em').forEach(el => el.classList.remove('d-none'));
  document.getElementById('eg-' + id).focus();
  document.getElementById('eg-' + id).select();
}

function cancelEdit(id) {
  const row = document.getElementById('row-' + id);
  row.querySelectorAll('.vm').forEach(el => el.classList.remove('d-none'));
  row.querySelectorAll('.em').forEach(el => el.classList.add('d-none'));
}

async function saveRow(id) {
  const grade    = document.getElementById('eg-'   + id).value.trim();
  const minScore = parseFloat(document.getElementById('emin-' + id).value);
  const maxScore = parseFloat(document.getElementById('emax-' + id).value);
  const remark   = document.getElementById('erm-'  + id).value.trim();

  if (!grade) {
    alert('Grade cannot be empty.');
    document.getElementById('eg-' + id).focus();
    return;
  }
  if (isNaN(minScore) || isNaN(maxScore)) {
    alert('Min and Max scores are required.');
    return;
  }
  if (minScore >= maxScore) {
    alert('Min score must be less than Max score.');
    document.getElementById('emin-' + id).focus();
    return;
  }

  const btn = document.getElementById('save-btn-' + id);
  const orig = btn.innerHTML;
  btn.innerHTML = '<span class="spinner-border spinner-border-sm"></span>';
  btn.disabled  = true;

  try {
    const res  = await fetch(`/grade-scales/${id}/update`, {
      method:  'POST',
      headers: { 'Content-Type': 'application/json' },
      body:    JSON.stringify({ grade, min_score: minScore, max_score: maxScore, remark }),
    });
    const data = await res.json();
    if (data.success) {
      window.location.reload();
    } else {
      alert(data.error ?? 'Update failed. Please try again.');
      btn.innerHTML = orig;
      btn.disabled  = false;
    }
  } catch (e) {
    alert('Network error. Please try again.');
    btn.innerHTML = orig;
    btn.disabled  = false;
  }
}

// Enter to save, Escape to cancel while in edit mode
document.addEventListener('keydown', function (e) {
  const active = document.activeElement;
  if (!active) return;
  const row = active.closest('tr[id^="row-"]');
  if (!row || !active.closest('.em')) return;
  const id = parseInt(row.id.replace('row-', ''));
  if (e.key === 'Enter')  { e.preventDefault(); saveRow(id); }
  if (e.key === 'Escape') { cancelEdit(id); }
});
</script>