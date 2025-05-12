package main

import (
	"emil-portfolio/internal/db"
	"emil-portfolio/internal/server"
)

func main() {
	// Initialize database
	if db_err := db.InitDB(); db_err != nil {
		panic(db_err)
	}

	router := server.SetupRouter()

	// Start Server
	err := router.Run(":3000")

	if err != nil {
		panic(err)
	}

}
