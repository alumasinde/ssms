-- School Management System — Full Schema Migration
-- Run this once against an empty database
-- Engine: MySQL 8+ / MariaDB 10.5+

SET FOREIGN_KEY_CHECKS = 0;

-- -------------------------------------------------------
-- Core / Multi-tenant
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS tenants (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  slug        VARCHAR(60)  NOT NULL UNIQUE,
  name        VARCHAR(150) NOT NULL,
  domain      VARCHAR(150) DEFAULT NULL,
  plan        ENUM('free','basic','enterprise') NOT NULL DEFAULT 'free',
  is_active   TINYINT(1) NOT NULL DEFAULT 1,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS schools (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  tenant_id   BIGINT UNSIGNED NOT NULL,
  name        VARCHAR(200) NOT NULL,
  code        VARCHAR(30)  NOT NULL,
  address     TEXT,
  phone       VARCHAR(20),
  email       VARCHAR(150),
  logo_url    VARCHAR(500),
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_schools_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS users (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  tenant_id     BIGINT UNSIGNED NOT NULL,
  name          VARCHAR(150) NOT NULL,
  email         VARCHAR(150) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role          ENUM('superadmin','admin','teacher','parent','student') NOT NULL DEFAULT 'student',
  is_active     TINYINT(1) NOT NULL DEFAULT 1,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uq_users_email_tenant (email, tenant_id),
  CONSTRAINT fk_users_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------
-- Academic Structure
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS academic_years (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id   BIGINT UNSIGNED NOT NULL,
  name        VARCHAR(60) NOT NULL,
  start_date  DATE NOT NULL,
  end_date    DATE NOT NULL,
  is_current  TINYINT(1) NOT NULL DEFAULT 0,
  CONSTRAINT fk_ay_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS terms (
  id               BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  academic_year_id BIGINT UNSIGNED NOT NULL,
  name             VARCHAR(60) NOT NULL,
  start_date       DATE NOT NULL,
  end_date         DATE NOT NULL,
  is_current       TINYINT(1) NOT NULL DEFAULT 0,
  CONSTRAINT fk_terms_ay FOREIGN KEY (academic_year_id) REFERENCES academic_years(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS classes (
  id        BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id BIGINT UNSIGNED NOT NULL,
  name      VARCHAR(80)  NOT NULL,
  level     VARCHAR(40)  NOT NULL,
  stream    VARCHAR(40)  DEFAULT NULL,
  CONSTRAINT fk_classes_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS subjects (
  id        BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id BIGINT UNSIGNED NOT NULL,
  name      VARCHAR(100) NOT NULL,
  code      VARCHAR(20)  NOT NULL,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  CONSTRAINT fk_subjects_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------
-- People
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS teachers (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id       BIGINT UNSIGNED NOT NULL,
  school_id     BIGINT UNSIGNED NOT NULL,
  employee_no   VARCHAR(40) NOT NULL,
  phone         VARCHAR(20),
  gender        ENUM('male','female','other') DEFAULT NULL,
  dob           DATE DEFAULT NULL,
  qualification VARCHAR(150),
  photo_url     VARCHAR(500),
  CONSTRAINT fk_teachers_user   FOREIGN KEY (user_id)   REFERENCES users(id),
  CONSTRAINT fk_teachers_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS teacher_subjects (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  teacher_id BIGINT UNSIGNED NOT NULL,
  subject_id BIGINT UNSIGNED NOT NULL,
  class_id   BIGINT UNSIGNED NOT NULL,
  UNIQUE KEY uq_ts (teacher_id, subject_id, class_id),
  CONSTRAINT fk_ts_teacher  FOREIGN KEY (teacher_id) REFERENCES teachers(id),
  CONSTRAINT fk_ts_subject  FOREIGN KEY (subject_id) REFERENCES subjects(id),
  CONSTRAINT fk_ts_class    FOREIGN KEY (class_id)   REFERENCES classes(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS students (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id    BIGINT UNSIGNED NOT NULL,
  class_id     BIGINT UNSIGNED NOT NULL,
  admission_no VARCHAR(40) NOT NULL,
  name         VARCHAR(150) NOT NULL,
  gender       ENUM('male','female','other') DEFAULT NULL,
  dob          DATE DEFAULT NULL,
  photo_url    VARCHAR(500),
  is_active    TINYINT(1) NOT NULL DEFAULT 1,
  enrolled_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uq_admission (school_id, admission_no),
  CONSTRAINT fk_students_school FOREIGN KEY (school_id) REFERENCES schools(id),
  CONSTRAINT fk_students_class  FOREIGN KEY (class_id)  REFERENCES classes(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS parents (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id    BIGINT UNSIGNED NOT NULL,
  school_id  BIGINT UNSIGNED NOT NULL,
  phone      VARCHAR(20),
  occupation VARCHAR(100),
  address    TEXT,
  CONSTRAINT fk_parents_user   FOREIGN KEY (user_id)   REFERENCES users(id),
  CONSTRAINT fk_parents_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS parent_student (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  parent_id    BIGINT UNSIGNED NOT NULL,
  student_id   BIGINT UNSIGNED NOT NULL,
  relationship VARCHAR(50) NOT NULL DEFAULT 'parent',
  UNIQUE KEY uq_ps (parent_id, student_id),
  CONSTRAINT fk_ps_parent  FOREIGN KEY (parent_id)  REFERENCES parents(id),
  CONSTRAINT fk_ps_student FOREIGN KEY (student_id) REFERENCES students(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------
-- Attendance
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS attendance (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  student_id  BIGINT UNSIGNED NOT NULL,
  class_id    BIGINT UNSIGNED NOT NULL,
  term_id     BIGINT UNSIGNED NOT NULL,
  recorded_by BIGINT UNSIGNED NOT NULL,
  date        DATE NOT NULL,
  status      ENUM('present','absent','late','excused') NOT NULL DEFAULT 'present',
  remark      VARCHAR(255),
  UNIQUE KEY uq_attendance (student_id, date, term_id),
  CONSTRAINT fk_att_student  FOREIGN KEY (student_id)  REFERENCES students(id),
  CONSTRAINT fk_att_class    FOREIGN KEY (class_id)    REFERENCES classes(id),
  CONSTRAINT fk_att_term     FOREIGN KEY (term_id)     REFERENCES terms(id),
  CONSTRAINT fk_att_recorder FOREIGN KEY (recorded_by) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------
-- Exams & Results
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS grade_scales (
  id        BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id BIGINT UNSIGNED NOT NULL,
  grade     VARCHAR(5)  NOT NULL,
  min_score DECIMAL(5,2) NOT NULL,
  max_score DECIMAL(5,2) NOT NULL,
  remark    VARCHAR(100),
  CONSTRAINT fk_gs_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS exams (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id  BIGINT UNSIGNED NOT NULL,
  term_id    BIGINT UNSIGNED NOT NULL,
  name       VARCHAR(150) NOT NULL,
  type       ENUM('midterm','endterm','cat','mock','opener') NOT NULL DEFAULT 'endterm',
  start_date DATE NOT NULL,
  end_date   DATE NOT NULL,
  CONSTRAINT fk_exams_school FOREIGN KEY (school_id) REFERENCES schools(id),
  CONSTRAINT fk_exams_term   FOREIGN KEY (term_id)   REFERENCES terms(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS exam_results (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  exam_id    BIGINT UNSIGNED NOT NULL,
  student_id BIGINT UNSIGNED NOT NULL,
  subject_id BIGINT UNSIGNED NOT NULL,
  graded_by  BIGINT UNSIGNED NOT NULL,
  marks      DECIMAL(6,2) NOT NULL DEFAULT 0,
  max_marks  DECIMAL(6,2) NOT NULL DEFAULT 100,
  grade      VARCHAR(5),
  remarks    VARCHAR(255),
  UNIQUE KEY uq_result (exam_id, student_id, subject_id),
  CONSTRAINT fk_er_exam    FOREIGN KEY (exam_id)    REFERENCES exams(id),
  CONSTRAINT fk_er_student FOREIGN KEY (student_id) REFERENCES students(id),
  CONSTRAINT fk_er_subject FOREIGN KEY (subject_id) REFERENCES subjects(id),
  CONSTRAINT fk_er_grader  FOREIGN KEY (graded_by)  REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------
-- Finance
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS fee_types (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id    BIGINT UNSIGNED NOT NULL,
  name         VARCHAR(100) NOT NULL,
  amount       DECIMAL(12,2) NOT NULL,
  frequency    ENUM('once','termly','monthly','annual') NOT NULL DEFAULT 'termly',
  is_mandatory TINYINT(1) NOT NULL DEFAULT 1,
  CONSTRAINT fk_ft_school FOREIGN KEY (school_id) REFERENCES schools(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS fee_invoices (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  student_id  BIGINT UNSIGNED NOT NULL,
  fee_type_id BIGINT UNSIGNED NOT NULL,
  term_id     BIGINT UNSIGNED NOT NULL,
  amount      DECIMAL(12,2) NOT NULL,
  status      ENUM('unpaid','partial','paid') NOT NULL DEFAULT 'unpaid',
  due_date    DATE NOT NULL,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_fi_student  FOREIGN KEY (student_id)  REFERENCES students(id),
  CONSTRAINT fk_fi_feetype  FOREIGN KEY (fee_type_id) REFERENCES fee_types(id),
  CONSTRAINT fk_fi_term     FOREIGN KEY (term_id)     REFERENCES terms(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS fee_payments (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  invoice_id  BIGINT UNSIGNED NOT NULL,
  amount_paid DECIMAL(12,2) NOT NULL,
  method      ENUM('cash','mpesa','bank','cheque','other') NOT NULL DEFAULT 'cash',
  ref_no      VARCHAR(100),
  paid_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_fp_invoice FOREIGN KEY (invoice_id) REFERENCES fee_invoices(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------
-- Notices
-- -------------------------------------------------------

CREATE TABLE IF NOT EXISTS notices (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  school_id    BIGINT UNSIGNED NOT NULL,
  author_id    BIGINT UNSIGNED NOT NULL,
  title        VARCHAR(255) NOT NULL,
  body         TEXT NOT NULL,
  audience     ENUM('all','teachers','parents','students') NOT NULL DEFAULT 'all',
  published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_notices_school  FOREIGN KEY (school_id) REFERENCES schools(id),
  CONSTRAINT fk_notices_author  FOREIGN KEY (author_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;

-- -------------------------------------------------------
-- Seed: default superadmin tenant
-- -------------------------------------------------------
INSERT IGNORE INTO tenants (slug, name, plan, is_active) VALUES ('default', 'Default Tenant', 'enterprise', 1);
