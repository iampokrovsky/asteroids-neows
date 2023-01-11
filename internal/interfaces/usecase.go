package interfaces

import "github.com/pokrovsky-io/neows-asteroids/internal/entity"

type Asteroids interface {
	Create(reports []entity.AsteroidsReport) error
	Get(dates []string) ([]entity.AsteroidsReport, error)
	Clear() error
}
