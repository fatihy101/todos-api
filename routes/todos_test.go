package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fatihy.space/todos-api/db"
	"github.com/stretchr/testify/assert"
)

const userId = "Do9030213a-8d03-4e8d-aae2-762989d54399tO"

func TestAddTodos(t *testing.T) {

	t.Run("should add todo", func(t *testing.T) {
		// Given
		mockdb := &mockDB{
			todos: []db.Todo{},
		}
		body := db.Todo{
			ID:        "10",
			Text:      "Buy some milk",
			Completed: false,
			UserId:    userId,
		}
		req, err := http.NewRequest("POST", "/todos", toJson(body))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", userId)

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, successMap, resBody, ErrUnexpectedResBody)
		assert.Equal(t, 1, len(mockdb.todos), "Number of todos are different")

	})

	t.Run("should return 401 unauthorized due to userId and todoItem mismatch", func(t *testing.T) {
		// Given
		mockdb := &mockDB{
			todos: []db.Todo{},
		}
		body := db.Todo{
			ID:        "10",
			Text:      "Buy some milk",
			Completed: false,
			UserId:    "different-user-id",
		}

		req, err := http.NewRequest("POST", "/todos", toJson(body))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", userId)

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		expectedRes := map[string]interface{}{"success": false, "error": ErrUserIdMismatch.Error()}
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)
		assert.Equal(t, http.StatusUnauthorized, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, expectedRes, resBody, ErrUnexpectedResBody)
		assert.Equal(t, 0, len(mockdb.todos), "Number of todos are different")
	})

}

func TestGetTodos(t *testing.T) {
	t.Run("should get todos", func(t *testing.T) {
		// Given
		mockdb := &mockDB{
			todos: []db.Todo{
				{
					ID:        "10",
					Text:      "Buy some milk",
					Completed: false,
					UserId:    userId,
				},
			},
		}

		req, err := http.NewRequest("GET", "/todos", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", userId)

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		var resBody []db.Todo
		json.NewDecoder(rr.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, mockdb.todos, resBody, ErrUnexpectedResBody)
		assert.Equal(t, 1, len(mockdb.todos), "Number of todos are different")

	})

	t.Run("should return 404 not found when there's no todo saved with userId", func(t *testing.T) {
		// Given
		mockdb := &mockDB{}

		req, err := http.NewRequest("GET", "/todos", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", userId)

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		expectedRes := map[string]interface{}{"success": false, "error": ErrNoTodosFound.Error()}
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)
		assert.Equal(t, http.StatusNotFound, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, expectedRes, resBody, ErrUnexpectedResBody)

	})

	t.Run("should return 400 when user id not present in headers", func(t *testing.T) {
		// Given
		mockdb := &mockDB{}

		req, err := http.NewRequest("GET", "/todos", nil)
		if err != nil {
			t.Fatal(err)
		}

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		expectedRes := map[string]interface{}{"success": false, "error": ErrUserIdRequired.Error()}
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)
		assert.Equal(t, http.StatusBadRequest, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, expectedRes, resBody, ErrUnexpectedResBody)

	})
}

func TestDeleteTodo(t *testing.T) {

	t.Run("should delete todo", func(t *testing.T) {
		// Given
		mockdb := &mockDB{
			todos: []db.Todo{
				{
					ID:        "10",
					Text:      "Buy some milk",
					Completed: false,
					UserId:    userId,
				},
			},
		}

		req, err := http.NewRequest("DELETE", "/todos/10", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", userId)

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		expectedRes := toJsonStr(map[string]bool{"success": true})
		assert.Equal(t, http.StatusOK, rr.Code, ErrUnexpectedStatus)
		assert.Contains(t, rr.Body.String(), expectedRes, ErrUnexpectedResBody)
		assert.Equal(t, 0, len(mockdb.todos), "Number of todos are different")

	})

	t.Run("should return 401 when header user id and item's user id mismatch", func(t *testing.T) {
		// Given
		mockdb := &mockDB{}
		req, err := http.NewRequest("DELETE", "/todos/10", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", "1")

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		expectedRes := map[string]interface{}{"success": false, "error": ErrUserIdMismatch.Error()}
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)
		assert.Equal(t, http.StatusUnauthorized, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, expectedRes, resBody, ErrUnexpectedResBody)
	})
}

func TestUpdateTodo(t *testing.T) {
	mockdb := &mockDB{
		todos: []db.Todo{
			{
				ID:        "10",
				Text:      "Buy some milk",
				Completed: false,
				UserId:    userId,
			},
		},
	}

	updated := db.Todo{
		ID:        "10",
		Text:      "Buy some milk",
		Completed: true, // change
		UserId:    userId,
	}

	t.Run("should update todo item", func(t *testing.T) {
		// Given
		req, err := http.NewRequest("PUT", "/todos", toJson(updated))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", userId)

		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// then
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)
		assert.Equal(t, http.StatusOK, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, successMap, resBody, ErrUnexpectedResBody)
		assert.Equal(t, updated, mockdb.todos[0], "Update failed. Data aren't same")
	})

	t.Run("Returns 401 when header's user id and todo id not matched", func(t *testing.T) {
		// Given
		req, err := http.NewRequest("PUT", "/todos", toJson(updated))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("X-User-Id", "fake")
		// When
		rr := httptest.NewRecorder()
		Routes(mockdb).ServeHTTP(rr, req)

		// Then
		expectedRes := map[string]interface{}{"success": false, "error": ErrUserIdMismatch.Error()}
		var resBody map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&resBody)
		assert.Equal(t, http.StatusUnauthorized, rr.Code, ErrUnexpectedStatus)
		assert.Equal(t, expectedRes, resBody, ErrUnexpectedResBody)
	})
}
