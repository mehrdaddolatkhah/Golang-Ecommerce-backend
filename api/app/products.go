package app

import (
	"cafekalaa/api/app/handler"
	"net/http"
)

func (a *App) getAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	handler.GetProduct(a.DB, w, r)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetProducts(a.DB, w, r)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

func (a *App) GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	handler.GetProductByCategory(a.DB, w, r)
}
