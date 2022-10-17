package logic

import (
	"follow_system/follow_service/internal/svc"
	"follow_system/follow_service/root/go/src/follow_system/pb/follow_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    svc.TokenContext
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx svc.TokenContext, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *follow_service.FollowRQ) (*follow_service.FollowRS, error) {
	// todo: add your logic here and delete this line

	return &follow_service.FollowRS{}, nil
}
