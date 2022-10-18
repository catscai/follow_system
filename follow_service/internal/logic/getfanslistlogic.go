package logic

import (
	"follow_system/follow_service/internal/svc"
	"follow_system/pb/Follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFansListLogic struct {
	ctx    svc.TokenContext
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFansListLogic(ctx svc.TokenContext, svcCtx *svc.ServiceContext) *GetFansListLogic {
	return &GetFansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFansListLogic) GetFansList(in *Follow.GetFansListRQ) (*Follow.GetFansListRS, error) {
	// todo: add your logic here and delete this line
	
	return &Follow.GetFansListRS{}, nil
}
