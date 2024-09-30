package visimisidesa

type VisiMisiInput struct {
	VisiMisi string `form:"visi_misi" validate:"required"`
}

type VisiMisiUpdate struct {
	ID       uint   `form:"id" validate:"required"`
	VisiMisi string `form:"visi_misi" validate:"required"`
}
