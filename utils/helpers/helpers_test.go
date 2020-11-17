package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Run("var is set", func(t *testing.T) {
		os.Setenv("TEST", "var")
		got := GetEnv("TEST", "default")
		assert.Equal(t, "var", got)
	})
	t.Run("var is not set", func(t *testing.T) {
		got := GetEnv("NOTSET", "default")
		assert.Equal(t, "default", got)
	})
}
