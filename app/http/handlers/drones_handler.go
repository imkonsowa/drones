package handlers

import (
	"drones/app/data/adapters"
	"drones/app/data/models"
	"drones/app/http/requests"
	"drones/app/http/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const MaxLoadingBatteryThreshold = 25

type DronesHandler struct {
	DronesAdapter      *adapters.DronesAdapter
	MedicationsAdapter *adapters.MedicationsAdapter
}

func NewDronesHandler(dronesAdapter *adapters.DronesAdapter, medicationsAdapter *adapters.MedicationsAdapter) *DronesHandler {
	return &DronesHandler{
		DronesAdapter:      dronesAdapter,
		MedicationsAdapter: medicationsAdapter,
	}
}

func (d *DronesHandler) Register(context *gin.Context) {
	ctxKey, _ := context.Get("request")
	request := ctxKey.(*requests.RegisterDrone)

	drone := &models.Drone{
		Model:           request.Model,
		SerialNumber:    request.SerialNumber,
		WeightLimit:     request.WeightLimit,
		BatteryCapacity: request.BatteryCapacity,
		Status:          models.Idle,
	}

	drone = d.DronesAdapter.Create(drone)

	responses.NewContextResponse(context).
		Success().
		Code(http.StatusOK).
		Message("drone registered successfully").
		Send()
}

func (d *DronesHandler) LoadMedications(context *gin.Context) {
	ctxKey, _ := context.Get("request")
	request := ctxKey.(*requests.LoadMedications)

	drone, droneErr := d.DronesAdapter.GetBySerialNumber(request.SerialNumber)
	if droneErr != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Send()
		return
	}

	if drone.BatteryCapacity < MaxLoadingBatteryThreshold {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusUnprocessableEntity).
			Message(fmt.Sprintf("can't load this drone with medications; Battery capacity: %d", drone.BatteryCapacity)).
			Send()
		return
	}

	var medications []models.Medication
	medicationsWeights := 0
	for _, medication := range request.Medications {
		medicationsWeights += medication.Weight

		m := models.Medication{
			DroneSerialNumber: drone.SerialNumber,
			Name:              medication.Name,
			Code:              medication.Code,
			Weight:            medication.Weight,
		}

		medications = append(medications, m)
	}

	loadedMedications, loadedMedicationsErr := d.MedicationsAdapter.GetDroneMedications(drone.SerialNumber)
	if loadedMedicationsErr != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Send()
	}

	loadedMedicationsWeights := 0
	for _, medication := range loadedMedications {
		loadedMedicationsWeights += medication.Weight
	}
	if loadedMedicationsWeights+medicationsWeights > drone.WeightLimit {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusUnprocessableEntity).
			Message("medications exceeds the max allowed weight limit.").
			Send()
		return
	}

	updateErr := d.DronesAdapter.UpdateStatus(request.SerialNumber, models.Loading)
	if updateErr != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Message("failed to update drone status to LOADING").
			Send()
	}

	medications, err := d.MedicationsAdapter.BatchCreate(medications)
	if err != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Message("failed to create medications").
			Send()
		return
	}

	// check if drone is fully loaded
	go func() {
		loadedMedications, _ = d.MedicationsAdapter.GetDroneMedications(drone.SerialNumber)
		if loadedMedicationsErr != nil {
			log.Fatalf("failed to fetch loaded medications")
			return
		}

		loadedMedicationsWeights = 0
		for _, medication := range loadedMedications {
			loadedMedicationsWeights += medication.Weight
		}
		if loadedMedicationsWeights >= drone.WeightLimit {
			// Update to LOADED if loaded items reached weight limit
			updateErr = d.DronesAdapter.UpdateStatus(request.SerialNumber, models.Loaded)
			if updateErr != nil {
				responses.NewContextResponse(context).
					Error().
					Code(http.StatusInternalServerError).
					Message("failed to update drone status to LOADED").
					Send()
			} else {
				// update to IDLE if still eligible for loading items
				updateErr = d.DronesAdapter.UpdateStatus(request.SerialNumber, models.Idle)
				responses.NewContextResponse(context).
					Error().
					Code(http.StatusInternalServerError).
					Message("failed to update drone status to LOADED").
					Send()
			}
		}
	}()

	responses.NewContextResponse(context).
		Success().
		Code(http.StatusOK).
		Message("medications loaded successfully").
		Send()
}

func (d *DronesHandler) UpdateStatus(context *gin.Context) {
	ctxKey, _ := context.Get("request")
	request := ctxKey.(*requests.UpdateDroneStatus)

	drone, droneErr := d.DronesAdapter.GetBySerialNumber(request.SerialNumber)
	if droneErr != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Send()
		return
	}

	if request.Status == models.Loading && drone.BatteryCapacity < MaxLoadingBatteryThreshold {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusUnprocessableEntity).
			Message(fmt.Sprintf("can't update satus to LOADING; Battery capacity: %d", drone.BatteryCapacity)).
			Send()
		return
	}

	err := d.DronesAdapter.UpdateStatus(request.SerialNumber, request.Status)
	if err != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Message("failed to update drone status").
			Send()
	}

	responses.NewContextResponse(context).
		Success().
		Code(http.StatusOK).
		Send()
}

func (d *DronesHandler) GetBatteryCapacity(context *gin.Context) {
	serial := context.Param("serialNumber")

	if len(serial) == 0 {
		responses.
			NewContextResponse(context).
			Error().
			Code(http.StatusBadRequest).
			Message("invalid serial number").
			Send()
		return
	}

	drone, err := d.DronesAdapter.GetBySerialNumber(serial)
	if err != nil {
		log.Print(err)
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Message("failed to resolve drone").
			Send()
		return
	}
	if drone == nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Message("invalid serial number").
			Send()
	}

	responses.NewContextResponse(context).
		Success().
		Code(http.StatusOK).
		Data(map[string]interface{}{
			"battery_capacity": drone.BatteryCapacity,
		}).
		Send()
}

func (d *DronesHandler) GetIdleDrones(context *gin.Context) {
	drones, dronesError := d.DronesAdapter.GetDronesByStatus(models.Idle)
	if dronesError != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Send()
		return
	}

	responses.NewContextResponse(context).
		Success().
		Code(http.StatusOK).
		Data(map[string]interface{}{
			"drones": drones,
		}).
		Send()
}
