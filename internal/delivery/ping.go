package delivery

import (
	"context"
	"github.com/Michael-Levitin/PingOmetr/internal/logic"
	pb "github.com/Michael-Levitin/PingOmetr/proto"
)

type PingServer struct {
	pb.UnimplementedPingOmetrServer
	logic logic.PingLogicI
}

func NewPingServer(logic logic.PingLogicI) *PingServer {
	return &PingServer{logic: logic}
}

func (p *PingServer) GetFastest(ctx context.Context, in *pb.GetFastestRequest) (*pb.GetResponse, error) {
	out, err := p.logic.GetFastest(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Ping: out.Msec, Site: out.Site, Error: out.Error}, nil
}

func (p *PingServer) GetSlowest(ctx context.Context, in *pb.GetSlowestRequest) (*pb.GetResponse, error) {
	out, err := p.logic.GetSlowest(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Ping: out.Msec, Site: out.Site, Error: out.Error}, nil
}

func (p *PingServer) GetSpecific(ctx context.Context, in *pb.GetSpecificRequest) (*pb.GetResponse, error) {
	out, err := p.logic.GetSpecific(ctx, in.GetSiteName())
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Ping: out.Msec, Site: out.Site, Error: out.Error}, nil
}

func (p *PingServer) GetAdminData(ctx context.Context, in *pb.GetAdminDataRequest) (*pb.GetAdminDataResponse, error) {
	out, err := p.logic.GetAdminData(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetAdminDataResponse{Max: out.Fastest, Min: out.Slowest, Specific: out.Specific}, nil
}
