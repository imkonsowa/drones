package adapters

import (
	"drones/app/data/models"
	"errors"
	"gorm.io/gorm"
)

type DronesAdapter struct {
	DB *gorm.DB
}

func NewDronesAdapter(db *gorm.DB) *DronesAdapter {
	return &DronesAdapter{db}
}

func (d *DronesAdapter) DronesList() []models.Drone {
	var drones []models.Drone

	d.DB.Table("drones").Find(&drones)

	return drones
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

	err := d.DB.Where("serial_number = ?", serial).First(&drone).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &drone, nil
}

func (d *DronesAdapter) UpdateStatus(serial string, status models.DroneStatus) error {
	if len(status) == 0 || len(status) == 0 {
		return errors.New("invalid params values")
	}

	err := d.DB.
		Where("serial_number = ?", serial).
		Updates(&models.Drone{
			Status: status,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (d *DronesAdapter) GetDronesByStatus(status models.DroneStatus) ([]models.Drone, error) {
	if len(status) == 0 {
		return nil, errors.New("invalid params values")
	}

	var drones []models.Drone

	err := d.DB.Where("status = ?", status).Find(&drones).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return drones, nil
}
