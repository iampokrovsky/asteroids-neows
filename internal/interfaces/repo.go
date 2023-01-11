package interfaces

import "github.com/pokrovsky-io/neows-asteroids/internal/entity"

type AsteroidsRepo interface {
	Create(reports []entity.AsteroidsReport) error
	Get(dates []string) ([]entity.AsteroidsReport, error)
	Clear() error
}

type AsteroidsWebAPI interface {
	Get(dates []string) ([]entity.AsteroidsReport, error)
}
