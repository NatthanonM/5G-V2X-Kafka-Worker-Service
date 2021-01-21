package repositories

import (
	"5g-v2x-kafka-worker-service/internal/config"
	"5g-v2x-kafka-worker-service/internal/infrastructures/grpc"
	proto "5g-v2x-kafka-worker-service/pkg/api"
	"context"
)

// DrowsinessRepository ...
type DrowsinessRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewDrowsinessRepository ...
func NewDrowsinessRepository(c *config.Config, g grpc.GRPC) *DrowsinessRepository {
	return &DrowsinessRepository{
		config: c,
		GRPC:   g,
	}
}

// CreateDrowsinessData ...
func (r *DrowsinessRepository) CreateDrowsinessData(request *proto.DrowsinessData) (*proto.CreateDrowsinessDataResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).CreateDrowsinessData(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
