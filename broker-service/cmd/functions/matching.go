package functions

import "fmt"

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
	for _, v := range people {
		for i, matchName := range v.matcher {
			if matchName == name {
				v.matcher = append(v.matcher[:i], v.matcher[i+1:]...)
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
func matchAndDate(people []Person, male Person, female Person) []Person {
	//corner case no matcher
	if len(male.matcher) == 0 || len(female.matcher) == 0 {
		return people
	}
	//matching search
	for _, v := range male.matcher {
		for _, matchName := range female.matcher {
			if v == matchName {
				if male.WantedDatesNumber <= 0 {
					fmt.Printf("Sorry, %s is not interested in you!\n", female.Name)
					RemoveSinglePerson(people, male.Name)
					return people
				}
				if female.WantedDatesNumber <= 0 {
					fmt.Printf("Sorry, %s is not interested in you!\n", male.Name)
					RemoveSinglePerson(people, female.Name)
					return people
				}
				fmt.Printf(" %s %s match and date!\n\n", male.Name, female.Name)
				male.WantedDatesNumber--
				female.WantedDatesNumber--
			}
		}
	}
	return people
}
