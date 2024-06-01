/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"recipeFinder/cmd"
	db "recipeFinder/database"
)

func main() {
	db.Connect()
	cmd.Execute()
}
