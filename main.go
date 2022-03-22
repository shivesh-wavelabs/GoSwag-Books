package main

import (
	"database/sql"
	"fmt"
	"GoSwag-Books/controllers"
	"GoSwag-Books/driver"
	"GoSwag-Books/models"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "GoSwag-Books/docs"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

// @title Books API
// @version 1.0
// @description This is a sample service for managing books
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email shivesh560@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	// Get Books
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	// Get a Book
	router.HandleFunc("/book/{id}", controller.GetBook(db)).Methods("GET")
	// Add a Book
	router.HandleFunc("/addBook", controller.AddBook(db)).Methods("POST")
	// Update a Book
	router.HandleFunc("/updateBook", controller.UpdateBook(db)).Methods("PUT")
	// Remove a Book
	router.HandleFunc("/removeBook/{id}", controller.RemoveBook(db)).Methods("DELETE")
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
