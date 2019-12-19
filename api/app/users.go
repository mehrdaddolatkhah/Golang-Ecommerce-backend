package app

import (
	"cafekalaa/api/app/handler"
	"net/http"
)

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) registerUser(w http.ResponseWriter, r *http.Request) {
	handler.RegisterUser(a.DB, w, r)
}

func (a *App) findUser(w http.ResponseWriter, r *http.Request) {
	handler.FindUser(a.DB, w, r)
}

func (a *App) findUsers(w http.ResponseWriter, r *http.Request) {
	handler.FindUsers(a.DB, w, r)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

func (a *App) loginUser(w http.ResponseWriter, r *http.Request) {
	handler.LoginUser(a.DB, w, r)
}

func (a *App) SendSmsVerfication(w http.ResponseWriter, r *http.Request) {
	handler.SendSmsVerfication(a.DB, w, r)
}

func (a *App) GetOtpFromUser(w http.ResponseWriter, r *http.Request) {
	handler.GetOtpFromUser(a.DB, w, r)
}

func (a *App) RefreshToken(w http.ResponseWriter, r *http.Request) {
	handler.RefreshToken(w, r)
}
