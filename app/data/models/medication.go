package models

import (
	"gorm.io/gorm"
	"strings"
)

type Medication struct {
	AppModel
	DroneSerialNumber string `json:"drone_serial_number"` // TODO: add foreign key
	Name              string `json:"name"`
	Weight            int    `json:"weight"`
	Code              string `json:"code" gorm:"index:index_drones_code_unique,unique"`
	ImageURL          string `json:"image_url" gorm:"column:image_url"`
}

func (*Medication) TableName() string {
	return "medications"
}

func (m *Medication) AfterFind(tx *gorm.DB) (err error) {
	if len(m.ImageURL) > 0 && !strings.HasPrefix(m.ImageURL, "/") {
		m.ImageURL = "/" + m.ImageURL
	}
	return
}
