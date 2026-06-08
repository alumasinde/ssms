<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1">
  <title>Sign In — SchoolMS</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
  <style>
    body { background: linear-gradient(135deg,#1e2a45 0%,#1a56db 100%); min-height:100vh; display:flex; align-items:center; justify-content:center; }
    .login-card { width:100%; max-width:420px; border-radius:1rem; box-shadow:0 20px 60px rgba(0,0,0,.25); }
    .login-brand { font-size:1.4rem; font-weight:700; color:#1a56db; }
  </style>
</head>
<body>
<div class="login-card card p-4 p-md-5 mx-3">
  <div class="text-center mb-4">
    <div class="fs-1 mb-1">🏫</div>
    <div class="login-brand">SchoolMS</div>
    <p class="text-muted small mt-1">School Management System</p>
  </div>

  <?php if ($err = \Core\Session::flash('error')): ?>
    <div class="alert alert-danger py-2 small">
      <i class="bi bi-exclamation-circle me-1"></i><?= htmlspecialchars($err) ?>
    </div>
  <?php endif; ?>

  <form method="POST" action="/auth/login">
    <div class="mb-3">
      <label class="form-label fw-semibold small">Email Address</label>
      <div class="input-group">
        <span class="input-group-text"><i class="bi bi-envelope"></i></span>
        <input type="email" name="email" class="form-control"
               placeholder="admin@school.ac.ke" required autofocus>
      </div>
    </div>
    <div class="mb-4">
      <label class="form-label fw-semibold small">Password</label>
      <div class="input-group">
        <span class="input-group-text"><i class="bi bi-lock"></i></span>
        <input type="password" name="password" class="form-control"
               placeholder="••••••••" required>
      </div>
    </div>
    <button type="submit" class="btn btn-primary w-100 py-2 fw-semibold">
      <i class="bi bi-box-arrow-in-right me-1"></i> Sign In
    </button>
  </form>
</div>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
