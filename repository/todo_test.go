package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/Thospol/go-todo-clean-arch/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository domain.ToDoRepository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = NewToDoRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetAllTodo() {
	// mockup sql
	rows := sqlmock.NewRows([]string{"id", "title", "description"}).AddRow(1, "Make some program", "Create TODO app for testing").AddRow(2, "Make some program", "Create TODO app for testing")
	// expected
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todo`")).WillReturnRows(rows)
	result := []domain.Todo{}
	// perform
	err := s.repository.GetAllTodo(&result)
	require.NoError(s.T(), err)
}
