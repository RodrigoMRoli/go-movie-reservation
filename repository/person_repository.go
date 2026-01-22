package repository

import (
	"context"
	"fmt"
	"go-movie-reservation/model"
	"go-movie-reservation/movie_resevation"
)

type PersonRepository struct {
	ctx     *context.Context
	queries *movie_resevation.Queries
}

func NewPersonRepository(ctx *context.Context, queries *movie_resevation.Queries) PersonRepository {
	return PersonRepository{
		ctx:     ctx,
		queries: queries,
	}
}

func (pr *PersonRepository) GetPeople() ([]model.Person, error) {

	rows, err := pr.queries.GetPeople(*pr.ctx)
	if err != nil {
		fmt.Println(err)
		return []model.Person{}, nil
	}

	var people []model.Person
	for _, p := range rows {
		people = append(people, model.Person{
			ID:        p.ID.String(),
			FirstName: p.FirstName.String,
			LastName:  p.LastName.String,
			Gender:    string(p.Gender.GenderEnum),
		})
	}

	return people, nil
}
