package logic

import (
	"follow_system/follow_service/internal/svc"
	"follow_system/follow_service/root/go/src/follow_system/pb/follow_service"

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

func (l *GetFansListLogic) GetFansList(in *follow_service.GetFansListRQ) (*follow_service.GetFansListRS, error) {
	// todo: add your logic here and delete this line

	return &follow_service.GetFansListRS{}, nil
}
