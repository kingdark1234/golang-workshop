package repository

import (
	"workshop-service/internals/entity"
)

type AuthorRepository struct {
	AuthorDb []*entity.Author
}

func (au *AuthorRepository) Add(author *entity.Author) error {
	au.AuthorDb = append(au.AuthorDb, author)
	return nil
}

func (au *AuthorRepository) GetAll() ([]*entity.Author, error) {
	return au.AuthorDb, nil
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{
		AuthorDb: []*entity.Author{},
	}
}
