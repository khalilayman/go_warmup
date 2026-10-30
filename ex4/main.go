package main

import (
	"errors"
	"fmt"
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

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{notes: make(map[string]*Note)}
}

func (s *MemoryStore) Save(n *Note) error {
	// validation
	if n.Title == "" {
		return fmt.Errorf("%w: title required", ErrValidation)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// doublon
	if _, exists := s.notes[n.ID]; exists {
		return fmt.Errorf("%w: id=%s", ErrDuplicate, n.ID)
	}

	s.notes[n.ID] = n
	return nil
}

func (s *MemoryStore) Get(id string) (*Note, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	note, ok := s.notes[id]
	if !ok {
		return nil, fmt.Errorf("%w: id=%s", ErrNotFound, id)
	}
	return note, nil
}

func (s *MemoryStore) All() []*Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := []*Note{}
	for _, n := range s.notes {
		result = append(result, n)
	}
	return result
}

// main de test
func main() {
	store := NewMemoryStore()

	store.Save(&Note{ID: "1", Title: "Go"})      // ok
	store.Save(&Note{ID: "1", Title: "Doublon"}) // ErrDuplicate
	store.Save(&Note{ID: "2"})                   // ErrValidation

	_, err := store.Get("999") // ErrNotFound
	fmt.Println("Get 999 :", err)
	fmt.Println("Est-ce ErrNotFound ?", errors.Is(err, ErrNotFound))

	fmt.Println("Total :", len(store.All()))
}
