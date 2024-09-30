package user

type RegisterUserInput struct {
	Nama           string `form:"nama" validate:"required"`
	Email          string `form:"email" validate:"required,email"`
	Password       string `form:"password" validate:"required"`
	PasswordRetype string `form:"password-Retype" validate:"required"`
	NoHp           string `form:"no_hp" validate:"required"`
	TanggalLahir   string `form:"tanggal_lahir" validate:"required"`
}

type LoginInput struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type UpdateUserInput struct {
	ID           uint   `form:"id" validate:"required"`
	Nama         string `form:"nama" validate:"required"`
	Email        string `form:"email"  validate:"required,email"`
	NoHp         string `form:"no_hp" validate:"required"`
	TanggalLahir string `form:"tanggal_lahir" validate:"required"`
	Role         string `form:"role" validate:"required"`
}

type UpdatePasswordInput struct {
	ID             uint   `form:"id" validate:"required"`
	Password       string `form:"password" validate:"required"`
	PasswordRetype string `form:"password-retype" validate:"required"`
}
