package services

import (
	"5g-v2x-data-management-service/internal/models"
	"5g-v2x-data-management-service/internal/repositories"
	"time"
)

type DrowsinessService struct {
	crud *repositories.CRUDRepository
}

func NewDrowsinessService(crud *repositories.CRUDRepository) *DrowsinessService {
	return &DrowsinessService{
		crud: crud,
	}
}

func (as *DrowsinessService) StoreData(username string,carID string,lat float64,lng float64,time time.Time,responseTime float64,workingHour string) error {
	var drowsiness models.Drowsiness
	drowsiness.Username = username
	drowsiness.CarID = carID
	drowsiness.Latitude = lat
	drowsiness.Longitude = lng
	drowsiness.Time = time
	drowsiness.WorkingHour = workingHour
	drowsiness.ResponseTime = responseTime
	_, err := as.crud.Create("drowsiness", &drowsiness)

	if err != nil {
		panic(err.Error())
	}
	return nil
}
