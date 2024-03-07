package functions

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
	for _, v := range people {
		if v.Gender.IsMale == true && person.Gender.IsMale == false && v.Height > person.Height {
			v.matcher = append(v.matcher, person.Name) // add person name to matcher list
		}
	}
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

func QuerySinglePeople() {

}
