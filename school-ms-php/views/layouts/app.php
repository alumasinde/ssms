<?php
/** @var string $title */
/** @var string $appName */
/** @var array $currentUser */
/** @var string $viewFile */
?>
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title><?= htmlspecialchars($title ?? 'SchoolMS') ?> — <?= htmlspecialchars($appName) ?></title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
  <style>
   :root {
  --sidebar-w: 240px;
  --primary: #1a56db;
  --primary-dark: #1040b0;
}
body { background: #f4f6fb; font-family: 'Segoe UI', system-ui, sans-serif; }
.sidebar {
  width: var(--sidebar-w); min-height: 100vh; background: #1e2a45;
  position: fixed; top: 0; left: 0; z-index: 100; display: flex; flex-direction: column;
  transition: transform .25s, width .25s;
  height: 100vh;        /* add this */
  overflow: hidden;     /* add this */
}

.sidebar-nav { 
  flex: 1; 
  padding: .75rem 0; 
  overflow-y: auto;     /* already there, keep it */
  overflow-x: hidden;   /* add this */
}
.sidebar-brand {
  padding: 1.25rem 1.2rem; color: #fff; font-size: 1.15rem; font-weight: 700;
  border-bottom: 1px solid rgba(255,255,255,.08); letter-spacing: .3px;
}
.sidebar-brand span { color: #60a5fa; }
.sidebar-section {
  padding: .35rem 1.2rem .1rem; font-size: .68rem; text-transform: uppercase;
  letter-spacing: .08em; color: rgba(255,255,255,.35); margin-top: .5rem;
}
.nav-link-side {
  display: flex; align-items: center; gap: .6rem; padding: .45rem 1.2rem;
  color: rgba(255,255,255,.72); text-decoration: none; border-radius: 0;
  font-size: .88rem; transition: background .15s, color .15s;
}
.nav-link-side:hover, .nav-link-side.active { background: rgba(255,255,255,.1); color: #fff; }
.nav-link-side.active { border-left: 3px solid #60a5fa; }
.nav-link-side i { font-size: 1rem; width: 18px; text-align: center; }
.sidebar-footer { padding: 1rem 1.2rem; border-top: 1px solid rgba(255,255,255,.08); }
.main-wrap {
  margin-left: var(--sidebar-w); min-height: 100vh; display: flex; flex-direction: column;
  transition: margin-left .25s;
}
.sidebar.collapsed { width: 60px; }
.sidebar.collapsed .sidebar-brand span { display: none; }
.sidebar.collapsed .sidebar-section { display: none; }
.sidebar.collapsed .nav-link-side span { display: none; }
.sidebar.collapsed .sidebar-footer .user-info { display: none; }
.sidebar.collapsed .nav-link-side { justify-content: center; padding: .45rem 0; }
.sidebar.collapsed .nav-link-side i { width: auto; font-size: 1.2rem; }
.main-wrap.expanded { margin-left: 60px; }
.topbar {
  background: #fff; border-bottom: 1px solid #e5e7eb; padding: .6rem 1.5rem;
  display: flex; align-items: center; justify-content: space-between; position: sticky; top: 0; z-index: 50;
}
.page-content { padding: 1.5rem; flex: 1; }
.card { border: 1px solid #e5e7eb; border-radius: .75rem; box-shadow: 0 1px 3px rgba(0,0,0,.04); }
.card-header { background: #fff; border-bottom: 1px solid #f0f2f5; font-weight: 600; }
.badge-role { font-size: .7rem; padding: .2em .55em; border-radius: .3rem; }
.stat-card { border-radius: .75rem; padding: 1.25rem; color: #fff; }
.stat-card .stat-num { font-size: 2rem; font-weight: 700; line-height: 1; }
.stat-card .stat-label { font-size: .8rem; opacity: .85; margin-top: .25rem; }
.table th {
  font-size: .78rem; text-transform: uppercase; letter-spacing: .05em;
  color: #6b7280; font-weight: 600; background: #f9fafb;
}
.btn-primary { background: var(--primary); border-color: var(--primary); }
.btn-primary:hover { background: var(--primary-dark); border-color: var(--primary-dark); }
.alert { border-radius: .6rem; }
@media (max-width: 768px) {
  .sidebar { transform: translateX(-100%); }
  .main-wrap { margin-left: 0; }
}
  </style>
</head>
<body>

<?php if (\Core\Session::isLoggedIn()): ?>
<?php $currentUser = \Core\Session::user(); $currentPath = strtok($_SERVER['REQUEST_URI'], '?'); ?>

<nav class="sidebar" id="sidebar">
  <div class="sidebar-brand">🏫 <span>School</span>MS</div>
  <div class="sidebar-nav">

    <div class="sidebar-section">Main</div>
    <a href="/dashboard" class="nav-link-side <?= str_starts_with($currentPath, '/dashboard') ? 'active' : '' ?>">
      <i class="bi bi-speedometer2"></i><span> Dashboard</span>
    </a>

    <div class="sidebar-section">Academics</div>
    <?php if (\Core\Session::can('students.view')): ?>
      <a href="/students" class="nav-link-side <?= str_starts_with($currentPath, '/students') ? 'active' : '' ?>">
        <i class="bi bi-person-badge"></i><span> Students</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('teachers.view')): ?>
    <a href="/teachers" class="nav-link-side <?= str_starts_with($currentPath, '/teachers') ? 'active' : '' ?>">
        <i class="bi bi-person-workspace"></i><span> Teachers</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('classes.view')): ?>
      <a href="/classes" class="nav-link-side <?= str_starts_with($currentPath, '/classes') ? 'active' : '' ?>">
        <i class="bi bi-door-open"></i><span> Classes</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('subjects.view')): ?>
      <a href="/subjects" class="nav-link-side <?= str_starts_with($currentPath, '/subjects') ? 'active' : '' ?>">
        <i class="bi bi-book"></i><span> Subjects</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('parents.view')): ?>
      <a href="/parents" class="nav-link-side <?= str_starts_with($currentPath, '/parents') ? 'active' : '' ?>">
        <i class="bi bi-people"></i><span> Parents</span>
      </a>
    <?php endif; ?>

    <div class="sidebar-section">Operations</div>
    <?php if (\Core\Session::can('attendance.view')): ?>
      <a href="/attendance" class="nav-link-side <?= str_starts_with($currentPath, '/attendance') ? 'active' : '' ?>">
        <i class="bi bi-calendar-check"></i><span> Attendance</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('exams.view')): ?>
      <a href="/exams" class="nav-link-side <?= str_starts_with($currentPath, '/exams') ? 'active' : '' ?>">
        <i class="bi bi-pencil-square"></i><span> Exams</span>
      </a>
    <?php endif; ?>

    <?php if (\Core\Session::can('finance.view')): ?>
      <a href="/finance" class="nav-link-side <?= str_starts_with($currentPath, '/finance') ? 'active' : '' ?>">
        <i class="bi bi-cash-coin"></i><span> Finance</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('notices.view')): ?>
      <a href="/notices" class="nav-link-side <?= str_starts_with($currentPath, '/notices') ? 'active' : '' ?>">
        <i class="bi bi-megaphone"></i><span> Notices</span>
      </a>
    <?php endif; ?>

    <div class="sidebar-section">Reports</div>
    <?php if (\Core\Session::can('reports.class_results.view')): ?>
      <a href="/reports/class-results" class="nav-link-side <?= str_starts_with($currentPath, '/reports/class') ? 'active' : '' ?>">
        <i class="bi bi-bar-chart-line"></i><span> Class Results</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('reports.attendance_summary.view')): ?>
      <a href="/reports/attendance-summary" class="nav-link-side <?= str_starts_with($currentPath, '/reports/attendance') ? 'active' : '' ?>">
        <i class="bi bi-pie-chart"></i><span> Attendance Report</span>
      </a>
    <?php endif; ?>
    <?php if (\Core\Session::can('reports.fee_collection.view')): ?>
      <a href="/reports/fee-collection" class="nav-link-side <?= str_starts_with($currentPath, '/reports/fee') ? 'active' : '' ?>">
        <i class="bi bi-receipt"></i><span> Fee Report</span>
      </a>
    <?php endif; ?>
  </div>

  <div class="sidebar-footer">
    <div class="d-flex align-items-center gap-2">
      <div class="rounded-circle bg-primary text-white d-flex align-items-center justify-content-center" style="width:32px;height:32px;font-size:.8rem;font-weight:700">
        <?= strtoupper(substr($currentUser['first_name'] ?? 'U', 0, 1)) ?>
      </div>
      <div style="flex:1;min-width:0">
        <div style="font-size:.78rem;color:#fff;font-weight:600;white-space:nowrap;overflow:hidden;text-overflow:ellipsis">
          <?= htmlspecialchars($currentUser['first_name'] ?? '') ?> <?= htmlspecialchars($currentUser['last_name'] ?? '') ?>
        </div>
        <div style="font-size:.68rem;color:rgba(255,255,255,.5)">
          <?= htmlspecialchars($currentUser['role'] ?? '') ?>
        </div>
      </div>
      <a href="/auth/logout" class="text-white-50" title="Logout"><i class="bi bi-box-arrow-right"></i></a>
    </div>
  </div>
