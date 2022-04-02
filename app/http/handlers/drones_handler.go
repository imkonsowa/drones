package handlers

import (
	"drones/app/data/adapters"
	"drones/app/data/models"
	"drones/app/http/requests"
	"drones/app/http/responses"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
		Status:          request.Status,
		BatteryCapacity: request.BatteryCapacity,
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

	medications, err := d.MedicationsAdapter.BatchCreate(medications)
	if err != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusInternalServerError).
			Message("failed to create medications").
			Send()
		return
	}

	responses.NewContextResponse(context).
		Success().
		Code(http.StatusOK).
		Message("medications loaded successfully").
		Send()
}

func (d *DronesHandler) UpdateStatus(context *gin.Context) {
	ctxKey, _ := context.Get("request")
	request := ctxKey.(*requests.UpdateDroneStatus)

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
