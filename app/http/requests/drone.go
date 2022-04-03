package requests

import (
	"drones/app/data/models"
)

type RegisterDrone struct {
	SerialNumber    string            `json:"serial_number" validate:"required,uniqueDB=drones.serial_number,lte=100"`
	Model           models.DroneModel `json:"model" validate:"required,oneof='LIGHT-WEIGHT' 'MIDDLE-WEIGHT' 'CRUISER-WEIGHT' 'HEAVY-WEIGHT'"`
	WeightLimit     int               `json:"weight_limit" validate:"required,lte=500"`
	BatteryCapacity int               `json:"battery_capacity" validate:"required,lte=100"`
}

type Medication struct {
	Name        string `json:"name" validate:"required,regexp=^[a-zA-Z0-9_-]*$"`
	Weight      int    `json:"weight" validate:"required"`
	Code        string `json:"code" validate:"required,uniqueDB=medications.code,regexp=^[A-Z0-9_]*$"`
	ImageBase64 string `json:"image_base64"`
}

type LoadMedications struct {
	SerialNumber string        `json:"serial_number" validate:"required,existsDB=drones.serial_number"`
	Medications  []*Medication `json:"medications" validate:"required,gte=1,dive"`
}

type UpdateDroneStatus struct {
	SerialNumber string             `json:"serial_number" validate:"required,existsDB=drones.serial_number"`
	Status       models.DroneStatus `json:"status" validate:"required,oneof='IDLE' 'LOADING' 'LOADED' 'DELIVERED' 'RETURNING'"`
}
