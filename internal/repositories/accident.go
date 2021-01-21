package repositories

import (
	"5g-v2x-kafka-worker-service/internal/config"
	"5g-v2x-kafka-worker-service/internal/infrastructures/grpc"
	proto "5g-v2x-kafka-worker-service/pkg/api"
	"context"
)

// AccidentRepository ...
type AccidentRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewAccidentRepository ...
func NewAccidentRepository(c *config.Config, g grpc.GRPC) *AccidentRepository {
	return &AccidentRepository{
		config: c,
		GRPC:   g,
	}
}

// CreateAccidentData ...
func (r *AccidentRepository) CreateAccidentData(request *proto.AccidentData) (*proto.CreateAccidentDataResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).CreateAccidentData(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
