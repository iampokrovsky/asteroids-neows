package interfaces

import "asteroids-neows/internal/entity"

type Asteroids interface {
	Create(reports []entity.AsteroidsReport) error
	Get(dates []string) ([]entity.AsteroidsReport, error)
	Clear() error
}
