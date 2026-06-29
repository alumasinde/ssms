<?php
use Core\Session;
$err = Session::flash('error');
?>
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>Sign In — SchoolMS</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
      background: #f0f2f5;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 1rem;
    }

    .card {
      background: #fff;
      border-radius: 12px;
      box-shadow: 0 2px 16px rgba(0,0,0,.10);
      padding: 2.5rem 2rem;
      width: 100%;
      max-width: 380px;
    }

    .brand {
      text-align: center;
      margin-bottom: 2rem;
    }
    .brand .icon { font-size: 2.2rem; }
    .brand h1 { font-size: 1.3rem; font-weight: 700; color: #1a56db; margin: .3rem 0 .2rem; }
    .brand p  { font-size: .82rem; color: #888; }

    .error {
      background: #fff5f5;
      border: 1px solid #fca5a5;
      border-radius: 8px;
      color: #b91c1c;
      font-size: .85rem;
      padding: .6rem .9rem;
      margin-bottom: 1.25rem;
      display: flex;
      align-items: center;
      gap: .5rem;
    }
    .error::before { content: '⚠'; flex-shrink: 0; }

    .field { margin-bottom: 1.1rem; }
    .field label {
      display: block;
      font-size: .82rem;
      font-weight: 600;
      color: #374151;
      margin-bottom: .4rem;
    }
    .field input {
      width: 100%;
      padding: .6rem .85rem;
      border: 1px solid #d1d5db;
      border-radius: 8px;
      font-size: .95rem;
      color: #111;
      outline: none;
      transition: border-color .15s;
    }
    .field input:focus { border-color: #1a56db; box-shadow: 0 0 0 3px rgba(26,86,219,.12); }

    .btn {
      width: 100%;
      padding: .7rem;
      background: #1a56db;
      color: #fff;
      border: none;
      border-radius: 8px;
      font-size: .95rem;
      font-weight: 600;
      cursor: pointer;
      margin-top: .5rem;
      transition: background .15s;
    }
    .btn:hover { background: #1648c0; }
  </style>
</head>
<body>
  <div class="card">

    <div class="brand">
      <div class="icon">🏫</div>
      <h1>SchoolMS</h1>
      <p>School Management System</p>
    </div>

    <?php if (!empty($err)): ?>
      <div class="error">
        <?= htmlspecialchars($err, ENT_QUOTES, 'UTF-8') ?>
      </div>
    <?php endif; ?>

    <form method="POST" action="/auth/login">
      <div class="field">
        <label for="email">Email Address</label>
        <input id="email" type="email" name="email"
               placeholder="admin@school.ac.ke"
               value="<?= htmlspecialchars($_POST['email'] ?? '', ENT_QUOTES, 'UTF-8') ?>"
               required autofocus>
      </div>
      <div class="field">
        <label for="password">Password</label>
        <input id="password" type="password" name="password"
               placeholder="••••••••" required>
      </div>
      <button type="submit" class="btn">Sign In</button>
    </form>

  </div>
</body>
</html>