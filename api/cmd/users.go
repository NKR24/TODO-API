package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Todos struct {
	ID        uuid.UUID `json:"id"`
	Note      string    `json:"note"`
	Completed bool      `json:"completed"`
}

func createNewNote(c echo.Context) error {
	newTodos := new(Todos)
	newTodos.ID = uuid.New()
	c.Bind(newTodos)
	key := fmt.Sprintf("Todos:%s", newTodos.ID)
	_, err := rh.JSONSet(key, ".", newTodos)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, newTodos)
}
