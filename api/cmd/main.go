package main

import (
	"log"

	"github.com/eander0105/emil-portfolio/api/internal/config"
	"github.com/eander0105/emil-portfolio/api/internal/db"
	"github.com/eander0105/emil-portfolio/api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg, conf_err := config.LoadConfig()
	if conf_err != nil {
		panic(conf_err)
	}

	log.Println(cfg)

	// Initialize database
	if db_err := db.InitDB(cfg.DB); db_err != nil {
		panic(db_err)
	}

	router := gin.Default()

	// Register routes
	routes.RegisterGeneralRoutes(router)

	// Start Server
	err := router.Run(":3000")

	if err != nil {
		panic(err)
	}

}
