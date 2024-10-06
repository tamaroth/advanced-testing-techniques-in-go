package setup

import "testing"

func cleanup(t *testing.T, db *DBConnection) func() {
	return func() {
		t.Log("Cleanup")
		db.Close()
	}

}

func TestWithCleanup(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{
			name: "with db",
		},
		{
			name: "without db",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db := setup(t)
			t.Cleanup(cleanup(t, db))

			// Test logic here
			t.Log("Test logic")
		})
	}
}
