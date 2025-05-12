package project

import (
	"emil-portfolio/internal/db"
	"emil-portfolio/internal/db/models"
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
