package db

import (
	"log"

	. "github.com/eander0105/emil-portfolio/api/internal/db/models"
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
