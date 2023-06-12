package logic

import (
	"context"
	ob "github.com/Michael-Levitin/PingOmetr/internal/objects"
)

type PingLogicI interface {
	GetFastest(ctx context.Context) (*ob.PingUser, error)
	GetSlowest(ctx context.Context) (*ob.PingUser, error)
	GetSpecific(ctx context.Context, site string) (*ob.PingUser, error)
	GetAdminData(ctx context.Context) (*ob.PingAdmin, error)
}
