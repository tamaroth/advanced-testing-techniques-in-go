package mocks

import "time"

type DB interface {
	Close() error
}

type DBConnection struct{}

func (db *DBConnection) Close() error {
	time.Sleep(1 * time.Hour)
	return nil
}

var _ DB = &DBConnection{}

type Wrapper struct {
	db    DB
	value int
}

func NewWrapper(db DB, value int) *Wrapper {
	return &Wrapper{db: db, value: value}
}

func (w *Wrapper) FuncToTest() (int, error) {
	return w.value * 2, w.db.Close()
}
