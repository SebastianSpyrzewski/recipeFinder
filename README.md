# Recipe Finder application
## Introduction
recipeFinder is a simple CLI application to get list of possible recipes from given ingredients using [Spoonacular API](https://spoonacular.com/food-api/) and local database. This repository contains both binary file and the source code (written in Go).
## Usage
To run the program, you need to execute command
```
./recipeFinder --ingredients=<list of ingredients> --numberOfRecipes=<number of recipes you want program to return>
```
where list of ingredients need to be given without spaces and separated by commas. Number of recipes can be ommitted (default value is 5). Example of use would be:
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
Package for simple tasks operating with user's request and formatting output. Its main goal is to connect **httpconnection** and **database** packages.
### database
The main purpose of this database is to avoid unnecessary API calls and use saved data if possible. It can be created with _create.sql_ file. The main tables of the database are _ingredients_, _requests_ and _recipes_, while the other ones describe relationships between them. We communicate with database via sqlite3. The whole database is included in this repository - in particular it can be not empty, but that is not a problem in our case.
### httpconnection
If necessary data is not present in the database, this package sends appropriate get request to Spoonacular page and saves the response in suitable structures.
