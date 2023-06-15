package delivery

import (
	"context"
	pb "github.com/Michael-Levitin/PingOmetr/proto"
)

type PingMockServer struct {
	pb.UnimplementedPingOmetrServer
}

func NewPingMockServer() *PingMockServer {
	return &PingMockServer{}
}

func (p *PingMockServer) GetFastest(ctx context.Context, in *pb.GetFastestRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Ping: 10, Site: "vk.com", Error: ""}, nil
}

func (p *PingMockServer) GetSlowest(ctx context.Context, in *pb.GetSlowestRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Ping: 456, Site: "uasodiu.so", Error: ""}, nil
}

func (p *PingMockServer) GetSpecific(ctx context.Context, in *pb.GetSpecificRequest) (*pb.GetResponse, error) {
	switch in.GetSiteName() {
	case "abdu.so":
		return &pb.GetResponse{Ping: 321, Site: "abdu.so", Error: ""}, nil
	case "wqwedas.hu":
		return &pb.GetResponse{Ping: 456, Site: "wqwedas.hu", Error: ""}, nil
	case "lskadjask":
		return &pb.GetResponse{Ping: 0, Site: "", Error: "site is not on the list"}, nil
	}
	return &pb.GetResponse{}, nil
}

func (p *PingMockServer) GetAdminData(ctx context.Context, in *pb.GetAdminDataRequest) (*pb.GetAdminDataResponse, error) {
	return &pb.GetAdminDataResponse{
		Min:      234234,
		Max:      456456,
		Specific: 456346,
	}, nil
}
