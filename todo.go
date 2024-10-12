package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	ComptedAt *time.Time `json:"completed_at"`
}

type Todos []Todo

func (todos *Todos) add(title string) {
	*todos = append(*todos, Todo{
		ID:        len(*todos) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	})
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) remove(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Completed = !(*todos)[index].Completed
	if (*todos)[index].Completed {
		t := time.Now()
		(*todos)[index].ComptedAt = &t
	} else {
		(*todos)[index].ComptedAt = nil
	}
	return nil
}

func (todos *Todos) update(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todos {
		completed := "✅"
		if !todo.Completed {
			completed = "❌"
		}
		completedAt := ""
		if todo.ComptedAt != nil {
			completedAt = todo.ComptedAt.Format("2006-01-02 15:04:05")
		}
		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format("2006-01-02 15:04:05"), completedAt)
	}

	table.Render()
}
