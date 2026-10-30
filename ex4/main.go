package main

import (
	"errors"
	"sync"
)

// Erreurs
var (
	ErrDuplicate  = errors.New("note already exists")
	ErrNotFound   = errors.New("note not found")
	ErrValidation = errors.New("validation failed")
)

// Note (avec ID)
type Note struct {
	ID      string
	Title   string
	Content string
	Tags    []string
}

// Interface
type NoteStore interface {
	Save(n *Note) error
	Get(id string) (*Note, error)
	All() []*Note
}

// MemoryStore
type MemoryStore struct {
	mu    sync.Mutex
	notes map[string]*Note
}
