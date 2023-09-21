/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mkCmd represents the mk command
var mkCmd = &cobra.Command{
	Use:   "mk",
	Short: "Create file [mdl] [rts]",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mk called")
	},
}

func init() {
	rootCmd.AddCommand(mkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
