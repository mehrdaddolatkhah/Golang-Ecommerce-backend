package app

import (
	"fmt"

	"log"
	"net/http"

	"cafekalaa/api/config"

	"database/sql"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

//App has router and db instances
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.DbName,
		config.DB.Password)

	var err error

	a.DB, err = sql.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatal("Could not connect database")
	}

	err = a.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	a.Router = mux.NewRouter()
	a.setRouters()

}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects

	a.Get("/products", a.getAllProducts)
	a.Post("/product", a.createProduct)
	a.Get("/product/findproducts", a.getProducts)
	a.Get("/product/bycategory", a.GetProductByCategory)
	a.Get("/product/findproduct", a.getProduct)
	a.Put("/product", a.updateProduct)
	a.Delete("/product", a.deleteProduct)

	a.Get("/users", a.getAllUsers)
	a.Post("/user", a.registerUser)
	a.Get("/user/finduser", a.findUser)
	a.Get("/user/findusers", a.findUsers)
	a.Put("/user", a.updateUser)
	a.Delete("/user", a.deleteUser)

	a.Post("/user/login", a.loginUser)
	a.Post("/user/sendmobile", a.SendSmsVerfication)
	a.Post("/user/sendotp", a.GetOtpFromUser)
	a.Get("/user/refresh", a.RefreshToken)

	a.Post("/category", a.createCategory)
	a.Get("/categories", a.getAllCategories)
	a.Get("/category", a.GetChildrenCategories)

	a.Post("/cart/add", a.AddToCart)
	a.Get("/cart/price", a.GetCartPrice)
	a.Delete("/cart/delete", a.DeleteFromCart)

}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// func handler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "Hello World!")
// }
