package main

import (
	"log"

	"github.com/eander0105/emil-portfolio/api/internal/config"
	"github.com/eander0105/emil-portfolio/api/internal/db"
	. "github.com/eander0105/emil-portfolio/api/internal/models"
	"github.com/eander0105/emil-portfolio/api/internal/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Register general routes
	routes.RegisterGeneralRoutes(router)

	return router
}

func migrate() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Initialize database
	if db_err := db.InitDB(cfg.DB); db_err != nil {
		panic(db_err)
	}

	db.DB.AutoMigrate(
		// text.go
		&Translation{},
		&Category{},

		// project.go
		&Project{},

		// blog.go
		&BlogPost{},
	)

	log.Println("Migration complete")
}

func main() {
	// Load config
	cfg, conf_err := config.LoadConfig()
	if conf_err != nil {
		panic(conf_err)
	}

	log.Println(cfg)

	// Run mrigration
	migrate()

	// Initialize database
	if db_err := db.InitDB(cfg.DB); db_err != nil {
		panic(db_err)
	}

	router := setupRouter()

	// Start Server
	err := router.Run(":3000")

	if err != nil {
		panic(err)
	}

}
