<?php /** @var array $records @var array $terms @var int $termID @var array $types */
$typeBadge = ['commendation'=>'bg-success','minor_offence'=>'bg-warning text-dark','major_offence'=>'bg-orange text-white','suspension'=>'bg-danger','expulsion'=>'bg-dark'];
?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-shield-exclamation me-2 text-danger"></i>Discipline Records</h5>
  <?php if (\Core\Session::can('discipline.create')): ?>
    <a href="/discipline/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Record Incident</a>
  <?php endif; ?>
</div>
<div class="card mb-3"><div class="card-body py-2">
  <form method="GET" class="row g-2 align-items-end">
    <div class="col-md-4">
      <label class="form-label small fw-semibold mb-1">Term</label>
      <select name="term_id" class="form-select form-select-sm" onchange="this.form.submit()">
        <option value="">All terms</option>
        <?php foreach ($terms as $t): ?><option value="<?= $t['id'] ?>" <?= $termID==$t['id']?'selected':'' ?>><?= htmlspecialchars($t['name']) ?></option><?php endforeach; ?>
      </select>
    </div>
  </form>
</div></div>
<div class="card"><div class="card-body p-0">
  <table class="table table-hover mb-0">
    <thead><tr><th>Date</th><th>Student</th><th>Type</th><th>Description</th><th>Action Taken</th><th>Term</th><th></th></tr></thead>
    <tbody>
      <?php if (empty($records)): ?>
        <tr><td colspan="7" class="text-center text-muted py-4">No records found.</td></tr>
      <?php else: foreach ($records as $rec): ?>
      <tr>
        <td class="small"><?= $rec['incident_date'] ?></td>
        <td class="fw-semibold small"><?= htmlspecialchars($rec['student_name']) ?>
          <div class="text-muted" style="font-size:.7rem"><?= htmlspecialchars($rec['admission_no']) ?></div></td>
        <td><span class="badge <?= $typeBadge[$rec['type']] ?? 'bg-secondary' ?>"><?= $types[$rec['type']] ?? $rec['type'] ?></span></td>
        <td class="small"><?= htmlspecialchars(mb_substr($rec['description'],0,80)) ?><?= strlen($rec['description'])>80?'…':'' ?></td>
        <td class="small text-muted"><?= htmlspecialchars(mb_substr($rec['action_taken']??'',0,60)) ?></td>
        <td class="small text-muted"><?= htmlspecialchars($rec['term_name']) ?></td>
        <td class="text-end">
          <?php if (\Core\Session::can('discipline.edit')): ?>
            <form method="POST" action="/discipline/<?= $rec['id'] ?>/delete" class="d-inline" onsubmit="return confirm('Delete?')">
              <button class="btn btn-sm btn-outline-danger"><i class="bi bi-trash"></i></button>
            </form>
          <?php endif; ?>
        </td>
      </tr>
      <?php endforeach; endif; ?>
    </tbody>
  </table>
</div></div>
