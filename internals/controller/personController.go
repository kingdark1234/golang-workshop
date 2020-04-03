package controller

import (
	"fmt"
	"workshop-service/internals/model"
)

// CreatePersons is function create person.
func CreatePersons(p *model.Person, arrPerson []model.Person) model.Person {
	fmt.Println(arrPerson)
	arrPerson = append(arrPerson, *p)
	fmt.Println(arrPerson)
	return model.Person{}
}

// ReadPersons is function read person by id.
func ReadPersons() model.Person {
	return model.Person{}
}

// UpdatePersons is function update person by id.
func UpdatePersons() model.Person {
	return model.Person{}
}

// DeletePersons is function delete person by id.
func DeletePersons() model.Person {
	return model.Person{}
}

// GetPersons is function get all
func GetPersons() model.Person {
	return model.Person{}
}
