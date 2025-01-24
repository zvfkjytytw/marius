package mulusgrpc

import (
	"context"

	"github.com/zvfkjytytw/marius/proto/mulus/api/v1"
)

func (s *ServerGRPC) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{
		Hello: "mus",
	}, nil
}

func (s *ServerGRPC) SaveData(ctx context.Context, request *api.SaveRequest) (*api.SaveResponse, error) {
	if err := s.storage.SaveData(request.Name, request.Data); err != nil {
		return nil, err
	}

	return &api.SaveResponse{
		Name: request.Name,
	}, nil
}

func (s *ServerGRPC) GetData(ctx context.Context, request *api.GetRequest) (*api.GetResponse, error) {
	data, err := s.storage.GetData(request.Name)
	if err != nil {
		return nil, err
	}

	return &api.GetResponse{
		Name: request.Name,
		Data: data,
	}, nil
}

func (s *ServerGRPC) DeleteData(ctx context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	if err := s.storage.DeleteData(request.Name); err != nil {
		return nil, err
	}

	return &api.DeleteResponse{
		Name: request.Name,
	}, nil
}
