package repository

import (
	"backend/persist/db"
	"backend/persist/model"

	"gorm.io/gorm"
)

type Project interface {
	CreateProject(project *model.Project) (uint, error)
	UpdateProject(project *model.Project) error
	FindByUser(userID int64) ([]model.Project, error)
	DelProject(projectID int64) error

	FindProjectRoleByUser(userID int64) ([]model.UserProject, error)
	FindProjectRoleByProject(projectID int64) ([]model.UserProject, error)
	UpsertRoleForProject(up *model.UserProject) error
	DelRoleForProject(userID, projectID int64) error
}

type ProjectImpl struct {
	db *gorm.DB
}

func (p *ProjectImpl) DelProject(projectID int64) error {
	return p.db.Delete(&model.Project{}, projectID).Error
}

func (p *ProjectImpl) UpsertRoleForProject(up *model.UserProject) error {
	tx := p.db.Begin()
	if t := tx.Delete(up); t.Error != nil {
		tx.Rollback()
		return t.Error
	}
	if t := tx.Create(up); t.Error != nil {
		tx.Rollback()
		return t.Error
	}
	tx.Commit()
	return tx.Error
}

func (p *ProjectImpl) DelRoleForProject(userID, projectID int64) error {
	t := p.db.Delete(&model.UserProject{
		UserId:    uint(userID),
		ProjectID: uint(projectID),
	})
	return t.Error
}

func (p *ProjectImpl) FindProjectRoleByUser(userID int64) ([]model.UserProject, error) {
	up := new([]model.UserProject)
	db := p.db.Where("user_id = ?", userID).Find(up)
	return *up, db.Error
}

func (p *ProjectImpl) FindProjectRoleByProject(projectID int64) ([]model.UserProject, error) {
	up := new([]model.UserProject)
	db := p.db.Where("project_id = ?", projectID).Find(up)
	return *up, db.Error
}

func (p *ProjectImpl) CreateProject(project *model.Project) (uint, error) {
	tx := p.db.Create(project)
	return project.ID, tx.Error
}

func (p *ProjectImpl) UpdateProject(project *model.Project) error {
	tx := p.db.Updates(project)
	return tx.Error
}

func (p *ProjectImpl) FindByUser(userID int64) ([]model.Project, error) {
	up, err := p.FindProjectRoleByUser(userID)
	if err != nil {
		return nil, err
	}
	pids := make([]uint, len(up))
	for i := range up {
		pids[i] = up[i].ProjectID
	}

	res := new([]model.Project)
	return *res, p.db.Find(res, pids).Error
}

func NewProject() Project {
	return &ProjectImpl{db: db.GetDB()}
}
