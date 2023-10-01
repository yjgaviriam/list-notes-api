package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"list-notes/app/controllers"
	"list-notes/config"
	"log"
)

type App struct {
	DB     *sql.DB
	Routes *gin.Engine
}

func (a *App) CreateConnection() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}

func (a *App) CreateRoutes() {
	routes := gin.Default()
	controller := controllers.NewNoteController(a.DB)
	routes.GET("/notes", controller.GetNote)
	routes.POST("/notes", controller.InsertNote)
	a.Routes = routes
}

func (a *App) Run() {
	err := a.Routes.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
