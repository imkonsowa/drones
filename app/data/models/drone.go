package models

type DroneModel string

const (
	LightWeight   DroneModel = "LIGHT-WEIGHT"
	MiddleWeight  DroneModel = "MIDDLE-WEIGHT"
	CruiserWeight DroneModel = "CRUISER-WEIGHT"
	HeavyWeight   DroneModel = "HEAVY-WEIGHT"
)

var _ = []DroneModel{
	LightWeight,
	MiddleWeight,
	CruiserWeight,
	HeavyWeight,
}

type DroneStatus string

const (
	Idle       DroneStatus = "IDLE"
	Loading    DroneStatus = "LOADING"
	Loaded     DroneStatus = "LOADED"
	Delivering DroneStatus = "DELIVERED"
	Delivered  DroneStatus = "RETURNING"
)

var _ = []DroneStatus{
	Idle,
	Loading,
	Loaded,
	Delivering,
	Delivered,
}

type Drone struct {
	AppModel
	SerialNumber    string      `json:"serial_number"`
	Model           DroneModel  `json:"model"`
	WeightLimit     int         `json:"weight_limit"`
	BatteryCapacity int         `json:"battery_capacity"`
	Status          DroneStatus `json:"state"`
}
