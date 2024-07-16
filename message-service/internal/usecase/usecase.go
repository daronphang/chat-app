package usecase

type UseCaseService struct {
	Repository interface{}
}

func NewUseCaseService() *UseCaseService {
	return &UseCaseService{}
}
