package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamaroth/advanced-testing-techniques-in-go/cases/mocks/mocks"
)

func TestWrapper(t *testing.T) {
	dbMock := mocks.NewDB(t)
	dbMock.EXPECT().Close().Return(nil)
	value := 1
	expected := 2
	wrapper := NewWrapper(dbMock, value)

	actual, err := wrapper.FuncToTest()

	require.NoError(t, err)
	assert.Equal(t, expected, actual)

}
