package dtos

import "strings"

type RegisterDTO struct {
FirstName string  `json:"first_name"`
LastName  string  `json:"last_name"`
Email     string  `json:"email"`
Password  string  `json:"password"`
Phone     *string `json:"phone"`
SchoolID  *int64  `json:"school_id"` // caller supplies or injected from context
RoleCode  string  `json:"role_code"`  // matched against roles.code
}

func (d *RegisterDTO) Normalize() {
d.Email     = strings.TrimSpace(strings.ToLower(d.Email))
d.FirstName = strings.TrimSpace(d.FirstName)
d.LastName  = strings.TrimSpace(d.LastName)
d.RoleCode  = strings.TrimSpace(strings.ToLower(d.RoleCode))
}

func (d *RegisterDTO) Validate() map[string]string {
errs := map[string]string{}
if d.FirstName == ""    { errs["first_name"] = "first name is required" }
if d.LastName == ""     { errs["last_name"]  = "last name is required" }
if d.Email == ""        { errs["email"]      = "email is required" }
if len(d.Password) < 8  { errs["password"]   = "password must be at least 8 characters" }
if d.RoleCode == ""     { errs["role_code"]  = "role_code is required" }
if len(errs) > 0 { return errs }
return nil
}

type UserResponse struct {
ID        int64    `json:"id"`
FirstName string   `json:"first_name"`
LastName  string   `json:"last_name"`
Name      string   `json:"name"` // computed: first_name + " " + last_name
Email     string   `json:"email"`
Phone     *string  `json:"phone,omitempty"`
SchoolID  *int64   `json:"school_id,omitempty"`
TenantID  int64    `json:"tenant_id"`
Roles     []string `json:"roles"`
IsActive  bool     `json:"is_active"`
}

type ChangePasswordDTO struct {
OldPassword string `json:"old_password"`
NewPassword string `json:"new_password"`
}

func (d *ChangePasswordDTO) Validate() map[string]string {
errs := map[string]string{}
if d.OldPassword == ""    { errs["old_password"] = "old password is required" }
if len(d.NewPassword) < 8 { errs["new_password"] = "new password must be at least 8 characters" }
if len(errs) > 0 { return errs }
return nil
}

type UpdateRoleDTO struct {
RoleCode string `json:"role_code"`
}