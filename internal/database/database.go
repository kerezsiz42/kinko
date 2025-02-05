package database

type Database interface {
	Load() ([]Record, error)
	Persist([]Record) error
}
