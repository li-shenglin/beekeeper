package repository

import (
	"backend/persist/db"
	"backend/persist/model"
	"backend/web/util"
	"testing"

	"gorm.io/gorm"

	"gorm.io/gorm/logger"

	"github.com/stretchr/testify/assert"
)

func TestUpsertRoleForProject(t *testing.T) {
	p := NewProject()
	err := p.UpsertRoleForProject(&model.UserProject{
		ProjectID: 1,
		UserId:    2,
		Role:      model.Owner,
	})
	assert.Nil(t, err)

	err = p.UpsertRoleForProject(&model.UserProject{
		ProjectID: 2,
		UserId:    2,
		Role:      model.Owner,
	})
	assert.Nil(t, err)
}

func TestDelRoleForProject(t *testing.T) {
	p := NewProject()
	err := p.DelRoleForProject(2, 1)
	assert.Nil(t, err)

	err = p.DelRoleForProject(2, 2)
	assert.Nil(t, err)
}

func TestFindProjectRoleByUser(t *testing.T) {
	p := NewProject()
	user, err := p.FindProjectRoleByUser(2)
	assert.Nil(t, err)
	assert.Len(t, user, 2)
}

func TestFindProjectRoleByProject(t *testing.T) {
	p := NewProject()
	user, err := p.FindProjectRoleByProject(1)
	assert.Nil(t, err)
	assert.Len(t, user, 1)
}

func TestCreateProject(t *testing.T) {
	p := NewProject()
	id, err := p.CreateProject(&model.Project{
		Name: "testProject",
		Note: "testProject",
	})
	assert.Nil(t, err)
	assert.NotEqual(t, id, 0)
}

func TestUpdateProject(t *testing.T) {
	p := NewProject()
	err := p.UpdateProject(&model.Project{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "testProject",
		Note: "testProject",
	})
	assert.Nil(t, err)
}

func TestDelProject(t *testing.T) {
	p := NewProject()
	err := p.DelProject(1)
	assert.Nil(t, err)
}

func TestFindByUser(t *testing.T) {
	p := NewProject()
	user, err := p.FindByUser(2)
	assert.Nil(t, err)
	assert.NotEmpty(t, t, user)
}

func TestProject(t *testing.T) {
	db.Url = util.GetEnv("testDB", "")
	db := db.GetDB()
	db.Logger = logger.Default.LogMode(logger.Info)
	//db.AutoMigrate(&model.UserProject{}, &model.Project{})

	t.Run("TestUpsertRoleForProject", TestUpsertRoleForProject)
	t.Run("TestFindProjectRoleByUser", TestFindProjectRoleByUser)

	t.Run("TestCreateProject", TestCreateProject)
	t.Run("TestUpdateProject", TestUpdateProject)
	t.Run("TestDelProject", TestDelProject)
	t.Run("TestFindByUser", TestFindByUser)

	t.Run("TestDelRoleForProject", TestDelRoleForProject)
}
