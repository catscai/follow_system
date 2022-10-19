package logic

import (
	"fmt"
	"follow_system/follow_service/internal/svc"
	"follow_system/pb/Follow"
	"math"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis/v8"
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

/*
create table follow_event_100 (
	event_id 	int not null auto increment,
	uid 		int not null,
	fans 		int not null,
	is_del 		tinyint not null default '0',
	update_time int not null default NOW(),
	index key(event_id,uid,fans),
	index key(uid,event_id,fans)
)engine=innodb,charset=utf8;

*/

type FollowListItem struct {
	eventID int32 `db:"event_id"`
	fans    int64 `db:"fans"`
}

func (l *FollowLogic) Follow(in *Follow.FollowRQ) (*Follow.FollowRS, error) {
	bucketNumKey := fmt.Sprintf("u:%d:follow:event:%d:bucket:num", in.GetUid(), in.GetEventType())
	const elementNumOfPerBucket = 10 // 每个桶中存储多少个元素
	bucketNum, err := l.svcCtx.CacheRedis.Get(l.ctx.Context, bucketNumKey).Int()
	if err != nil && err != redis.ErrNil {
		l.Error("redis get bucket failed", err)
		return nil, err
	}

	tableName := fmt.Sprintf("follow_event_%d", in.GetEventType())
	fansList := make([]*FollowListItem, 0, 3)
	if err == redis.ErrNil {
		queryStr := fmt.Sprintf("select event_id,fans from %s where uid=? and is_del=?", tableName)
		if err = l.svcCtx.FollowMysql.QueryRows(fansList, queryStr, in.GetUid(), 0); err != nil {
			l.Error(fmt.Sprintf("QueryRows follow list failed", queryStr, err))
			return nil, err
		}
		// 将查询的列表分段写入redis
		bucketNum = int(math.Floor((float64(len(fansList)) / float64(elementNumOfPerBucket))))
		wg := &sync.WaitGroup{}
		wg.Add(bucketNum)
		for i := 0; i < bucketNum; i++ {
			items := make([]*redis.Z, 0, elementNumOfPerBucket)
			for j := i * elementNumOfPerBucket; j < (i+1)*elementNumOfPerBucket && j < len(fansList); j++ {
				items = append(items, &redis.Z{
					Score:  float64(fansList[j].eventID),
					Member: fansList[j].fans,
				})
			}

			fansKey := fmt.Sprintf("u:%d:fans:list:event:%d:bucket:%d", in.GetUid(), in.GetEventType(), i)
			go func() {
				defer wg.Done()
				err := l.svcCtx.CacheRedis.ZAdd(l.ctx.Context, fansKey, items...).Err()
				if err != nil {
					l.Error("redis zadd failed", fansKey, err, len(items))
				}
			}()
		}
		wg.Wait()
		go func() {
			if err := l.svcCtx.CacheRedis.Set(l.ctx.Context, bucketNumKey, bucketNum, -1).Err(); err != nil {
				l.Error("redis set bucket number failed", err, bucketNumKey, bucketNum)
			}
		}()
	}

	incrEventIDKey := fmt.Sprintf("u:%d:type:%d:incr:event:id", in.GetUid(), in.GetEventType())
	eventID, err := l.svcCtx.CacheRedis.Incr(l.ctx.Context, incrEventIDKey).Result()
	if err != nil {
		l.Error("redis incr event id failed", incrEventIDKey, err)
		return nil, err
	}

	fansKey := fmt.Sprintf("u:%d:fans:list:event:%d:bucket:%d", in.GetUid(), in.GetEventType(), bucketNum)
	script := "local length = redis.call('ZCARD',KEYS[1]); if tonumber(length) < tonumer(ARGV[1]) then redis.call('ZADD',KEYS[1],ARGV[2],ARGV[3]);return 0; else return 1;"

	n, err := l.svcCtx.CacheRedis.Eval(l.ctx.Context, script, []string{fansKey}, elementNumOfPerBucket, eventID, in.GetFans()).Int()
	if err != nil {
		l.Error("redis eval failed", fansKey, err)
		return nil, err
	}
	if n == 1 {
		newBucket, _ := l.svcCtx.CacheRedis.Incr(l.ctx.Context, bucketNumKey).Result()
		fansKey = fmt.Sprintf("u:%d:fans:list:event:%d:bucket:%d", in.GetUid(), in.GetEventType(), newBucket)
		l.svcCtx.CacheRedis.Eval(l.ctx.Context, script, []string{fansKey}, elementNumOfPerBucket, eventID, in.GetFans()).Int()
	}

	go func() {
		l.svcCtx.FollowMysql.Exec(fmt.Sprintf("insert into %s values(null,?,?,?,?,?)", tableName), eventID, in.GetUid(), in.GetFans(), 0, time.Now().Unix())
	}()

	return &Follow.FollowRS{}, nil
}
