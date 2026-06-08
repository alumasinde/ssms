package dtos

type CreateTenantDTO struct {
	Slug   string `json:"slug" validate:"required,min=3,max=50"`
	Name   string `json:"name" validate:"required"`
	Domain string `json:"domain"`
	Plan   string `json:"plan" validate:"required,oneof=free basic enterprise"`
}

type UpdateTenantDTO struct {
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Plan     string `json:"plan"`
	IsActive *bool  `json:"is_active"`
}

type TenantResponse struct {
	ID       int64  `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Plan     string `json:"plan"`
	IsActive bool   `json:"is_active"`
}
