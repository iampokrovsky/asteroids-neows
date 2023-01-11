package interfaces

import "github.com/pokrovsky-io/neows-asteroids/internal/entity"

type Asteroids interface {
	// TODO: Подумать, что должны возвращать методы
	Get(dates ...string) ([]entity.AsteroidsReport, error)
	Create(reports ...entity.AsteroidsReport) error
	// TODO: Нужно ли делать метод Update?
}
