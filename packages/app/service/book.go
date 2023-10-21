package service

import (
	"strconv"

	"github.com/seojoo21/dbconnector-test/packages/app/dto"
	"github.com/seojoo21/dbconnector-test/packages/app/repository/interfaces"
)

type BookService struct {
	repos interfaces.Repository
}

type BookRequest struct {
	Id     string
	Title  string
	Author string
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

func (b *BookService) GetBook(id string) (*dto.BookDto, error) {
	res, err := b.repos.FindById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BookService) SaveBook(req BookRequest) (int, error) {
	bookDto, err := createBookDtoFromRequest(req)

	if err != nil {
		return -1, err
	}

	res, err := b.repos.Save(bookDto)

	if err != nil {
		return -1, err
	}
	return res, nil
}

func (b *BookService) UpdateBook(req BookRequest) (int, error) {
	bookDto, err := createBookDtoFromRequest(req)

	if err != nil {
		return -1, err
	}

	res, err := b.repos.Update(bookDto)

	if err != nil {
		return -1, err
	}

	return res, nil
}

func (b *BookService) DeleteBook(id string) (int, error) {
	res, err := b.repos.Delete(id)

	if err != nil {
		return -1, err
	}

	return res, nil
}

func createBookDtoFromRequest(req BookRequest) (*dto.BookDto, error) {
	id, err := strconv.Atoi(req.Id)

	if err != nil {
		return nil, err
	}

	return &dto.BookDto{
		Id:     id,
		Title:  req.Title,
		Author: req.Author,
	}, nil
}
