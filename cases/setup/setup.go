package setup

type DBConnection struct{}

func (db *DBConnection) Close() error {
	return nil
}
