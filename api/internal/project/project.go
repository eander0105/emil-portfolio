package project

import (
	"github.com/eander0105/emil-portfolio/api/internal/db"
	"github.com/eander0105/emil-portfolio/api/internal/db/models"
)

func GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := db.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := db.DB.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func CreateProject(project *models.Project) error {
	if err := db.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}
