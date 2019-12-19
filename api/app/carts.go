package app

import (
	"cafekalaa/api/app/handler"
	"net/http"
)

func (a *App) AddToCart(w http.ResponseWriter, r *http.Request) {
	handler.AddToCart(a.DB, w, r)
}

func (a *App) GetCartPrice(w http.ResponseWriter, r *http.Request) {
	handler.GetCartPrice(a.DB, w, r)
}

func (a *App) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	handler.DeleteFromCart(a.DB, w, r)
}
