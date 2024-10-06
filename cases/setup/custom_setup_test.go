package setup

import "testing"

func setup(t *testing.T) *DBConnection {
	t.Log("Setup")
	return &DBConnection{}
}

func teardown(t *testing.T, db *DBConnection) {
	t.Log("Teardown")
	db.Close()
}

func TestCustomSetup(t *testing.T) {
	testCases := []struct {
		name   string
		withDb bool
	}{
		{
			name:   "with db",
			withDb: true,
		},
		{
			name:   "without db",
			withDb: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var db *DBConnection
			if testCase.withDb {
				db = setup(t)
				defer teardown(t, db)
			}

			t.Log("Test logic")
		})
	}
}
