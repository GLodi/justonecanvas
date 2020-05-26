package canvas

type UseCase interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return &usecase{
		repo: r,
	}
}

func (u *usecase) Get() (*Canvas, error) {
	return u.repo.Get()
}

func (u *usecase) Update(pos int, color uint8) error {
	return u.repo.Update(pos, color)
}
