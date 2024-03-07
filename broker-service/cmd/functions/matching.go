package functions

import (
	"fmt"
)

type Person struct {
	Name              string   `json:"name"`
	Height            int      `json:"height"`
	Gender            Gender   `json:"gender"`
	WantedDatesNumber int      `json:"wantedDatesNumber"`
	matcher           []string `json:"matcher"`
}

type Gender struct {
	IsMale bool `json:"isMale"`
}

func AddSinglePersonAndMatch(people []Person, person Person) []Person {
	people = append(people, person)
	for i, v := range people {
		if v.Gender.IsMale == true && person.Gender.IsMale == false && v.Height > person.Height {
			people[i].matcher = append(v.matcher, person.Name) // add person name to matcher list
			//fmt.Println(people[i].matcher)
			fmt.Printf(" add person name %s to matcher list\n", person.Name)
		}
		if v.Gender.IsMale == false && person.Gender.IsMale == true && v.Height < person.Height {
			people[i].matcher = append(v.matcher, person.Name) // add person name to matcher list
			//fmt.Println(people[i].matcher)
			fmt.Printf(" add person name %s to matcher list\n", person.Name)
		}
	}
	//fmt.Println(people)
	return people
}

func RemoveSinglePerson(people []Person, name string) []Person {

	//corner case
	if len(people) <= 0 {
		return people
	}
	//base case
	if len(people) == 1 && people[0].Name == name {
		people = []Person{}
		return people
	}
	//normal case
	for i, v := range people {
		if v.Name == name {
			people = append(people[:i], people[i+1:]...)
		}
	}
	for x, v := range people {
		for i, matchName := range v.matcher {
			if matchName == name {
				people[x].matcher = append(v.matcher[:i], v.matcher[i+1:]...)
			}
		}
	}
	return people
}

func QuerySinglePeople(people []Person, name string) []string {
	for _, v := range people {
		if v.Name == name {
			return v.matcher
		}
	}
	return []string{}
}

//matching algorithm
func matchAndDate(people []Person, male string, female string) []Person {
	indexMale := -1
	indexFemale := -1
	// find 2 people from the list
	for k, v := range people {
		if v.Name == male {
			indexMale = k
		}
		if v.Name == female {
			indexFemale = k
		}
	}
	//corner case no matcher
	if indexMale == -1 || indexFemale == -1 {
		fmt.Println("do not find the matcing people")
		return people
	}
	//matching search
	for _, v := range people[indexMale].matcher {
		for _, matchName := range people[indexFemale].matcher {
			if v == people[indexFemale].Name && matchName == people[indexMale].Name {
				if people[indexMale].WantedDatesNumber <= 0 {
					fmt.Printf("Sorry, %s is not interested in you!\n", people[indexFemale].Name)
					RemoveSinglePerson(people, male)
					return people
				}
				if people[indexFemale].WantedDatesNumber <= 0 {
					fmt.Printf("Sorry, %s is not interested in you!\n", people[indexMale].Name)
					RemoveSinglePerson(people, female)
					return people
				}
				// find the matcher, match and date
				fmt.Printf(" %s %s match and date!\n\n", people[indexMale].Name, people[indexFemale].Name)
				//Number of dates -1
				people[indexMale].WantedDatesNumber--
				people[indexFemale].WantedDatesNumber--
				//both use up one date, so remove both from matching list
				for i, v := range people[indexMale].matcher {
					if v == female {
						people[indexMale].matcher = append(
							people[indexMale].matcher[:i], people[indexMale].matcher[i+1:]...,
						)
					}
				}
				for i, v := range people[indexFemale].matcher {
					if v == male {
						people[indexFemale].matcher = append(
							people[indexFemale].matcher[:i], people[indexFemale].matcher[i+1:]...,
						)
					}
				}
				return people
			}
		}
	}
	// if did not find the matcher in 2 person, print and return list
	fmt.Printf(" %s %s male and female dosen't match!\n\n", people[indexMale].Name, people[indexFemale].Name)
	return people
}
