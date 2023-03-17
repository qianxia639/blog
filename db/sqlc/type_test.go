package db

import (
	"Blog/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func randomType(t *testing.T) Type {
	typeName := utils.RandomString(6)

	types, err := testQueries.InsertType(context.Background(), typeName)
	require.NoError(t, err)

	require.NotZero(t, types.ID)
	require.Equal(t, typeName, types.TypeName)

	return types
}

func TestInsertType(t *testing.T) {
	randomType(t)
}
