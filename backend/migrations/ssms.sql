-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               9.4.0 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.11.0.7065
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table SCHOOL-DB.academic_years
CREATE TABLE IF NOT EXISTS `academic_years` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(60) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `is_current` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_ay_school` (`school_id`),
  CONSTRAINT `fk_ay_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.academic_years: ~0 rows (approximately)
REPLACE INTO `academic_years` (`id`, `school_id`, `name`, `start_date`, `end_date`, `is_current`) VALUES
	(1, 1, '2026 Academic Year', '2026-01-06', '2026-11-30', 1);

-- Dumping structure for table SCHOOL-DB.assignments
CREATE TABLE IF NOT EXISTS `assignments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `teacher_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text,
  `due_date` date NOT NULL,
  `max_marks` decimal(6,2) NOT NULL DEFAULT '100.00',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_as_school` (`school_id`),
  KEY `fk_as_class` (`class_id`),
  KEY `fk_as_subject` (`subject_id`),
  KEY `fk_as_teacher` (`teacher_id`),
  KEY `fk_as_term` (`term_id`),
  CONSTRAINT `fk_as_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_as_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_as_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`),
  CONSTRAINT `fk_as_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`),
  CONSTRAINT `fk_as_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.assignments: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.attendance
CREATE TABLE IF NOT EXISTS `attendance` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `recorded_by` bigint unsigned NOT NULL,
  `date` date NOT NULL,
  `status` enum('present','absent','late','excused') NOT NULL DEFAULT 'present',
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_attendance` (`student_id`,`date`,`term_id`),
  KEY `fk_att_class` (`class_id`),
  KEY `fk_att_term` (`term_id`),
  KEY `fk_att_recorder` (`recorded_by`),
  KEY `idx_attendance_class_date` (`class_id`,`date`),
  CONSTRAINT `fk_att_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_att_recorder` FOREIGN KEY (`recorded_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_att_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_att_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.attendance: ~2 rows (approximately)
REPLACE INTO `attendance` (`id`, `student_id`, `class_id`, `term_id`, `recorded_by`, `date`, `status`, `remark`) VALUES
	(6, 9, 2, 2, 1, '2026-06-20', 'present', 'Early'),
	(7, 11, 2, 2, 1, '2026-06-20', 'present', 'Early'),
	(8, 10, 2, 2, 1, '2026-06-20', 'present', 'Early');

-- Dumping structure for table SCHOOL-DB.audit_logs
CREATE TABLE IF NOT EXISTS `audit_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `school_id` bigint unsigned NOT NULL,
  `actor_id` bigint unsigned DEFAULT NULL,
  `action` varchar(50) NOT NULL,
  `entity` varchar(80) NOT NULL,
  `entity_id` bigint unsigned DEFAULT NULL,
  `meta` json DEFAULT NULL,
  `ip_address` varchar(45) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_al_tenant` (`tenant_id`),
  KEY `idx_al_school` (`school_id`),
  KEY `idx_al_actor` (`actor_id`),
  KEY `idx_al_entity` (`entity`,`entity_id`),
  KEY `idx_al_created` (`created_at`),
  KEY `idx_audit_tenant_entity` (`tenant_id`,`entity`,`entity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.audit_logs: ~3 rows (approximately)
REPLACE INTO `audit_logs` (`id`, `tenant_id`, `school_id`, `actor_id`, `action`, `entity`, `entity_id`, `meta`, `ip_address`, `created_at`) VALUES
	(1, 1, 1, 1, 'create', 'student', 9, '{"name": "Albert Masinde", "admission_no": "ADM001"}', NULL, '2026-06-09 13:47:59'),
	(2, 1, 1, 1, 'create', 'student', 10, '{"name": "Nancy Njeru", "admission_no": "ADM002"}', NULL, '2026-06-10 14:41:51'),
	(3, 1, 1, 1, 'create', 'student', 11, '{"name": "Allan Onyango", "admission_no": "ADM003"}', NULL, '2026-06-10 14:51:06');

-- Dumping structure for table SCHOOL-DB.classes
CREATE TABLE IF NOT EXISTS `classes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(80) NOT NULL,
  `level` varchar(40) NOT NULL,
  `stream` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_classes_school` (`school_id`),
  CONSTRAINT `fk_classes_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.classes: ~4 rows (approximately)
REPLACE INTO `classes` (`id`, `school_id`, `name`, `level`, `stream`) VALUES
	(1, 1, 'Form 1 East', 'Form 1', 'East'),
	(2, 1, 'Form 2 East', 'Form 2', 'East'),
	(3, 1, 'Form 3 East', 'Form 3', 'East'),
	(4, 1, 'Form 4 East', 'Form 4', 'East');

-- Dumping structure for table SCHOOL-DB.class_subjects
CREATE TABLE IF NOT EXISTS `class_subjects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `class_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `is_compulsory` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_cs` (`class_id`,`subject_id`),
  KEY `fk_cs_subject` (`subject_id`),
  CONSTRAINT `fk_cs_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_cs_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.class_subjects: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.class_teachers
CREATE TABLE IF NOT EXISTS `class_teachers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `class_id` bigint unsigned NOT NULL,
  `teacher_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `assigned_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_ct` (`class_id`,`term_id`),
  KEY `fk_ct_teacher` (`teacher_id`),
  KEY `fk_ct_term` (`term_id`),
  CONSTRAINT `fk_ct_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_ct_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`),
  CONSTRAINT `fk_ct_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.class_teachers: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.discipline_records
CREATE TABLE IF NOT EXISTS `discipline_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `incident_date` date NOT NULL,
  `type` enum('commendation','minor_offence','major_offence','suspension','expulsion') NOT NULL,
  `description` text NOT NULL,
  `action_taken` text,
  `recorded_by` bigint unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_dr_school` (`school_id`),
  KEY `fk_dr_student` (`student_id`),
  KEY `fk_dr_term` (`term_id`),
  KEY `fk_dr_recorder` (`recorded_by`),
  CONSTRAINT `fk_dr_recorder` FOREIGN KEY (`recorded_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_dr_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_dr_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_dr_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.discipline_records: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.exams
