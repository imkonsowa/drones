package handlers

type DroneModel string

const (
	LightWeight   DroneModel = "LIGHT-WEIGHT"
	MiddleWeight  DroneModel = "MIDDLE-WEIGHT"
	CruiserWeight DroneModel = "CRUISER-WEIGHT"
	HeavyWeight   DroneModel = "HEAVY-WEIGHT"
)

type DroneState string

const (
	Idle       DroneState = "IDLE"
	Loading    DroneState = "LOADING"
	Loaded     DroneState = "LOADED"
	Delivering DroneState = "DELIVERED"
	Delivered  DroneState = "RETURNING"
)

type Drone struct {
	SerialNumber    string     `json:"serial_number"`
	Model           DroneModel `json:"model"`
	WeightLimit     int        `json:"weight_limit"`
	BatteryCapacity int        `json:"battery_capacity"`
	State           DroneState `json:"state"`
}

type Medication struct {
	Name     string
	Weight   int
	Code     string
	ImageURL string
}
