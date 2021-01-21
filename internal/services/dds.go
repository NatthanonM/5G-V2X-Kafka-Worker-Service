package services

import (
	"5g-v2x-kafka-worker-service/internal/repositories"
	proto "5g-v2x-kafka-worker-service/pkg/api"
	"time"

	"github.com/golang/protobuf/ptypes"
)

type DrowsinessService struct {
	DrowsinessRepository *repositories.DrowsinessRepository
}

func NewDrowsinessService(drowsinessRepository *repositories.DrowsinessRepository) *DrowsinessService {
	return &DrowsinessService{
		DrowsinessRepository: drowsinessRepository,
	}
}

func (as *DrowsinessService) StoreData(username string, carID string, lat float64, lng float64, time time.Time, responseTime float64, workingHour float64) error {

	protoTime, err := ptypes.TimestampProto(time)

	if err != nil {
		panic(err.Error())
	}

	drowsiness := &proto.DrowsinessData{
		Username:     username,
		CarId:        carID,
		Latitude:     lat,
		Longitude:    lng,
		Time:         protoTime,
		WorkingHour:  workingHour,
		ResponseTime: responseTime,
	}

	_, err = as.DrowsinessRepository.CreateDrowsinessData(drowsiness)

	if err != nil {
		panic(err.Error())
	}

	return nil
}
