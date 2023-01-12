package repository

import (
	"backend/persist/db"
	"backend/persist/model"
	"backend/web/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

func NewDocTest() Doc {
	db.Url = util.GetEnv("testDB")
	db := db.GetDB()
	db.Logger = logger.Default.LogMode(logger.Error)
	db.AutoMigrate(&model.Doc{}, &model.DocVersion{})
	db.Logger = logger.Default.LogMode(logger.Info)

	return NewDoc()
}

func TestDocImpl_CreateDir(t *testing.T) {
	doc := NewDocTest()
	tests := []struct {
		name string
		args *model.Doc
	}{
		{
			name: "dir1",
			args: &model.Doc{
				ProjectID: 1,
				Name:      "path1",
				Order:     1,
				ParentID:  0,
				Type:      model.DIR,
				Owner:     2,
			},
		},
		{
			name: "dir2",
			args: &model.Doc{
				ProjectID: 1,
				ParentID:  1,
				Name:      "path1",
				Order:     1,
				Type:      model.DIR,
				Owner:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doc.CreateDir(tt.args)
			assert.Nil(t, err)
			assert.NotEqual(t, got, 0)
		})
	}
}

func TestDocImpl_CreateDoc(t *testing.T) {
	doc := NewDocTest()
	createDoc, err := doc.CreateDoc(&model.DocDTO{
		DocVersion: model.DocVersion{
			RequestContentType:  "application/json",
			ResponseContentType: "application/json",
			Uri:                 "/api/v1/doc",
			Note:                "This is a test",
			Headers: []model.Scheme{{
				Name:    "testH",
				Note:    "",
				Default: "string",
			}},
			Params: []model.Scheme{{
				Name:    "testP",
				Note:    "",
				Default: "string",
			}},
			Result: []model.Scheme{{
				Name:    "testR",
				Note:    "",
				Default: "string",
			}},
			Type:  model.GET,
			Owner: 1,
		},
		ProjectID: 1,
		Order:     1,
		ParentID:  0,
	})
	assert.Nil(t, err)
	assert.NotEqual(t, 0, createDoc)
}

func TestDocImpl_GetTreeMenuByProject(t *testing.T) {

}

func TestDocImpl_UpdateDir(t *testing.T) {
	doc := NewDocTest()
	tests := []struct {
		name string
		args *model.Doc
	}{
		{
			name: "dir2",
			args: &model.Doc{
				Model: gorm.Model{
					ID: 1,
				},
				ProjectID: 1,
				ParentID:  1,
				Name:      "up_path1",
				Order:     1,
				Type:      model.DIR,
				Owner:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := doc.UpdateDir(tt.args)
			assert.Nil(t, err)
		})
	}
}

func TestDocImpl_UpdateDoc(t *testing.T) {
	project, err := NewDocTest().GetTreeMenuByProject(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, project)
}
