package logic

import (
	"context"

	"github.com/tqtcloud/mall/service/user/model"
	"github.com/tqtcloud/mall/service/user/rpc/internal/svc"
	"github.com/tqtcloud/mall/service/user/rpc/types/github.com/tqtcloud/mall/service/user"
	"google.golang.org/grpc/status"

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

	// 查询用户是否存在
	resp, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     resp.Id,
		Name:   resp.Name,
		Gender: resp.Gender,
		Mobile: resp.Mobile,
	}, nil
}
