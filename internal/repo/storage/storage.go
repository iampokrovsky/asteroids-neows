package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/pokrovsky-io/neows-asteroids/internal/entity"
)

const (
	defaultCapacity = 10
)

type Storage struct {
	cache
	psql
}

func (stg *Storage) restoreCache() error {
	reports, err := stg.psql.getAll()
	if err != nil {
		return err
	}

	stg.cache.create(reports)

	return nil
}

func New(db *sqlx.DB) (*Storage, error) {
	stg := &Storage{
		cache{make(map[string]entity.AsteroidsReport, defaultCapacity)},
		psql{db},
	}

	if len(stg.cache.data) == 0 {
		if err := stg.restoreCache(); err != nil {
			return nil, err
		}
	}

	return stg, nil
}

func (stg *Storage) getDifference(reports []entity.AsteroidsReport) (create, update []entity.AsteroidsReport) {
	create = make([]entity.AsteroidsReport, 0, len(reports))
	update = make([]entity.AsteroidsReport, 0, len(reports))

	for _, rep := range reports {
		_, isExist := stg.cache.data[rep.Date]

		if isExist {
			update = append(update, rep)
		} else {
			create = append(create, rep)
		}
	}

	return
}

func (stg *Storage) Create(reports []entity.AsteroidsReport) error {
	create, update := stg.getDifference(reports)

	if err := stg.psql.create(create); err != nil {
		return err
	}

	if err := stg.psql.update(update); err != nil {
		return err
	}

	stg.cache.create(reports)

	return nil
}

func (stg *Storage) Get(dates []string) ([]entity.AsteroidsReport, error) {
	return stg.cache.get(dates)
}

func (stg *Storage) GetLen() int {
	return len(stg.cache.data)
}

func (stg *Storage) Clear() error {
	stg.cache.clear()

	return stg.psql.clear()
}
