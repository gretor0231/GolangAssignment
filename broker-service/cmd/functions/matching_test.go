package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSinglePersonAndMatch(t *testing.T) {
	people := []Person{
		{
			Name:              "John",
			Height:            180,
			Gender:            Gender{IsMale: true},
			WantedDatesNumber: 3,
			matcher:           []string{},
		},
		{
			Name:              "Jane",
			Height:            160,
			Gender:            Gender{IsMale: false},
			WantedDatesNumber: 2,
			matcher:           []string{}, // add person name to matcher list
		},
	}
	person := Person{
		Name:              "Bob",
		Height:            170,
		Gender:            Gender{IsMale: true},
		WantedDatesNumber: 1,
		matcher:           []string{},
	}
	person2 := Person{
		Name:              "Alice",
		Height:            150,
		Gender:            Gender{IsMale: false},
		WantedDatesNumber: 1,
		matcher:           []string{},
	}
	newPeople := AddSinglePersonAndMatch(people, person)
	newPeople2 := AddSinglePersonAndMatch(people, person2)
	assert.Equal(t, len(newPeople), 3)
	assert.Equal(t, newPeople[1].matcher, []string{"Bob"})
	assert.Equal(t, newPeople2[0].matcher, []string{"Alice"})
}

func TestRemoveSinglePerson(t *testing.T) {
	people := []Person{
		{
			Name:              "John",
			Height:            180,
			Gender:            Gender{IsMale: true},
			WantedDatesNumber: 3,
			matcher:           []string{},
		},
		{
			Name:              "Jane",
			Height:            160,
			Gender:            Gender{IsMale: false},
			WantedDatesNumber: 2,
			matcher:           []string{"Bob"}, // add person name to matcher list
		},
		{
			Name:              "Bob",
			Height:            170,
			Gender:            Gender{IsMale: true},
			WantedDatesNumber: 1,
			matcher:           []string{},
		},
	}
	newPeople := RemoveSinglePerson(people, "Bob")
	newPeople2 := RemoveSinglePerson(newPeople, "Alice")
	assert.Equal(t, len(newPeople), 2)
	assert.Equal(t, len(newPeople[1].matcher), 0)
	assert.Equal(t, len(newPeople2), 2)
}

func TestQuerySinglePeople(t *testing.T) {
	people := []Person{
		{
			Name:              "Jane",
			Height:            160,
			Gender:            Gender{IsMale: false},
			WantedDatesNumber: 2,
			matcher:           []string{"Bob"}, // add person name to matcher list
		},
	}
	query := QuerySinglePeople(people, "Jane")
	assert.Equal(t, query, []string{"Bob"})
}
