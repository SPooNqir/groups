package groups

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	groups_grpc "github.com/slavayssiere-spoon/groups"
)

type MockGrpcGroup struct {
	cc *grpc.ClientConn
}

func (s MockGrpcGroup) Get(ctx context.Context, in *groups_grpc.Group, opts ...grpc.CallOption) (*groups_grpc.Group, error) {
	if in.Id == 401 {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	} else if in.Id == 404 {
		return nil, status.Error(codes.NotFound, "group not found")
	} else if in.Id == 500 {
		return nil, status.Error(codes.Internal, "internal error")
	} else if in.Id == 501 {
		return nil, status.Error(codes.Unavailable, "Unavailable")
	}
	return &groups_grpc.Group{
		Id:    in.Id,
		Name:  "group-" + fmt.Sprint(in.Id),
		Paths: []string{"group-" + fmt.Sprint(in.Id)},
	}, nil
}

func (s MockGrpcGroup) GetAll(ctx context.Context, in *groups_grpc.Groups, opts ...grpc.CallOption) (*groups_grpc.Groups, error) {
	return nil, nil
}

func (s MockGrpcGroup) GetByName(ctx context.Context, in *groups_grpc.Group, opts ...grpc.CallOption) (*groups_grpc.Group, error) {
	return nil, nil
}

func (s MockGrpcGroup) GetGraph(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*groups_grpc.Groups, error) {
	return nil, nil
}

func (s MockGrpcGroup) Add(ctx context.Context, in *groups_grpc.Group, opts ...grpc.CallOption) (*groups_grpc.Group, error) {
	return nil, nil
}

func (s MockGrpcGroup) Update(ctx context.Context, in *groups_grpc.Group, opts ...grpc.CallOption) (*groups_grpc.Group, error) {
	return nil, nil
}

func (s MockGrpcGroup) AddSubGroup(ctx context.Context, in *groups_grpc.Group, opts ...grpc.CallOption) (*groups_grpc.Group, error) {
	return nil, nil
}
