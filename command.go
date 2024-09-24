package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlag {
	cf := CmdFlag{}

	flag.StringVar(&cf.Add, "add", "", "Add new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "Show todo list")

	flag.Parse()

	return &cf
}

func (cf *CmdFlag) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	}
}
