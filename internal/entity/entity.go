package entity

type AsteroidsReport struct {
	// TODO: Изменить Date на time.Time
	Date  string `json:"date"`
	Count int    `json:"count"`
}
