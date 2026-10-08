package store

import (
	"errors"
	"sort"

	"github.com/victorkovalyov-teletec/linkforge/internal/link"
)

var (
	ErrDuplicateID = errors.New("store: duplicate id")
)

type Store struct {
	byID map[uint64]link.Link
}

// create map
func New() Store {
	return Store{
		byID: make(map[uint64]link.Link),
	}
}

// add link
func Add(s Store, l link.Link) error {

	// check duplicate
	if _, exists := s.byID[l.ID]; exists {
		return ErrDuplicateID
	}

	// add
	s.byID[l.ID] = l

	return nil
}

// get link by id
func Get(s Store, id uint64) (link.Link, bool) {
	l, ok := s.byID[id]

	return l, ok
}

// get links
func All(s Store) []link.Link {

	// create empty map + projected capacity
	links := make([]link.Link, 0, len(s.byID))

	// iterate over all objs
	for _, l := range s.byID {
		links = append(links, l)
	}

	// sort
	sort.Slice(links, func(i, j int) bool {
		return links[i].ID < links[j].ID
	})

	return links
}

func Count(s Store) int {
	return len(s.byID)
}
