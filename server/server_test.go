package server

import (
	"testing"

	"github.com/andygodish/go-mongo-starter/mocks"
	mongostore "github.com/andygodish/go-mongo-starter/stores/mongodb"
	"github.com/stretchr/testify/assert"
)

//Constructor is used to create a new server instance
func TestConstructor(t *testing.T) {
	t.Run("Check if a new server is created", func(t *testing.T) {
		dbHelper := &mocks.DatabaseHelper{}
		db := mongostore.NewStore(dbHelper)
		s := Constructor("api", db, 100000)
		assert.NotNil(t, s, "Expected server")
		assert.NotNil(t, s.Database, "Expected database")
		assert.Equal(t, s.appPath, "api")
	})

}
