package store

import (
	"errors"
	"github.com/sofc-t/mereb_simple_go_api/internals/models"
)

var (
	ErrPersonNotFound = errors.New("person not found")
)

type InMemoryStore struct {
	data map[string]*models.Person
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{data: make(map[string]*models.Person)}
}

func (s *InMemoryStore) GetAllPersons() []*models.Person {
	persons := make([]*models.Person, 0, len(s.data))
	for _, person := range s.data {
		persons = append(persons, person)
	}
	return persons
}

func (s *InMemoryStore) GetPerson(id string) (*models.Person, error) {
	person, exists := s.data[id]
	if !exists {
		return nil, ErrPersonNotFound
	}
	return person, nil
}

func (s *InMemoryStore) CreatePerson(person *models.Person) {
	s.data[person.ID] = person
}

func (s *InMemoryStore) UpdatePerson(id string, updatedPerson *models.Person) error {
	_, exists := s.data[id]
	if !exists {
		return ErrPersonNotFound
	}
	s.data[id] = updatedPerson
	return nil
}

func (s *InMemoryStore) DeletePerson(id string) error {
	_, exists := s.data[id]
	if !exists {
		return ErrPersonNotFound
	}
	delete(s.data, id)
	return nil
}
