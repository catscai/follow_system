package logic

import (
	"follow_system/follow_service/internal/svc"
	"follow_system/pb/Follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    svc.TokenContext
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx svc.TokenContext, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *Follow.GetFollowListRQ) (*Follow.GetFollowListRS, error) {
	// todo: add your logic here and delete this line

	return &Follow.GetFollowListRS{}, nil
}
