package main

// // A simple program demonstrating the text input component from the Bubbles
// // component library.

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sandepten/todo-cli/internal/cli"
	"github.com/sandepten/todo-cli/internal/database"
)

func main() {
	database, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	p := tea.NewProgram(cli.NewWelcomeModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
