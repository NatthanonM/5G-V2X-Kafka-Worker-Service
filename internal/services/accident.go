package services

import (
	"5g-v2x-kafka-worker-service/internal/repositories"
	proto "5g-v2x-kafka-worker-service/pkg/api"
	"time"

	"github.com/golang/protobuf/ptypes"
)

type AccidentService struct {
	AccidentRepository *repositories.AccidentRepository
}

func NewAccidentService(accidentRepository *repositories.AccidentRepository) *AccidentService {
	return &AccidentService{
		AccidentRepository: accidentRepository,
	}
}

func (as *AccidentService) StoreData(username string, carID string, lat float64, lng float64, time time.Time) error {
	protoTime, err := ptypes.TimestampProto(time)

	if err != nil {
		panic(err.Error())
	}

	accident := &proto.AccidentData{
		Username:  username,
		CarId:     carID,
		Latitude:  lat,
		Longitude: lng,
		Time:      protoTime,
	}

	_, err = as.AccidentRepository.CreateAccidentData(accident)

	if err != nil {
		panic(err.Error())
	}

	return nil
}
