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

	queryString := `SELECT 
						id, title, author 
					FROM BOOKS`

	rows, err := b.db.QueryMultiRows(queryString)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

func (b *BookRepository) FindById(id string) (*dto.BookDto, error) {
	var bookDomain domain.Book

	queryString := `SELECT 
						id, title, author 
					FROM 
						BOOKS
					WHERE 
						id = ?`
	err := b.db.QueryRow(queryString, &id).Scan(&bookDomain.Id, &bookDomain.Title, &bookDomain.Author)

	if err != nil {
		return nil, err
	}

	bookDto := dto.CreateBookDtoFromDomain(&bookDomain)

	return bookDto, nil
}

func (b *BookRepository) Save(d *dto.BookDto) (int, error) {
	queryString := `INSERT INTO BOOKS 
						(title, author) 
					VALUES (?, ?)`

	res, err := b.db.TExec(queryString, &d.Title, &d.Author)

	if err != nil {
		return -1, err
	}

	lastId, _ := res.LastInsertId()

	return int(lastId), nil
}

func (b *BookRepository) Update(d *dto.BookDto) (int, error) {
	queryString := `UPDATE
						BOOKS
					SET
						title = ?,
						author = ?
					WHERE 
						id = ?`

	res, err := b.db.TExec(queryString, &d.Title, &d.Author, &d.Id)

	if err != nil {
		return -1, err
	}

	count, _ := res.RowsAffected()

	return int(count), nil
}

func (b *BookRepository) Delete(id string) (int, error) {
	queryString := `DELETE FROM BOOKS WHERE id = ?`

	res, err := b.db.TExec(queryString, &id)

	if err != nil {
		return -1, err
	}

	count, _ := res.RowsAffected()

	return int(count), nil
}
