package adapters

import (
	"drones/app/data/models"
	"errors"
	"gorm.io/gorm"
)

type DronesAdapter struct {
	DB *gorm.DB
}

func NewDronesAdapter(db *gorm.DB) DronesAdapter {
	return DronesAdapter{db}
}

func (d *DronesAdapter) Create(drone *models.Drone) *models.Drone {
	d.DB.Create(&drone)
	return drone
}

func (d *DronesAdapter) GetBySerialNumber(serial string) (*models.Drone, error) {
	if len(serial) == 0 {
		return nil, errors.New("invalid serial number value")
	}

	var drone models.Drone

	err := d.DB.First(&drone).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &drone, nil
}
