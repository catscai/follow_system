package logic

import (
	"follow_system/follow_service/internal/svc"
	"follow_system/follow_service/root/go/src/follow_system/pb/follow_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    svc.TokenContext
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx svc.TokenContext, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *follow_service.FollowReq) (*follow_service.FollowRsp, error) {
	// todo: add your logic here and delete this line

	return &follow_service.FollowRsp{}, nil
}
