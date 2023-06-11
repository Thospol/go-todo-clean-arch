package usecase

import (
	"testing"

	"github.com/Thospol/go-todo-clean-arch/domain"
	"github.com/Thospol/go-todo-clean-arch/mocks"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	TodoRepository *mocks.ToDoRepository
}

func (s *Suite) SetupSuite() {
	s.TodoRepository = mocks.NewToDoRepository(s.T())
}

func (s *Suite) AfterTest(_, _ string) {}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetAllTodo() {
	var entities []domain.Todo
	s.TodoRepository.
		On("GetAllTodo", &entities).
		Return(nil).
		Once()

	uc := NewToDoUseCase(s.TodoRepository)
	_, err := uc.GetAllTodo()
	require.NoError(s.T(), err)
}
