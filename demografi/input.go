package demografi

type DemografiInput struct {
	Demografi string `form:"demografi" validate:"required"`
}

type DemografiUpdate struct {
	ID        uint   `form:"id" validate:"required"`
	Demografi string `form:"demografi" validate:"required"`
}
