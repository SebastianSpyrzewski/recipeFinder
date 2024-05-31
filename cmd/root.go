/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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
			os.Exit(1)
		}
		request.HandleRequest(ingredients, n)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.recipeFinder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	var ingredients []string
	var numberOfRecipes = 5
	rootCmd.Flags().StringSliceP("ingredients", "i", ingredients[:], "ingredients contained in a fridge")
	rootCmd.Flags().IntP("numberOfRecipes", "n", numberOfRecipes, "number of recipes to be shown")
}
