package handler_test

import (
	"go-gin-ent-rest/handler"
	"go-gin-ent-rest/mock"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	c.Request = httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(mock.UserJson))
	c.Request.Header.Set("Content-Type", "application/json")

	h := handler.CreateUserHandler(clientMock)

	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `profiles`")).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.CreateUser(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetUser(t *testing.T) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	c.Request = httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	h := handler.CreateUserHandler(clientMock)

	query := "SELECT DISTINCT `users`.`id`, `users`.`name`, `users`.`created_at`, `users`.`updated_at` FROM `users` WHERE `users`.`id` = ?"

	sqlMock.MatchExpectationsInOrder(false)
	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(mock.UserColumns).AddRow(mock.User.CreatedAt, mock.User.ID, mock.User.Name, mock.User.UpdatedAt))

	query = "SELECT DISTINCT `profiles`.`id`, `profiles`.`gender`, `profiles`.`age`, `profiles`.`user_id` FROM `profiles` WHERE `user_id` IN (?)"

	sqlMock.MatchExpectationsInOrder(false)
	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(mock.ProfileColumns).AddRow(mock.Profile.Age, mock.Profile.Gender, mock.Profile.ID, 1))

	h.GetUser(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, mock.UserJson, rec.Body.String())
}
