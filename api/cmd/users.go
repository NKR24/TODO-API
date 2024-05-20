package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Note      string    `json:"note"`
	Completed bool      `json:"completed"`
}

func postNewNote(c echo.Context) error {
	newTodos := new(Todo)
	newTodos.ID = uuid.New()
	c.Bind(newTodos)
	key := fmt.Sprintf("Todos:%s", newTodos.ID)
	_, err := rh.JSONSet(key, ".", newTodos)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, newTodos.ID)
}

func getNoteById(c echo.Context) error {
	id := c.Param("id")
	key := fmt.Sprintf("Todos:%s", id)
	todoJSON, err := redis.Bytes(rh.JSONGet(key, "."))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	todo := new(Todo)
	json.Unmarshal(todoJSON, &todo)
	return c.JSON(http.StatusOK, &todo)
}

func deleteNoteById(c echo.Context) error {
	var log string = "Deleted"

	id := c.Param("id")
	key := fmt.Sprintf("Todos:%s", id)
	_, err := rh.JSONDel(key, ".")
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, log)
}
