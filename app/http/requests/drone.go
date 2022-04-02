package requests

import (
	"drones/app/data/models"
)

type RegisterDrone struct {
	SerialNumber    string             `json:"serial_number" validate:"required,lte=100"`
	Model           models.DroneModel  `json:"model" validate:"required,oneof='LIGHT-WEIGHT' 'MIDDLE-WEIGHT' 'CRUISER-WEIGHT' 'HEAVY-WEIGHT'"`
	WeightLimit     int                `json:"weight_limit" validate:"required,lte=500"`
	BatteryCapacity int                `json:"battery_capacity" validate:"required,lte=100"`
	Status          models.DroneStatus `json:"status" validate:"required,oneof='IDLE' 'LOADING' 'LOADED' 'DELIVERED' 'RETURNING'"`
}

func NewRegisterDroneRequest() *RegisterDrone {
	return &RegisterDrone{}
}

func InjectableRegisterDroneRequest() interface{} {
	return NewRegisterDroneRequest()
}