CREATE TABLE IF NOT EXISTS `exams` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned DEFAULT NULL,
  `name` varchar(150) NOT NULL,
  `type` enum('midterm','endterm','cat','mock','opener') NOT NULL DEFAULT 'endterm',
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_exams_school` (`school_id`),
  KEY `fk_exams_term` (`term_id`),
  KEY `fk_exams_class` (`class_id`),
  CONSTRAINT `fk_exams_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_exams_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_exams_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.exams: ~1 rows (approximately)
REPLACE INTO `exams` (`id`, `school_id`, `term_id`, `class_id`, `name`, `type`, `start_date`, `end_date`) VALUES
	(2, 1, 1, NULL, 'CAT One Term One', 'cat', '2026-06-20', '2026-06-24');

-- Dumping structure for table SCHOOL-DB.exam_results
CREATE TABLE IF NOT EXISTS `exam_results` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `exam_id` bigint unsigned NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL DEFAULT '0',
  `graded_by` bigint unsigned NOT NULL,
  `marks` decimal(6,2) NOT NULL DEFAULT '0.00',
  `max_marks` decimal(6,2) NOT NULL DEFAULT '100.00',
  `grade` varchar(5) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_result` (`exam_id`,`student_id`,`subject_id`),
  KEY `fk_er_student` (`student_id`),
  KEY `fk_er_subject` (`subject_id`),
  KEY `fk_er_grader` (`graded_by`),
  KEY `fk_er_class` (`class_id`),
  CONSTRAINT `fk_er_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_er_exam` FOREIGN KEY (`exam_id`) REFERENCES `exams` (`id`),
  CONSTRAINT `fk_er_grader` FOREIGN KEY (`graded_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_er_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_er_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.exam_results: ~27 rows (approximately)
REPLACE INTO `exam_results` (`id`, `exam_id`, `student_id`, `subject_id`, `class_id`, `graded_by`, `marks`, `max_marks`, `grade`, `remarks`) VALUES
	(6, 2, 9, 4, 2, 1, 60.00, 100.00, 'N/A', ''),
	(7, 2, 9, 5, 2, 1, 55.00, 100.00, 'N/A', ''),
	(8, 2, 9, 9, 2, 1, 78.00, 100.00, 'N/A', ''),
	(9, 2, 9, 2, 2, 1, 42.00, 100.00, 'N/A', ''),
	(10, 2, 9, 8, 2, 1, 52.00, 100.00, 'N/A', ''),
	(11, 2, 9, 7, 2, 1, 66.00, 100.00, 'N/A', ''),
	(12, 2, 9, 1, 2, 1, 68.00, 100.00, 'N/A', ''),
	(52, 2, 9, 3, 2, 1, 55.00, 100.00, 'N/A', ''),
	(53, 2, 9, 6, 2, 1, 69.00, 100.00, 'N/A', ''),
	(63, 2, 11, 4, 2, 1, 59.00, 100.00, 'N/A', ''),
	(64, 2, 11, 5, 2, 1, 69.00, 100.00, 'N/A', ''),
	(65, 2, 11, 9, 2, 1, 55.00, 100.00, 'N/A', ''),
	(66, 2, 11, 2, 2, 1, 78.00, 100.00, 'N/A', ''),
	(67, 2, 11, 8, 2, 1, 69.00, 100.00, 'N/A', ''),
	(68, 2, 11, 7, 2, 1, 58.00, 100.00, 'N/A', ''),
	(69, 2, 11, 3, 2, 1, 84.00, 100.00, 'N/A', ''),
	(70, 2, 11, 1, 2, 1, 67.00, 100.00, 'N/A', ''),
	(71, 2, 11, 6, 2, 1, 62.00, 100.00, 'N/A', ''),
	(72, 2, 10, 4, 2, 1, 69.00, 100.00, 'N/A', ''),
	(73, 2, 10, 5, 2, 1, 44.00, 100.00, 'N/A', ''),
	(74, 2, 10, 9, 2, 1, 69.00, 100.00, 'N/A', ''),
	(75, 2, 10, 2, 2, 1, 82.00, 100.00, 'N/A', ''),
	(76, 2, 10, 8, 2, 1, 68.00, 100.00, 'N/A', ''),
	(77, 2, 10, 7, 2, 1, 68.00, 100.00, 'N/A', ''),
	(78, 2, 10, 3, 2, 1, 67.00, 100.00, 'N/A', ''),
	(79, 2, 10, 1, 2, 1, 66.00, 100.00, 'N/A', ''),
	(80, 2, 10, 6, 2, 1, 66.00, 100.00, 'N/A', '');

-- Dumping structure for table SCHOOL-DB.exam_schedules
CREATE TABLE IF NOT EXISTS `exam_schedules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `exam_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `exam_date` date NOT NULL,
  `start_time` time NOT NULL,
  `end_time` time NOT NULL,
  `venue` varchar(100) DEFAULT NULL,
  `invigilator` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_es` (`exam_id`,`subject_id`,`class_id`),
  KEY `fk_es_subject` (`subject_id`),
  KEY `fk_es_class` (`class_id`),
  KEY `fk_es_invig` (`invigilator`),
  CONSTRAINT `fk_es_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_es_exam` FOREIGN KEY (`exam_id`) REFERENCES `exams` (`id`),
  CONSTRAINT `fk_es_invig` FOREIGN KEY (`invigilator`) REFERENCES `teachers` (`id`),
  CONSTRAINT `fk_es_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.exam_schedules: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.fee_discounts
