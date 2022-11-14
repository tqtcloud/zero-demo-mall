package logic

import (
	"context"

	"github.com/tqtcloud/mall/service/user/rpc/types/github.com/tqtcloud/mall/service/user"
	"github.com/tqtcloud/mall/service/user/rpcgoct/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResponse{}, nil
}
