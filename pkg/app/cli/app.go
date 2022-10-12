package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/4strodev/todo_app/pkg/todo/core/database/sqlite"
	"github.com/4strodev/todo_app/pkg/todo/task/application"
	"github.com/4strodev/todo_app/pkg/todo/task/infrastructure"
	"github.com/spf13/cobra"
)

func init() {
	sqlite.InitPool()
}

var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "togo is a simple and easy to use task cli manager",
	Long:  `Built with go and sqlite togo helps you to manage your tasks from your cli`,
	Run: func(cmd *cobra.Command, args []string) {
		repository := infrastructure.SqliteTaskRepository{}
		repository.SetPool(sqlite.SqlitePool)

		taskList := application.TaskList{}
		taskList.SetRepository(&repository)
		tasks, err := taskList.Run()
		if err != nil {
			log.Fatal(err)
		}

		for _, task := range tasks {
			fmt.Printf("Name: %s\nDescription: %s\n", task.Name, task.Description)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
