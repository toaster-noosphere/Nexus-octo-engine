package app

// Storage is a generic interface for a database.
type Storage interface {
	// TODO: describe interface methods
}

// Engine is the central core-logic struct.
type Engine struct {
  db Storage //nolint:structcheck,unused
}
