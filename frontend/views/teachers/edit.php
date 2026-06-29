<?php
/** @var array $teacher */
?>

<div class="row justify-content-center">
    <div class="col-lg-8">
        <div class="card shadow-sm">
            <div class="card-header py-3">
                <i class="bi bi-person-badge me-2"></i>Edit Teacher
            </div>

            <div class="card-body">
                <form action="/teachers/<?= (int)$teacher['id'] ?>" method="POST">

                    <input type="hidden" name="_method" value="PUT">
                    <input type="hidden" name="teacher_id" value="<?= (int)$teacher['id'] ?>">

                    <?php if (function_exists('csrf_token')): ?>
                        <input type="hidden" name="csrf_token" value="<?= csrf_token() ?>">
                    <?php endif; ?>

                    <div class="mb-3">
                        <label class="form-label fw-semibold">
                            Employee No <span class="text-danger">*</span>
                        </label>
                        <input
                            type="text"
                            name="employee_no"
                            class="form-control"
                            required
                            value="<?= htmlspecialchars($teacher['employee_no'] ?? '') ?>">
                    </div>

                    <div class="row g-3">

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">TSC No</label>
                            <input
                                type="text"
                                name="tsc_no"
                                class="form-control"
                                placeholder="e.g. TSC/12345"
                                value="<?= htmlspecialchars($teacher['tsc_no'] ?? '') ?>">
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Phone</label>
                            <input
                                type="tel"
                                name="phone"
                                class="form-control"
                                placeholder="+254..."
                                value="<?= htmlspecialchars($teacher['phone'] ?? '') ?>">
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Gender</label>
                            <select name="gender" class="form-select">
                                <option value="">Select Gender</option>
                                <option value="male" <?= ($teacher['gender'] ?? '') === 'male' ? 'selected' : '' ?>>Male</option>
                                <option value="female" <?= ($teacher['gender'] ?? '') === 'female' ? 'selected' : '' ?>>Female</option>
                            </select>
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Date of Birth</label>
                            <input
                                type="date"
                                name="dob"
                                class="form-control"
                                value="<?= htmlspecialchars($teacher['dob'] ?? '') ?>">
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Qualification</label>
                            <input
                                type="text"
                                name="qualification"
                                class="form-control"
                                placeholder="e.g. B.Ed Mathematics"
                                value="<?= htmlspecialchars($teacher['qualification'] ?? '') ?>">
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Specialization</label>
                            <input
                                type="text"
                                name="specialization"
                                class="form-control"
                                placeholder="e.g. Mathematics"
                                value="<?= htmlspecialchars($teacher['specialization'] ?? '') ?>">
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Hire Date</label>
                            <input
                                type="date"
                                name="hire_date"
                                class="form-control"
                                value="<?= htmlspecialchars($teacher['hire_date'] ?? '') ?>">
                        </div>

                        <div class="col-md-6">
                            <label class="form-label fw-semibold">Employment Type</label>
                            <select name="employment_type" class="form-select">
                                <option value="permanent" <?= ($teacher['employment_type'] ?? '') === 'permanent' ? 'selected' : '' ?>>
                                    Permanent
                                </option>
                                <option value="contract" <?= ($teacher['employment_type'] ?? '') === 'contract' ? 'selected' : '' ?>>
                                    Contract
                                </option>
                                <option value="part_time" <?= ($teacher['employment_type'] ?? '') === 'part_time' ? 'selected' : '' ?>>
                                    Part Time
                                </option>
                            </select>
                        </div>

                    </div>

                    <hr class="my-4">

                    <div class="d-flex justify-content-end gap-2">
                        <a href="/teachers" class="btn btn-light">
                            <i class="bi bi-x-circle me-1"></i>Cancel
                        </a>

                        <button type="submit" class="btn btn-primary">
                            <i class="bi bi-check-circle me-1"></i>Update Teacher
                        </button>
                    </div>

                </form>
            </div>
        </div>
    </div>
</div>