package handler

import (
	"cafekalaa/api/app/model"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	uuid "github.com/google/uuid"
)

func CreateCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	category := model.Category{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO categories(id,parent,title,level) VALUES($1,$2,$3,$4);")

	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewUUID()

	res, err := stmt.Exec(uuid, category.Parent, category.Title, category.Level)

	if err != nil && res != nil {
		panic(err)
	}

	respondJSON(w, http.StatusCreated, category)
}

func GetAllCategories(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "select * from categories")

	if err != nil {
		panic(err)
	}

	categories := make([]model.Category, 0)
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Parent, &category.Title, &category.Level); err != nil {

		}

		categories = append(categories, category)
	}

	respondJSON(w, http.StatusOK, categories)
}

func GetChildrenCategories(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	type ParentID struct {
		ParentUUID uuid.UUID `json:"parentuuid"`
	}

	parentID := ParentID{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&parentID); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	query := `
	WITH RECURSIVE subordinates AS (
		SELECT
		   id,
		   parent,
		   title,
		   level
		FROM
		   categories
		WHERE
		parent = $1 
		UNION
		   SELECT
			  e.id,
			  e.parent,
			  e.title,
			  e.level
		   FROM
		   categories e
		   INNER JOIN subordinates s ON s.id = e.parent
	 ) SELECT
		*
	 FROM
		subordinates;
	`

	rows, err := db.QueryContext(ctx, query, parentID.ParentUUID)

	if err != nil {
		panic(err)
	}

	categories := make([]model.Category, 0)
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Parent, &category.Title, &category.Level); err != nil {

		}

		categories = append(categories, category)
	}

	respondJSON(w, http.StatusOK, categories)
}
