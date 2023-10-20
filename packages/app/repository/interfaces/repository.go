package interfaces

import "github.com/seojoo21/dbconnector-test/packages/app/dto"

type Repository interface {
	FindAll() ([]*dto.BookDto, error)
	FindById()
	Save()
	Update()
	Delete()
}
