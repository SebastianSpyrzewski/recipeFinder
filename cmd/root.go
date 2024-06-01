/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"recipeFinder/request"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "recipeFinder",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		ingredients, err1 := cmd.Flags().GetStringSlice("ingredients")
		n, err2 := cmd.Flags().GetInt("numberOfRecipes")
		if err1 != nil || err2 != nil {
			fmt.Println("error parsing flags")
			println(err1.Error())
			println(err2.Error())
			os.Exit(1)
		}
		request.HandleRequest(ingredients, n)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("error exetucing command")
		println(err.Error())
		os.Exit(1)
	}
}

func init() {
	var ingredients []string
	var numberOfRecipes = 5
	rootCmd.Flags().StringSliceP("ingredients", "i", ingredients[:], "ingredients contained in a fridge")
	rootCmd.Flags().IntP("numberOfRecipes", "n", numberOfRecipes, "number of recipes to be shown")
}
