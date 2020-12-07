package stores

import "github.com/andygodish/go-mongo-starter/utils/mongodb"

// Storer comment
type Storer interface {
	RawDB() mongodb.DatabaseHelper
}
