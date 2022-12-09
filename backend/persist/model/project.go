package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"gorm.io/gorm"
)

type UserProject struct {
	UserId    uint `gorm:"autoIncrement:false;primaryKey;type:bigint"`
	ProjectID uint `gorm:"autoIncrement:false;primaryKey;type:bigint"`
	Role      Role `gorm:"type:varchar(16);not null"`
}

type Project struct {
	gorm.Model
	Name string `gorm:"type:varchar(128);not null"`
	Note string `gorm:"type:varchar(1024)"`
}

type Doc struct {
	gorm.Model
	ProjectID uint    `gorm:"not null"`
	ParentID  uint    `gorm:"not null"`
	Name      string  `gorm:"type:varchar(256);not null"`
	Order     int     `gorm:"type:int;not null;default:0"`
	Type      DocType `gorm:"<-:create;type:varchar(256);not null"`
	Owner     uint
}

type DocVersion struct {
	ID                  uint `gorm:"primaryKey"`
	DocID               uint `gorm:"not null"`
	CreatedAt           time.Time
	Name                string  `gorm:"type:varchar(256);not null"`
	RequestContentType  string  `gorm:"type:varchar(256);not null"`
	ResponseContentType string  `gorm:"type:varchar(256);not null"`
	Uri                 string  `gorm:"type:varchar(256);not null"`
	Note                string  `gorm:"type:varchar(1024);not null"`
	Headers             Schemes `gorm:"type:json"`
	Params              Schemes `gorm:"type:json"`
	Result              Schemes `gorm:"type:json"`
	Type                DocType `gorm:"<-:create;type:varchar(16);not null"`
	Owner               uint
}

type DocDTO struct {
	DocVersion
	ProjectID uint
	ParentID  uint
	Order     int
}

type DirDTO struct {
	ID       uint     `json:"ID,omitempty"`
	Name     string   `json:"name,omitempty"`
	Type     DocType  `json:"type,omitempty"`
	Order    int      `json:"order,omitempty"`
	Children []DirDTO `json:"children,omitempty"`
}

type Scheme struct {
	Name    string
	Type    string
	Note    string
	Default string
}

type Schemes []Scheme

func (s *Schemes) Add(scheme Scheme) {
	for i := range *s {
		if (*s)[i].Name == scheme.Name {
			(*s)[i] = scheme
			return
		}
	}
	*s = append(*s, scheme)
	sort.SliceStable(*s, func(i, j int) bool {
		return (*s)[i].Name > (*s)[j].Name
	})
}

func (s *Schemes) Remove(name string) {
	for i := range *s {
		if (*s)[i].Name == name {
			*s = append((*s)[0:i], (*s)[i+1:]...)
			break
		}
	}

	sort.SliceStable(*s, func(i, j int) bool {
		return (*s)[i].Name > (*s)[j].Name
	})
}

func (s *Schemes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := make([]Scheme, 0)
	err := json.Unmarshal(bytes, &result)
	*s = result
	return err
}

func (s Schemes) Value() (driver.Value, error) {
	return json.Marshal(s)
}