</nav>

<div class="main-wrap">
  <div class="topbar">
   <button class="btn btn-sm btn-light" id="sidebarToggle">
  <i class="bi bi-list fs-5"></i>
</button>
    <div class="fw-semibold text-secondary" style="font-size:.9rem"><?= htmlspecialchars($title ?? '') ?></div>
    <div class="d-flex align-items-center gap-2">
      <span class="badge bg-primary-subtle text-primary badge-role"><?= htmlspecialchars($currentUser['role'] ?? '') ?></span>
    </div>
  </div>

  <div class="page-content">

    <?php if ($flashError = \Core\Session::flash('error')): ?>
      <div class="alert alert-danger alert-dismissible fade show" role="alert">
        <i class="bi bi-exclamation-triangle-fill me-2"></i><?= htmlspecialchars($flashError) ?>
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
      </div>
    <?php endif; ?>

    <?php if ($flashSuccess = \Core\Session::flash('success')): ?>
      <div class="alert alert-success alert-dismissible fade show" role="alert">
        <i class="bi bi-check-circle-fill me-2"></i><?= htmlspecialchars($flashSuccess) ?>
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
      </div>
    <?php endif; ?>

    <?php include $viewFile; ?>

  </div>
</div>

<?php else: ?>
  <?php include $viewFile; ?>
<?php endif; ?>

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
<script>
  const sidebar    = document.getElementById('sidebar');
  const mainWrap   = document.querySelector('.main-wrap');
  const toggleBtn  = document.getElementById('sidebarToggle');

  // Restore state from localStorage
  if (localStorage.getItem('sidebarCollapsed') === 'true') {
    sidebar.classList.add('collapsed');
    mainWrap.classList.add('expanded');
  }

  toggleBtn.addEventListener('click', function () {
    sidebar.classList.toggle('collapsed');
    mainWrap.classList.toggle('expanded');
    localStorage.setItem('sidebarCollapsed', sidebar.classList.contains('collapsed'));
  });

  // Mobile: close sidebar when clicking outside
  document.addEventListener('click', function (e) {
    if (window.innerWidth < 768) {
      if (!sidebar.contains(e.target) && !toggleBtn.contains(e.target)) {
        sidebar.classList.add('collapsed');
        mainWrap.classList.add('expanded');
      }
    }
  });
</script>
</body>
</html>
