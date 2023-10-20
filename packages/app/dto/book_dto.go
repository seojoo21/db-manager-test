package dto

import "github.com/seojoo21/dbconnector-test/packages/app/domain"

type BookDto struct {
	Id     int
	Title  string
	Author string
}

func CreateBookDtoFromDomain(d *domain.Book) *BookDto {
	return &BookDto{
		Id:     d.Id,
		Title:  d.Title,
		Author: d.Author,
	}
}

func CreateBookDomainFromDto(d *BookDto) *domain.Book {
	return &domain.Book{
		Title:  d.Title,
		Author: d.Author,
	}
}
