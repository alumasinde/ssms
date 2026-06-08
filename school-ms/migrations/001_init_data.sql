-- ssms-data.sql
-- Sample data for Highway Secondary School

SET FOREIGN_KEY_CHECKS = 0;

-- -------------------------------------------------------
-- Tenant
-- -------------------------------------------------------

INSERT INTO tenants
(id, slug, name, domain, plan, is_active)
VALUES
(1, 'highway-secondary', 'Highway Secondary School', 'highwaysecondary.ac.ke', 'enterprise', 1);

-- -------------------------------------------------------
-- School
-- -------------------------------------------------------

INSERT INTO schools
(id, tenant_id, name, code, address, phone, email)
VALUES
(
    1,
    1,
    'Highway Secondary School',
    'HSS001',
    'Mombasa Road, Nairobi, Kenya',
    '+254712345678',
    'info@highwaysecondary.ac.ke'
);

-- -------------------------------------------------------
-- Users
-- Password: Password@123
-- BCrypt hash generated sample
-- -------------------------------------------------------

INSERT INTO users
(id, tenant_id, name, email, password_hash, role)
VALUES
(1,1,'System Administrator','admin@highwaysecondary.ac.ke','$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i','superadmin'),
(2,1,'John Mwangi','jmwangi@highwaysecondary.ac.ke','$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i','teacher'),
(3,1,'Mary Wanjiku','mwanjiku@highwaysecondary.ac.ke','$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i','teacher'),
(4,1,'Peter Otieno','potieno@gmail.com','$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i','parent'),
(5,1,'Grace Akinyi','gakinyi@gmail.com','$2y$10$w4S4A6G7M4sP6X7YxFQ0He7sl2JQ2IYlQhDRD6E6V0r5F0Jm8Gg9i','parent');

-- -------------------------------------------------------
-- Academic Year
-- -------------------------------------------------------

INSERT INTO academic_years
(id, school_id, name, start_date, end_date, is_current)
VALUES
(1,1,'2026 Academic Year','2026-01-06','2026-11-30',1);

-- -------------------------------------------------------
-- Terms
-- -------------------------------------------------------

INSERT INTO terms
(id, academic_year_id, name, start_date, end_date, is_current)
VALUES
(1,1,'Term 1','2026-01-06','2026-04-04',0),
(2,1,'Term 2','2026-05-04','2026-08-07',1),
(3,1,'Term 3','2026-09-01','2026-11-30',0);

-- -------------------------------------------------------
-- Classes
-- -------------------------------------------------------

INSERT INTO classes
(id, school_id, name, level, stream)
VALUES
(1,1,'Form 1 East','Form 1','East'),
(2,1,'Form 2 East','Form 2','East'),
(3,1,'Form 3 East','Form 3','East'),
(4,1,'Form 4 East','Form 4','East');

-- -------------------------------------------------------
-- Subjects
-- -------------------------------------------------------

INSERT INTO subjects
(id, school_id, name, code)
VALUES
(1,1,'Mathematics','MAT'),
(2,1,'English','ENG'),
(3,1,'Kiswahili','KIS'),
(4,1,'Biology','BIO'),
(5,1,'Chemistry','CHE'),
(6,1,'Physics','PHY'),
(7,1,'History','HIS'),
(8,1,'Geography','GEO'),
(9,1,'Computer Studies','COM');

-- -------------------------------------------------------
-- Teachers
-- -------------------------------------------------------

INSERT INTO teachers
(id, user_id, school_id, employee_no, phone, gender, qualification)
VALUES
(1,2,1,'TCH001','0712000001','male','B.Ed Mathematics'),
(2,3,1,'TCH002','0712000002','female','B.Ed Languages');

-- -------------------------------------------------------
-- Teacher Subjects
-- -------------------------------------------------------

INSERT INTO teacher_subjects
(teacher_id, subject_id, class_id)
VALUES
(1,1,1),
(1,1,2),
(1,6,3),
(2,2,1),
(2,2,2),
(2,3,3);

-- -------------------------------------------------------
-- Students
-- -------------------------------------------------------