CREATE TABLE IF NOT EXISTS `fee_discounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `fee_type_id` bigint unsigned DEFAULT NULL,
  `term_id` bigint unsigned DEFAULT NULL,
  `label` varchar(100) NOT NULL,
  `discount_pct` decimal(5,2) DEFAULT NULL,
  `discount_amt` decimal(12,2) DEFAULT NULL,
  `approved_by` bigint unsigned NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_fd_school` (`school_id`),
  KEY `fk_fd_student` (`student_id`),
  KEY `fk_fd_approver` (`approved_by`),
  CONSTRAINT `fk_fd_approver` FOREIGN KEY (`approved_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_fd_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_fd_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.fee_discounts: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.fee_invoices
CREATE TABLE IF NOT EXISTS `fee_invoices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `fee_type_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `status` enum('unpaid','partial','paid') NOT NULL DEFAULT 'unpaid',
  `due_date` date NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_fi_student` (`student_id`),
  KEY `fk_fi_feetype` (`fee_type_id`),
  KEY `fk_fi_term` (`term_id`),
  KEY `idx_invoices_student` (`student_id`,`status`),
  CONSTRAINT `fk_fi_feetype` FOREIGN KEY (`fee_type_id`) REFERENCES `fee_types` (`id`),
  CONSTRAINT `fk_fi_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_fi_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.fee_invoices: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.fee_payments
CREATE TABLE IF NOT EXISTS `fee_payments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `invoice_id` bigint unsigned NOT NULL,
  `amount_paid` decimal(12,2) NOT NULL,
  `method` enum('cash','mpesa','bank','cheque','other') NOT NULL DEFAULT 'cash',
  `ref_no` varchar(100) DEFAULT NULL,
  `receipt_no` varchar(60) DEFAULT NULL,
  `collected_by` bigint unsigned DEFAULT NULL,
  `mpesa_code` varchar(20) DEFAULT NULL,
  `notes` varchar(255) DEFAULT NULL,
  `paid_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_fp_invoice` (`invoice_id`),
  CONSTRAINT `fk_fp_invoice` FOREIGN KEY (`invoice_id`) REFERENCES `fee_invoices` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.fee_payments: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.fee_types
CREATE TABLE IF NOT EXISTS `fee_types` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `frequency` enum('once','termly','monthly','annual') NOT NULL DEFAULT 'termly',
  `is_mandatory` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `fk_ft_school` (`school_id`),
  CONSTRAINT `fk_ft_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.fee_types: ~4 rows (approximately)
REPLACE INTO `fee_types` (`id`, `school_id`, `name`, `amount`, `frequency`, `is_mandatory`) VALUES
	(4, 1, 'Admission Fee', 1000.00, 'once', 1),
	(5, 1, 'Tuition Fee', 12000.00, 'termly', 1),
	(6, 1, 'Laboratory Fee', 1500.00, 'annual', 1),
	(7, 1, 'Computer Fee', 500.00, 'termly', 1);

-- Dumping structure for table SCHOOL-DB.grade_scales
CREATE TABLE IF NOT EXISTS `grade_scales` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `grade` varchar(5) NOT NULL,
  `min_score` decimal(5,2) NOT NULL,
  `max_score` decimal(5,2) NOT NULL,
  `remark` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gs_school` (`school_id`),
  CONSTRAINT `fk_gs_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.grade_scales: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.library_books
CREATE TABLE IF NOT EXISTS `library_books` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `isbn` varchar(20) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `author` varchar(200) NOT NULL,
  `publisher` varchar(150) DEFAULT NULL,
  `category` varchar(80) DEFAULT NULL,
  `total_copies` int NOT NULL DEFAULT '1',
  `available` int NOT NULL DEFAULT '1',
  `added_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_lb_school` (`school_id`),
  CONSTRAINT `fk_lb_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.library_books: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.library_issues
CREATE TABLE IF NOT EXISTS `library_issues` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `book_id` bigint unsigned NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `issued_by` bigint unsigned NOT NULL,
  `issued_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `due_date` date NOT NULL,
  `returned_at` datetime DEFAULT NULL,
  `fine_amount` decimal(8,2) NOT NULL DEFAULT '0.00',
  `fine_paid` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_li_book` (`book_id`),
  KEY `fk_li_student` (`student_id`),
  KEY `fk_li_issuer` (`issued_by`),
  CONSTRAINT `fk_li_book` FOREIGN KEY (`book_id`) REFERENCES `library_books` (`id`),
  CONSTRAINT `fk_li_issuer` FOREIGN KEY (`issued_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_li_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.library_issues: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.notices
CREATE TABLE IF NOT EXISTS `notices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `author_id` bigint unsigned NOT NULL,
  `title` varchar(255) NOT NULL,
  `body` text NOT NULL,
  `audience` enum('all','teachers','parents','students') NOT NULL DEFAULT 'all',
  `published_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_notices_school` (`school_id`),
  KEY `fk_notices_author` (`author_id`),
  CONSTRAINT `fk_notices_author` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_notices_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.notices: ~0 rows (approximately)
REPLACE INTO `notices` (`id`, `school_id`, `author_id`, `title`, `body`, `audience`, `published_at`) VALUES
	(4, 1, 1, 'Test Notice', 'This is a Test Notice to everyone', 'all', '2026-06-09 13:49:30');

-- Dumping structure for table SCHOOL-DB.notifications
CREATE TABLE IF NOT EXISTS `notifications` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `school_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `type` varchar(50) NOT NULL,
  `title` varchar(255) NOT NULL,
  `body` text NOT NULL,
  `channel` enum('in_app','sms','email') NOT NULL DEFAULT 'in_app',
  `is_read` tinyint(1) NOT NULL DEFAULT '0',
  `sent_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_notif_user` (`user_id`,`is_read`),
  KEY `idx_notif_school` (`school_id`),
  CONSTRAINT `fk_notif_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.notifications: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.parents
CREATE TABLE IF NOT EXISTS `parents` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `school_id` bigint unsigned NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `occupation` varchar(100) DEFAULT NULL,
  `address` text,
  PRIMARY KEY (`id`),
  KEY `fk_parents_user` (`user_id`),
  KEY `fk_parents_school` (`school_id`),
  CONSTRAINT `fk_parents_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_parents_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.parents: ~0 rows (approximately)
REPLACE INTO `parents` (`id`, `user_id`, `school_id`, `phone`, `occupation`, `address`) VALUES
	(4, 4, 1, '', 'Business Man', 'Nairobi, Kenya');

-- Dumping structure for table SCHOOL-DB.parent_student
CREATE TABLE IF NOT EXISTS `parent_student` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` bigint unsigned NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `relationship` varchar(50) NOT NULL DEFAULT 'parent',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_ps` (`parent_id`,`student_id`),
  KEY `fk_ps_student` (`student_id`),
  CONSTRAINT `fk_ps_parent` FOREIGN KEY (`parent_id`) REFERENCES `parents` (`id`),
  CONSTRAINT `fk_ps_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.parent_student: ~0 rows (approximately)
REPLACE INTO `parent_student` (`id`, `parent_id`, `student_id`, `relationship`) VALUES
	(6, 4, 9, 'father');

-- Dumping structure for table SCHOOL-DB.pending_mpesa_pushes
CREATE TABLE IF NOT EXISTS `pending_mpesa_pushes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `invoice_id` bigint unsigned NOT NULL,
  `checkout_request_id` varchar(100) NOT NULL,
  `merchant_request_id` varchar(100) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `status` enum('pending','success','failed') NOT NULL DEFAULT 'pending',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_checkout` (`checkout_request_id`),
  KEY `fk_push_invoice` (`invoice_id`),
  CONSTRAINT `fk_push_invoice` FOREIGN KEY (`invoice_id`) REFERENCES `fee_invoices` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.pending_mpesa_pushes: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.permissions
CREATE TABLE IF NOT EXISTS `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `module` varchar(60) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.permissions: ~59 rows (approximately)
REPLACE INTO `permissions` (`id`, `name`, `description`, `module`, `created_at`) VALUES
	(25, 'students.view', 'View students', 'students', '2026-06-09 13:31:01'),
	(26, 'students.create', 'Enrol new students', 'students', '2026-06-09 13:31:01'),
	(27, 'students.edit', 'Edit student records', 'students', '2026-06-09 13:31:01'),
	(28, 'students.delete', 'Deactivate students', 'students', '2026-06-09 13:31:01'),
	(29, 'teachers.view', 'View teachers', 'teachers', '2026-06-09 13:31:01'),
	(30, 'teachers.create', 'Add teachers', 'teachers', '2026-06-09 13:31:01'),
	(31, 'teachers.edit', 'Edit teacher records', 'teachers', '2026-06-09 13:31:01'),
	(32, 'teachers.delete', 'Remove teachers', 'teachers', '2026-06-09 13:31:01'),
	(33, 'parents.view', 'View parents', 'parents', '2026-06-09 13:31:01'),
	(34, 'parents.create', 'Add parents', 'parents', '2026-06-09 13:31:01'),
	(35, 'parents.edit', 'Edit parent records', 'parents', '2026-06-09 13:31:01'),
	(36, 'classes.view', 'View classes', 'classes', '2026-06-09 13:31:01'),
	(37, 'classes.create', 'Create classes', 'classes', '2026-06-09 13:31:01'),
	(38, 'classes.edit', 'Edit classes', 'classes', '2026-06-09 13:31:01'),
	(39, 'classes.delete', 'Delete classes', 'classes', '2026-06-09 13:31:01'),
	(40, 'subjects.view', 'View subjects', 'subjects', '2026-06-09 13:31:01'),
	(41, 'subjects.create', 'Create subjects', 'subjects', '2026-06-09 13:31:01'),
	(42, 'subjects.edit', 'Edit subjects', 'subjects', '2026-06-09 13:31:01'),
	(43, 'subjects.delete', 'Delete subjects', 'subjects', '2026-06-09 13:31:01'),
	(44, 'attendance.view', 'View attendance', 'attendance', '2026-06-09 13:31:01'),
	(45, 'attendance.mark', 'Mark attendance', 'attendance', '2026-06-09 13:31:01'),
	(46, 'exams.view', 'View exams and results', 'exams', '2026-06-09 13:31:01'),
	(47, 'exams.create', 'Create exams', 'exams', '2026-06-09 13:31:01'),
	(48, 'exams.grade', 'Submit exam results', 'exams', '2026-06-09 13:31:01'),
	(49, 'exams.delete', 'Delete exams', 'exams', '2026-06-09 13:31:01'),
	(50, 'finance.view', 'View financial records', 'finance', '2026-06-09 13:31:01'),
	(51, 'finance.create', 'Record payments and invoices', 'finance', '2026-06-09 13:31:01'),
	(52, 'finance.edit', 'Edit financial records', 'finance', '2026-06-09 13:31:01'),
	(53, 'notices.view', 'View notices', 'notices', '2026-06-09 13:31:01'),
	(54, 'notices.create', 'Post notices', 'notices', '2026-06-09 13:31:01'),
	(55, 'notices.delete', 'Delete notices', 'notices', '2026-06-09 13:31:01'),
	(56, 'schools.view', 'View school details', 'schools', '2026-06-09 13:31:01'),
	(57, 'schools.create', 'Create schools', 'schools', '2026-06-09 13:31:01'),
	(58, 'schools.edit', 'Edit school details', 'schools', '2026-06-09 13:31:01'),
	(59, 'academic_years.view', 'View academic years', 'academic_years', '2026-06-09 13:31:01'),
	(60, 'academic_years.create', 'Create academic years', 'academic_years', '2026-06-09 13:31:01'),
	(61, 'academic_years.edit', 'Edit academic years', 'academic_years', '2026-06-09 13:31:01'),
	(62, 'reports.view', 'Generate reports', 'reports', '2026-06-09 13:31:01'),
	(63, 'timetable.view', 'View timetable', 'timetable', '2026-06-09 13:31:01'),
	(64, 'timetable.edit', 'Manage timetable', 'timetable', '2026-06-09 13:31:01'),
	(65, 'library.view', 'View library', 'library', '2026-06-09 13:31:01'),
	(66, 'library.manage', 'Manage library', 'library', '2026-06-09 13:31:01'),
	(67, 'discipline.view', 'View discipline records', 'discipline', '2026-06-09 13:31:01'),
	(68, 'discipline.manage', 'Manage discipline records', 'discipline', '2026-06-09 13:31:01'),
	(69, 'users.view', 'View users', 'users', '2026-06-09 13:31:01'),
	(70, 'users.create', 'Create users', 'users', '2026-06-09 13:31:01'),
	(71, 'users.edit', 'Edit users', 'users', '2026-06-09 13:31:01'),
	(72, 'terms.view', 'View terms', 'terms', '2026-06-20 20:27:12'),
	(73, 'terms.create', 'Create terms', 'terms', '2026-06-20 20:27:12'),
	(74, 'terms.edit', 'Edit terms', 'terms', '2026-06-20 20:27:12'),
	(75, 'finance.discount', 'Manage fee discounts', 'finance', '2026-06-20 20:27:12'),
	(76, 'students.promote', 'Promote students between classes', 'students', '2026-06-20 20:27:12'),
	(77, 'assignments.view', 'View assignments', 'assignments', '2026-06-21 15:33:20'),
	(78, 'assignments.create', 'Create assignments', 'assignments', '2026-06-21 15:33:20'),
	(79, 'assignments.edit', 'Delete assignments', 'assignments', '2026-06-21 15:33:20'),
	(80, 'discipline.create', 'Record discipline incident', 'discipline', '2026-06-21 15:33:20'),
	(81, 'discipline.edit', 'Delete discipline records', 'discipline', '2026-06-21 15:33:20'),
	(82, 'staff_attendance.view', 'View staff attendance', 'staff_attendance', '2026-06-21 15:33:20'),
	(83, 'staff_attendance.mark', 'Mark staff attendance', 'staff_attendance', '2026-06-21 15:33:20');

-- Dumping structure for event SCHOOL-DB.purge_token_blacklist
DELIMITER //
CREATE EVENT `purge_token_blacklist` ON SCHEDULE EVERY 1 HOUR STARTS '2026-06-27 14:04:55' ON COMPLETION NOT PRESERVE ENABLE DO DELETE FROM token_blacklist WHERE expires_at < NOW()//
DELIMITER ;

-- Dumping structure for table SCHOOL-DB.roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned DEFAULT NULL,
  `school_id` bigint unsigned DEFAULT NULL,
  `name` varchar(80) NOT NULL,
  `code` varchar(50) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `is_system` tinyint(1) NOT NULL DEFAULT '0',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_role_code` (`tenant_id`,`school_id`,`code`),
  KEY `idx_roles_school` (`school_id`),
  KEY `idx_roles_tenant` (`tenant_id`),
  CONSTRAINT `FK_roles_schools` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `FK_roles_tenants` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.roles: ~0 rows (approximately)
REPLACE INTO `roles` (`id`, `tenant_id`, `school_id`, `name`, `code`, `description`, `is_system`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 1, 1, 'Super Admin', 'superadmin', NULL, 1, 1, '2026-06-27 11:55:31', '2026-06-27 11:58:01'),
	(2, NULL, NULL, 'Admin', 'admin', NULL, 1, 1, '2026-06-27 11:55:31', '2026-06-27 11:55:31'),
	(3, NULL, NULL, 'Teacher', 'teacher', NULL, 1, 1, '2026-06-27 11:55:31', '2026-06-27 11:55:31'),
	(4, NULL, NULL, 'Parent', 'parent', NULL, 1, 1, '2026-06-27 11:55:31', '2026-06-27 11:55:31'),
	(5, NULL, NULL, 'Student', 'student', NULL, 1, 1, '2026-06-27 11:55:31', '2026-06-27 11:55:31');

-- Dumping structure for table SCHOOL-DB.role_permissions
CREATE TABLE IF NOT EXISTS `role_permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role` varchar(50) NOT NULL,
  `permission_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_role_perm` (`role`,`permission_id`),
  KEY `permission_id` (`permission_id`),
  CONSTRAINT `role_permissions_ibfk_1` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=282 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.role_permissions: ~148 rows (approximately)
REPLACE INTO `role_permissions` (`id`, `role`, `permission_id`) VALUES
	(88, 'superadmin', 60),
	(89, 'superadmin', 61),
	(90, 'superadmin', 59),
	(91, 'superadmin', 45),
	(92, 'superadmin', 44),
	(93, 'superadmin', 37),
	(94, 'superadmin', 39),
	(95, 'superadmin', 38),
	(96, 'superadmin', 36),
	(97, 'superadmin', 68),
	(98, 'superadmin', 67),
	(99, 'superadmin', 47),
	(100, 'superadmin', 49),
	(101, 'superadmin', 48),
	(102, 'superadmin', 46),
	(103, 'superadmin', 51),
	(104, 'superadmin', 52),
	(105, 'superadmin', 50),
	(106, 'superadmin', 66),
	(107, 'superadmin', 65),
	(108, 'superadmin', 54),
	(109, 'superadmin', 55),
	(110, 'superadmin', 53),
	(111, 'superadmin', 34),
	(112, 'superadmin', 35),
	(113, 'superadmin', 33),
	(114, 'superadmin', 62),
	(115, 'superadmin', 57),
	(116, 'superadmin', 58),
	(117, 'superadmin', 56),
	(118, 'superadmin', 26),
	(119, 'superadmin', 28),
	(120, 'superadmin', 27),
	(121, 'superadmin', 25),
	(122, 'superadmin', 41),
	(123, 'superadmin', 43),
	(124, 'superadmin', 42),
	(125, 'superadmin', 40),
	(126, 'superadmin', 30),
	(127, 'superadmin', 32),
	(128, 'superadmin', 31),
	(129, 'superadmin', 29),
	(130, 'superadmin', 64),
	(131, 'superadmin', 63),
	(132, 'superadmin', 70),
	(133, 'superadmin', 71),
	(134, 'superadmin', 69),
	(151, 'admin', 60),
	(152, 'admin', 61),
	(153, 'admin', 59),
	(154, 'admin', 45),
	(155, 'admin', 44),
	(156, 'admin', 37),
	(157, 'admin', 39),
	(158, 'admin', 38),
	(159, 'admin', 36),
	(160, 'admin', 68),
	(161, 'admin', 67),
	(162, 'admin', 47),
	(163, 'admin', 49),
	(164, 'admin', 48),
	(165, 'admin', 46),
	(166, 'admin', 51),
	(167, 'admin', 52),
	(168, 'admin', 50),
	(169, 'admin', 66),
	(170, 'admin', 65),
	(171, 'admin', 54),
	(172, 'admin', 55),
	(173, 'admin', 53),
	(174, 'admin', 34),
	(175, 'admin', 35),
	(176, 'admin', 33),
	(177, 'admin', 62),
	(178, 'admin', 58),
	(179, 'admin', 56),
	(180, 'admin', 26),
	(181, 'admin', 28),
	(182, 'admin', 27),
	(183, 'admin', 25),
	(184, 'admin', 41),
	(185, 'admin', 43),
	(186, 'admin', 42),
	(187, 'admin', 40),
	(188, 'admin', 30),
	(189, 'admin', 32),
	(190, 'admin', 31),
	(191, 'admin', 29),
	(192, 'admin', 64),
	(193, 'admin', 63),
	(194, 'admin', 70),
	(195, 'admin', 71),
	(196, 'admin', 69),
	(214, 'teacher', 45),
	(215, 'teacher', 44),
	(216, 'teacher', 36),
	(217, 'teacher', 68),
	(218, 'teacher', 67),
	(219, 'teacher', 48),
	(220, 'teacher', 46),
	(221, 'teacher', 65),
	(222, 'teacher', 54),
	(223, 'teacher', 53),
	(224, 'teacher', 62),
	(225, 'teacher', 25),
	(226, 'teacher', 40),
	(227, 'teacher', 63),
	(229, 'parent', 44),
	(230, 'parent', 46),
	(231, 'parent', 50),
	(232, 'parent', 53),
	(233, 'parent', 62),
	(234, 'parent', 25),
	(235, 'parent', 63),
	(236, 'student', 44),
	(237, 'student', 46),
	(238, 'student', 53),
	(239, 'student', 63),
	(240, 'superadmin', 75),
	(241, 'superadmin', 76),
	(242, 'superadmin', 73),
	(243, 'superadmin', 74),
	(244, 'superadmin', 72),
	(247, 'admin', 75),
	(248, 'admin', 76),
	(249, 'admin', 73),
	(250, 'admin', 74),
	(251, 'admin', 72),
	(254, 'teacher', 72),
	(255, 'parent', 72),
	(256, 'superadmin', 77),
	(257, 'superadmin', 78),
	(258, 'superadmin', 79),
	(259, 'superadmin', 80),
	(260, 'superadmin', 81),
	(261, 'superadmin', 82),
	(262, 'superadmin', 83),
	(263, 'admin', 77),
	(264, 'admin', 78),
	(265, 'admin', 79),
	(266, 'admin', 80),
	(267, 'admin', 81),
	(268, 'admin', 82),
	(269, 'admin', 83),
	(270, 'teacher', 78),
	(271, 'teacher', 77),
	(272, 'teacher', 82),
	(277, 'parent', 77),
	(280, 'student', 77);

-- Dumping structure for table SCHOOL-DB.schools
CREATE TABLE IF NOT EXISTS `schools` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `name` varchar(200) NOT NULL,
  `code` varchar(30) NOT NULL,
  `address` text,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `motto` varchar(255) DEFAULT NULL,
  `website` varchar(300) DEFAULT NULL,
  `county` varchar(80) DEFAULT NULL,
  `sub_county` varchar(80) DEFAULT NULL,
  `knec_code` varchar(20) DEFAULT NULL,
  `school_type` enum('day','boarding','mixed') NOT NULL DEFAULT 'day',
  `school_level` enum('primary','secondary','both') NOT NULL DEFAULT 'secondary',
  `principal_name` varchar(150) DEFAULT NULL,
  `mpesa_paybill` varchar(20) DEFAULT NULL,
  `mpesa_account` varchar(50) DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `logo_url` varchar(500) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_schools_tenant` (`tenant_id`),
  CONSTRAINT `fk_schools_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.schools: ~0 rows (approximately)
REPLACE INTO `schools` (`id`, `tenant_id`, `name`, `code`, `address`, `phone`, `email`, `motto`, `website`, `county`, `sub_county`, `knec_code`, `school_type`, `school_level`, `principal_name`, `mpesa_paybill`, `mpesa_account`, `is_active`, `logo_url`, `created_at`, `updated_at`) VALUES
	(1, 1, 'Highway Secondary School', 'HSS001', 'Mombasa Road, Nairobi, Kenya', '+254712345678', 'info@highwaysecondary.ac.ke', NULL, NULL, NULL, NULL, NULL, 'day', 'secondary', NULL, NULL, NULL, 1, NULL, '2026-06-09 13:32:54', '2026-06-09 13:32:54');

-- Dumping structure for table SCHOOL-DB.staff_attendance
CREATE TABLE IF NOT EXISTS `staff_attendance` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` bigint unsigned NOT NULL,
  `school_id` bigint unsigned NOT NULL,
  `date` date NOT NULL,
  `status` enum('present','absent','late','on_leave','sick_leave') NOT NULL DEFAULT 'present',
  `check_in` time DEFAULT NULL,
  `check_out` time DEFAULT NULL,
  `recorded_by` bigint unsigned NOT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_sa` (`teacher_id`,`date`),
  KEY `fk_sa_school` (`school_id`),
  KEY `fk_sa_recorder` (`recorded_by`),
  CONSTRAINT `fk_sa_recorder` FOREIGN KEY (`recorded_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_sa_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_sa_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.staff_attendance: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.students
CREATE TABLE IF NOT EXISTS `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `class_id` bigint unsigned NOT NULL,
  `admission_no` varchar(40) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `middle_name` varchar(100) NOT NULL DEFAULT '',
  `last_name` varchar(100) NOT NULL,
  `gender` enum('male','female','other') DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `nationality` varchar(60) DEFAULT 'Kenyan',
  `national_id` varchar(20) DEFAULT NULL,
  `religion` varchar(50) DEFAULT NULL,
  `blood_group` varchar(5) DEFAULT NULL,
  `address` text,
  `medical_notes` text,
  `photo_url` varchar(500) NOT NULL DEFAULT '',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `enrolled_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `left_date` date DEFAULT NULL,
  `left_reason` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_admission` (`school_id`,`admission_no`),
  KEY `fk_students_class` (`class_id`),
  KEY `idx_students_school_active` (`school_id`,`is_active`),
  CONSTRAINT `fk_students_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_students_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.students: ~3 rows (approximately)
REPLACE INTO `students` (`id`, `school_id`, `user_id`, `class_id`, `admission_no`, `first_name`, `middle_name`, `last_name`, `gender`, `dob`, `nationality`, `national_id`, `religion`, `blood_group`, `address`, `medical_notes`, `photo_url`, `is_active`, `enrolled_at`, `left_date`, `left_reason`) VALUES
	(9, 1, NULL, 2, 'ADM001', 'Albert', 'Wanjala', 'Masinde', 'male', '1999-08-05', 'Kenyan', NULL, NULL, NULL, NULL, NULL, '', 1, '2026-06-09 13:47:59', NULL, NULL),
	(10, 1, NULL, 2, 'ADM002', 'Nancy', 'Gitau', 'Njeru', 'female', '2005-09-05', 'Kenyan', NULL, NULL, NULL, NULL, NULL, '', 1, '2026-06-10 14:41:51', NULL, NULL),
	(11, 1, NULL, 2, 'ADM003', 'Allan', 'Omondi', 'Onyango', 'male', '2007-08-06', 'Kenyan', NULL, NULL, NULL, NULL, NULL, '', 1, '2026-06-10 14:51:06', NULL, NULL);

-- Dumping structure for table SCHOOL-DB.student_class_history
CREATE TABLE IF NOT EXISTS `student_class_history` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `academic_year_id` bigint unsigned NOT NULL,
  `promoted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `promoted_by` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sch_student` (`student_id`),
  KEY `fk_sch_class` (`class_id`),
  KEY `fk_sch_year` (`academic_year_id`),
  CONSTRAINT `fk_sch_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_sch_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_sch_year` FOREIGN KEY (`academic_year_id`) REFERENCES `academic_years` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.student_class_history: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.student_transfers
CREATE TABLE IF NOT EXISTS `student_transfers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `from_school_id` bigint unsigned NOT NULL,
  `to_school_id` bigint unsigned DEFAULT NULL,
  `transfer_date` date NOT NULL,
  `reason` varchar(255) DEFAULT NULL,
  `approved_by` bigint unsigned NOT NULL,
  `status` enum('pending','approved','rejected') NOT NULL DEFAULT 'pending',
  `notes` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_str_student` (`student_id`),
  KEY `fk_str_from` (`from_school_id`),
  KEY `fk_str_approver` (`approved_by`),
  CONSTRAINT `fk_str_approver` FOREIGN KEY (`approved_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_str_from` FOREIGN KEY (`from_school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_str_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.student_transfers: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.subjects
CREATE TABLE IF NOT EXISTS `subjects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `code` varchar(20) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `fk_subjects_school` (`school_id`),
  CONSTRAINT `fk_subjects_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.subjects: ~9 rows (approximately)
REPLACE INTO `subjects` (`id`, `school_id`, `name`, `code`, `is_active`) VALUES
	(1, 1, 'Mathematics', 'MAT', 1),
	(2, 1, 'English', 'ENG', 1),
	(3, 1, 'Kiswahili', 'KIS', 1),
	(4, 1, 'Biology', 'BIO', 1),
	(5, 1, 'Chemistry', 'CHE', 1),
	(6, 1, 'Physics', 'PHY', 1),
	(7, 1, 'History', 'HIS', 1),
	(8, 1, 'Geography', 'GEO', 1),
	(9, 1, 'Computer Studies', 'COM', 1);

-- Dumping structure for table SCHOOL-DB.teachers
CREATE TABLE IF NOT EXISTS `teachers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `school_id` bigint unsigned NOT NULL,
  `employee_no` varchar(40) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `gender` enum('male','female','other') DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `qualification` varchar(150) DEFAULT NULL,
  `tsc_no` varchar(40) DEFAULT NULL,
  `specialization` varchar(150) DEFAULT NULL,
  `hire_date` date DEFAULT NULL,
  `employment_type` enum('permanent','contract','part_time') NOT NULL DEFAULT 'permanent',
  `is_class_teacher` tinyint(1) NOT NULL DEFAULT '0',
  `national_id` varchar(20) DEFAULT NULL,
  `address` text,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `photo_url` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_teachers_user` (`user_id`),
  KEY `fk_teachers_school` (`school_id`),
  CONSTRAINT `fk_teachers_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_teachers_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.teachers: ~2 rows (approximately)
REPLACE INTO `teachers` (`id`, `user_id`, `school_id`, `employee_no`, `phone`, `gender`, `dob`, `qualification`, `tsc_no`, `specialization`, `hire_date`, `employment_type`, `is_class_teacher`, `national_id`, `address`, `is_active`, `photo_url`) VALUES
	(1, 2, 1, 'TCH001', '0712000001', 'male', NULL, 'B.Ed Mathematics', NULL, NULL, NULL, 'permanent', 0, NULL, NULL, 1, NULL),
	(2, 3, 1, 'TCH002', '0712000002', 'female', NULL, 'B.Ed Languages', NULL, NULL, NULL, 'permanent', 0, NULL, NULL, 1, NULL);

-- Dumping structure for table SCHOOL-DB.teacher_subjects
CREATE TABLE IF NOT EXISTS `teacher_subjects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_ts` (`teacher_id`,`subject_id`,`class_id`),
  KEY `fk_ts_subject` (`subject_id`),
  KEY `fk_ts_class` (`class_id`),
  CONSTRAINT `fk_ts_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_ts_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`),
  CONSTRAINT `fk_ts_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.teacher_subjects: ~6 rows (approximately)
REPLACE INTO `teacher_subjects` (`id`, `teacher_id`, `subject_id`, `class_id`) VALUES
	(7, 1, 1, 1),
	(8, 1, 1, 2),
	(9, 1, 6, 3),
	(10, 2, 2, 1),
	(11, 2, 2, 2),
	(12, 2, 3, 3);

-- Dumping structure for table SCHOOL-DB.tenants
CREATE TABLE IF NOT EXISTS `tenants` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `slug` varchar(60) NOT NULL,
  `name` varchar(150) NOT NULL,
  `domain` varchar(150) DEFAULT NULL,
  `plan` enum('free','basic','enterprise') NOT NULL DEFAULT 'free',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.tenants: ~0 rows (approximately)
REPLACE INTO `tenants` (`id`, `slug`, `name`, `domain`, `plan`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 'highway-secondary', 'Highway Secondary School', 'highwaysecondary.ac.ke', 'enterprise', 1, '2026-06-09 13:32:54', '2026-06-09 13:32:54');

-- Dumping structure for table SCHOOL-DB.terms
CREATE TABLE IF NOT EXISTS `terms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `academic_year_id` bigint unsigned NOT NULL,
  `name` varchar(60) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `is_current` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_terms_ay` (`academic_year_id`),
  CONSTRAINT `fk_terms_ay` FOREIGN KEY (`academic_year_id`) REFERENCES `academic_years` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.terms: ~3 rows (approximately)
REPLACE INTO `terms` (`id`, `academic_year_id`, `name`, `start_date`, `end_date`, `is_current`) VALUES
	(1, 1, 'Term 1', '2026-01-06', '2026-04-04', 0),
	(2, 1, 'Term 2', '2026-05-04', '2026-08-07', 1),
	(3, 1, 'Term 3', '2026-09-01', '2026-11-30', 0);

-- Dumping structure for table SCHOOL-DB.timetable_slots
CREATE TABLE IF NOT EXISTS `timetable_slots` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `teacher_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `day_of_week` tinyint NOT NULL COMMENT '1=Mon 2=Tue 3=Wed 4=Thu 5=Fri',
  `start_time` time NOT NULL,
  `end_time` time NOT NULL,
  `room` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tt_class_day` (`class_id`,`day_of_week`),
  KEY `fk_tt_school` (`school_id`),
  KEY `fk_tt_subject` (`subject_id`),
  KEY `fk_tt_teacher` (`teacher_id`),
  KEY `fk_tt_term` (`term_id`),
  CONSTRAINT `fk_tt_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_tt_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_tt_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`),
  CONSTRAINT `fk_tt_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`),
  CONSTRAINT `fk_tt_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.timetable_slots: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.token_blacklist
CREATE TABLE IF NOT EXISTS `token_blacklist` (
  `jti` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `expires_at` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`jti`),
  KEY `idx_tbl_expires` (`expires_at`),
  KEY `idx_tbl_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table SCHOOL-DB.token_blacklist: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-DB.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `first_name` varchar(150) NOT NULL DEFAULT '',
  `last_name` varchar(150) NOT NULL DEFAULT '',
  `name` varchar(150) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `role` enum('superadmin','admin','teacher','parent','student') NOT NULL DEFAULT 'student',
  `phone` varchar(20) DEFAULT NULL,
  `school_id` bigint unsigned DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_login_at` datetime DEFAULT NULL,
  `avatar_url` varchar(500) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_users_email_tenant` (`email`,`tenant_id`),
  KEY `fk_users_tenant` (`tenant_id`),
  KEY `idx_users_tenant_active` (`tenant_id`,`is_active`),
  CONSTRAINT `fk_users_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.users: ~5 rows (approximately)
REPLACE INTO `users` (`id`, `tenant_id`, `first_name`, `last_name`, `name`, `email`, `password_hash`, `role`, `phone`, `school_id`, `is_active`, `created_at`, `updated_at`, `last_login_at`, `avatar_url`, `deleted_at`, `deleted_by`) VALUES
	(1, 1, 'Albert', 'Masinde', 'Albert Masinde', 'alumasinde@gmail.com', '$2y$10$M0S9L.dHQTb8NSgSfCk1quXbQzW9/AB7j2FOTXh7AYdP6jb8RV6TC', 'superadmin', NULL, NULL, 1, '2026-06-09 13:32:54', '2026-06-27 15:00:26', '2026-06-27 15:00:26', NULL, NULL, NULL),
	(2, 1, 'John', 'Mwangi', 'John Mwangi', 'jmwangi@highwaysecondary.ac.ke', '$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i', 'teacher', NULL, NULL, 1, '2026-06-09 13:32:54', '2026-06-27 11:55:31', NULL, NULL, NULL, NULL),
	(3, 1, 'Mary', 'Wanjiku', 'Mary Wanjiku', 'mwanjiku@highwaysecondary.ac.ke', '$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i', 'teacher', NULL, NULL, 1, '2026-06-09 13:32:54', '2026-06-27 11:55:31', NULL, NULL, NULL, NULL),
	(4, 1, 'Peter', 'Otieno', 'Peter Otieno', 'potieno@gmail.com', '$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i', 'parent', NULL, NULL, 1, '2026-06-09 13:32:54', '2026-06-27 11:55:31', NULL, NULL, NULL, NULL),
	(5, 1, 'Grace', 'Akinyi', 'Grace Akinyi', 'gakinyi@gmail.com', '$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i', 'parent', NULL, NULL, 1, '2026-06-09 13:32:54', '2026-06-27 11:55:31', NULL, NULL, NULL, NULL);

-- Dumping structure for table SCHOOL-DB.user_roles
CREATE TABLE IF NOT EXISTS `user_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  `assigned_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `assigned_by` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_user_role` (`user_id`,`role_id`),
  KEY `idx_user_roles_role` (`role_id`),
  KEY `fk_user_roles_assigned` (`assigned_by`),
  CONSTRAINT `fk_user_roles_assigned` FOREIGN KEY (`assigned_by`) REFERENCES `school-ms-backup`.`users` (`id`),
  CONSTRAINT `fk_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `school-ms-backup`.`users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-DB.user_roles: ~0 rows (approximately)
REPLACE INTO `user_roles` (`id`, `user_id`, `role_id`, `assigned_at`, `assigned_by`) VALUES
	(1, 1, 1, '2026-06-27 11:55:31', NULL),
	(2, 2, 3, '2026-06-27 11:55:31', NULL),
	(3, 3, 3, '2026-06-27 11:55:31', NULL),
	(4, 4, 4, '2026-06-27 11:55:31', NULL),
	(5, 5, 4, '2026-06-27 11:55:31', NULL);

-- Dumping structure for view SCHOOL-DB.v_current_terms
-- Creating temporary table to overcome VIEW dependency errors
CREATE TABLE `v_current_terms` (
	`id` BIGINT UNSIGNED NOT NULL,
	`academic_year_id` BIGINT UNSIGNED NOT NULL,
	`school_id` BIGINT UNSIGNED NOT NULL,
	`name` VARCHAR(1) NOT NULL COLLATE 'utf8mb4_0900_ai_ci',
	`start_date` DATE NOT NULL,
	`end_date` DATE NOT NULL,
	`is_current` TINYINT(1) NOT NULL
);

-- Removing temporary table and create final VIEW structure
DROP TABLE IF EXISTS `v_current_terms`;
CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `v_current_terms` AS select `t`.`id` AS `id`,`t`.`academic_year_id` AS `academic_year_id`,`ay`.`school_id` AS `school_id`,`t`.`name` AS `name`,`t`.`start_date` AS `start_date`,`t`.`end_date` AS `end_date`,`t`.`is_current` AS `is_current` from (`terms` `t` join `academic_years` `ay` on((`ay`.`id` = `t`.`academic_year_id`))) where (`t`.`is_current` = 1)
;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
