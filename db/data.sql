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

-- Dumping structure for event school_ms.purge_token_blacklist
DELIMITER //
CREATE EVENT `purge_token_blacklist` ON SCHEDULE EVERY 1 HOUR STARTS '2026-06-27 14:04:55' ON COMPLETION NOT PRESERVE ENABLE DO DELETE FROM token_blacklist WHERE expires_at < NOW()//
DELIMITER ;

-- Dumping structure for table SCHOOL-MS-BACKUP.academic_years
CREATE TABLE IF NOT EXISTS `academic_years` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(60) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `is_current` tinyint(1) NOT NULL DEFAULT '0',
  `current_year_school` bigint GENERATED ALWAYS AS ((case when (`is_current` = 1) then `school_id` else NULL end)) STORED,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_school_year` (`school_id`,`name`),
  UNIQUE KEY `uq_current_year_school` (`current_year_school`),
  KEY `fk_ay_school` (`school_id`),
  CONSTRAINT `fk_ay_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.academic_years: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.assignments
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

-- Dumping data for table SCHOOL-MS-BACKUP.assignments: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.attendance
CREATE TABLE IF NOT EXISTS `attendance` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `recorded_by` bigint unsigned NOT NULL,
  `date` date NOT NULL,
  `status` enum('present','absent','late','excused') NOT NULL DEFAULT 'present',
  `remark` varchar(255) DEFAULT NULL,
  `school_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_attendance` (`student_id`,`date`,`term_id`),
  KEY `fk_att_class` (`class_id`),
  KEY `fk_att_term` (`term_id`),
  KEY `fk_att_recorder` (`recorded_by`),
  KEY `idx_att_school` (`school_id`),
  CONSTRAINT `fk_att_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_att_recorder` FOREIGN KEY (`recorded_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_att_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_att_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_att_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.attendance: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.audit_logs
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
  KEY `idx_al_created` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.audit_logs: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.classes
CREATE TABLE IF NOT EXISTS `classes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(80) NOT NULL,
  `level` varchar(40) NOT NULL,
  `stream` varchar(40) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_class_stream` (`school_id`,`name`,`stream`),
  KEY `fk_classes_school` (`school_id`),
  CONSTRAINT `fk_classes_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.classes: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.class_subjects
CREATE TABLE IF NOT EXISTS `class_subjects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `class_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `is_compulsory` tinyint(1) NOT NULL DEFAULT '1',
  `school_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_cs` (`class_id`,`subject_id`),
  KEY `fk_cs_subject` (`subject_id`),
  KEY `idx_cs_school` (`school_id`),
  CONSTRAINT `fk_cs_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_cs_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_cs_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.class_subjects: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.class_teachers
CREATE TABLE IF NOT EXISTS `class_teachers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `class_id` bigint unsigned NOT NULL,
  `teacher_id` bigint unsigned NOT NULL,
  `term_id` bigint unsigned NOT NULL,
  `assigned_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `school_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_ct` (`class_id`,`term_id`),
  KEY `fk_ct_teacher` (`teacher_id`),
  KEY `fk_ct_term` (`term_id`),
  KEY `idx_ct_school` (`school_id`),
  CONSTRAINT `fk_ct_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_ct_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_ct_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`),
  CONSTRAINT `fk_ct_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.class_teachers: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.discipline_records
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

-- Dumping data for table SCHOOL-MS-BACKUP.discipline_records: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.exams
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
  KEY `idx_exam_school_class` (`school_id`,`class_id`),
  CONSTRAINT `fk_exams_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_exams_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_exams_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.exams: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.exam_results
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
  `school_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_result` (`exam_id`,`student_id`,`subject_id`),
  KEY `fk_er_student` (`student_id`),
  KEY `fk_er_subject` (`subject_id`),
  KEY `fk_er_grader` (`graded_by`),
  KEY `fk_er_class` (`class_id`),
  KEY `idx_exam_school` (`school_id`),
  KEY `idx_results_student` (`student_id`,`exam_id`),
  CONSTRAINT `fk_er_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_er_exam` FOREIGN KEY (`exam_id`) REFERENCES `exams` (`id`),
  CONSTRAINT `fk_er_grader` FOREIGN KEY (`graded_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_er_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_er_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`),
  CONSTRAINT `fk_exam_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `chk_marks` CHECK ((`marks` >= 0)),
  CONSTRAINT `chk_max_marks` CHECK ((`max_marks` > 0))
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.exam_results: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.exam_schedules
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

-- Dumping data for table SCHOOL-MS-BACKUP.exam_schedules: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.fee_discounts
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

-- Dumping data for table SCHOOL-MS-BACKUP.fee_discounts: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.fee_invoices
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
  CONSTRAINT `fk_fi_feetype` FOREIGN KEY (`fee_type_id`) REFERENCES `fee_types` (`id`),
  CONSTRAINT `fk_fi_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  CONSTRAINT `fk_fi_term` FOREIGN KEY (`term_id`) REFERENCES `terms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.fee_invoices: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.fee_payments
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
  `school_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_fp_invoice` (`invoice_id`),
  KEY `idx_payment_school` (`school_id`),
  CONSTRAINT `fk_fp_invoice` FOREIGN KEY (`invoice_id`) REFERENCES `fee_invoices` (`id`),
  CONSTRAINT `fk_payment_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `chk_amount_paid` CHECK ((`amount_paid` >= 0))
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.fee_payments: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.fee_types
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

-- Dumping data for table SCHOOL-MS-BACKUP.fee_types: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.grade_scales
CREATE TABLE IF NOT EXISTS `grade_scales` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `grade` varchar(5) NOT NULL,
  `min_score` decimal(5,2) NOT NULL,
  `max_score` decimal(5,2) NOT NULL,
  `remark` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_school_grade` (`school_id`,`grade`),
  KEY `fk_gs_school` (`school_id`),
  CONSTRAINT `fk_gs_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.grade_scales: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.library_books
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

-- Dumping data for table SCHOOL-MS-BACKUP.library_books: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.library_issues
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

-- Dumping data for table SCHOOL-MS-BACKUP.library_issues: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.notices
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

-- Dumping data for table SCHOOL-MS-BACKUP.notices: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.notifications
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
  CONSTRAINT `fk_notif_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_notif_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.notifications: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.parents
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

-- Dumping data for table SCHOOL-MS-BACKUP.parents: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.parent_student
CREATE TABLE IF NOT EXISTS `parent_student` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` bigint unsigned NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `relationship` varchar(50) NOT NULL DEFAULT 'parent',
  `school_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_ps` (`parent_id`,`student_id`),
  KEY `fk_ps_student` (`student_id`),
  KEY `idx_ps_school` (`school_id`),
  CONSTRAINT `fk_ps_parent` FOREIGN KEY (`parent_id`) REFERENCES `parents` (`id`),
  CONSTRAINT `fk_ps_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_ps_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.parent_student: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.pending_mpesa_pushes
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

-- Dumping data for table SCHOOL-MS-BACKUP.pending_mpesa_pushes: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.permissions
CREATE TABLE IF NOT EXISTS `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `module` varchar(60) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.permissions: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.role_permissions
CREATE TABLE IF NOT EXISTS `role_permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `permission_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_role_permission` (`role_id`,`permission_id`),
  KEY `permission_id` (`permission_id`),
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `school-db`.`roles` (`id`),
  CONSTRAINT `role_permissions_ibfk_1` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=283 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.role_permissions: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.schools
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

-- Dumping data for table SCHOOL-MS-BACKUP.schools: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.staff_attendance
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

-- Dumping data for table SCHOOL-MS-BACKUP.staff_attendance: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.students
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
  `deleted_at` datetime DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_admission` (`school_id`,`admission_no`),
  KEY `fk_students_class` (`class_id`),
  KEY `idx_students_school_active` (`school_id`,`is_active`),
  CONSTRAINT `fk_students_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_students_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.students: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.student_class_history
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

-- Dumping data for table SCHOOL-MS-BACKUP.student_class_history: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.student_transfers
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

-- Dumping data for table SCHOOL-MS-BACKUP.student_transfers: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.subjects
CREATE TABLE IF NOT EXISTS `subjects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `code` varchar(20) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `deleted_at` datetime DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_subject_code` (`school_id`,`code`),
  KEY `fk_subjects_school` (`school_id`),
  CONSTRAINT `fk_subjects_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.subjects: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.teachers
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
  `deleted_at` datetime DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_teacher_employee` (`school_id`,`employee_no`),
  KEY `fk_teachers_user` (`user_id`),
  KEY `fk_teachers_school` (`school_id`),
  KEY `idx_teachers_school_active` (`school_id`,`is_active`),
  CONSTRAINT `fk_teachers_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_teachers_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.teachers: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.teacher_subjects
CREATE TABLE IF NOT EXISTS `teacher_subjects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` bigint unsigned NOT NULL,
  `subject_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `school_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_ts` (`teacher_id`,`subject_id`,`class_id`),
  KEY `fk_ts_subject` (`subject_id`),
  KEY `fk_ts_class` (`class_id`),
  KEY `idx_ts_school` (`school_id`),
  CONSTRAINT `fk_ts_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_ts_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_ts_subject` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`),
  CONSTRAINT `fk_ts_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.teacher_subjects: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.tenants
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

-- Dumping data for table SCHOOL-MS-BACKUP.tenants: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.terms
CREATE TABLE IF NOT EXISTS `terms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `academic_year_id` bigint unsigned NOT NULL,
  `name` varchar(60) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `is_current` tinyint(1) NOT NULL DEFAULT '0',
  `current_term_year` bigint GENERATED ALWAYS AS ((case when (`is_current` = 1) then `academic_year_id` else NULL end)) STORED,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_current_term_year` (`current_term_year`),
  KEY `fk_terms_ay` (`academic_year_id`),
  CONSTRAINT `fk_terms_ay` FOREIGN KEY (`academic_year_id`) REFERENCES `academic_years` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.terms: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.timetable_slots
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

-- Dumping data for table SCHOOL-MS-BACKUP.timetable_slots: ~0 rows (approximately)

-- Dumping structure for table SCHOOL-MS-BACKUP.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `first_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `last_name` varchar(150) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
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
  KEY `idx_users_school` (`school_id`),
  KEY `idx_users_tenant_school` (`tenant_id`,`school_id`),
  CONSTRAINT `fk_users_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`),
  CONSTRAINT `fk_users_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table SCHOOL-MS-BACKUP.users: ~0 rows (approximately)

-- Dumping structure for view SCHOOL-MS-BACKUP.v_current_terms
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
CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `school-ms-backup`.`v_current_terms` AS select `t`.`id` AS `id`,`t`.`academic_year_id` AS `academic_year_id`,`ay`.`school_id` AS `school_id`,`t`.`name` AS `name`,`t`.`start_date` AS `start_date`,`t`.`end_date` AS `end_date`,`t`.`is_current` AS `is_current` from (`school-ms-backup`.`terms` `t` join `school-ms-backup`.`academic_years` `ay` on((`ay`.`id` = `t`.`academic_year_id`))) where (`t`.`is_current` = 1)
;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
