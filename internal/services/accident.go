package services

import (
	"5g-v2x-data-management-service/internal/models"
	"5g-v2x-data-management-service/internal/repositories"
	"time"
)

type AccidentService struct {
	crud *repositories.CRUDRepository
}

func NewAccidentService(crud *repositories.CRUDRepository) *AccidentService {
	return &AccidentService{
		crud: crud,
	}
}

func (as *AccidentService) StoreData(username string,carID string,lat float64,lng float64,time time.Time) error {
	var accident models.Accident
	accident.Username = username
	accident.CarID = carID
	accident.Latitude = lat
	accident.Longitude = lng
	accident.Time = time

	_, err := as.crud.Create("accident", &accident)

	if err != nil {
		panic(err.Error())
	}
	return nil
}
