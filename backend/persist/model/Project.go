package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sort"
)

type UserProject struct {
	UserId    uint `gorm:"autoIncrement:false;primaryKey"`
	ProjectID uint `gorm:"autoIncrement:false;primaryKey"`
	Role      Role
}

type Project struct {
	gorm.Model
	Name  string
	Note  string
	Owner uint
}

type Doc struct {
	gorm.Model
	Name                string
	RequestContentType  string
	ResponseContentType string
	Method              string
	Uri                 string
	Note                string
	Headers             Schemes `gorm:"type:json"`
	Params              Schemes `gorm:"type:json"`
	Result              Schemes `gorm:"type:json"`
	Type                DocType `gorm:"<-:create"`
	Owner               uint
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
