package db

import (
	"log"

	. "emil-portfolio/internal/db/models"
)

func migrate() {
	// Initialize database
	if db_err := InitDB(); db_err != nil {
		panic(db_err)
	}

	DB.AutoMigrate(
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
