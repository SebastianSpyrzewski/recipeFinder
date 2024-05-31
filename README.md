# Recipe Finder application
## Introduction
recipeFinder is a simple CLI application to get list of possible recipes from given ingredients with help of [Spoonacular API](https://spoonacular.com/food-api/). This repository contains both binary file and the source code (written in Go).
## Usage
To run the program, you need to execute command
```
./recipeFinder --ingredients=<list of ingredients> --numberOfRecipes=<number of recipes you want program to return>
```
where list of ingredients need to be given without spaces, separated by commas, and number of recipes can be ommitted - default value is 5. Example of use would be:
```
./recipeFinder --ingredients=tomatoes,eggs,pasta --numberOfRecipes=3
```
Program will then output the results.
## Code
For clarity and working convenience, code is divided into packages, each with different functionality:
### structures
This package defines few useful structures, necessary for other packages.
### cmd
Package for parsing command line arguments, created with [Cobra library](https://github.com/spf13/cobra).
### request
Package for simple tasks operating with user's request and formatting output. In future its goal will be to connect **httpconnection** and **database** packages.
### httpconnection
This package sends appropriate get request to Spoonacular page and saves the response in suitable structures.
