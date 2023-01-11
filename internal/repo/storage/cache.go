package storage

import "github.com/pokrovsky-io/neows-asteroids/internal/entity"

type cache struct {
	data map[string]entity.AsteroidsReport
}

func (c *cache) create(reports []entity.AsteroidsReport) {
	for _, rep := range reports {
		c.data[rep.Date] = rep
	}
}

func (c *cache) check(dates []string) bool {
	for _, date := range dates {
		_, isExist := c.data[date]
		if !isExist {
			return false
		}
	}

	return true
}

// TODO Кэш отсортировать массив
func (c *cache) get(dates []string) ([]entity.AsteroidsReport, error) {
	if !c.check(dates) {
		return nil, ErrReportsNotFound
	}

	res := make([]entity.AsteroidsReport, 0, len(dates))

	for _, date := range dates {
		res = append(res, c.data[date])
	}

	return res, nil
}

func (c *cache) clear() {
	c.data = make(map[string]entity.AsteroidsReport, defaultCapacity)
}
