package logic

import (
	"context"

	"github.com/tqtcloud/mall/common/cryptx"
	"github.com/tqtcloud/mall/service/user/model"
	"github.com/tqtcloud/mall/service/user/rpc/internal/svc"
	"github.com/tqtcloud/mall/service/user/rpc/types/github.com/tqtcloud/mall/service/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != resp.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.LoginResponse{
		Id:     resp.Id,
		Name:   resp.Name,
		Gender: resp.Gender,
		Mobile: resp.Mobile,
	}, nil
}
