package entity

type AsteroidsReport struct {
	Date  string `json:"date" db:"date"`
	Count int    `json:"count" db:"count"`
}
