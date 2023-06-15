package logic

import (
	"context"
	ob "github.com/Michael-Levitin/PingOmetr/internal/objects"
)

type PingLogicMock struct {
}

func NewPingLogicMock() *PingLogicMock {
	return &PingLogicMock{}
}

func (l PingLogicMock) GetFastest(ctx context.Context) (*ob.PingUser, error) {
	return &ob.PingUser{
		Msec:  9,
		Site:  "vk.com",
		Error: "",
	}, nil
}

func (l PingLogicMock) GetSlowest(ctx context.Context) (*ob.PingUser, error) {
	return &ob.PingUser{
		Msec:  230,
		Site:  "sfggs.ni",
		Error: "",
	}, nil
}

func (l PingLogicMock) GetSpecific(ctx context.Context, site string) (*ob.PingUser, error) {
	switch site {
	case "sdsdfs.rh":
		return &ob.PingUser{Msec: 156, Site: "sdsdfs.rh", Error: ""}, nil
	case "grert.wu":
		return &ob.PingUser{Msec: 56, Site: "grert.wu", Error: ""}, nil
	}
	return &ob.PingUser{Msec: 0, Site: "", Error: "site is not on the list"}, nil
}

func (l PingLogicMock) GetAdminData(ctx context.Context) (*ob.PingAdmin, error) {
	return &ob.PingAdmin{
		456,
		1234,
		789,
	}, nil
}