INSERT INTO students
(id, school_id, class_id, admission_no, name, gender, dob)
VALUES
(1,1,1,'HSS2026001','Brian Mwangi','male','2011-05-10'),
(2,1,1,'HSS2026002','Faith Achieng','female','2011-09-18'),
(3,1,2,'HSS2025001','Kevin Otieno','male','2010-03-22'),
(4,1,3,'HSS2024001','Mercy Wambui','female','2009-07-15'),
(5,1,4,'HSS2023001','Dennis Kiptoo','male','2008-11-28');

-- -------------------------------------------------------
-- Parents
-- -------------------------------------------------------

INSERT INTO parents
(id, user_id, school_id, phone, occupation, address)
VALUES
(1,4,1,'0722111111','Business Owner','Nairobi'),
(2,5,1,'0722222222','Teacher','Machakos');

-- -------------------------------------------------------
-- Parent Student
-- -------------------------------------------------------

INSERT INTO parent_student
(parent_id, student_id, relationship)
VALUES
(1,1,'father'),
(1,3,'father'),
(2,2,'mother');

-- -------------------------------------------------------
-- Attendance
-- -------------------------------------------------------

INSERT INTO attendance
(student_id, class_id, term_id, recorded_by, date, status)
VALUES
(1,1,2,2,'2026-06-01','present'),
(2,1,2,2,'2026-06-01','late'),
(3,2,2,2,'2026-06-01','present'),
(4,3,2,2,'2026-06-01','absent'),
(5,4,2,2,'2026-06-01','present');

-- -------------------------------------------------------
-- Grade Scale
-- -------------------------------------------------------

INSERT INTO grade_scales
(school_id, grade, min_score, max_score, remark)
VALUES
(1,'A',80,100,'Excellent'),
(1,'B',70,79.99,'Very Good'),
(1,'C',60,69.99,'Good'),
(1,'D',50,59.99,'Average'),
(1,'E',0,49.99,'Needs Improvement');

-- -------------------------------------------------------
-- Exams
-- -------------------------------------------------------

INSERT INTO exams
(id, school_id, term_id, name, type, start_date, end_date)
VALUES
(1,1,2,'Term 2 Mid-Term Examination','midterm','2026-06-10','2026-06-14');

-- -------------------------------------------------------
-- Exam Results
-- -------------------------------------------------------

INSERT INTO exam_results
(exam_id, student_id, subject_id, graded_by, marks, max_marks, grade)
VALUES
(1,1,1,2,82,100,'A'),
(1,1,2,3,75,100,'B'),
(1,2,1,2,68,100,'C'),
(1,2,2,3,73,100,'B'),
(1,3,1,2,88,100,'A');

-- -------------------------------------------------------
-- Fee Types
-- -------------------------------------------------------

INSERT INTO fee_types
(id, school_id, name, amount, frequency)
VALUES
(1,1,'Tuition Fee',25000,'termly'),
(2,1,'Development Fee',5000,'termly'),
(3,1,'Examination Fee',2000,'termly');

-- -------------------------------------------------------
-- Fee Invoices
-- -------------------------------------------------------

INSERT INTO fee_invoices
(id, student_id, fee_type_id, term_id, amount, status, due_date)
VALUES
(1,1,1,2,25000,'partial','2026-06-30'),
(2,2,1,2,25000,'paid','2026-06-30'),
(3,3,1,2,25000,'unpaid','2026-06-30');

-- -------------------------------------------------------
-- Fee Payments
-- -------------------------------------------------------

INSERT INTO fee_payments
(invoice_id, amount_paid, method, ref_no)
VALUES
(1,15000,'mpesa','QWE123XYZ'),
(2,25000,'bank','BANK56789');

-- -------------------------------------------------------
-- Notices
-- -------------------------------------------------------

INSERT INTO notices
(school_id, author_id, title, body, audience)
VALUES
(
    1,
    1,
    'Mid-Term Examination Notice',
    'Mid-term examinations will begin on 10 June 2026. All students should clear outstanding fees and report on time.',
    'all'
),
(
    1,
    1,
    'Parents Meeting',
    'Parents meeting scheduled for 20 June 2026 in the school hall.',
    'parents'
);

SET FOREIGN_KEY_CHECKS = 1;