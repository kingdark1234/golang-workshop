package service

import (
	"workshop-service/internals/repository"
)

// AuthorService ...
type AuthorService struct {
	AuthorRepository *repository.AuthorRepository
}

// ListAuthors ...
func (au *AuthorService) ListAuthors() (string, error) {
	// authors, err := au.AuthorRepository.GetAll()
	// if err != nil {
	// 	return []*entity.Author{}, err
	// }
	return "test", nil
}

// NewAuthorService ...
func NewAuthorService() *AuthorService {
	return &AuthorService{
		AuthorRepository: repository.NewAuthorRepository(),
	}
}
