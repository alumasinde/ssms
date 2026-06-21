-- ============================================================
-- SSMS Enterprise Upgrade Migration
-- Run AFTER 001_init_schema.sql
-- Engine: MySQL 8+ / MariaDB 10.5+
-- ============================================================

SET FOREIGN_KEY_CHECKS = 0;

-- ============================================================
-- 1. AUDIT LOGS
-- ============================================================
CREATE TABLE IF NOT EXISTS audit_logs (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  tenant_id  BIGINT UNSIGNED NOT NULL,
  school_id  BIGINT UNSIGNED NOT NULL,
  actor_id   BIGINT UNSIGNED,                      -- NULL = system action
  action     VARCHAR(50)  NOT NULL,                -- create|update|delete|login|mark|grade …
  entity     VARCHAR(80)  NOT NULL,                -- student|teacher|payment|attendance …
  entity_id  BIGINT UNSIGNED,
  meta       JSON,
  ip_address VARCHAR(45),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_al_tenant   (tenant_id),
  INDEX idx_al_school   (school_id),
  INDEX idx_al_actor    (actor_id),
  INDEX idx_al_entity   (entity, entity_id),
  INDEX idx_al_created  (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 2. PERMISSIONS TABLES (already exist in ssms.sql — add
--    missing columns and ensure correct structure)
-- ============================================================
ALTER TABLE permissions
  ADD COLUMN IF NOT EXISTS module     VARCHAR(60) NOT NULL DEFAULT '' AFTER description,
  ADD COLUMN IF NOT EXISTS created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP AFTER module;

-- ============================================================
-- 3. SCHOOLS — extra fields for Kenyan school context
-- ============================================================
ALTER TABLE schools
  ADD COLUMN IF NOT EXISTS motto          VARCHAR(255)  DEFAULT NULL AFTER email,
  ADD COLUMN IF NOT EXISTS website        VARCHAR(300)  DEFAULT NULL AFTER motto,
  ADD COLUMN IF NOT EXISTS county         VARCHAR(80)   DEFAULT NULL AFTER website,
  ADD COLUMN IF NOT EXISTS sub_county     VARCHAR(80)   DEFAULT NULL AFTER county,
  ADD COLUMN IF NOT EXISTS knec_code      VARCHAR(20)   DEFAULT NULL AFTER sub_county,
  ADD COLUMN IF NOT EXISTS school_type    ENUM('day','boarding','mixed') NOT NULL DEFAULT 'day' AFTER knec_code,
  ADD COLUMN IF NOT EXISTS school_level   ENUM('primary','secondary','both') NOT NULL DEFAULT 'secondary' AFTER school_type,
  ADD COLUMN IF NOT EXISTS principal_name VARCHAR(150)  DEFAULT NULL AFTER school_level,
  ADD COLUMN IF NOT EXISTS mpesa_paybill  VARCHAR(20)   DEFAULT NULL AFTER principal_name,
  ADD COLUMN IF NOT EXISTS mpesa_account  VARCHAR(50)   DEFAULT NULL AFTER mpesa_paybill,
  ADD COLUMN IF NOT EXISTS is_active      TINYINT(1)   NOT NULL DEFAULT 1 AFTER mpesa_account;

-- ============================================================
-- 4. USERS — additional profile fields
-- ============================================================
ALTER TABLE users
  ADD COLUMN IF NOT EXISTS phone         VARCHAR(20)  DEFAULT NULL AFTER role,
  ADD COLUMN IF NOT EXISTS school_id     BIGINT UNSIGNED DEFAULT NULL AFTER phone,
  ADD COLUMN IF NOT EXISTS last_login_at DATETIME     DEFAULT NULL AFTER updated_at,
  ADD COLUMN IF NOT EXISTS avatar_url    VARCHAR(500) DEFAULT NULL AFTER last_login_at;

-- ============================================================
-- 5. TEACHERS — enriched HR fields
-- ============================================================
ALTER TABLE teachers
  ADD COLUMN IF NOT EXISTS tsc_no            VARCHAR(40)  DEFAULT NULL AFTER qualification,
  ADD COLUMN IF NOT EXISTS specialization    VARCHAR(150) DEFAULT NULL AFTER tsc_no,
  ADD COLUMN IF NOT EXISTS hire_date         DATE         DEFAULT NULL AFTER specialization,
  ADD COLUMN IF NOT EXISTS employment_type   ENUM('permanent','contract','part_time') NOT NULL DEFAULT 'permanent' AFTER hire_date,
  ADD COLUMN IF NOT EXISTS is_class_teacher  TINYINT(1) NOT NULL DEFAULT 0 AFTER employment_type,
  ADD COLUMN IF NOT EXISTS national_id       VARCHAR(20)  DEFAULT NULL AFTER is_class_teacher,
  ADD COLUMN IF NOT EXISTS address           TEXT         DEFAULT NULL AFTER national_id,
  ADD COLUMN IF NOT EXISTS is_active         TINYINT(1) NOT NULL DEFAULT 1 AFTER address;

-- ============================================================
-- 6. STUDENTS — additional fields from ssms.sql + extras
-- ============================================================
-- ssms.sql already split name into first/middle/last — align migration schema
ALTER TABLE students
  ADD COLUMN IF NOT EXISTS user_id       BIGINT UNSIGNED DEFAULT NULL AFTER school_id,
  ADD COLUMN IF NOT EXISTS term_id       BIGINT UNSIGNED DEFAULT NULL AFTER class_id,
  ADD COLUMN IF NOT EXISTS nationality   VARCHAR(60)  DEFAULT 'Kenyan' AFTER dob,
  ADD COLUMN IF NOT EXISTS national_id   VARCHAR(20)  DEFAULT NULL AFTER nationality,
  ADD COLUMN IF NOT EXISTS religion      VARCHAR(50)  DEFAULT NULL AFTER national_id,
  ADD COLUMN IF NOT EXISTS blood_group   VARCHAR(5)   DEFAULT NULL AFTER religion,
  ADD COLUMN IF NOT EXISTS address       TEXT         DEFAULT NULL AFTER blood_group,
  ADD COLUMN IF NOT EXISTS medical_notes TEXT         DEFAULT NULL AFTER address,
  ADD COLUMN IF NOT EXISTS left_date     DATE         DEFAULT NULL AFTER enrolled_at,
  ADD COLUMN IF NOT EXISTS left_reason   VARCHAR(255) DEFAULT NULL AFTER left_date;

-- ============================================================
-- 7. CLASS TEACHER ASSIGNMENT
-- ============================================================
CREATE TABLE IF NOT EXISTS class_teachers (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  class_id    BIGINT UNSIGNED NOT NULL,
  teacher_id  BIGINT UNSIGNED NOT NULL,
  term_id     BIGINT UNSIGNED NOT NULL,
  assigned_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uq_ct (class_id, term_id),
  CONSTRAINT fk_ct_class   FOREIGN KEY (class_id)   REFERENCES classes(id),
  CONSTRAINT fk_ct_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id),
  CONSTRAINT fk_ct_term    FOREIGN KEY (term_id)    REFERENCES terms(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 8. TIMETABLE
-- ============================================================
CREATE TABLE IF NOT EXISTS timetable_slots (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id   BIGINT UNSIGNED NOT NULL,
  class_id    BIGINT UNSIGNED NOT NULL,
  subject_id  BIGINT UNSIGNED NOT NULL,
  teacher_id  BIGINT UNSIGNED NOT NULL,
  term_id     BIGINT UNSIGNED NOT NULL,
  day_of_week TINYINT NOT NULL COMMENT '1=Mon … 5=Fri',
  start_time  TIME NOT NULL,
  end_time    TIME NOT NULL,
  room        VARCHAR(60) DEFAULT NULL,
  CONSTRAINT fk_tt_school  FOREIGN KEY (school_id)  REFERENCES schools(id),
  CONSTRAINT fk_tt_class   FOREIGN KEY (class_id)   REFERENCES classes(id),
  CONSTRAINT fk_tt_subject FOREIGN KEY (subject_id) REFERENCES subjects(id),
  CONSTRAINT fk_tt_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id),
  CONSTRAINT fk_tt_term    FOREIGN KEY (term_id)    REFERENCES terms(id),
  INDEX idx_tt_class_day (class_id, day_of_week)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 9. EXAM SCHEDULES (per subject scheduling within an exam)
-- ============================================================
CREATE TABLE IF NOT EXISTS exam_schedules (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  exam_id      BIGINT UNSIGNED NOT NULL,
  subject_id   BIGINT UNSIGNED NOT NULL,
  class_id     BIGINT UNSIGNED NOT NULL,
  exam_date    DATE NOT NULL,
  start_time   TIME NOT NULL,
  end_time     TIME NOT NULL,
  venue        VARCHAR(100) DEFAULT NULL,
  invigilator  BIGINT UNSIGNED DEFAULT NULL,
  UNIQUE KEY uq_es (exam_id, subject_id, class_id),
  CONSTRAINT fk_es_exam      FOREIGN KEY (exam_id)     REFERENCES exams(id),
  CONSTRAINT fk_es_subject   FOREIGN KEY (subject_id)  REFERENCES subjects(id),
  CONSTRAINT fk_es_class     FOREIGN KEY (class_id)    REFERENCES classes(id),
  CONSTRAINT fk_es_invig     FOREIGN KEY (invigilator) REFERENCES teachers(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 10. STAFF ATTENDANCE
-- ============================================================
CREATE TABLE IF NOT EXISTS staff_attendance (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  teacher_id  BIGINT UNSIGNED NOT NULL,
  school_id   BIGINT UNSIGNED NOT NULL,
  date        DATE NOT NULL,
  status      ENUM('present','absent','late','on_leave','sick_leave') NOT NULL DEFAULT 'present',
  check_in    TIME DEFAULT NULL,
  check_out   TIME DEFAULT NULL,
  recorded_by BIGINT UNSIGNED NOT NULL,
  remark      VARCHAR(255) DEFAULT NULL,
  UNIQUE KEY uq_sa (teacher_id, date),
  CONSTRAINT fk_sa_teacher  FOREIGN KEY (teacher_id)  REFERENCES teachers(id),
  CONSTRAINT fk_sa_school   FOREIGN KEY (school_id)   REFERENCES schools(id),
  CONSTRAINT fk_sa_recorder FOREIGN KEY (recorded_by) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 11. LIBRARY
-- ============================================================
CREATE TABLE IF NOT EXISTS library_books (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id     BIGINT UNSIGNED NOT NULL,
  isbn          VARCHAR(20) DEFAULT NULL,
  title         VARCHAR(255) NOT NULL,
  author        VARCHAR(200) NOT NULL,
  publisher     VARCHAR(150) DEFAULT NULL,
  category      VARCHAR(80) DEFAULT NULL,
  total_copies  INT NOT NULL DEFAULT 1,
  available     INT NOT NULL DEFAULT 1,
  added_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_lb_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS library_issues (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  book_id     BIGINT UNSIGNED NOT NULL,
  student_id  BIGINT UNSIGNED NOT NULL,
  issued_by   BIGINT UNSIGNED NOT NULL,
  issued_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  due_date    DATE NOT NULL,
  returned_at DATETIME DEFAULT NULL,
  fine_amount DECIMAL(8,2) NOT NULL DEFAULT 0,
  fine_paid   TINYINT(1) NOT NULL DEFAULT 0,
  CONSTRAINT fk_li_book    FOREIGN KEY (book_id)    REFERENCES library_books(id),
  CONSTRAINT fk_li_student FOREIGN KEY (student_id) REFERENCES students(id),
  CONSTRAINT fk_li_issuer  FOREIGN KEY (issued_by)  REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 12. STUDENT TRANSFERS
-- ============================================================
CREATE TABLE IF NOT EXISTS student_transfers (
  id               BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  student_id       BIGINT UNSIGNED NOT NULL,
  from_school_id   BIGINT UNSIGNED NOT NULL,
  to_school_id     BIGINT UNSIGNED DEFAULT NULL,  -- NULL = left school system
  transfer_date    DATE NOT NULL,
  reason           VARCHAR(255) DEFAULT NULL,
  approved_by      BIGINT UNSIGNED NOT NULL,
  status           ENUM('pending','approved','rejected') NOT NULL DEFAULT 'pending',
  notes            TEXT DEFAULT NULL,
  created_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_st_student FOREIGN KEY (student_id)     REFERENCES students(id),
  CONSTRAINT fk_st_from    FOREIGN KEY (from_school_id) REFERENCES schools(id),
  CONSTRAINT fk_st_approver FOREIGN KEY (approved_by)  REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 13. NOTIFICATIONS
-- ============================================================
CREATE TABLE IF NOT EXISTS notifications (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  tenant_id   BIGINT UNSIGNED NOT NULL,
  school_id   BIGINT UNSIGNED NOT NULL,
  user_id     BIGINT UNSIGNED NOT NULL,
  type        VARCHAR(50) NOT NULL,              -- fee_reminder|exam_result|attendance|notice …
  title       VARCHAR(255) NOT NULL,
  body        TEXT NOT NULL,
  channel     ENUM('in_app','sms','email') NOT NULL DEFAULT 'in_app',
  is_read     TINYINT(1) NOT NULL DEFAULT 0,
  sent_at     DATETIME DEFAULT NULL,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_notif_user   (user_id, is_read),
  INDEX idx_notif_school (school_id),
  CONSTRAINT fk_notif_user FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 14. FEE DISCOUNTS / BURSARIES
-- ============================================================
CREATE TABLE IF NOT EXISTS fee_discounts (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id    BIGINT UNSIGNED NOT NULL,
  student_id   BIGINT UNSIGNED NOT NULL,
  fee_type_id  BIGINT UNSIGNED DEFAULT NULL,  -- NULL = applies to all fee types
  term_id      BIGINT UNSIGNED DEFAULT NULL,  -- NULL = recurring
  label        VARCHAR(100) NOT NULL,          -- 'Bursary','Scholarship','Staff Child' …
  discount_pct DECIMAL(5,2) DEFAULT NULL,      -- % off
  discount_amt DECIMAL(12,2) DEFAULT NULL,     -- fixed KES amount
  approved_by  BIGINT UNSIGNED NOT NULL,
  is_active    TINYINT(1) NOT NULL DEFAULT 1,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_fd_school   FOREIGN KEY (school_id)   REFERENCES schools(id),
  CONSTRAINT fk_fd_student  FOREIGN KEY (student_id)  REFERENCES students(id),
  CONSTRAINT fk_fd_approver FOREIGN KEY (approved_by) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 15. FEE PAYMENTS — add receipt & collected_by
-- ============================================================
ALTER TABLE fee_payments
  ADD COLUMN IF NOT EXISTS receipt_no   VARCHAR(60) DEFAULT NULL AFTER ref_no,
  ADD COLUMN IF NOT EXISTS collected_by BIGINT UNSIGNED DEFAULT NULL AFTER receipt_no,
  ADD COLUMN IF NOT EXISTS mpesa_code   VARCHAR(20) DEFAULT NULL AFTER collected_by,
  ADD COLUMN IF NOT EXISTS notes        VARCHAR(255) DEFAULT NULL AFTER mpesa_code;

-- ============================================================
-- 16. BEHAVIOUR / DISCIPLINARY RECORDS
-- ============================================================
CREATE TABLE IF NOT EXISTS discipline_records (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id    BIGINT UNSIGNED NOT NULL,
  student_id   BIGINT UNSIGNED NOT NULL,
  term_id      BIGINT UNSIGNED NOT NULL,
  incident_date DATE NOT NULL,
  type         ENUM('commendation','minor_offence','major_offence','suspension','expulsion') NOT NULL,
  description  TEXT NOT NULL,
  action_taken TEXT DEFAULT NULL,
  recorded_by  BIGINT UNSIGNED NOT NULL,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_dr_school   FOREIGN KEY (school_id)   REFERENCES schools(id),
  CONSTRAINT fk_dr_student  FOREIGN KEY (student_id)  REFERENCES students(id),
  CONSTRAINT fk_dr_term     FOREIGN KEY (term_id)     REFERENCES terms(id),
  CONSTRAINT fk_dr_recorder FOREIGN KEY (recorded_by) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 17. HOMEWORK / ASSIGNMENTS
-- ============================================================
CREATE TABLE IF NOT EXISTS assignments (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id    BIGINT UNSIGNED NOT NULL,
  class_id     BIGINT UNSIGNED NOT NULL,
  subject_id   BIGINT UNSIGNED NOT NULL,
  teacher_id   BIGINT UNSIGNED NOT NULL,
  term_id      BIGINT UNSIGNED NOT NULL,
  title        VARCHAR(255) NOT NULL,
  description  TEXT DEFAULT NULL,
  due_date     DATE NOT NULL,
  max_marks    DECIMAL(6,2) NOT NULL DEFAULT 100,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_as_school  FOREIGN KEY (school_id)  REFERENCES schools(id),
  CONSTRAINT fk_as_class   FOREIGN KEY (class_id)   REFERENCES classes(id),
  CONSTRAINT fk_as_subject FOREIGN KEY (subject_id) REFERENCES subjects(id),
  CONSTRAINT fk_as_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id),
  CONSTRAINT fk_as_term    FOREIGN KEY (term_id)    REFERENCES terms(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 18. PERMISSIONS — full enterprise seed
-- ============================================================
INSERT IGNORE INTO permissions (name, description, module) VALUES
  -- Students
  ('students.view',    'View students',          'students'),
  ('students.create',  'Enrol new students',     'students'),
  ('students.edit',    'Edit student records',   'students'),
  ('students.delete',  'Deactivate students',    'students'),
  -- Teachers
  ('teachers.view',    'View teachers',          'teachers'),
  ('teachers.create',  'Add teachers',           'teachers'),
  ('teachers.edit',    'Edit teacher records',   'teachers'),
  ('teachers.delete',  'Remove teachers',        'teachers'),
  -- Parents
  ('parents.view',     'View parents',           'parents'),
  ('parents.create',   'Add parents',            'parents'),
  ('parents.edit',     'Edit parent records',    'parents'),
  -- Classes
  ('classes.view',     'View classes',           'classes'),
  ('classes.create',   'Create classes',         'classes'),
  ('classes.edit',     'Edit classes',           'classes'),
  ('classes.delete',   'Delete classes',         'classes'),
  -- Subjects
  ('subjects.view',    'View subjects',          'subjects'),
  ('subjects.create',  'Create subjects',        'subjects'),
  ('subjects.edit',    'Edit subjects',          'subjects'),
  ('subjects.delete',  'Delete subjects',        'subjects'),
  -- Attendance
  ('attendance.view',  'View attendance',        'attendance'),
  ('attendance.mark',  'Mark attendance',        'attendance'),
  -- Exams
  ('exams.view',       'View exams & results',   'exams'),
  ('exams.create',     'Create exams',           'exams'),
  ('exams.grade',      'Submit exam results',    'exams'),
  ('exams.delete',     'Delete exams',           'exams'),
  -- Finance
  ('finance.view',     'View financial records', 'finance'),
  ('finance.create',   'Record payments/invoices','finance'),
  ('finance.edit',     'Edit financial records', 'finance'),
  -- Notices
  ('notices.view',     'View notices',           'notices'),
  ('notices.create',   'Post notices',           'notices'),
  ('notices.delete',   'Delete notices',         'notices'),
  -- Schools
  ('schools.view',     'View school details',    'schools'),
  ('schools.create',   'Create schools',         'schools'),
  ('schools.edit',     'Edit school details',    'schools'),
  -- Academic Years
  ('academic_years.view',   'View academic years',  'academic_years'),
  ('academic_years.create', 'Create academic years','academic_years'),
  ('academic_years.edit',   'Edit academic years',  'academic_years'),
  -- Reports
  ('reports.view',     'Generate reports',       'reports'),
  -- Timetable
  ('timetable.view',   'View timetable',         'timetable'),
  ('timetable.edit',   'Manage timetable',       'timetable'),
  -- Library
  ('library.view',     'View library',           'library'),
  ('library.manage',   'Manage library',         'library'),
  -- Discipline
  ('discipline.view',  'View discipline records','discipline'),
  ('discipline.manage','Manage discipline',      'discipline'),
  -- Users
  ('users.view',       'View users',             'users'),
  ('users.create',     'Create users',           'users'),
  ('users.edit',       'Edit users',             'users');

-- ============================================================
-- 19. ROLE PERMISSIONS — enterprise defaults
-- ============================================================
-- superadmin: everything
INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'superadmin', id FROM permissions;

-- admin: all except tenant-level destructive ops
INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'admin', id FROM permissions
  WHERE name NOT IN ('schools.create');

-- teacher: scoped read + their own operations
INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'teacher', id FROM permissions
  WHERE name IN (
    'students.view','attendance.view','attendance.mark',
    'exams.view','exams.grade',
    'subjects.view','classes.view',
    'notices.view','notices.create',
    'timetable.view',
    'reports.view',
    'library.view',
    'discipline.view','discipline.manage',
    'assignments.view','assignments.manage'
  );

-- parent: read-only access to their children's data
INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'parent', id FROM permissions
  WHERE name IN (
    'students.view','attendance.view',
    'exams.view','notices.view',
    'finance.view','reports.view',
    'timetable.view'
  );

-- student: minimal self-service
INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'student', id FROM permissions
  WHERE name IN (
    'notices.view','exams.view','attendance.view','timetable.view'
  );

SET FOREIGN_KEY_CHECKS = 1;
