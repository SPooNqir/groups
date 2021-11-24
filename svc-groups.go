package groups

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	lib "gitlab.com/SpoonQIR/Cloud/library/golang-common.git"
)

type GroupService struct {
	Groupsconn *grpc.ClientConn
	Groupssvc  GroupsClient
	Groupreco  chan bool

	Id *lib.Identity
}

// InitGroups tst
func (s *GroupService) InitGroups(groupsHost string, tracer opentracing.Tracer, logger *logrus.Logger) chan bool {
	logentry := logrus.NewEntry(logger)
	logopts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	otopts := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(tracer),
	}

	var err error

	connect := make(chan bool)

	go func(lconn chan bool) {
		for {
			logrus.Info("Wait for connect")
			r := <-lconn
			logrus.WithFields(logrus.Fields{"reconn": r}).Info("conn chan receive")
			if r {
				for i := 1; i < 5; i++ {
					s.Groupsconn, err = grpc.Dial(groupsHost,
						grpc.WithInsecure(),
						grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
							grpc_logrus.UnaryClientInterceptor(logentry, logopts...),
							grpc_opentracing.UnaryClientInterceptor(otopts...),
							grpc_prometheus.UnaryClientInterceptor,
						)),
						grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
							grpc_logrus.StreamClientInterceptor(logentry, logopts...),
							grpc_opentracing.StreamClientInterceptor(otopts...),
							grpc_prometheus.StreamClientInterceptor,
						)),
					)
					if err != nil {
						logger.Fatalf("did not connect: %v, try : %d - sleep 5s", err, i)
						time.Sleep(2 * time.Second)
					} else {
						s.Groupssvc = NewGroupsClient(s.Groupsconn)
						break
					}
				}
			} else {
				logrus.Info("end of goroutine - reconnect")
				return
			}
		}
	}(connect)

	logger.WithFields(logrus.Fields{"host": groupsHost}).Info("Connexion au service gRPC 'Groups'")

	//Identity
	s.Id = &lib.Identity{}
	go s.Id.Launch()

	connect <- true
	return connect
}

func (s *GroupService) GetGroup(ctx context.Context, groupid uint64) (*Group, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+s.Id.Tokens.JWTToken)
	for i := 0; i <= 3; i++ {
		nctx, cancel := context.WithTimeout(ctx, time.Duration(500)*time.Millisecond)
		defer cancel()
		grp, err := s.Groupssvc.Get(nctx, &Group{
			Id: groupid,
		})
		logrus.WithFields(logrus.Fields{"ctx.err": ctx.Err(), "err": err}).Trace("error ctx get group")

		if err != nil {
			logrus.WithFields(logrus.Fields{"err": err}).Error("error get group")
			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.Unavailable {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.Canceled {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.DeadlineExceeded {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.Aborted {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.Unauthenticated {
				return nil, status.Error(codes.Unauthenticated, "user Unauthenticated.")
			} else if errStatus.Code() == codes.InvalidArgument {
				return nil, status.Errorf(codes.InvalidArgument, "argument invalid %v", err)
			} else if errStatus.Code() == codes.NotFound {
				return nil, nil
			}
			// errStatus.Code() == codes.Internal = retry
		} else if ctx.Err() != nil {
			s.Groupreco <- true
		} else {
			return grp, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Group not found")
}

func (s *GroupService) GetGraph(ctx context.Context) (*Groups, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+s.Id.Tokens.JWTToken)
	for i := 1; i <= 5; i++ {
		nctx, cancel := context.WithTimeout(ctx, time.Duration(500)*time.Millisecond)
		defer cancel()
		grp, err := s.Groupssvc.GetGraph(nctx, &empty.Empty{})
		logrus.WithFields(logrus.Fields{"ctx.err": ctx.Err(), "err": err}).Trace("error ctx get graph")

		if err != nil {
			logrus.WithFields(logrus.Fields{"err": err}).Error("error get group")
			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.Unavailable {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.Canceled {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.DeadlineExceeded {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.Aborted {
				s.Groupreco <- true
			} else if errStatus.Code() == codes.Unauthenticated {
				return nil, status.Error(codes.Unauthenticated, "user Unauthenticated.")
			} else if errStatus.Code() == codes.InvalidArgument {
				return nil, status.Errorf(codes.InvalidArgument, "argument invalid %v", err)
			} else if errStatus.Code() == codes.NotFound {
				return nil, nil
			}
			// errStatus.Code() == codes.Internal = retry
		} else if ctx.Err() != nil {
			s.Groupreco <- true
		} else {
			return grp, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Group not found")
}
