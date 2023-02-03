package interfaces

import "asteroids-neows/internal/entity"

type AsteroidsRepo interface {
	Create(reports []entity.AsteroidsReport) error
	Get(dates []string) ([]entity.AsteroidsReport, error)
	Clear() error
}

type AsteroidsWebAPI interface {
	Get(dates []string) ([]entity.AsteroidsReport, error)
}
