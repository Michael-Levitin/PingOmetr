package main

import (
	"context"
	"github.com/Michael-Levitin/PingOmetr/internal/delivery"
	"log"
	"net"
	"testing"

	pb "github.com/Michael-Levitin/PingOmetr/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	srv := delivery.NewPingMockServer()
	pb.RegisterPingOmetrServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("PingMockServer exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestServer_GetFastest(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewPingOmetrClient(conn)
	in := &pb.GetFastestRequest{}
	resp, err := client.GetFastest(ctx, in)
	assert.NoError(t, err)
	site := pb.GetResponse{Ping: 10, Site: "vk.com", Error: ""}
	assert.Equal(t, site.Ping, resp.Ping)
	assert.Equal(t, site.Site, resp.Site)
	assert.Equal(t, site.Error, resp.Error)
}

func TestServer_GetSlowest(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewPingOmetrClient(conn)
	in := &pb.GetSlowestRequest{}
	resp, err := client.GetSlowest(ctx, in)
	assert.NoError(t, err)
	site := pb.GetResponse{Ping: 456, Site: "uasodiu.so", Error: ""}
	assert.Equal(t, site.Ping, resp.Ping)
	assert.Equal(t, site.Site, resp.Site)
	assert.Equal(t, site.Error, resp.Error)
}

func TestServer_GetSpecific(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	tests := []struct {
		name    string
		site    string
		want    *pb.GetResponse
		wantErr error
	}{
		{
			name: "site 1",
			site: "abdu.so",
			want: &pb.GetResponse{
				Ping:  321,
				Site:  "abdu.so",
				Error: "",
			},
			wantErr: nil,
		},
		{
			name: "site 2",
			site: "wqwedas.hu",
			want: &pb.GetResponse{
				Ping:  456,
				Site:  "wqwedas.hu",
				Error: "",
			},
			wantErr: nil,
		},
		{
			name: "not on the list",
			site: "lskadjask",
			want: &pb.GetResponse{
				Ping:  0,
				Site:  "",
				Error: "site is not on the list",
			},
			wantErr: nil,
		},
	}

	client := pb.NewPingOmetrClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.GetSpecific(ctx, &pb.GetSpecificRequest{SiteName: tt.site})
			if err != tt.wantErr {
				t.Errorf("GetSpecific() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.Ping, resp.Ping)
			assert.Equal(t, tt.want.Site, resp.Site)
			assert.Equal(t, tt.want.Error, resp.Error)
		})
	}

}

func TestServer_GetSAdminData(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewPingOmetrClient(conn)
	in := &pb.GetAdminDataRequest{}
	resp, err := client.GetAdminData(ctx, in)
	assert.NoError(t, err)
	site := pb.GetAdminDataResponse{
		Min:      234234,
		Max:      456456,
		Specific: 456346,
	}
	assert.Equal(t, site.Max, resp.Max)
	assert.Equal(t, site.Max, resp.Max)
	assert.Equal(t, site.Specific, resp.Specific)
}
