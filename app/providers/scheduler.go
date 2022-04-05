package providers

import (
	"drones/app"
	"drones/app/data/adapters"
	"drones/app/scheduler"
	"errors"
	"github.com/robfig/cron/v3"
	"os"
)

const DronesLogsFilePath = "storage/logs/drones.log"

type SchedulerProvider struct {
	App  *app.App
	Cron *cron.Cron
}

func NewSchedulerProvider(a *app.App) *SchedulerProvider {
	return &SchedulerProvider{
		App:  a,
		Cron: cron.New(),
	}
}

func (s *SchedulerProvider) Boot() error {
	dronesAdapter := adapters.NewDronesAdapter(s.App.DB)

	dronesCronService := scheduler.NewDronesCronService(dronesAdapter)

	if _, err := os.Stat(DronesLogsFilePath); errors.Is(err, os.ErrNotExist) {
		_, createErr := os.Create(DronesLogsFilePath)
		if createErr != nil {
			return createErr
		}
	}

	_, err := s.Cron.AddFunc("@every 1m", dronesCronService.LogBatteryLevels(DronesLogsFilePath))
	if err != nil {
		return err
	}

	go func() {
		s.Cron.Run()
	}()

	return nil
}
