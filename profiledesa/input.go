package profiledesa

type ProfileDesaInput struct {
	// Nama        string `form:"nama" validate:"required"`
	ProfileDesa string `form:"profile" validate:"required"`
}

type ProfileDesaUpdate struct {
	ID uint `form:"id" validate:"required"`
	// Nama        string `form:"nama" validate:"required"`
	ProfileDesa string `form:"profile" validate:"required"`
}
