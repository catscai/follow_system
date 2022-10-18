package logic

import (
	"follow_system/follow_service/internal/svc"
	"follow_system/pb/Follow"

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

func (l *FollowLogic) Follow(in *Follow.FollowRQ) (*Follow.FollowRS, error) {
	// todo: add your logic here and delete this line
	
	return &Follow.FollowRS{}, nil
}
