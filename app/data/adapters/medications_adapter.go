package adapters

import (
	"drones/app/data/models"
	"errors"
	"gorm.io/gorm"
)

type MedicationsAdapter struct {
	DB *gorm.DB
}

func NewMedicationsAdapter(db *gorm.DB) *MedicationsAdapter {
	return &MedicationsAdapter{db}
}

func (m *MedicationsAdapter) BatchCreate(medications []models.Medication) ([]models.Medication, error) {
	m.DB.Create(&medications)

	return medications, nil
}

func (m *MedicationsAdapter) GetDroneMedications(serial string) ([]models.Medication, error) {
	if len(serial) == 0 {
		return nil, errors.New("invalid serial param value")
	}

	var medications []models.Medication

	err := m.DB.Where("drone_serial_number = ?", serial).Find(&medications).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return medications, nil
}
