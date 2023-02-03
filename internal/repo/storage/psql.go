package storage

import (
	"asteroids-neows/internal/entity"
	"github.com/jmoiron/sqlx"
)

type psql struct {
	db *sqlx.DB
}

func (psql *psql) create(reports []entity.AsteroidsReport) error {
	if len(reports) == 0 {
		return nil
	}

	query := "INSERT INTO neo_count (date, count) VALUES (:date, :count)"
	_, err := psql.db.NamedExec(query, reports)

	return err
}

func (psql *psql) getAll() ([]entity.AsteroidsReport, error) {
	res := make([]entity.AsteroidsReport, 0, defaultCapacity)

	if err := psql.db.Select(&res, "SELECT * FROM neo_count"); err != nil {
		return nil, err
	}

	return res, nil
}

func (psql *psql) update(reports []entity.AsteroidsReport) error {
	if len(reports) == 0 {
		return nil
	}

	tx, err := psql.db.Begin()
	if err != nil {
		return err
	}

	query := "UPDATE neo_count SET count=$1 WHERE date=$2"
	for _, rep := range reports {
		_, err := psql.db.Exec(query, rep.Count, rep.Date)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (psql *psql) clear() error {
	query := "DELETE FROM neo_count"
	_, err := psql.db.Exec(query)

	return err
}
