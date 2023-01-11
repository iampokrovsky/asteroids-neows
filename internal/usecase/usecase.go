package usecase

import (
	"github.com/pokrovsky-io/neows-asteroids/internal/entity"
	"github.com/pokrovsky-io/neows-asteroids/internal/interfaces"
)

type AsteroidsUseCase struct {
	repo   interfaces.AsteroidsRepo
	webAPI interfaces.AsteroidsWebAPI
}

func New(r interfaces.AsteroidsRepo, w interfaces.AsteroidsWebAPI) *AsteroidsUseCase {
	return &AsteroidsUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *AsteroidsUseCase) Get(dates []string) ([]entity.AsteroidsReport, error) {
	data, err := uc.repo.Get(dates)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		data, err = uc.webAPI.Get(dates)
		if err != nil {
			return nil, err
		}

		if err := uc.repo.Create(data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (uc *AsteroidsUseCase) Create(reports []entity.AsteroidsReport) error {
	if err := uc.repo.Create(reports); err != nil {
		return err
	}

	return nil
}
