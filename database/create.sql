DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS requests;
DROP TABLE IF EXISTS ingredients_requests;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS request_answers;
DROP TABLE IF EXISTS owned_ingredients;
DROP TABLE IF EXISTS missing_ingredients;
CREATE TABLE IF NOT EXISTS ingredients (id INTEGER PRIMARY KEY, name VARCHAR(255) UNIQUE);
CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, ingredients INTEGER, number_of_recipes INTEGER);
CREATE TABLE IF NOT EXISTS ingredients_requests (ingredient INTEGER REFERENCES ingredients, request INTEGER REFERENCES requests);
CREATE TABLE IF NOT EXISTS recipes (id INTEGER PRIMARY KEY, title VARCHAR(255), carbs VARCHAR(255), protein VARCHAR(255), calories VARCHAR(255));
CREATE TABLE IF NOT EXISTS request_answers (id INTEGER PRIMARY KEY, request INTEGER REFERENCES requests, recipe INTEGER REFERENCES RECIPES, number INTEGER, UNIQUE(request, number));
CREATE TABLE IF NOT EXISTS owned_ingredients (ingredient INTEGER REFERENCES INGREDIENTS, answer INTEGER REFERENCES request_answers);
CREATE TABLE IF NOT EXISTS missing_ingredients (ingredient INTEGER REFERENCES INGREDIENTS, answer INTEGER REFERENCES request_answers);


