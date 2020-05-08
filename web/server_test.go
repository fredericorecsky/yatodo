package web

import (
	"bytes"
	"encoding/json"
	"github.com/fredericorecsky/yatodo/app"
	"github.com/fredericorecsky/yatodo/models/todo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testapp = new(app.Todo)
var router = SetRouter()
var user todo.User

func init() {
	testapp.Migrate()
}

func performRequest(r http.Handler, method, path string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestUserPOST(t *testing.T) {
	body := todo.User{Username: "frederico2"}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	w := performRequest(router, "POST", "/users", buf)
	assert.Equal(t, http.StatusOK, w.Code)

	err := json.NewDecoder(w.Body).Decode(&user)

	if err != nil {
		t.Errorf("Error processing json %+v", err)
	}

	assert.NotEmpty(t, user.SecretToken)
}

func TestTodolistPOST(t *testing.T) {
	body := todo.TodoList{Name: "Testing"}

	// Get the first user from Db

}
