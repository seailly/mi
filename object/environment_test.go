package object

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnvironment_Get(t *testing.T) {
	env := NewEnvironment()
	val := Integer{Value: 1}
	env.Set("a", &val)
	obj, ok := env.Get("a")

	require.True(t, ok)
	require.Equal(t, &val, obj)
}

func TestEnvironment_Set(t *testing.T) {
	env := NewEnvironment()
	val := Integer{Value: 1}
	obj := env.Set("a", &val)

	require.Equal(t, &val, obj)
}
