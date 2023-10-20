package repository

import (
	"github.com/seojoo21/dbconnector-test/packages/app/domain"
	"github.com/seojoo21/dbconnector-test/packages/app/dto"
	dbmanager "github.com/seojoo21/dbconnector-test/packages/lib/db"
)

type BookRepository struct {
	db *dbmanager.DBManger
}

func InitBookRepository() *BookRepository {
	return &BookRepository{
		db: dbmanager.GetConnection(),
	}
}

func (b *BookRepository) FindAll() ([]*dto.BookDto, error) {
	var bookList []*dto.BookDto

	queryString := `SELECT id, title, author FROM BOOKS`

	rows, err := b.db.QueryMultiRows(queryString)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookDomain domain.Book

		err := rows.Scan(&bookDomain.Id, &bookDomain.Title, &bookDomain.Author)

		if err != nil {
			return nil, err
		}

		bookDto := dto.CreateBookDtoFromDomain(&bookDomain)
		bookList = append(bookList, bookDto)
	}

	return bookList, nil
}

func (b *BookRepository) FindById() {

}

func (b *BookRepository) Save() {

}

func (b *BookRepository) Update() {

}

func (b *BookRepository) Delete() {

}
