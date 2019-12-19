package app

import (
	"cafekalaa/api/app/handler"
	"net/http"
)

func (a *App) createCategory(w http.ResponseWriter, r *http.Request) {
	handler.CreateCategory(a.DB, w, r)
}

func (a *App) getAllCategories(w http.ResponseWriter, r *http.Request) {
	handler.GetAllCategories(a.DB, w, r)
}

func (a *App) GetChildrenCategories(w http.ResponseWriter, r *http.Request) {
	handler.GetChildrenCategories(a.DB, w, r)
}
