// Code generated by goctl. DO NOT EDIT!
// Source: Follow.proto

package server

import (
	"context"
	"encoding/json"

	"follow_system/follow_service/internal/logic"
	"follow_system/follow_service/internal/svc"
	"follow_system/pb/Follow"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type FollowServiceServer struct {
	svcCtx *svc.ServiceContext
	Follow.UnimplementedFollowServiceServer
}

func NewFollowServiceServer(svcCtx *svc.ServiceContext) *FollowServiceServer {
	return &FollowServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *FollowServiceServer) Follow(ctx context.Context, in *Follow.FollowRQ) (*Follow.FollowRS, error) {
	tokenCtx := svc.TokenContext{
		Context: ctx,
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if datas, ok := md["context"]; ok && len(datas) > 0 {
			if err := json.Unmarshal([]byte(datas[0]), &tokenCtx.COMM); err != nil {
				logx.WithContext(ctx).Error("context common data json Unmarshal failed, data:", datas[0], "err:", err)
			}

		}
	}
	l := logic.NewFollowLogic(tokenCtx, s.svcCtx)
	return l.Follow(in)
}

func (s *FollowServiceServer) GetFollowList(ctx context.Context, in *Follow.GetFollowListRQ) (*Follow.GetFollowListRS, error) {
	tokenCtx := svc.TokenContext{
		Context: ctx,
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if datas, ok := md["context"]; ok && len(datas) > 0 {
			if err := json.Unmarshal([]byte(datas[0]), &tokenCtx.COMM); err != nil {
				logx.WithContext(ctx).Error("context common data json Unmarshal failed, data:", datas[0], "err:", err)
			}

		}
	}
	l := logic.NewGetFollowListLogic(tokenCtx, s.svcCtx)
	return l.GetFollowList(in)
}

func (s *FollowServiceServer) GetFansList(ctx context.Context, in *Follow.GetFansListRQ) (*Follow.GetFansListRS, error) {
	tokenCtx := svc.TokenContext{
		Context: ctx,
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if datas, ok := md["context"]; ok && len(datas) > 0 {
			if err := json.Unmarshal([]byte(datas[0]), &tokenCtx.COMM); err != nil {
				logx.WithContext(ctx).Error("context common data json Unmarshal failed, data:", datas[0], "err:", err)
			}

		}
	}
	l := logic.NewGetFansListLogic(tokenCtx, s.svcCtx)
	return l.GetFansList(in)
}
