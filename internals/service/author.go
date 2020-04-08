package service

import (
	"workshop-service/internals/entity"
	"workshop-service/internals/repository"
)

type AuthorService struct {
	AuthorRepository *repository.AuthorRepository
}

func (au *AuthorService) ListAuthors() ([]*entity.Author, error) {
	authors, err := au.AuthorRepository.GetAll()
	if err != nil {
		return []*entity.Author{}, err
	}
	return authors, nil
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		AuthorRepository: repository.NewAuthorRepository(),
	}
}
