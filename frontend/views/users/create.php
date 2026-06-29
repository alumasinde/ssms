<?php /** @var array $roles */ ?>
<div class="row justify-content-center"><div class="col-lg-5">
<div class="card"><div class="card-header py-3"><i class="bi bi-person-plus me-2 text-primary"></i>Create User</div>
<div class="card-body">
  <form method="POST" action="/users">
    <div class="mb-3"><label class="form-label small fw-semibold">Full Name *</label>
      <input type="text" name="name" class="form-control" required></div>
    <div class="mb-3"><label class="form-label small fw-semibold">Email *</label>
      <input type="email" name="email" class="form-control" required></div>
    <div class="mb-3"><label class="form-label small fw-semibold">Password *</label>
      <input type="password" name="password" class="form-control" required minlength="8"></div>
    <div class="mb-3"><label class="form-label small fw-semibold">Role *</label>
      <select name="role" class="form-select" required>
        <?php foreach ($roles as $r): ?><option value="<?= $r ?>"><?= ucfirst($r) ?></option><?php endforeach; ?>
      </select></div>
    <div class="d-flex gap-2">
      <button type="submit" class="btn btn-primary"><i class="bi bi-check-lg me-1"></i>Create</button>
      <a href="/users" class="btn btn-outline-secondary">Cancel</a>
    </div>
  </form>
</div></div></div></div>
