package sejarah

type SejarahInput struct {
	Sejarah string `form:"sejarah" validate:"required"`
}

type SejarahUpdate struct {
	ID      uint   `form:"id" validate:"required"`
	Sejarah string `form:"sejarah" validate:"required"`
}
