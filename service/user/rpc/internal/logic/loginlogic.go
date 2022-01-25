package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/model"
	"mall/service/user/rpc/user"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/userclient"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LoginLogic) Login(in *userclient.LoginRequest) (*userclient.LoginResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "user not exist")
		}

		return nil,status.Error(500, err.Error())
	}

	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if res.Password != password {
		return nil, status.Error(100, "password error")
	}

	return &user.LoginResponse{
		Id: res.Id,
		Gender: res.Gender,
		Name: res.Name,
		Mobile: res.Mobile,
	},nil
}
