package delivery

import (
	"context"
	"github.com/Michael-Levitin/PingOmetr/internal/logic"
	pb "github.com/Michael-Levitin/PingOmetr/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var p = NewPingServer(logic.NewPingLogicMock())

func TestPingServer_GetFastest(t *testing.T) {
	got, err := p.GetFastest(context.Background(), &pb.GetFastestRequest{})
	assert.NoError(t, err)
	site := pb.GetResponse{Ping: 9, Site: "vk.com", Error: ""}
	assert.Equal(t, site.Ping, got.Ping)
	assert.Equal(t, site.Site, got.Site)
	assert.Equal(t, site.Error, got.Error)

}

func TestPingServer_GetSlowest(t *testing.T) {
	got, err := p.GetSlowest(context.Background(), &pb.GetSlowestRequest{})
	assert.NoError(t, err)
	site := pb.GetResponse{Ping: 230, Site: "sfggs.ni", Error: ""}
	assert.Equal(t, site.Ping, got.Ping)
	assert.Equal(t, site.Site, got.Site)
	assert.Equal(t, site.Error, got.Error)
}

func TestPingServer_GetSpecific(t *testing.T) {
	tests := []struct {
		name    string
		site    string
		want    *pb.GetResponse
		wantErr error
	}{
		{
			name: "site 1",
			site: "sdsdfs.rh",
			want: &pb.GetResponse{
				Ping:  156,
				Site:  "sdsdfs.rh",
				Error: "",
			},
			wantErr: nil,
		},
		{
			name: "site 2",
			site: "grert.wu",
			want: &pb.GetResponse{
				Ping:  56,
				Site:  "grert.wu",
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

	//got, err := p.GetSlowest(context.Background(), &pb.GetSlowestRequest{})
	//assert.NoError(t, err)
	//site := pb.GetResponse{Ping: 230, Site: "sfggs.ni", Error: ""}
	//assert.Equal(t, site.Ping, got.Ping)
	//assert.Equal(t, site.Site, got.Site)
	//assert.Equal(t, site.Error, got.Error)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.GetSpecific(context.Background(), &pb.GetSpecificRequest{SiteName: tt.site})
			if err != tt.wantErr {
				t.Errorf("GetSpecific() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.Ping, got.Ping)
			assert.Equal(t, tt.want.Site, got.Site)
			assert.Equal(t, tt.want.Error, got.Error)
		})
	}
}

func TestPingServer_GetAdminData(t *testing.T) {
	got, err := p.GetAdminData(context.Background(), &pb.GetAdminDataRequest{})
	assert.NoError(t, err)
	data := pb.GetAdminDataResponse{
		Min:      456,
		Max:      1234,
		Specific: 789,
	}
	assert.Equal(t, data.Min, got.Min)
	assert.Equal(t, data.Max, got.Max)
	assert.Equal(t, data.Specific, got.Specific)
}
