package logic

import (
	"fmt"
	"follow_system/follow_service/internal/svc"
	"follow_system/pb/Follow"

	"github.com/garyburd/redigo/redis"
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
	bucketNumKey := fmt.Sprintf("u:%d:follow:event:%d:bucket:num", in.GetUid(), in.GetEventType())
	const elementNumOfPerBucket = 10 // 每个桶中存储多少个元素
	bucketNum,err := l.svcCtx.CacheRedis.Get(l.ctx.Context, bucketNumKey).Int()
	if err != nil && err != redis.ErrNil {
		l.Error("redis get bucket failed", err)
		return nil, err
	}
	if err == redis.ErrNil {
		l.svcCtx.FollowMysql.QueryRows()
	}

	return &Follow.FollowRS{}, nil
}
