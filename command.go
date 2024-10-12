package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Delete int
	Toggle int
	Edit   string
	List   bool
	Help   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.Help, "help", false, "Show help")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit format. Use index:title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index")
			os.Exit(1)
		}

		if err := todos.update(index, parts[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case cf.Delete != -1:
		if err := todos.remove(cf.Delete); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case cf.Toggle != -1:
		if err := todos.toggle(cf.Toggle); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case cf.Help:
		flag.PrintDefaults()

	default:
		fmt.Println("No command provided")
		os.Exit(1)

	}

}
