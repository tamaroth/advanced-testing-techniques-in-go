package setup

import "testing"

func TestSubtestSetupTeardown(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(t *testing.T) *DBConnection
		teardown func(db *DBConnection)
	}{
		{
			name: "with db",
			setup: func(t *testing.T) *DBConnection {
				t.Log("Setup")
				return &DBConnection{}
			},
			teardown: func(db *DBConnection) {
				t.Log("Teardown")
				db.Close()
			},
		},
		{
			name: "without db",
			setup: func(t *testing.T) *DBConnection {
				t.Log("Setup")
				return nil
			},
			teardown: func(db *DBConnection) {
				t.Log("Teardown")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db := testCase.setup(t)
			if db != nil {
				defer testCase.teardown(db)
			}

			// Test logic here
			t.Log("Test logic")
		})
	}
}
