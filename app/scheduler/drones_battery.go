package scheduler

import (
	"drones/app/data/adapters"
	"fmt"
	"os"
	"time"
)

type DronesCronService struct {
	DronesAdapter *adapters.DronesAdapter
}

func NewDronesCronService(adapter *adapters.DronesAdapter) *DronesCronService {
	return &DronesCronService{
		adapter,
	}
}

func (d *DronesCronService) LogBatteryLevels(dronesLogsFilePath string) func() {
	return func() {
		file, openErr := os.OpenFile(dronesLogsFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if openErr != nil {
			fmt.Println(fmt.Errorf("failed to open file, %v", openErr))
		}

		logRecord := fmt.Sprintf("[%s] start: \n", time.Now().String())

		drones := d.DronesAdapter.DronesList()

		for _, drone := range drones {
			logRecord += fmt.Sprintf("-> DroneID: %d - Battery Level: %d\n", drone.ID, drone.BatteryCapacity)
		}

		logRecord += fmt.Sprintf("[%s] end; \n", time.Now().String())
		fmt.Println(logRecord)

		file.WriteString(logRecord)
	}
}
