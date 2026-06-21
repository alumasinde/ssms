-- ============================================================
-- Migration 003 — SchoolMS improvements
-- Run AFTER 001_init_schema.sql and 002_enterprise_upgrade.sql
-- ============================================================

SET FOREIGN_KEY_CHECKS = 0;

-- 1. Add class_id to exams (NULL = school-wide exam like mock)
ALTER TABLE exams
  ADD COLUMN IF NOT EXISTS class_id BIGINT UNSIGNED DEFAULT NULL AFTER term_id,
  ADD CONSTRAINT fk_exams_class FOREIGN KEY (class_id) REFERENCES classes(id);

-- 2. Add class_id to exam_results and update unique key
ALTER TABLE exam_results
  ADD COLUMN IF NOT EXISTS class_id BIGINT UNSIGNED NOT NULL DEFAULT 0 AFTER subject_id;

-- Drop old unique key and add new one including class_id
ALTER TABLE exam_results DROP INDEX IF EXISTS uq_result;
ALTER TABLE exam_results
  ADD UNIQUE KEY uq_result (exam_id, student_id, subject_id);

ALTER TABLE exam_results
  ADD CONSTRAINT fk_er_class FOREIGN KEY (class_id) REFERENCES classes(id);

-- 3. Ensure students table has split name columns (002 may have missed adding them)
ALTER TABLE students
  ADD COLUMN IF NOT EXISTS first_name  VARCHAR(80)  DEFAULT '' AFTER admission_no,
  ADD COLUMN IF NOT EXISTS middle_name VARCHAR(80)  DEFAULT NULL AFTER first_name,
  ADD COLUMN IF NOT EXISTS last_name   VARCHAR(80)  DEFAULT '' AFTER middle_name;

-- 4. Add school_id to terms for efficient current-term lookups (denormalised)
-- (Terms join via academic_years, but a direct index helps)
CREATE OR REPLACE VIEW v_current_terms AS
  SELECT t.id, t.academic_year_id, ay.school_id, t.name, t.start_date, t.end_date, t.is_current
  FROM terms t
  JOIN academic_years ay ON ay.id = t.academic_year_id
  WHERE t.is_current = 1;

-- 5. class_subjects — links subjects to classes (required subjects per class level)
CREATE TABLE IF NOT EXISTS class_subjects (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  class_id     BIGINT UNSIGNED NOT NULL,
  subject_id   BIGINT UNSIGNED NOT NULL,
  is_compulsory TINYINT(1) NOT NULL DEFAULT 1,
  UNIQUE KEY uq_cs (class_id, subject_id),
  CONSTRAINT fk_cs_class   FOREIGN KEY (class_id)   REFERENCES classes(id),
  CONSTRAINT fk_cs_subject FOREIGN KEY (subject_id) REFERENCES subjects(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 6. student_class_history — tracks class per academic year for historical report cards
CREATE TABLE IF NOT EXISTS student_class_history (
  id               BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  student_id       BIGINT UNSIGNED NOT NULL,
  class_id         BIGINT UNSIGNED NOT NULL,
  academic_year_id BIGINT UNSIGNED NOT NULL,
  promoted_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  promoted_by      BIGINT UNSIGNED DEFAULT NULL,
  INDEX idx_sch_student (student_id),
  CONSTRAINT fk_sch_student FOREIGN KEY (student_id)       REFERENCES students(id),
  CONSTRAINT fk_sch_class   FOREIGN KEY (class_id)         REFERENCES classes(id),
  CONSTRAINT fk_sch_year    FOREIGN KEY (academic_year_id) REFERENCES academic_years(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 7. pending_mpesa_pushes — tracks STK push CheckoutRequestID → invoice mapping
CREATE TABLE IF NOT EXISTS pending_mpesa_pushes (
  id                 BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  invoice_id         BIGINT UNSIGNED NOT NULL,
  checkout_request_id VARCHAR(100) NOT NULL,
  merchant_request_id VARCHAR(100) NOT NULL,
  phone              VARCHAR(20) NOT NULL,
  amount             DECIMAL(12,2) NOT NULL,
  status             ENUM('pending','success','failed') NOT NULL DEFAULT 'pending',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uq_checkout (checkout_request_id),
  CONSTRAINT fk_push_invoice FOREIGN KEY (invoice_id) REFERENCES fee_invoices(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 8. Add new permissions for terms, discounts, class promotion
INSERT IGNORE INTO permissions (name, description, module) VALUES
  ('terms.view',      'View terms',            'terms'),
  ('terms.create',    'Create terms',          'terms'),
  ('terms.edit',      'Edit terms',            'terms'),
  ('finance.discount','Manage fee discounts',  'finance'),
  ('students.promote','Promote students',       'students');

-- Grant new permissions to appropriate roles
INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'superadmin', id FROM permissions
  WHERE name IN ('terms.view','terms.create','terms.edit','finance.discount','students.promote');

INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'admin', id FROM permissions
  WHERE name IN ('terms.view','terms.create','terms.edit','finance.discount','students.promote');

INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'teacher', id FROM permissions WHERE name = 'terms.view';

INSERT IGNORE INTO role_permissions (role, permission_id)
  SELECT 'parent', id FROM permissions WHERE name = 'terms.view';

SET FOREIGN_KEY_CHECKS = 1;
