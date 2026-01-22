package usecase

import (
	"go-movie-reservation/model"
	"go-movie-reservation/repository"
)

type PersonUseCase struct {
	repository repository.PersonRepository
}

func NewPersonUseCase(repository repository.PersonRepository) PersonUseCase {
	return PersonUseCase{
		repository: repository,
	}
}

func (pu *PersonUseCase) GetPeople() ([]model.Person, error) {
	return pu.repository.GetPeople()
}
