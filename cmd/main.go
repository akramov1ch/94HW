package main

import (
	"94HW/api/handlers"
	"94HW/config"
	"94HW/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfig()

	dbConn, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	r := gin.Default()

	r.GET("/books", handlers.GetBooks(dbConn))
	r.POST("/books", handlers.CreateBook(dbConn))
	r.GET("/books/:id", handlers.GetBookByID(dbConn))
	r.PUT("/books/:id", handlers.UpdateBook(dbConn))
	r.DELETE("/books/:id", handlers.DeleteBook(dbConn))

	r.GET("/metrics", gin.WrapH(http.DefaultServeMux))

	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
