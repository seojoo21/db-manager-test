package repository_test

import (
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/seojoo21/dbconnector-test/packages/app/dto"
	"github.com/seojoo21/dbconnector-test/packages/app/repository"
	dbmanager "github.com/seojoo21/dbconnector-test/packages/lib/db"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	// initialize a book repository with the database
	bookRepos := repository.InitBookRepository()

	for i := 0; i < 10; i++ {
		testBook := &dto.BookDto{
			Title:  "title" + strconv.Itoa(i),
			Author: "author" + strconv.Itoa(i),
		}

		// call the save function
		insertedId, err := bookRepos.Save(testBook)

		if err != nil {
			t.Fatalf("Saving Test Data Error: %v", err)
		}

		t.Logf("Saving Test Data Success!! InsertedId :: %v", insertedId)
	}
}

func TestSaveBySqlMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set up your repository with the mock database
	dbmanager := dbmanager.NewDBManager(db)
	repo := repository.NewBookRepository(dbmanager)

	for i := 0; i < 10; i++ {
		query := `INSERT INTO BOOKS`
		mock.ExpectExec(query).WithArgs("Test Title", "Test Author").WillReturnResult(sqlmock.NewResult(int64(i), int64(i)))

		// Call the method you want to test
		bookDto := &dto.BookDto{Title: "Test Title", Author: "Test Author"}
		insertedId, err := repo.Save(bookDto)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, i, insertedId)
	}
}
