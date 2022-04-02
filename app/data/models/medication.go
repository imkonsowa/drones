package models

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
