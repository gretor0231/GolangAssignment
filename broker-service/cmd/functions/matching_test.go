package functions

import (
	"fmt"
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

	newPeople := AddSinglePersonAndMatch(people, person)

	assert.Equal(t, len(newPeople), 3)
	fmt.Println(newPeople[0].matcher)
	fmt.Println(newPeople[1].matcher)
	fmt.Println(newPeople[2].matcher)
	//assert.Equal(t, people[1].matcher, "Bob")

}
