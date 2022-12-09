package repository

import (
	"backend/persist/db"
	"backend/persist/model"

	"gorm.io/gorm"
)

type Doc interface {
	CreateDir(dir *model.Doc) (uint, error)
	UpdateDir(dir *model.Doc) error
	CreateDoc(api *model.DocDTO) (uint, error)
	UpdateDoc(api *model.DocDTO) error

	GetTreeMenuByProject(projectID int64) ([]model.DirDTO, error)
}

func NewDoc() Doc {
	return &DocImpl{db: db.GetDB()}
}

type DocImpl struct {
	db *gorm.DB
}

func (d *DocImpl) CreateDir(dir *model.Doc) (uint, error) {
	tx := d.db.Create(dir)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return dir.ID, nil
}

func (d *DocImpl) UpdateDir(dir *model.Doc) error {
	tx := d.db.Updates(dir)
	return tx.Error
}

func (d *DocImpl) CreateDoc(api *model.DocDTO) (uint, error) {
	doc := model.Doc{
		ProjectID: api.ProjectID,
		ParentID:  api.ParentID,
		Name:      api.Uri,
		Order:     api.Order,
		Type:      api.Type,
		Owner:     api.Owner,
	}
	tx := d.db.Begin()
	tx.Create(&doc)
	if tx.Error != nil {
		tx.Rollback()
		return 0, tx.Error
	}

	version := model.DocVersion{
		DocID:               doc.ID,
		Name:                api.Uri,
		RequestContentType:  api.RequestContentType,
		ResponseContentType: api.ResponseContentType,
		Uri:                 api.Uri,
		Note:                api.Note,
		Headers:             api.Headers,
		Params:              api.Params,
		Result:              api.Result,
		Type:                api.Type,
		Owner:               api.Owner,
	}
	tx.Create(&version)
	if tx.Error != nil {
		tx.Rollback()
		return 0, tx.Error
	}
	return doc.ID, tx.Commit().Error
}

func (d *DocImpl) UpdateDoc(api *model.DocDTO) error {
	doc := model.Doc{
		ProjectID: api.ProjectID,
		ParentID:  api.ParentID,
		Name:      api.Uri,
		Order:     api.Order,
		Type:      api.Type,
		Owner:     api.Owner,
	}
	doc.ID = api.ID
	tx := d.db.Updates(&doc)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	version := model.DocVersion{
		DocID:               doc.ID,
		Name:                api.Uri,
		RequestContentType:  api.RequestContentType,
		ResponseContentType: api.ResponseContentType,
		Uri:                 api.Uri,
		Note:                api.Note,
		Headers:             api.Headers,
		Params:              api.Params,
		Result:              api.Result,
		Type:                api.Type,
		Owner:               api.Owner,
	}
	tx = tx.Create(&version)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return tx.Commit().Error
}

func (d *DocImpl) GetTreeMenuByProject(projectID int64) ([]model.DirDTO, error) {
	list := make([]model.Doc, 0)
	tx := d.db.Find(&list, &model.Doc{
		ProjectID: uint(projectID),
	})
	if tx.Error != nil {
		return nil, tx.Error
	}
	dirMap := make(map[uint]model.DirDTO) // ID -> children
	for i := range list {
		dto := model.DirDTO{
			ID:    list[i].ID,
			Name:  list[i].Name,
			Type:  list[i].Type,
			Order: list[i].Order,
		}
		dirMap[list[i].ID] = dto
	}
	tree := make([]model.DirDTO, 0) // parentID -> children
	for i := range list {
		id := list[i].ID
		pid := list[i].ParentID
		if pid == 0 {
			tree = append(tree, dirMap[id])
			continue
		}
		if dto, ok := dirMap[pid]; ok {
			if dto.Children == nil {
				dto.Children = make([]model.DirDTO, 0)
			}
			dto.Children = append(dto.Children, dirMap[pid])
		}
	}
	return tree, nil
}
