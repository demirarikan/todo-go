/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"fmt"

	"fmt"
	"log"

	"github.com/demirarikan/todo-go/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create and add a new todo item to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(datafile)
		if err != nil {
			log.Printf("%v", err)
		}
		for _, x := range args {
			items = append(items, todo.Item{Text: x})
		}

		err = todo.SaveItems(datafile, items)
		if err != nil {
			fmt.Errorf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
