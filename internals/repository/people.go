package repository

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"
	"workshop-service/internals/entity"
	"workshop-service/internals/infrastructure/database"
)

// PeopleRepository ...
type PeopleRepository struct {
	PeopleDB        *entity.People
	AddPeopleEntity *entity.AddPeople
	sb              database.ServerBase
	inputAddPeople  struct {
		firstName string
		lastName  string
		age       int
	}
}

// Prepare ...
func (p *PeopleRepository) Prepare() {
	p.PeopleDB.ID = 0
	p.PeopleDB.FirstName = html.EscapeString(strings.TrimSpace(p.PeopleDB.FirstName))
	p.PeopleDB.LastName = html.EscapeString(strings.TrimSpace(p.PeopleDB.LastName))
	p.PeopleDB.Age = 0
	p.PeopleDB.CreatedAt = time.Now()
	p.PeopleDB.UpdatedAt = time.Now()
}

// ValidateAddPeople ...
func (p *PeopleRepository) ValidateAddPeople() error {
	if p.inputAddPeople.firstName == "" {
		return errors.New("Required Title")
	}
	if p.inputAddPeople.lastName == "" {
		return errors.New("Required Content")
	}
	if p.inputAddPeople.age < 0 {
		return errors.New("Required Author")
	}
	return nil
}

// Add ...
func (p *PeopleRepository) Add(added *entity.AddPeople) error {
	var err error
	p.sb.Initialize()
	db := p.sb.DB
	fmt.Println("&db IS ", db)
	fmt.Println("&added IS ", &added)
	err = db.Debug().Model(*p.PeopleDB).Create(&added).Error
	if err != nil {
		return err
	}
	return nil
}

// // AddPeople ...
// func (p *PeopleRepository) AddPeople(db *gorm.DB) (*p.AddPeopleEntity, error) {

// }

// NewPeopleRepository ...
func NewPeopleRepository(sb database.ServerBase) *PeopleRepository {
	return &PeopleRepository{
		sb: sb,
	}
}
