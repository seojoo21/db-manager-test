package interfaces

import "github.com/seojoo21/dbconnector-test/packages/app/dto"

type Repository interface {
	FindAll() ([]*dto.BookDto, error)
	FindById(id string) (*dto.BookDto, error)
	Save(d *dto.BookDto) (int, error)
	Update(d *dto.BookDto) (int, error)
	Delete(id string) (int, error)
}
