package usecase

import (
	"asteroids-neows/internal/entity"
	"asteroids-neows/internal/errors"
	"asteroids-neows/internal/interfaces"
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

func (uc *AsteroidsUseCase) Create(reports []entity.AsteroidsReport) error {
	if err := uc.repo.Create(reports); err != nil {
		return err
	}

	return nil
}

func (uc *AsteroidsUseCase) Get(dates []string) ([]entity.AsteroidsReport, error) {
	data, err := uc.repo.Get(dates)

	switch {
	case err == errors.ErrReportsNotFound:
		data, err = uc.webAPI.Get(dates)
		if err != nil {
			return nil, err
		}

		if err = uc.repo.Create(data); err != nil {
			return nil, err
		}
	case err != nil:
		return nil, err
	}

	return data, nil
}

func (uc *AsteroidsUseCase) Clear() error {
	return uc.repo.Clear()
}
