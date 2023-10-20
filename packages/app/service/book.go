package service

import (
	"github.com/seojoo21/dbconnector-test/packages/app/dto"
	"github.com/seojoo21/dbconnector-test/packages/app/repository/interfaces"
)

type BookService struct {
	repos interfaces.Repository
}

func NewBookService(repos interfaces.Repository) *BookService {
	return &BookService{
		repos: repos,
	}
}

func (b *BookService) GetBookList() ([]*dto.BookDto, error) {
	res, err := b.repos.FindAll()

	if err != nil {
		return nil, err
	}

	return res, nil
}
