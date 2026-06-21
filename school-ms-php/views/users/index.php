<?php /** @var array $users @var array $roles */ ?>
<div class="d-flex justify-content-between align-items-center mb-3">
  <h5 class="fw-bold mb-0"><i class="bi bi-person-gear me-2 text-primary"></i>Users</h5>
  <a href="/users/create" class="btn btn-primary btn-sm"><i class="bi bi-plus-lg me-1"></i>Create User</a>
</div>
<div class="card"><div class="card-body p-0">
  <table class="table table-hover mb-0">
    <thead><tr><th>Name</th><th>Email</th><th>Role</th><th class="text-center">Active</th><th></th></tr></thead>
    <tbody>
      <?php if (empty($users)): ?>
        <tr><td colspan="5" class="text-center text-muted py-4">No users found.</td></tr>
      <?php else: foreach ($users as $u): ?>
      <tr>
        <td class="fw-semibold"><?= htmlspecialchars($u['name']) ?></td>
        <td class="small text-muted"><?= htmlspecialchars($u['email']) ?></td>
        <td>
          <form method="POST" action="/users/<?= $u['id'] ?>/role" class="d-inline-flex gap-1">
            <select name="role" class="form-select form-select-sm" style="width:auto" onchange="this.form.submit()">
              <?php foreach ($roles as $r): ?><option value="<?= $r ?>" <?= $u['role']===$r?'selected':'' ?>><?= ucfirst($r) ?></option><?php endforeach; ?>
            </select>
          </form>
        </td>
        <td class="text-center">
          <?php if ($u['is_active']): ?>
            <span class="badge bg-success-subtle text-success">Active</span>
          <?php else: ?>
            <span class="badge bg-secondary-subtle text-secondary">Inactive</span>
          <?php endif; ?>
        </td>
        <td class="text-end">
          <?php if ($u['is_active']): ?>
            <form method="POST" action="/users/<?= $u['id'] ?>/deactivate" class="d-inline" onsubmit="return confirm('Deactivate this user?')">
              <button class="btn btn-sm btn-outline-danger"><i class="bi bi-person-x"></i></button>
            </form>
          <?php else: ?>
            <form method="POST" action="/users/<?= $u['id'] ?>/activate" class="d-inline">
              <button class="btn btn-sm btn-outline-success"><i class="bi bi-person-check"></i></button>
            </form>
          <?php endif; ?>
        </td>
      </tr>
      <?php endforeach; endif; ?>
    </tbody>
  </table>
</div></div>
