package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thospol/go-todo-clean-arch/domain"
	"github.com/Thospol/go-todo-clean-arch/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	TodoUseCase *mocks.ToDoUseCase
}

func (s *Suite) SetupSuite() {
	s.TodoUseCase = mocks.NewToDoUseCase(s.T())
}

func (s *Suite) AfterTest(_, _ string) {
	s.TodoUseCase.AssertExpectations(s.T())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetAllTodo() {
	reply := []domain.Todo{
		{
			Title:       "test",
			Description: "test",
		},
	}
	s.TodoUseCase.On("GetAllTodo").Return(reply, nil).Once()
	gin := gin.Default()

	uc := NewToDoController(s.TodoUseCase)
	gin.GET("/todo", uc.GetAllTodo)
	r := httptest.NewRequest(http.MethodGet, "/todo", nil)
	w := httptest.NewRecorder()
	gin.ServeHTTP(w, r)

	body, err := json.Marshal(reply)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), string(body), w.Body.String())
}
