package main

import (
    "ginbook/controllers"
    "ginbook/models"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Connect to the database
    db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&models.Book{})

    bookController := &controllers.BookController{DB: db}

    // Create a Gin router
    r := gin.Default()

    // Define the routes
    r.GET("/books", bookController.GetBooks)
    r.GET("/books/:id", bookController.GetBookByID)
    r.POST("/books", bookController.CreateBook)
    r.PUT("/books/:id", bookController.UpdateBook)
    r.DELETE("/books/:id", bookController.DeleteBook)

    // Start the server
    r.Run(":8080")
}
