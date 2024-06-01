package database

import (
	"database/sql"
	"fmt"
	"os"
	st "recipeFinder/structures"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const create string = `CREATE TABLE IF NOT EXISTS ingredients (id INTEGER PRIMARY KEY, name VARCHAR(255) UNIQUE);
CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, ingredients INTEGER, number_of_recipes INTEGER);
CREATE TABLE IF NOT EXISTS ingredients_requests (ingredient INTEGER REFERENCES ingredients, request INTEGER REFERENCES requests);
CREATE TABLE IF NOT EXISTS recipes (id INTEGER PRIMARY KEY, title VARCHAR(255), carbs VARCHAR(255), protein VARCHAR(255), calories VARCHAR(255));
CREATE TABLE IF NOT EXISTS request_answers (id INTEGER PRIMARY KEY, request INTEGER REFERENCES requests, recipe INTEGER REFERENCES RECIPES, number INTEGER, UNIQUE(request, number));
CREATE TABLE IF NOT EXISTS owned_ingredients (ingredient INTEGER REFERENCES INGREDIENTS, answer INTEGER REFERENCES request_answers);
CREATE TABLE IF NOT EXISTS missing_ingredients (ingredient INTEGER REFERENCES INGREDIENTS, answer INTEGER REFERENCES request_answers);`

func Search(req st.Request) []st.Recipe {
	list := "'" + strings.Join(req.Ingredients, "', '") + "'"
	l := len(req.Ingredients)
	query := fmt.Sprintf("select r.id, r.number_of_recipes from ingredients i join ingredients_requests ir on i.id = ir.ingredient  join requests r on r.id=ir.request where i.name in (%s) and r.ingredients=%v group by r.id having count(*) = %v", list, l, l)
	row := db.QueryRow(query)
	var id int
	var n int
	err := row.Scan(&id, &n)
	if err == sql.ErrNoRows {
		return nil
	}
	if n < req.NumberOfRecipes {
		fmt.Println("less")
		db.Exec("DELETE FROM requests WHERE id=?", id) //we will insert new request with same ingredients but more recipes, so the old one becomes useless
		return nil
	}
	var recipes = make([]st.Recipe, req.NumberOfRecipes)
	var rid int
	var aid int
	i := 0
	rows, err := db.Query("SELECT id, recipe FROM request_answers WHERE request=? ORDER BY number LIMIT ?", id, req.NumberOfRecipes)
	for rows.Next() {
		rows.Scan(&aid, &rid)
		row = db.QueryRow("SELECT * FROM recipes WHERE id = ?", rid)
		row.Scan(&recipes[i].Id, &recipes[i].Title, &recipes[i].Carbs, &recipes[i].Protein, &recipes[i].Calories)
		var ingr st.Ingredient
		ingr_rows, _ := db.Query("SELECT i.id, i.name FROM ingredients i JOIN owned_ingredients oi ON i.id=oi.ingredient WHERE oi.answer=?", aid)
		for ingr_rows.Next() {
			ingr_rows.Scan(&ingr.Id, &ingr.Name)
			recipes[i].UsedIngredients = append(recipes[i].UsedIngredients, ingr)
		}
		ingr_rows, _ = db.Query("SELECT i.id, i.name FROM ingredients i JOIN missing_ingredients mi ON i.id=mi.ingredient WHERE mi.answer=?", aid)
		for ingr_rows.Next() {
			ingr_rows.Scan(&ingr.Id, &ingr.Name)
			recipes[i].MissedIngredients = append(recipes[i].MissedIngredients, ingr)
		}
		i += 1
	}
	if err != nil {
		fmt.Println("database communication error")
		println(err.Error())
		os.Exit(1)
	}
	return recipes
}

func Update(req st.Request, recipes []st.Recipe) {
	db.Exec("INSERT INTO requests (ingredients, number_of_recipes) VALUES (?, ?)", len(req.Ingredients), req.NumberOfRecipes)
	row := db.QueryRow("SELECT last_insert_rowid()")
	var reqid int
	row.Scan(&reqid)
	for _, ing := range req.Ingredients {
		db.Exec("INSERT OR IGNORE INTO ingredients (name) VALUES (?)", ing)
		row = db.QueryRow("SELECT id FROM ingredients WHERE name=?", ing)
		var ingid int
		row.Scan(&ingid)
		db.Exec("INSERT INTO ingredients_requests VALUES (?, ?)", ingid, reqid)
	}
	for i, rec := range recipes {
		db.Exec("INSERT OR IGNORE INTO recipes VALUES (?, ?, ?, ?, ?)", rec.Id, rec.Title, rec.Carbs, rec.Protein, rec.Calories)
		db.Exec("INSERT INTO request_answers (request, recipe, number) VALUES(?, ?, ?)", reqid, rec.Id, i+1)
		row = db.QueryRow("SELECT last_insert_rowid()")
		var ansid int
		row.Scan(&ansid)
		for j, ing := range rec.UsedIngredients {
			db.Exec("INSERT OR IGNORE INTO ingredients (name) VALUES (?)", ing.Name)
			row = db.QueryRow("SELECT id FROM ingredients WHERE name=?", ing.Name)
			var ingid int
			row.Scan(&ingid)
			rec.UsedIngredients[j].Id = ingid
			db.Exec("INSERT INTO owned_ingredients VALUES (?, ?)", ingid, ansid)
		}
		for j, ing := range rec.MissedIngredients {
			db.Exec("INSERT OR IGNORE INTO ingredients (name) VALUES (?)", ing.Name)
			row = db.QueryRow("SELECT id FROM ingredients WHERE name=?", ing.Name)
			var ingid int
			row.Scan(&ingid)
			rec.MissedIngredients[j].Id = ingid
			db.Exec("INSERT INTO missing_ingredients VALUES (?, ?)", ingid, ansid)
		}
	}
}

func Connect() {
	database, err := sql.Open("sqlite3", "database/recipefinder.db")
	db = database
	if err != nil {
		fmt.Println("error connecting with database")
		println(err.Error())
		os.Exit(1)
	}
	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("error creating database")
		println(err.Error())
		os.Exit(1)
	}
}
