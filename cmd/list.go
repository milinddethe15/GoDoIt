// Package cmd provides command-line subcommands for the GoDoIt application.
package cmd

import (
	"flag"

	"github.com/milinddethe15/GoDoIt/todo"
)

// Print the tasks table
func ListTasks(todos *todo.Todos, args []string) {
	// Define the "list" subcommand to list todo items
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listDone := listCmd.Int("done", 2, "The status of todo to be printed")
	listCat := listCmd.String("cat", "", "The category of tasks to be listed")

	// Parse the arguments for the "list" subcommand
	listCmd.Parse(args)
	todos.Print(*listDone, *listCat)
}
